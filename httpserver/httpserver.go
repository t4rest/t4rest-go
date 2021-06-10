package httpserver

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	"github.com/rs/cors"
)

// HTTPServer .
type (
	HTTPServer struct {
		router *httprouter.Router
		server *http.Server
		routs  []Rout
		cfg    Conf
	}

	Rout struct {
		Method      string
		Path        string
		Handler     httprouter.Handle
		Middlewares []Middleware
	}

	Middleware = func(next httprouter.Handle, name string) httprouter.Handle
)

// New .
func New(cfg Conf, routs []Rout) *HTTPServer {
	return &HTTPServer{
		cfg:   cfg,
		routs: routs,
	}
}

// Start .
func (srv *HTTPServer) Start() error {
	srv.router = httprouter.New()

	for _, r := range srv.routs {
		srv.register(r)
	}

	// todo: move to config
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
	})

	srv.server = &http.Server{
		Addr:         srv.cfg.ListenOnPort,
		Handler:      c.Handler(srv.router),
		ReadTimeout:  srv.cfg.ServerReadTimeoutSec,
		WriteTimeout: srv.cfg.ServerWriteTimeoutSec,
		IdleTimeout:  srv.cfg.ServerIdleTimeoutSec,
	}

	if err := srv.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return errors.New("Error can't launch the server on port: " + srv.cfg.ListenOnPort)
	}

	return nil
}

// Stop .
func (srv *HTTPServer) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), srv.cfg.GracefulShutdownSec)
	defer cancel()

	err := srv.server.Shutdown(ctx)
	if err != nil {
		return errors.Wrap(err, "srv.server.Shutdown")
	}

	return nil
}

// Title .
func (srv *HTTPServer) Title() string {
	return "HTTPServer"
}

func (srv *HTTPServer) register(rout Rout) {
	for _, mw := range rout.Middlewares {
		rout.Handler = mw(rout.Handler, rout.Path)
	}

	srv.router.Handle(rout.Method, rout.Path, rout.Handler)
}
