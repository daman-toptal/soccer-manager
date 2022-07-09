package model

import (
	"protobuf-v1/golang"
	grpcTxn "protobuf-v1/golang/transaction"
	"soccer-manager/util/id"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Transaction struct {
	Id          id.TransactionID        `bson:"_id"`
	TeamId      id.TeamID               `bson:"teamId"`
	PlayerId    id.PlayerID             `bson:"playerId"`
	Title       string                  `bson:"title"`
	Description string                  `bson:"description"`
	Amount      int64                   `bson:"amount"`
	Budget      int64                   `bson:"budget"`
	Type        grpcTxn.TransactionType `bson:"type"`
	CreatedAt   time.Time               `bson:"createdAt"`
	Currency    golang.Currency         `bson:"currency"`
}

func (t Transaction) ToProto() *grpcTxn.Transaction {
	return &grpcTxn.Transaction{
		Id:          t.Id.String(),
		Title:       t.Title,
		Description: t.Description,
		TeamId:      t.TeamId.String(),
		Amount:      t.Amount,
		Budget:      t.Budget,
		PlayerId:    t.PlayerId.String(),
		Type:        t.Type,
		CreatedAt:   timestamppb.New(t.CreatedAt),
		Currency:    t.Currency,
	}
}
