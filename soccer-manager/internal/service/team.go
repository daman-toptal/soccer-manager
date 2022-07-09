package service

import (
	"context"
	"soccer-manager/internal/model"

	"protobuf-v1/golang"
	grpcTeam "protobuf-v1/golang/team"
	"soccer-manager/internal/db"
	grpcError "soccer-manager/util/error"
	"soccer-manager/util/id"

	"go.mongodb.org/mongo-driver/mongo"
)

type team struct {
	collection *mongo.Collection
	grpcTeam.UnimplementedTeamServiceServer
}

func NewTeamService(collection *mongo.Collection) grpcTeam.TeamServiceServer {
	return team{
		collection: collection,
	}
}

func (t team) Get(ctx context.Context, req *grpcTeam.GetRequest) (*grpcTeam.Team, error) {

	teamId, err := id.ParseTeamID(req.Id)
	if err != nil {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, err.Error())
	}

	teamResp, err := db.NewTeamDbManager(t.collection).Get(ctx, teamId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, grpcError.NewError(ctx, golang.Error_ERROR_NOT_FOUND)
		}
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, err.Error())
	}

	return teamResp.ToProto(), nil
}

func (t team) Update(ctx context.Context, req *grpcTeam.UpdateRequest) (*grpcTeam.Team, error) {

	teamId, err := id.ParseTeamID(req.Id)
	if err != nil {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, err.Error())
	}

	if teamId.IsZero() {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INVALID_ARGS, "teamId can not be blank")
	}

	updateModel := &model.Team{
		Id:      teamId,
		Name:    req.Name,
		Country: req.Country,
	}

	if req.Value != nil {
		updateModel.Value = &req.Value.Value
	}

	if req.Budget != nil {
		updateModel.Budget = &req.Budget.Value
	}

	teamResp, err := db.NewTeamDbManager(t.collection).Update(ctx, updateModel)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, grpcError.NewError(ctx, golang.Error_ERROR_NOT_FOUND)
		}
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, err.Error())
	}

	return teamResp.ToProto(), nil
}
