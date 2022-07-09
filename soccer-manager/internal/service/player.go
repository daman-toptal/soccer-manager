package service

import (
	"context"
	"protobuf-v1/golang"
	grpcPlayer "protobuf-v1/golang/player"
	"soccer-manager/internal/db"
	"soccer-manager/internal/model"
	grpcError "soccer-manager/util/error"
	"soccer-manager/util/id"

	"go.mongodb.org/mongo-driver/mongo"
)

type player struct {
	collection *mongo.Collection
	grpcPlayer.UnimplementedPlayerServiceServer
}

func NewPlayerService(collection *mongo.Collection) grpcPlayer.PlayerServiceServer {
	return player{
		collection: collection,
	}
}

func (p player) Get(ctx context.Context, req *grpcPlayer.GetRequest) (*grpcPlayer.Player, error) {

	playerId, err := id.ParsePlayerID(req.Id)
	if err != nil {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, err.Error())
	}

	playerResp, err := db.NewPlayerDbManager(p.collection).Get(ctx, playerId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, grpcError.NewError(ctx, golang.Error_ERROR_NOT_FOUND)
		}
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, err.Error())
	}

	return playerResp.ToProto(), nil
}

func (p player) GetByTeam(ctx context.Context, req *grpcPlayer.GetByTeamRequest) (*grpcPlayer.Players, error) {

	teamId, err := id.ParseTeamID(req.TeamId)
	if err != nil {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INVALID_ID, err.Error())
	}

	where := map[string]interface{}{}
	where["teamId"] = teamId

	playerResp, err := db.NewPlayerDbManager(p.collection).Find(ctx, where)
	if err != nil {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, err.Error())
	}

	playersResp := &grpcPlayer.Players{}
	for _, player := range playerResp {
		playersResp.Players = append(playersResp.Players, player.ToProto())
		playersResp.Total++
	}
	return playersResp, nil
}

func (p player) GetListed(ctx context.Context, req *grpcPlayer.GetListedRequest) (*grpcPlayer.Players, error) {

	where := map[string]interface{}{}
	where["isListed"] = true

	playerResp, err := db.NewPlayerDbManager(p.collection).Find(ctx, where)
	if err != nil {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, err.Error())
	}

	playersResp := &grpcPlayer.Players{}
	for _, player := range playerResp {
		playersResp.Players = append(playersResp.Players, player.ToProto())
		playersResp.Total++
	}
	return playersResp, nil
}

func (p player) Update(ctx context.Context, req *grpcPlayer.UpdateRequest) (*grpcPlayer.Player, error) {

	playerId, err := id.ParsePlayerID(req.Id)
	if err != nil {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, err.Error())
	}

	if playerId.IsZero() {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INVALID_ARGS, "playerId can not be blank")
	}

	player, err := p.Get(ctx, &grpcPlayer.GetRequest{Id: playerId.String()})
	if err != nil {
		return nil, err
	}

	if req.IsListed != nil && req.IsListed.GetValue() && req.AskValue == nil && player.AskValue == nil {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INVALID_ARGS, "askValue can not be blank")
	}

	updateModel := &model.Player{
		Id:        playerId,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Country:   req.Country,
	}

	if req.AskValue != nil {
		updateModel.AskValue = &req.AskValue.Value
		if req.AskValue.GetValue() <= 0 {
			return nil, grpcError.NewError(ctx, golang.Error_ERROR_INVALID_ARGS, "askValue should be greater than 0")
		}
	}

	if req.IsListed != nil {
		updateModel.IsListed = &req.IsListed.Value
	}

	if req.Value != nil {
		updateModel.Value = &req.Value.Value
	}

	playerResp, err := db.NewPlayerDbManager(p.collection).Update(ctx, updateModel)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, grpcError.NewError(ctx, golang.Error_ERROR_NOT_FOUND)
		}
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, err.Error())
	}

	return playerResp.ToProto(), nil
}
