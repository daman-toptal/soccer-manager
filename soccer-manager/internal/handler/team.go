package handler

import (
	"io/ioutil"
	"net/http"
	grpcRoot "protobuf-v1/golang"
	grpcTeamApi "protobuf-v1/golang/external/team"
	grpcTeam "protobuf-v1/golang/team"
	"soccer-manager/util"
	grpcError "soccer-manager/util/error"
	"soccer-manager/util/id"
	"soccer-manager/util/router"

	"github.com/go-chi/chi"
	grpcCodes "google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
)

func (c clientController) GetTeam(w http.ResponseWriter, r *http.Request) {
	req := new(grpcTeam.GetRequest)
	req.Id = chi.URLParam(r, ParamTeamID)

	headerTeamId := router.NewHeader(r.Context()).GetTeamID()

	//access check
	if req.Id != headerTeamId.String() {
		router.RenderHttpError(w, r, grpcError.NewHttpErrorFromError(r.Method, grpcError.NewError(r.Context(), grpcRoot.Error_ERROR_AUTH_ERROR)))
		return
	}

	team, err := c.tc.Get(r.Context(), req)
	if err != nil {
		router.RenderHttpError(w, r, grpcError.NewHttpErrorFromError(r.Method, err))
		return
	}

	apiResp := c.getTeamApiResponse(team)
	resp := router.Response{
		Writer:   w,
		Status:   http.StatusOK,
		GRPCData: apiResp,
	}

	router.RenderJSON(resp)
}

func (c clientController) UpdateTeam(w http.ResponseWriter, r *http.Request) {
	req := new(grpcTeamApi.UpdateRequest)
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

	teamId, err := id.ParseTeamID(chi.URLParam(r, ParamTeamID))
	if err != nil {
		router.RenderHttpError(w, r, grpcError.NewRESTError(r.Method, grpcCodes.InvalidArgument, grpcRoot.Error_ERROR_INVALID_ID, err.Error()))
		return
	}

	headerTeamId := router.NewHeader(r.Context()).GetTeamID()

	//access check
	if headerTeamId != teamId {
		router.RenderHttpError(w, r, grpcError.NewHttpErrorFromError(r.Method, grpcError.NewError(r.Context(), grpcRoot.Error_ERROR_AUTH_ERROR)))
		return
	}

	grpcReq := &grpcTeam.UpdateRequest{
		Id:      teamId.String(),
		Name:    req.Name,
		Country: req.Country,
	}

	team, err := c.tc.Update(r.Context(), grpcReq)
	if err != nil {
		router.RenderHttpError(w, r, grpcError.NewHttpErrorFromError(r.Method, err))
		return
	}

	apiResp := c.getTeamApiResponse(team)
	resp := router.Response{
		Writer:   w,
		Status:   http.StatusOK,
		GRPCData: apiResp,
	}

	router.RenderJSON(resp)
}

func (c clientController) getTeamApiResponse(team *grpcTeam.Team) *grpcTeamApi.Team {
	return &grpcTeamApi.Team{
		Id:       team.Id,
		Name:     team.Name,
		Country:  team.Country,
		Value:    util.ParseAmountToString(team.Value),
		Budget:   util.ParseAmountToString(team.Budget),
		UserId:   team.UserId,
		Currency: string(util.CurrencyFromProto[team.Currency]),
	}
}
