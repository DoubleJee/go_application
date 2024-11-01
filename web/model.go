package web

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Model = fx.Options(

	fx.Provide(NewHttpServer),
	fx.Invoke(func(*gin.Engine) {}),
)
