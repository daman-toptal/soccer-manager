package handler

import (
	"io/ioutil"
	"net/http"
	grpcRoot "protobuf-v1/golang"
	grpcPlayerApi "protobuf-v1/golang/external/player"
	grpcPlayer "protobuf-v1/golang/player"
	"soccer-manager/util"
	grpcError "soccer-manager/util/error"
	"soccer-manager/util/id"
	"soccer-manager/util/router"

	"github.com/go-chi/chi"
	grpcCodes "google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (c clientController) GetPlayer(w http.ResponseWriter, r *http.Request) {
	req := new(grpcPlayer.GetRequest)
	req.Id = chi.URLParam(r, ParamPlayerID)

	headerTeamId := router.NewHeader(r.Context()).GetTeamID()

	player, err := c.pc.Get(r.Context(), req)
	if err != nil {
		router.RenderHttpError(w, r, grpcError.NewHttpErrorFromError(r.Method, err))
		return
	}

	//access check
	if player.TeamId != headerTeamId.String() {
		router.RenderHttpError(w, r, grpcError.NewHttpErrorFromError(r.Method, grpcError.NewError(r.Context(), grpcRoot.Error_ERROR_AUTH_ERROR)))
		return
	}

	apiResp := c.getPlayerApiResponse(player)
	resp := router.Response{
		Writer:   w,
		Status:   http.StatusOK,
		GRPCData: apiResp,
	}

	router.RenderJSON(resp)
}

func (c clientController) GetListedPlayers(w http.ResponseWriter, r *http.Request) {
	req := new(grpcPlayer.GetListedRequest)

	players, err := c.pc.GetListed(r.Context(), req)
	if err != nil {
		router.RenderHttpError(w, r, grpcError.NewHttpErrorFromError(r.Method, err))
		return
	}

	apiResp := &grpcPlayerApi.Players{}
	apiResp.Total = players.Total
	for _, player := range players.Players {
		apiResp.Players = append(apiResp.Players, c.getPlayerApiResponse(player))
	}
	resp := router.Response{
		Writer:   w,
		Status:   http.StatusOK,
		GRPCData: apiResp,
	}

	router.RenderJSON(resp)
}

func (c clientController) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	req := new(grpcPlayerApi.UpdateRequest)
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

	headerTeamId := router.NewHeader(r.Context()).GetTeamID()

	playerId, err := id.ParsePlayerID(chi.URLParam(r, ParamPlayerID))
	if err != nil {
		router.RenderHttpError(w, r, grpcError.NewRESTError(r.Method, grpcCodes.InvalidArgument, grpcRoot.Error_ERROR_INVALID_ID, err.Error()))
		return
	}

	player, err := c.pc.Get(r.Context(), &grpcPlayer.GetRequest{Id: playerId.String()})
	if err != nil {
		router.RenderHttpError(w, r, grpcError.NewHttpErrorFromError(r.Method, err))
		return
	}

	//access check
	if player.TeamId != headerTeamId.String() {
		router.RenderHttpError(w, r, grpcError.NewHttpErrorFromError(r.Method, grpcError.NewError(r.Context(), grpcRoot.Error_ERROR_AUTH_ERROR)))
		return
	}

	grpcReq := &grpcPlayer.UpdateRequest{
		Id:        playerId.String(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Country:   req.Country,
		IsListed:  req.IsListed,
	}

	if req.AskValue != nil {
		askValue, err := util.ParseAmountString(req.AskValue.Value)
		if err != nil {
			router.RenderHttpError(w, r, grpcError.NewRESTError(r.Method, grpcCodes.InvalidArgument, grpcRoot.Error_ERROR_INVALID_ARGS, "invalid ask value"))
			return
		}
		if askValue <= 0 {
			router.RenderHttpError(w, r, grpcError.NewRESTError(r.Method, grpcCodes.InvalidArgument, grpcRoot.Error_ERROR_INVALID_ARGS, "ask value should be greater than 0"))
			return
		}
		grpcReq.AskValue = &wrapperspb.Int64Value{Value: askValue}
	}

	player, err = c.pc.Update(r.Context(), grpcReq)
	if err != nil {
		router.RenderHttpError(w, r, grpcError.NewHttpErrorFromError(r.Method, err))
		return
	}

	apiResp := c.getPlayerApiResponse(player)
	resp := router.Response{
		Writer:   w,
		Status:   http.StatusOK,
		GRPCData: apiResp,
	}

	router.RenderJSON(resp)
}

func (c clientController) GetPlayersByTeam(w http.ResponseWriter, r *http.Request) {
	req := new(grpcPlayer.GetByTeamRequest)
	teamId, err := id.ParseTeamID(chi.URLParam(r, ParamTeamID))
	if err != nil {
		router.RenderHttpError(w, r, grpcError.NewRESTError(r.Method, grpcCodes.InvalidArgument, grpcRoot.Error_ERROR_INVALID_ID, err.Error()))
		return
	}

	req.TeamId = teamId.String()
	headerTeamId := router.NewHeader(r.Context()).GetTeamID()

	//access check
	if req.TeamId != headerTeamId.String() {
		router.RenderHttpError(w, r, grpcError.NewHttpErrorFromError(r.Method, grpcError.NewError(r.Context(), grpcRoot.Error_ERROR_AUTH_ERROR)))
		return
	}

	players, err := c.pc.GetByTeam(r.Context(), req)
	if err != nil {
		router.RenderHttpError(w, r, grpcError.NewHttpErrorFromError(r.Method, err))
		return
	}

	apiResp := &grpcPlayerApi.Players{}
	apiResp.Total = players.Total
	for _, player := range players.Players {
		apiResp.Players = append(apiResp.Players, c.getPlayerApiResponse(player))
	}

	resp := router.Response{
		Writer:   w,
		Status:   http.StatusOK,
		GRPCData: apiResp,
	}

	router.RenderJSON(resp)
}

func (c clientController) getPlayerApiResponse(player *grpcPlayer.Player) *grpcPlayerApi.Player {
	playerResp := &grpcPlayerApi.Player{
		Id:        player.Id,
		FirstName: player.FirstName,
		LastName:  player.LastName,
		Age:       player.Age,
		Type:      string(util.PlayerTypeFromProto[player.Type]),
		Country:   player.Country,
		TeamId:    player.TeamId,
		Value:     util.ParseAmountToString(player.Value),
		IsListed:  player.IsListed,
		Currency:  string(util.CurrencyFromProto[player.Currency]),
	}

	if player.AskValue != nil {
		playerResp.AskValue = &wrapperspb.StringValue{Value: util.ParseAmountToString(player.AskValue.Value)}
	}

	return playerResp
}
