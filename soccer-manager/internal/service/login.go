package service

import (
	"context"
	"math/rand"
	"net/mail"
	"protobuf-v1/golang"
	grpcLogin "protobuf-v1/golang/login"
	grpcPlayer "protobuf-v1/golang/player"
	"soccer-manager/internal/db"
	"soccer-manager/internal/model"
	"soccer-manager/util"
	"soccer-manager/util/config"
	grpcError "soccer-manager/util/error"
	"soccer-manager/util/id"
	"soccer-manager/util/jwt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type login struct {
	userCollection   *mongo.Collection
	playerCollection *mongo.Collection
	teamCollection   *mongo.Collection
	mongoClient      *mongo.Client
	asyncWaitGroup   AsyncWaitGroup
	grpcLogin.UnimplementedLoginServiceServer
}

func NewLoginService(userCollection *mongo.Collection, playerCollection *mongo.Collection, teamCollection *mongo.Collection, mongoClient *mongo.Client, asyncWaitGroup AsyncWaitGroup) grpcLogin.LoginServiceServer {
	return login{
		userCollection:   userCollection,
		teamCollection:   teamCollection,
		playerCollection: playerCollection,
		mongoClient:      mongoClient,
		asyncWaitGroup:   asyncWaitGroup,
	}
}

func (l login) Login(ctx context.Context, req *grpcLogin.LoginRequest) (*grpcLogin.LoginResponse, error) {
	_, err := mail.ParseAddress(req.Email)
	if err != nil {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INVALID_ARGS, "invalid email")
	}

	if len(req.Password) < 6 {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INVALID_ARGS, "password should be at least 6 characters")
	}

	user, err := l.getUser(ctx, req)
	if err != nil {
		return nil, err
	}

	if user == nil {
		user, err = l.createUser(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_AUTH_ERROR, "user exists; password not matching")
	}

	tokenExpirationSeconds := config.GetInt32("jwt.expirationSeconds")
	token, err := jwt.GenerateToken(ctx, config.GetString("JWT_KEY"), user.Id, user.TeamId, tokenExpirationSeconds)
	if err != nil {
		return nil, err
	}

	return &grpcLogin.LoginResponse{
		AccessToken:   token,
		TokenType:     "Bearer",
		UserId:        user.Id.String(),
		TokenValidity: timestamppb.New(time.Now().Add(time.Second * time.Duration(tokenExpirationSeconds))),
	}, nil
}

func (l login) ValidateJWT(ctx context.Context, req *grpcLogin.ValidateJWTRequest) (*grpcLogin.ValidateJWTResponse, error) {
	if req.Jwt == "" {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_TOKEN_INVALID, "token missing")
	}

	claims := &jwt.Claims{}
	err := jwt.ParseToken(ctx, req.Jwt, claims, config.GetString("JWT_KEY"))
	if err != nil {
		return nil, err
	}
	return &grpcLogin.ValidateJWTResponse{
		UserId: claims.UserId,
		TeamId: claims.TeamId,
	}, nil
}

func (l login) getUser(ctx context.Context, req *grpcLogin.LoginRequest) (*model.User, error) {
	where := map[string]interface{}{}
	where["email"] = req.Email
	userResp, err := db.NewUserDbManager(l.userCollection).Find(ctx, where)
	if err != nil {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, "unable to find users")
	}

	if len(userResp) > 1 {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, "multiple users exists with this email")
	}

	if len(userResp) == 0 {
		return nil, nil
	}

	return userResp[0], nil
}

func (l login) createUser(ctx context.Context, req *grpcLogin.LoginRequest) (*model.User, error) {
	userID, err := id.NewUserID()
	if err != nil {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, err.Error())
	}

	teamID, err := id.NewTeamID()
	if err != nil {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, err.Error())
	}

	password, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, err.Error())
	}

	user := &model.User{
		Id:       userID,
		TeamId:   teamID,
		Email:    req.Email,
		Password: password,
	}
	userResp, err := db.NewUserDbManager(l.userCollection).Create(ctx, user)
	if err != nil {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, "unable to create user")
	}

	l.asyncWaitGroup.Add(1)
	go l.createTeam(ctx, userID, teamID)
	return userResp, nil
}

func (l login) createTeam(ctx context.Context, userID id.UserID, teamID id.TeamID) error {
	defer l.asyncWaitGroup.Done()

	//create team
	teamValue := config.GetInt64("team.value")
	teamBudget := config.GetInt64("team.budget")
	teamModel := &model.Team{
		Id:       teamID,
		UserId:   userID,
		Value:    &teamValue,
		Budget:   &teamBudget,
		Currency: golang.Currency_CURRENCY_USD,
	}
	_, err := db.NewTeamDbManager(l.teamCollection).Create(ctx, teamModel)
	if err != nil {
		return err
	}

	//create players
	err = l.createTeamPlayers(ctx, teamID)
	if err != nil {
		return err
	}

	return nil
}

func (l login) createTeamPlayers(ctx context.Context, teamID id.TeamID) error {
	var err error
	for i := 1; i <= config.GetInt("team.goalKeepers"); i++ {
		err = l.createPlayer(ctx, teamID, grpcPlayer.PlayerType_PT_GOAL_KEEPER)
		if err != nil {
			return err
		}
	}
	for i := 1; i <= config.GetInt("team.defenders"); i++ {
		err = l.createPlayer(ctx, teamID, grpcPlayer.PlayerType_PT_DEFENDER)
		if err != nil {
			return err
		}
	}
	for i := 1; i <= config.GetInt("team.midFielders"); i++ {
		err = l.createPlayer(ctx, teamID, grpcPlayer.PlayerType_PT_MID_FIELDER)
		if err != nil {
			return err
		}
	}
	for i := 1; i <= config.GetInt("team.attackers"); i++ {
		err = l.createPlayer(ctx, teamID, grpcPlayer.PlayerType_PT_ATTACKER)
		if err != nil {
			return err
		}
	}
	return nil
}

func (l login) createPlayer(ctx context.Context, teamID id.TeamID, playerType grpcPlayer.PlayerType) error {
	//create team
	playerValue := config.GetInt64("player.value")
	playerMaxAge := config.GetInt32("player.maxAge")
	playerMinAge := config.GetInt32("player.minAge")
	playerId, err := id.NewPlayerID()
	if err != nil {
		return err
	}
	playerModel := &model.Player{
		Id:       playerId,
		TeamId:   teamID,
		Age:      rand.Int31n(playerMaxAge-playerMinAge) + playerMinAge,
		Value:    &playerValue,
		Type:     playerType,
		Currency: golang.Currency_CURRENCY_USD,
	}
	_, err = db.NewPlayerDbManager(l.playerCollection).Create(ctx, playerModel)
	return err
}
