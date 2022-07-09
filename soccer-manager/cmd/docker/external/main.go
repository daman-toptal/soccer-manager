package main

import (
	"fmt"
	"log"
	"net/http"
	grpcLogin "protobuf-v1/golang/login"
	grpcPlayer "protobuf-v1/golang/player"
	grpcTeam "protobuf-v1/golang/team"
	grpcTxn "protobuf-v1/golang/transaction"
	grpcUser "protobuf-v1/golang/user"
	"soccer-manager/internal/handler"
	"soccer-manager/util/config"
	"soccer-manager/util/id"
	"soccer-manager/util/logging"
	"soccer-manager/util/router"

	ggrpc "google.golang.org/grpc"

	"go.elastic.co/apm/module/apmgrpc"
)

func main() {
	config.SetupConfig()
	port := config.GetString("grpc.port")
	cnString := fmt.Sprintf("%s:%s", config.GetString("grpc.name"), port)

	logging.SetupLogging(config.GetString("log.level"))

	serviceConn, err := ggrpc.Dial(cnString, ggrpc.WithInsecure(), ggrpc.WithUnaryInterceptor(apmgrpc.NewUnaryClientInterceptor()))
	if err != nil {
		log.Fatalf("Error initializing grpc service client, err=%s", err.Error())
	}

	defer serviceConn.Close()

	clients := &handler.Clients{
		Lc:  grpcLogin.NewLoginServiceClient(serviceConn),
		Uc:  grpcUser.NewUserServiceClient(serviceConn),
		Tc:  grpcTeam.NewTeamServiceClient(serviceConn),
		Pc:  grpcPlayer.NewPlayerServiceClient(serviceConn),
		Trc: grpcTxn.NewTransactionServiceClient(serviceConn),
	}

	r := registerRoutes(handler.NewClientController(clients), clients.Lc)

	err = r.ListenAndServeTLS(config.GetString("server.httpPort"), nil)
	if err != nil {
		log.Fatalf("Something went wrong with the server")
	}
}

func registerRoutes(clientCntrl handler.ClientController, lc grpcLogin.LoginServiceClient) router.Router {
	r := router.NewApiRouter(lc)

	r.Route("/health", func(r router.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			a := map[string]string{"status": "OK"}
			router.RenderJSON(
				router.Response{
					Status: 200,
					Data:   a,
					Writer: w,
				})
		})
	})

	r.Route(clientCntrl.GetAPIVersionPath("/login"), func(r router.Router) {
		r.Post("/", clientCntrl.Login)

	})

	r.Route(clientCntrl.GetAPIVersionPath("/user"), func(r router.Router) {
		r.Route(fmt.Sprintf("/{userId:%s}", id.IDPrefixUser.REMatch()), func(r router.Router) {
			r.Get("/", clientCntrl.GetUser)
			r.Patch("/", clientCntrl.UpdateUser)
		})

	})

	r.Route(clientCntrl.GetAPIVersionPath("/team"), func(r router.Router) {
		r.Route(fmt.Sprintf("/{teamId:%s}", id.IDPrefixTeam.REMatch()), func(r router.Router) {
			r.Get("/", clientCntrl.GetTeam)
			r.Patch("/", clientCntrl.UpdateTeam)
			r.Get("/players", clientCntrl.GetPlayersByTeam)
			r.Get("/transactions", clientCntrl.GetTransactionsByTeam)
		})

	})

	r.Route(clientCntrl.GetAPIVersionPath("/player"), func(r router.Router) {
		r.Get("/listed", clientCntrl.GetListedPlayers)
		r.Post("/buy", clientCntrl.BuyPlayer)

		r.Route(fmt.Sprintf("/{playerId:%s}", id.IDPrefixPlayer.REMatch()), func(r router.Router) {
			r.Get("/", clientCntrl.GetPlayer)
			r.Patch("/", clientCntrl.UpdatePlayer)
		})

	})

	r.Route(clientCntrl.GetAPIVersionPath("/transaction"), func(r router.Router) {
		r.Route(fmt.Sprintf("/{txnId:%s}", id.IDPrefixTransaction.REMatch()), func(r router.Router) {
			r.Get("/", clientCntrl.GetTransaction)
		})

	})
	return r
}
