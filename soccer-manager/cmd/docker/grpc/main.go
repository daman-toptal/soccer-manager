package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"soccer-manager/internal/service"
	"soccer-manager/util/config"
	"soccer-manager/util/logging"
	"soccer-manager/util/signal"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoClient *mongo.Client
	asyncWg     service.AsyncWaitGroup
)

func setup() {
	config.SetupConfig()
	asyncWg = service.NewAsyncWaitGroupService()
	signal.SetupSignals()
	logging.SetupLogging(config.GetString("log.level"))
	setupMongo()
}

func initialise() {
	initCollections()
	initGRPCServices()
	initGRPCServer()
}

func run() {
	signal.CleanupOnSignal(cleanUp)
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", config.GetString("server.grpcPort")))
	if err != nil {
		log.Fatal("Something went wrong with the service")
	}
	err = server.Serve(l)
	if err != nil {
		log.Fatal("Something went wrong with the service")
	}
}

func main() {
	setup()
	initialise()
	run()
}

func cleanUp() {
	asyncWg.Wait()
	mongoClient.Disconnect(context.Background())
}
