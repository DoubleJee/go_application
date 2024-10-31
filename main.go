package main

import (
	"bytes"
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
			NewHttpServer,
			NewEchoHandler,
			NewServeMux,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}

func NewHttpServer(lc fx.Lifecycle, mux *http.ServeMux) *http.Server {

	srv := &http.Server{Addr: ":8080", Handler: mux}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			fmt.Println("启动HTTP服务在", srv.Addr)
			srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}

type EchoHandler struct{}

func NewEchoHandler() *EchoHandler {
	return &EchoHandler{}
}

func (*EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	// jsonByte, err := json.Marshal(query)
	// if err != nil {
	// 	fmt.Println("json marshal error")
	// }

	// w.Write(jsonByte)

	queryStr := fmt.Sprintf("%v", query)

	if _, err := io.Copy(w, bytes.NewBufferString(queryStr)); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to handle request:", err)
	}
}

func NewServeMux(echo *EchoHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/echo", echo)
	return mux
}
