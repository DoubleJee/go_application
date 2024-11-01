package main

import (
	"go_application/easy"
	"net/http"

	"go.uber.org/fx"
)

func main() {

	fx.New(
		fx.Provide(
			easy.NewHttpServer,
			easy.NewEchoHandler,
			easy.NewServeMux,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
