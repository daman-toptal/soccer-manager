package handler

import (
	"io/ioutil"
	"net/http"
	grpcRoot "protobuf-v1/golang"
	grpcUserApi "protobuf-v1/golang/external/user"
	grpcUser "protobuf-v1/golang/user"
	grpcError "soccer-manager/util/error"
	"soccer-manager/util/id"
	"soccer-manager/util/router"

	"github.com/go-chi/chi"
	grpcCodes "google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
)

func (c clientController) GetUser(w http.ResponseWriter, r *http.Request) {
	req := new(grpcUser.GetRequest)
	req.Id = chi.URLParam(r, ParamUserID)

	headerUserID := router.NewHeader(r.Context()).GetUserID()

	//access check
	if req.Id != headerUserID.String() {
		router.RenderHttpError(w, r, grpcError.NewHttpErrorFromError(r.Method, grpcError.NewError(r.Context(), grpcRoot.Error_ERROR_AUTH_ERROR)))
		return
	}

	user, err := c.uc.Get(r.Context(), req)
	if err != nil {
		router.RenderHttpError(w, r, grpcError.NewHttpErrorFromError(r.Method, err))
		return
	}

	apiResp := c.getUserApiResponse(user)
	resp := router.Response{
		Writer:   w,
		Status:   http.StatusOK,
		GRPCData: apiResp,
	}

	router.RenderJSON(resp)
}

func (c clientController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	req := new(grpcUserApi.UpdateRequest)
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

	userId, err := id.ParseUserID(chi.URLParam(r, ParamUserID))
	if err != nil {
		router.RenderHttpError(w, r, grpcError.NewRESTError(r.Method, grpcCodes.InvalidArgument, grpcRoot.Error_ERROR_INVALID_ID, err.Error()))
		return
	}

	headerUserID := router.NewHeader(r.Context()).GetUserID()

	//access check
	if userId != headerUserID {
		router.RenderHttpError(w, r, grpcError.NewHttpErrorFromError(r.Method, grpcError.NewError(r.Context(), grpcRoot.Error_ERROR_AUTH_ERROR)))
		return
	}

	grpcReq := &grpcUser.UpdateRequest{
		Id:   userId.String(),
		Name: req.Name,
	}

	user, err := c.uc.Update(r.Context(), grpcReq)
	if err != nil {
		router.RenderHttpError(w, r, grpcError.NewHttpErrorFromError(r.Method, err))
		return
	}

	apiResp := c.getUserApiResponse(user)
	resp := router.Response{
		Writer:   w,
		Status:   http.StatusOK,
		GRPCData: apiResp,
	}

	router.RenderJSON(resp)
}

func (c clientController) getUserApiResponse(user *grpcUser.User) *grpcUserApi.User {
	return &grpcUserApi.User{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		TeamId:    user.TeamId,
		CreatedAt: user.CreatedAt,
	}
}
