package handler

import (
	"net/http"
	grpcLogin "protobuf-v1/golang/login"
	grpcPlayer "protobuf-v1/golang/player"
	grpcTeam "protobuf-v1/golang/team"
	grpcTxn "protobuf-v1/golang/transaction"
	grpcUser "protobuf-v1/golang/user"
)

const (
	clientApiVersion = "v1"
	ParamUserID      = "userId"
	ParamTeamID      = "teamId"
	ParamTxnID       = "txnId"
	ParamPlayerID    = "playerId"
)

type ClientController interface {

	//metadata
	GetAPIVersion() string
	GetAPIVersionPath(string) string

	//authorization
	Login(http.ResponseWriter, *http.Request)

	//user
	GetUser(http.ResponseWriter, *http.Request)
	UpdateUser(http.ResponseWriter, *http.Request)

	//team
	GetTeam(http.ResponseWriter, *http.Request)
	UpdateTeam(http.ResponseWriter, *http.Request)

	//player
	BuyPlayer(http.ResponseWriter, *http.Request)
	GetPlayer(http.ResponseWriter, *http.Request)
	GetPlayersByTeam(http.ResponseWriter, *http.Request)
	GetListedPlayers(http.ResponseWriter, *http.Request)
	UpdatePlayer(http.ResponseWriter, *http.Request)

	//transaction
	GetTransaction(http.ResponseWriter, *http.Request)
	GetTransactionsByTeam(http.ResponseWriter, *http.Request)
}

type clientController struct {
	lc grpcLogin.LoginServiceClient
	uc grpcUser.UserServiceClient
	tc grpcTeam.TeamServiceClient
	pc grpcPlayer.PlayerServiceClient
	trc grpcTxn.TransactionServiceClient
}

type Clients struct {
	Lc grpcLogin.LoginServiceClient
	Uc grpcUser.UserServiceClient
	Tc grpcTeam.TeamServiceClient
	Pc grpcPlayer.PlayerServiceClient
	Trc grpcTxn.TransactionServiceClient
}

func NewClientController(clients *Clients) ClientController {
	return &clientController{
		lc: clients.Lc,
		uc: clients.Uc,
		tc: clients.Tc,
		pc: clients.Pc,
		trc: clients.Trc,
	}
}

func (cntrl *clientController) GetAPIVersion() string {
	return clientApiVersion
}

func (cntrl *clientController) GetAPIVersionPath(p string) string {
	return "/" + clientApiVersion + p
}
