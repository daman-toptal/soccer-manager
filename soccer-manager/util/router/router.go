package router

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"os"
	grpcLogin "protobuf-v1/golang/login"

	"go.elastic.co/apm/module/apmchi"

	"github.com/go-chi/chi"
	chimw "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

const (
	LocalMetadataKey    = "metadata"
	HeaderRequestID     = "sd-request-id"
	HeaderAuthorization = "authorization"
	HeaderUserID        = "sd-user-id"
	HeaderTeamID        = "sd-team-id"

	defaultCertLocation = "./ssl/cert.pem"
	defaultKeyLocation  = "./ssl/key.pem"

	defaultHealthCheckPath = "/healthcheck.html"
)

type Router interface {
	Delete(string, http.HandlerFunc, ...func(http.Handler) http.Handler)
	Get(string, http.HandlerFunc, ...func(http.Handler) http.Handler)
	Head(string, http.HandlerFunc, ...func(http.Handler) http.Handler)
	Patch(string, http.HandlerFunc, ...func(http.Handler) http.Handler)
	Post(string, http.HandlerFunc, ...func(http.Handler) http.Handler)
	Put(string, http.HandlerFunc, ...func(http.Handler) http.Handler)
	Options(string, http.HandlerFunc, ...func(http.Handler) http.Handler)

	Route(string, func(r Router)) Router
	Mount(string, http.Handler)
	Handle(string, http.Handler)
	HandleFunc(string, http.HandlerFunc)
	With(middlewares ...func(http.Handler) http.Handler) Router

	HandleHealthCheck()
	ListenAndServeTLS(string, *tls.Config) error
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type router struct {
	chi *chi.Mux
}

func (r router) With(middlewares ...func(http.Handler) http.Handler) Router {
	r.chi = r.chi.With(middlewares...).(*chi.Mux)
	return r
}

const gzipPerf = 5

var gzipMime = []string{"application/json", "text/html", "text/css", "text/plain", "application/pdf", "application/csv"}

func NewApiRouter(lc grpcLogin.LoginServiceClient) Router {
	rchi := chi.NewRouter()

	compressor := chimw.NewCompressor(gzipPerf, gzipMime...)
	rchi.Use(compressor.Handler)

	// CORS
	rchi.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "PUT"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300, 
	}))

	rchi.Use(
		apmchi.Middleware(),
		accessTokenAuthMiddleware(lc),
	)

	return router{
		chi: rchi,
	}
}

func (r router) Delete(p string, h http.HandlerFunc, middlewares ...func(http.Handler) http.Handler) {
	r.chi.With(middlewares...).Delete(p, h)
}

func (r router) Get(p string, h http.HandlerFunc, middlewares ...func(http.Handler) http.Handler) {
	r.chi.With(middlewares...).Get(p, h)
}

func (r router) Head(p string, h http.HandlerFunc, middlewares ...func(http.Handler) http.Handler) {
	r.chi.With(middlewares...).Head(p, h)
}

func (r router) Patch(p string, h http.HandlerFunc, middlewares ...func(http.Handler) http.Handler) {
	r.chi.With(middlewares...).Patch(p, h)
}

func (r router) Post(p string, h http.HandlerFunc, middlewares ...func(http.Handler) http.Handler) {
	r.chi.With(middlewares...).Post(p, h)
}

func (r router) Put(p string, h http.HandlerFunc, middlewares ...func(http.Handler) http.Handler) {
	r.chi.With(middlewares...).Put(p, h)
}

func (r router) Options(p string, h http.HandlerFunc, middlewares ...func(http.Handler) http.Handler) {
	r.chi.With(middlewares...).Options(p, h)
}

func (r router) Route(p string, fn func(r Router)) Router {
	nr := router{chi.NewRouter()}

	if fn != nil {
		fn(nr)
	}

	r.Mount(p, nr)
	return nr
}

func (r router) Mount(p string, h http.Handler) {
	r.chi.Mount(p, h)
}

func (r router) Handle(p string, h http.Handler) {
	r.chi.Handle(p, h)
}

func (r router) HandleFunc(p string, h http.HandlerFunc) {
	r.chi.HandleFunc(p, h)
}

func (r router) HandleHealthCheck() {
	r.chi.Get(defaultHealthCheckPath, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})
}

func (r router) ListenAndServeTLS(listenPort string, config *tls.Config) error {
	if listenPort == "" {
		return errors.New("invalid or missing listen port")
	}

	server := &http.Server{
		Addr:      fmt.Sprintf(":%s", listenPort),
		Handler:   r.chi,
		TLSConfig: config,
	}

	if config != nil {
		server.TLSConfig.BuildNameToCertificate()
	}

	if _, err := os.Stat(defaultCertLocation); os.IsNotExist(err) {
		return server.ListenAndServe()
	}

	return server.ListenAndServe()
}

func (r router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.chi.ServeHTTP(w, req)
}
