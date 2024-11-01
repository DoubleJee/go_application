package easy

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"go.uber.org/fx"
)

func NewHttpServer(lc fx.Lifecycle, mux *http.ServeMux) *http.Server {

	srv := &http.Server{Addr: ":8080", Handler: mux}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)

			if err != nil {
				return err
			}

			fmt.Println("Starting Http Server at", srv.Addr)
			go srv.Serve(ln)
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

	fmt.Fprintf(w, "查询参数: \n%s", r.URL.Query())

	// if _, err := io.Copy(w, r.Body); err != nil {
	// 	fmt.Fprintln(os.Stderr, "Failed to handle request:", err)
	// }
}

func NewServeMux(echo *EchoHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/echo", echo)
	return mux
}
