package handler

import (
	"io/ioutil"
	"net/http"
	grpcRoot "protobuf-v1/golang"
	grpcTxnApi "protobuf-v1/golang/external/transaction"
	grpcPlayer "protobuf-v1/golang/player"
	grpcTxn "protobuf-v1/golang/transaction"
	"soccer-manager/util"
	grpcError "soccer-manager/util/error"
	"soccer-manager/util/id"
	"soccer-manager/util/router"

	"github.com/go-chi/chi"
	grpcCodes "google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
)

func (c clientController) GetTransaction(w http.ResponseWriter, r *http.Request) {
	req := new(grpcTxn.GetRequest)
	req.Id = chi.URLParam(r, ParamTxnID)

	txn, err := c.trc.Get(r.Context(), req)
	if err != nil {
		router.RenderHttpError(w, r, grpcError.NewHttpErrorFromError(r.Method, err))
		return
	}

	headerTeamId := router.NewHeader(r.Context()).GetTeamID()

	//access check
	if headerTeamId.String() != txn.TeamId {
		router.RenderHttpError(w, r, grpcError.NewHttpErrorFromError(r.Method, grpcError.NewError(r.Context(), grpcRoot.Error_ERROR_AUTH_ERROR)))
		return
	}

	apiResp := c.getTxnApiResponse(txn)
	resp := router.Response{
		Writer:   w,
		Status:   http.StatusOK,
		GRPCData: apiResp,
	}

	router.RenderJSON(resp)
}

func (c clientController) GetTransactionsByTeam(w http.ResponseWriter, r *http.Request) {
	req := new(grpcTxn.GetByTeamRequest)
	teamId, err := id.ParseTeamID(chi.URLParam(r, ParamTeamID))
	if err != nil {
		router.RenderHttpError(w, r, grpcError.NewRESTError(r.Method, grpcCodes.InvalidArgument, grpcRoot.Error_ERROR_INVALID_ID, err.Error()))
		return
	}
	req.TeamId = teamId.String()
	headerTeamId := router.NewHeader(r.Context()).GetTeamID()

	//access check
	if headerTeamId != teamId {
		router.RenderHttpError(w, r, grpcError.NewHttpErrorFromError(r.Method, grpcError.NewError(r.Context(), grpcRoot.Error_ERROR_AUTH_ERROR)))
		return
	}

	txns, err := c.trc.GetByTeam(r.Context(), req)
	if err != nil {
		router.RenderHttpError(w, r, grpcError.NewHttpErrorFromError(r.Method, err))
		return
	}

	apiResp := &grpcTxnApi.Transactions{}
	apiResp.Total = txns.Total
	for _, txn := range txns.Transactions {
		apiResp.Transactions = append(apiResp.Transactions, c.getTxnApiResponse(txn))
	}

	resp := router.Response{
		Writer:   w,
		Status:   http.StatusOK,
		GRPCData: apiResp,
	}

	router.RenderJSON(resp)
}

func (c clientController) BuyPlayer(w http.ResponseWriter, r *http.Request) {
	req := new(grpcTxnApi.BuyRequest)
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

	playerId, err := id.ParsePlayerID(req.PlayerId)
	if err != nil {
		router.RenderHttpError(w, r, grpcError.NewRESTError(r.Method, grpcCodes.InvalidArgument, grpcRoot.Error_ERROR_INVALID_ID, err.Error()))
		return
	}

	player, err := c.pc.Get(r.Context(), &grpcPlayer.GetRequest{Id: playerId.String()})
	if err != nil {
		router.RenderHttpError(w, r, grpcError.NewHttpErrorFromError(r.Method, err))
		return
	}

	headerTeamId := router.NewHeader(r.Context()).GetTeamID()

	//access check
	if player.IsListed != true {
		router.RenderHttpError(w, r, grpcError.NewHttpErrorFromError(r.Method, grpcError.NewError(r.Context(), grpcRoot.Error_ERROR_AUTH_ERROR, "player not listed")))
		return
	}

	if player.TeamId == headerTeamId.String() {
		router.RenderHttpError(w, r, grpcError.NewHttpErrorFromError(r.Method, grpcError.NewError(r.Context(), grpcRoot.Error_ERROR_AUTH_ERROR, "cannot buy own player")))
		return
	}

	txn, err := c.trc.Buy(r.Context(), &grpcTxn.BuyRequest{PlayerId: playerId.String(), TeamId: headerTeamId.String(), Description: req.Description})
	if err != nil {
		router.RenderHttpError(w, r, grpcError.NewHttpErrorFromError(r.Method, err))
		return
	}

	apiResp := c.getTxnApiResponse(txn)
	resp := router.Response{
		Writer:   w,
		Status:   http.StatusCreated,
		GRPCData: apiResp,
	}

	router.RenderJSON(resp)
}

func (c clientController) getTxnApiResponse(txn *grpcTxn.Transaction) *grpcTxnApi.Transaction {
	return &grpcTxnApi.Transaction{
		Id:          txn.Id,
		Title:       txn.Title,
		Description: txn.Description,
		TeamId:      txn.TeamId,
		Amount:      util.ParseAmountToString(txn.Amount),
		Budget:      util.ParseAmountToString(txn.Budget),
		PlayerId:    txn.PlayerId,
		CreatedAt:   txn.CreatedAt,
		Type:        string(util.TransactionTypeFromProto[txn.Type]),
		Currency:    string(util.CurrencyFromProto[txn.Currency]),
	}
}
