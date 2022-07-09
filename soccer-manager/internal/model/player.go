package model

import (
	"protobuf-v1/golang"
	grpcPlayer "protobuf-v1/golang/player"
	"soccer-manager/util/id"
	"time"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Player struct {
	Id                        id.PlayerID             `bson:"_id"`
	TeamId                    id.TeamID               `bson:"teamId"`
	FirstName                 string                  `bson:"firstName"`
	LastName                  string                  `bson:"lastName"`
	Country                   string                  `bson:"country"`
	Age                       int32                   `bson:"age"`
	Value                     *int64                  `bson:"value"`
	Type                      grpcPlayer.PlayerType   `bson:"type"`
	IsListed                  *bool                   `bson:"isListed"`
	AskValue                  *int64                  `bson:"askValue"`
	Currency                  golang.Currency         `bson:"currency"`
	CreatedAt                 time.Time               `bson:"createdAt"`
}

func (p Player) ToProto() *grpcPlayer.Player {
	player := &grpcPlayer.Player{
		Id:        p.Id.String(),
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Age:       p.Age,
		Type:      p.Type,
		Country:   p.Country,
		TeamId:    p.TeamId.String(),
		Currency:  p.Currency,
	}

	if p.Value != nil {
		player.Value = *p.Value
	}

	if p.IsListed != nil {
		player.IsListed = *p.IsListed
	}

	if p.AskValue != nil {
		player.AskValue = &wrapperspb.Int64Value{Value: *p.AskValue}
	}

	return player
}
