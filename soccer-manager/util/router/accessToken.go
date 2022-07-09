package router

import (
	"context"
	"net/http"
	grpcLogin "protobuf-v1/golang/login"
	grpcError "soccer-manager/util/error"
	"strings"

	"google.golang.org/grpc/metadata"
)

var PathAuthRequired = map[string]bool{
	"/v1/login": false, // public url
	"health":    false, // public url
}

func IsAuthRequired(url string) bool {
	path := strings.TrimRight(url, "/")

	if v, ok := PathAuthRequired[path]; ok {
		return v
	}

	return true
}

func accessTokenAuthMiddleware(lc grpcLogin.LoginServiceClient) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			var mdPairs []string
			var err error
			ctx := r.Context()
			if !IsAuthRequired(r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}
			mdPairs, err = validateJWT(lc, r)
			if err != nil {
				httpErr := grpcError.NewHttpErrorFromError(r.Method, err)
				RenderHttpError(w, r, httpErr)
				return
			}
			ctx = getContext(ctx, mdPairs)
			updatedReq := r.WithContext(ctx)
			next.ServeHTTP(w, updatedReq)
		}
		return http.HandlerFunc(fn)
	}
}

func validateJWT(lc grpcLogin.LoginServiceClient, r *http.Request) ([]string, error) {
	var mdPairs []string
	authHeader := r.Header.Get(HeaderAuthorization)
	req := &grpcLogin.ValidateJWTRequest{
		Jwt: strings.TrimPrefix(strings.ReplaceAll(authHeader, " ", ""), "Bearer"),
	}

	resp, err := lc.ValidateJWT(r.Context(), req)
	if err != nil {
		return mdPairs, err
	}

	mdPairs = append(
		mdPairs,
		HeaderAuthorization, authHeader,
		HeaderUserID, resp.UserId,
		HeaderTeamID, resp.TeamId,
	)
	return mdPairs, nil
}

func getContext(ctx context.Context, mdPairs []string) context.Context {
	md := metadata.Join(MetadataFromIncoming(ctx), metadata.Pairs(mdPairs...))
	ctx = context.WithValue(ctx, LocalMetadataKey, md)
	ctx = metadata.AppendToOutgoingContext(ctx, mdPairs...)
	return ctx
}

func MetadataFromIncoming(ctx context.Context) metadata.MD {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md, ok = ctx.Value(LocalMetadataKey).(metadata.MD)
		if !ok || md == nil {
			return metadata.New(map[string]string{})
		}
	}

	return md
}
