package model

import (
	grpcUser "protobuf-v1/golang/user"
	"soccer-manager/util/id"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type User struct {
	Id                        id.UserID               `bson:"_id"`
	TeamId                    id.TeamID               `bson:"teamId"`
	Name                      string                  `bson:"name"`
	Email                     string                  `bson:"email"`
	Password                  string                  `bson:"password"`
	CreatedAt                 time.Time               `bson:"createdAt"`
}

func (u User) ToProto() *grpcUser.User {
	return &grpcUser.User{
		Id:          u.Id.String(),
		Name:        u.Name,
		Email:       u.Email,
		TeamId:      u.TeamId.String(),
		CreatedAt:   timestamppb.New(u.CreatedAt),
	}
}
