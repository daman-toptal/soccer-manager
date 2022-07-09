package main

import (
	"context"
	grpcLogin "protobuf-v1/golang/login"
	grpcPlayer "protobuf-v1/golang/player"
	grpcTeam "protobuf-v1/golang/team"
	grpcTransaction "protobuf-v1/golang/transaction"
	grpcUser "protobuf-v1/golang/user"
	"soccer-manager/internal/service"
	"time"

	ggrpc "google.golang.org/grpc"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/keepalive"
)

var (
	mongoDatabase         *mongo.Database
	userCollection        *mongo.Collection
	teamCollection        *mongo.Collection
	playerCollection      *mongo.Collection
	transactionCollection *mongo.Collection
	server                *ggrpc.Server
	loginServer           grpcLogin.LoginServiceServer
	userServer            grpcUser.UserServiceServer
	playerServer          grpcPlayer.PlayerServiceServer
	teamServer            grpcTeam.TeamServiceServer
	transactionService    grpcTransaction.TransactionServiceServer
)

func initCollections() {
	mongoDatabase = mongoClient.Database("soccer")

	userCollection = mongoDatabase.Collection("users")
	mongoDatabase.RunCommand(context.TODO(), bson.D{{"create", "users"}})

	playerCollection = mongoDatabase.Collection("players")
	mongoDatabase.RunCommand(context.TODO(), bson.D{{"create", "players"}})

	teamCollection = mongoDatabase.Collection("teams")
	mongoDatabase.RunCommand(context.TODO(), bson.D{{"create", "teams"}})

	transactionCollection = mongoDatabase.Collection("transactions")
	mongoDatabase.RunCommand(context.TODO(), bson.D{{"create", "transactions"}})
}

func initGRPCServices() {
	loginServer = service.NewLoginService(userCollection, playerCollection, teamCollection, mongoClient, asyncWg)
	userServer = service.NewUserService(userCollection)
	playerServer = service.NewPlayerService(playerCollection)
	teamServer = service.NewTeamService(teamCollection)
	transactionService = service.NewTransactionService(transactionCollection, playerCollection, teamCollection, mongoClient)
}

func initGRPCServer() {
	server = ggrpc.NewServer(
		ggrpc.KeepaliveParams(keepalive.ServerParameters{
			Timeout: 100 * time.Second,
		}),
	)
	registerGRPCServerServices()
}

func registerGRPCServerServices() {
	grpcLogin.RegisterLoginServiceServer(server, loginServer)
	grpcUser.RegisterUserServiceServer(server, userServer)
	grpcPlayer.RegisterPlayerServiceServer(server, playerServer)
	grpcTeam.RegisterTeamServiceServer(server, teamServer)
	grpcTransaction.RegisterTransactionServiceServer(server, transactionService)
}
