package service

import (
	"context"
	"soccer-manager/internal/model"

	"protobuf-v1/golang"
	grpcUser "protobuf-v1/golang/user"
	"soccer-manager/internal/db"
	grpcError "soccer-manager/util/error"
	"soccer-manager/util/id"

	"go.mongodb.org/mongo-driver/mongo"
)

type user struct {
	collection *mongo.Collection
	grpcUser.UnimplementedUserServiceServer
}

func NewUserService(collection *mongo.Collection) grpcUser.UserServiceServer {
	return user{
		collection: collection,
	}
}

func (u user) Get(ctx context.Context, req *grpcUser.GetRequest) (*grpcUser.User, error) {

	userId, err := id.ParseUserID(req.Id)
	if err != nil {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, err.Error())
	}

	userResp, err := db.NewUserDbManager(u.collection).Get(ctx, userId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, grpcError.NewError(ctx, golang.Error_ERROR_NOT_FOUND)
		}
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, err.Error())
	}

	return userResp.ToProto(), nil
}

func (u user) Update(ctx context.Context, req *grpcUser.UpdateRequest) (*grpcUser.User, error) {

	userId, err := id.ParseUserID(req.Id)
	if err != nil {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, err.Error())
	}

	if userId.IsZero() {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INVALID_ARGS, "userId can not be blank")
	}

	userResp, err := db.NewUserDbManager(u.collection).Update(ctx, &model.User{Id: userId, Name: req.Name})
	if err != nil {
		return nil, grpcError.NewError(ctx, golang.Error_ERROR_INTERNAL_ERROR, err.Error())
	}

	return userResp.ToProto(), nil
}
