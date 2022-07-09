package handler

import (
	"io/ioutil"
	"net/http"
	grpcRoot "protobuf-v1/golang"
	grpcLoginApi "protobuf-v1/golang/external/login"
	grpcLogin "protobuf-v1/golang/login"
	grpcError "soccer-manager/util/error"
	"soccer-manager/util/router"

	grpcCodes "google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
)

func (c clientController) Login(w http.ResponseWriter, r *http.Request) {
	req := new(grpcLoginApi.LoginRequest)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		router.RenderHttpError(w, r, grpcError.NewRESTError(r.Method, grpcCodes.InvalidArgument, grpcRoot.Error_ERROR_INTERNAL_ERROR, err.Error()))
		return
	}

	err = protojson.UnmarshalOptions{}.Unmarshal(body, req)
	if err != nil {
		router.RenderHttpError(w, r, grpcError.NewRESTError(r.Method, grpcCodes.InvalidArgument, grpcRoot.Error_ERROR_INTERNAL_ERROR, err.Error()))
		return
	}

	if req.Email == "" || req.Password == "" {
		router.RenderHttpError(w, r, grpcError.NewRESTError(r.Method, grpcCodes.InvalidArgument, grpcRoot.Error_ERROR_AUTH_ERROR, "invalid email/ password"))
		return
	}

	grpcReq := &grpcLogin.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	loginResp, err := c.lc.Login(r.Context(), grpcReq)
	if err != nil {
		router.RenderHttpError(w, r, grpcError.NewHttpErrorFromError(r.Method, err))
		return
	}

	apiResp := c.getLoginApiResponse(loginResp)
	resp := router.Response{
		Writer:   w,
		Status:   http.StatusOK,
		GRPCData: apiResp,
	}

	router.RenderJSON(resp)
}

func (c clientController) getLoginApiResponse(resp *grpcLogin.LoginResponse) *grpcLoginApi.LoginResponse {
	return &grpcLoginApi.LoginResponse{
		AccessToken:   resp.AccessToken,
		TokenType:     resp.TokenType,
		UserId:        resp.UserId,
		TokenValidity: resp.TokenValidity,
	}
}
