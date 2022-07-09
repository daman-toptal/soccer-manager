package model

import (
	"protobuf-v1/golang"
	grpcTeam "protobuf-v1/golang/team"
	"soccer-manager/util/id"
	"time"
)

type Team struct {
	Id                        id.TeamID               `bson:"_id"`
	UserId                    id.UserID               `bson:"userId"`
	Name                      string                  `bson:"name"`
	Country                   string                  `bson:"country"`
	Value                     *int64                  `bson:"value"`
	Budget                    *int64                  `bson:"budget"`
	Currency                  golang.Currency         `bson:"currency"`
	CreatedAt                 time.Time               `bson:"createdAt"`
}

func (t Team) ToProto() *grpcTeam.Team {
	team := &grpcTeam.Team{
		Id:       t.Id.String(),
		Name:     t.Name,
		Country:  t.Country,
		UserId:   t.UserId.String(),
		Currency: t.Currency,
	}

	if t.Value != nil {
		team.Value = *t.Value
	}

	if t.Budget != nil {
		team.Budget = *t.Budget
	}

	return team
}
