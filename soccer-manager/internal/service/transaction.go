package service

import (
	"context"
	"math/rand"
	"protobuf-v1/golang"
	grpcTxn "protobuf-v1/golang/transaction"
	"soccer-manager/internal/db"
	"soccer-manager/internal/model"
	"soccer-manager/util"
	grpcError "soccer-manager/util/error"
	"soccer-manager/util/id"
	"soccer-manager/util/logging"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type transaction struct {
	txnCollection    *mongo.Collection
	playerCollection *mongo.Collection
	teamCollection   *mongo.Collection
	mongoClient      *mongo.Client
	grpcTxn.UnimplementedTransactionServiceServer
}

type createTransactionRequest struct {
	teamModel   *model.Team
	amount      int64
	txnType     grpcTxn.TransactionType
	playerModel *model.Player
	description string
}

type updatePlayersAndTeamResponse struct {
	newPlayer   *model.Player
	newDestTeam *model.Team
	newSrcTeam  *model.Team
}

func NewTransactionService(txnCollection *mongo.Collection, playerCollection *mongo.Collection, teamCollection *mongo.Collection, mongoClient *mongo.Client) grpcTxn.TransactionServiceServer {
	return transaction{
		txnCollection:    txnCollection,
		teamCollection:   teamCollection,
		playerCollection: playerCollection,
		mongoClient:      mongoClient,
	}
}

func (t transaction) Get(ctx context.Context, req *grpcTxn.GetRequest) (*grpcTxn.Transaction, error) {

	transactionId, err := id.ParseTransactionID(req.Id)
	if err != nil {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, err.Error())
	}

	transactionResp, err := db.NewTransactionDbManager(t.txnCollection).Get(ctx, transactionId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, grpcError.NewError(ctx, golang.Error_ERROR_NOT_FOUND)
		}
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, err.Error())
	}

	return transactionResp.ToProto(), nil
}

func (t transaction) GetByTeam(ctx context.Context, req *grpcTxn.GetByTeamRequest) (*grpcTxn.Transactions, error) {

	teamId, err := id.ParseTeamID(req.TeamId)
	if err != nil {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INVALID_ID, err.Error())
	}

	where := map[string]interface{}{}
	where["teamId"] = teamId

	transactionResp, err := db.NewTransactionDbManager(t.txnCollection).Find(ctx, where)
	if err != nil {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, err.Error())
	}

	transactionsResp := &grpcTxn.Transactions{}
	for _, transaction := range transactionResp {
		transactionsResp.Transactions = append(transactionsResp.Transactions, transaction.ToProto())
		transactionsResp.Total++
	}
	return transactionsResp, nil
}

func (t transaction) Buy(ctx context.Context, req *grpcTxn.BuyRequest) (*grpcTxn.Transaction, error) {

	playerId, err := id.ParsePlayerID(req.PlayerId)
	if err != nil {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INVALID_ID, err.Error())
	}

	destTeamId, err := id.ParseTeamID(req.TeamId)
	if err != nil {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INVALID_ID, err.Error())
	}

	//checks - player status, ask value and team budget
	oldPlayer, err := db.NewPlayerDbManager(t.playerCollection).Get(ctx, playerId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, grpcError.NewError(ctx, golang.Error_ERROR_NOT_FOUND)
		}
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, err.Error())
	}

	if oldPlayer.IsListed == nil || *oldPlayer.IsListed == false || oldPlayer.AskValue == nil || *oldPlayer.AskValue == 0 {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_AUTH_ERROR, "player is not listed")
	}

	if oldPlayer.TeamId.String() == req.TeamId {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_AUTH_ERROR, "cannot buy own player")
	}

	destTeam, err := db.NewTeamDbManager(t.teamCollection).Get(ctx, destTeamId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, grpcError.NewError(ctx, golang.Error_ERROR_NOT_FOUND)
		}
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, err.Error())
	}

	if destTeam.Budget != nil && *destTeam.Budget < *oldPlayer.AskValue {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, "not enough budget")
	}

	srcTeam, err := db.NewTeamDbManager(t.teamCollection).Get(ctx, oldPlayer.TeamId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, grpcError.NewError(ctx, golang.Error_ERROR_NOT_FOUND)
		}
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, err.Error())
	}

	updatePlayerAndTeamsResp, err := t.updatePlayerAndTeams(ctx, oldPlayer, destTeam, srcTeam)
	if err != nil {
		return nil, err
	}

	_, destTxn, err := t.createTransactions(ctx, updatePlayerAndTeamsResp.newSrcTeam, updatePlayerAndTeamsResp.newDestTeam, updatePlayerAndTeamsResp.newPlayer, req.Description, *oldPlayer.AskValue)
	if err != nil {
		return nil, err
	}

	return destTxn.ToProto(), nil
}

func getNewPlayerValue(value int64) int64 {
	lo := value + value/10
	hi := value * 2
	return rand.Int63n(hi-lo) + lo
}

func (t transaction) updatePlayerAndTeams(ctx context.Context, oldPlayer *model.Player, destTeam *model.Team, srcTeam *model.Team) (*updatePlayersAndTeamResponse, error) {
	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Snapshot()
	txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)

	mongoSession, err := t.mongoClient.StartSession()
	if err != nil {
		logging.Error("failed to get mongo session", logging.Fields{"error": err.Error()})
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, err.Error())
	}
	defer mongoSession.EndSession(ctx)

	callback := func(sessionContext mongo.SessionContext) (interface{}, error) {

		//update - player status (check old listed, ask value, team and value) (update listed, value, team)
		playerNewValue := getNewPlayerValue(*oldPlayer.Value)
		playerNewListed := false
		playerUpdateModel := &model.Player{
			Id:       oldPlayer.Id,
			TeamId:   destTeam.Id,
			Value:    &playerNewValue,
			IsListed: &playerNewListed,
		}
		playerFilters := map[string]interface{}{}
		playerFilters["isListed"] = oldPlayer.IsListed
		playerFilters["askValue"] = oldPlayer.AskValue
		playerFilters["teamId"] = oldPlayer.TeamId
		playerFilters["value"] = oldPlayer.Value

		newPlayer, err := db.NewPlayerDbManager(t.playerCollection).Update(sessionContext, playerUpdateModel, playerFilters)
		if err != nil {
			logging.Error("failed to update player", logging.Fields{"playerId": oldPlayer.Id.String()})
			return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, err.Error())
		}

		//update - new team (check budget, value) (update budget, value)
		destTeamNewValue := *destTeam.Value + *newPlayer.Value
		destTeamNewBudget := *destTeam.Budget - *oldPlayer.AskValue
		destTeamUpdateModel := &model.Team{
			Id:     destTeam.Id,
			Value:  &destTeamNewValue,
			Budget: &destTeamNewBudget,
		}
		destTeamFilters := map[string]interface{}{}
		destTeamFilters["value"] = destTeam.Value
		destTeamFilters["budget"] = destTeam.Budget

		newDestTeam, err := db.NewTeamDbManager(t.teamCollection).Update(sessionContext, destTeamUpdateModel, destTeamFilters)
		if err != nil {
			logging.Error("failed to update dest team", logging.Fields{"teamId": destTeam.Id.String()})
			return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, err.Error())
		}

		//update - old team (check value) (update value)
		srcTeamNewValue := *srcTeam.Value - *oldPlayer.Value
		srcTeamNewBudget := *srcTeam.Budget + *oldPlayer.AskValue
		srcTeamUpdateModel := &model.Team{
			Id:     srcTeam.Id,
			Value:  &srcTeamNewValue,
			Budget: &srcTeamNewBudget,
		}
		srcTeamFilters := map[string]interface{}{}
		srcTeamFilters["value"] = srcTeam.Value
		srcTeamFilters["budget"] = srcTeam.Budget

		newSrcTeam, err := db.NewTeamDbManager(t.teamCollection).Update(sessionContext, srcTeamUpdateModel, srcTeamFilters)
		if err != nil {
			logging.Error("failed to update src team", logging.Fields{"teamId": srcTeam.Id.String()})
			return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, err.Error())
		}

		return &updatePlayersAndTeamResponse{
			newPlayer:   newPlayer,
			newDestTeam: newDestTeam,
			newSrcTeam:  newSrcTeam,
		}, nil
	}

	result, err := mongoSession.WithTransaction(ctx, callback, txnOpts)
	if err != nil {
		logging.Error("transaction failed to error", logging.Fields{"error": err.Error()})
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, "buy failed due to internal error")
	}

	resp, ok := result.(*updatePlayersAndTeamResponse)
	if !ok || resp.newPlayer == nil || resp.newSrcTeam == nil || resp.newDestTeam == nil {
		logging.Error("transaction returned invalid response")
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, "buy failed due to internal error")
	}

	return resp, nil
}

func (t transaction) createTransactions(ctx context.Context, newSrcTeam *model.Team, newDestTeam *model.Team, newPlayer *model.Player, description string, askValue int64) (*model.Transaction, *model.Transaction, error) {

	// create source transaction
	srcTxn, err := t.createTransaction(ctx, &createTransactionRequest{
		teamModel:   newSrcTeam,
		amount:      askValue,
		txnType:     grpcTxn.TransactionType_TT_SELL,
		playerModel: newPlayer,
		description: description,
	})
	if err != nil {
		return nil, nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, err.Error())
	}

	// create dest transaction
	destTxn, err := t.createTransaction(ctx, &createTransactionRequest{
		teamModel:   newDestTeam,
		amount:      askValue,
		txnType:     grpcTxn.TransactionType_TT_BUY,
		playerModel: newPlayer,
		description: description,
	})
	if err != nil {
		return nil, nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, err.Error())
	}

	return srcTxn, destTxn, nil
}

func (t transaction) createTransaction(ctx context.Context, req *createTransactionRequest) (*model.Transaction, error) {
	txnId, err := id.NewTransactionID()
	if err != nil {
		return nil, err
	}
	txnModel := &model.Transaction{
		Id:          txnId,
		TeamId:      req.teamModel.Id,
		PlayerId:    req.playerModel.Id,
		Title:       string(util.TransactionTypeFromProto[req.txnType]) + " Player",
		Description: req.description,
		Amount:      req.amount,
		Budget:      *req.teamModel.Budget,
		Type:        req.txnType,
		Currency:    req.teamModel.Currency,
	}
	txnModel, err = db.NewTransactionDbManager(t.txnCollection).Create(ctx, txnModel)
	if err != nil {
		return nil, err
	}
	return txnModel, nil
}
