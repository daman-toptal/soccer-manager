package main

import (
	"context"
	"log"
	"os"
	"soccer-manager/util/logging"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateClient(ctx context.Context, uri string, retry int32) (*mongo.Client, error) {
	time.Sleep(time.Second * 5)
	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err := conn.Ping(ctx, nil); err != nil {
		logging.Error("db connection failed", logging.Fields{"error": err.Error()})
		if retry == 0 {
			return nil, err
		}
		retry = retry - 1
		return CreateClient(ctx, uri, retry)
	}

	return conn, err
}

func setupMongo() {
	uri := os.Getenv("DB_HOST")
	var err error
	mongoClient, err = CreateClient(context.Background(), uri, 1)
	if err != nil {
		log.Panic(err)
	}
}
