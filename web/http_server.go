package web

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewHttpServer(lc fx.Lifecycle) *gin.Engine {

	r := gin.Default()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {

			r.GET("/ping", func(ctx *gin.Context) {
				ctx.JSON(200, gin.H{
					"id": "123",
				})
			})

			r.GET("/hello", func(ctx *gin.Context) {
				ctx.Data(200, "text/html; charset=utf-8", []byte("<h1>Hello</h1>"))
			})

			go func() {
				r.Run()
			}()

			return nil
		},
	})

	return r
}
