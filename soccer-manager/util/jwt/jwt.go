package jwt

import (
	"context"
	"protobuf-v1/golang"
	grpcError "soccer-manager/util/error"
	"soccer-manager/util/id"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserId string `json:"userId"`
	TeamId string `json:"teamId"`
	jwt.StandardClaims
}

func GenerateToken(ctx context.Context, jwtKey string, userId id.UserID, teamId id.TeamID, expirationSeconds int32) (string, error) {
	expirationTime := time.Now().Add(time.Second * time.Duration(expirationSeconds))
	claims := &Claims{
		UserId: userId.String(),
		TeamId: teamId.String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return tokenString, grpcError.NewError(ctx, golang.Error_ERROR_TOKEN_ERROR, err.Error())
	}
	return tokenString, nil
}

func ParseToken(ctx context.Context, jwtToken string, claims *Claims, jwtKey string) error {
	token, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return grpcError.NewError(ctx, golang.Error_ERROR_TOKEN_INVALID, err.Error())
		}
		return grpcError.NewError(ctx, golang.Error_ERROR_TOKEN_ERROR)
	}
	if !token.Valid {
		return grpcError.NewError(ctx, golang.Error_ERROR_TOKEN_INVALID, err.Error())
	}

	return nil
}
