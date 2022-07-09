package router

import (
	"context"
	"soccer-manager/util/id"

	"google.golang.org/grpc/metadata"
)

type Header interface {
	GetTeamID() id.TeamID
	GetUserID() id.UserID
}

type header struct {
	ctx context.Context
}

func (h *header) GetUserID() id.UserID {
	key, err := id.ParseUserID(HeaderValueFromIncoming(h.ctx, HeaderUserID))
	if err != nil {
		return id.UserID{}
	}
	return key
}

func (h *header) GetTeamID() id.TeamID {
	key, err := id.ParseTeamID(HeaderValueFromIncoming(h.ctx, HeaderTeamID))
	if err != nil {
		return id.TeamID{}
	}
	return key
}

func NewHeader(ctx context.Context) Header {
	return &header{ctx}
}

func HeaderValueFromIncoming(ctx context.Context, key string) string {
	var s string
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md, ok = ctx.Value(LocalMetadataKey).(metadata.MD)
		if !ok || md == nil {
			return s
		}
	}

	val := md[key]
	if len(val) > 0 {
		return val[0]
	}

	return s
}
