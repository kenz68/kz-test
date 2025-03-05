package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			NewHTTPServer,
			NewEchoHandler,
			NewServeMux,
		),
		fx.Invoke(func(server *http.Server) {}),
	).Run()
}

func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux) *http.Server {
	srv := &http.Server{Addr: ":8080", Handler: mux}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			fmt.Printf("Starting HTTP server at %s", srv.Addr)
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

// EchoHandler is an http.Handler that copies its request body
// back to the response.
type EchoHandler struct{}

// NewEchoHandler builds a new EchoHandler.
func NewEchoHandler() *EchoHandler {
	return &EchoHandler{}
}

// ServeHTTP handles an HTTP request to the /echo endpoint.
func (*EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to handle request:", err)
	}
}

// NewServeMux builds a ServerMux that will route requests
// to the given EchoHandler.
func NewServeMux(echo *EchoHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/echo", echo)
	return mux
}
