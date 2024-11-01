package easy

import (
	"net/http"

	"go.uber.org/fx"
)

var Model = fx.Options(
	fx.Provide(
		NewHttpServer,
		// NewEchoHandler作为Route类型依赖
		// 解耦注册，你依赖一个Route，那我NewEchoHandler出来的对象，就注入给你
		fx.Annotate(
			NewEchoHandler,
			fx.As(new(Route)),
			// 我的创建出来的依赖对象标签是name:echo，解决依赖Route，有多个Route对象，该如何选择的事情
			// fx.ResultTags(`name:"echo"`),

			// 使用组标签
			fx.ResultTags(`group:"routes"`),
		),
		fx.Annotate(
			NewHelloHandler,
			fx.As(new(Route)),
			// fx.ResultTags(`name:"hello"`),
			fx.ResultTags(`group:"routes"`),
		),

		// 我创建的时候依赖的对象标签是name:echo、name:hello
		fx.Annotate(
			NewServeMux,
			// 依赖组标签
			fx.ParamTags(`group:"routes"`),
		),
	),
	// 执行器，发起现在要依赖某对象行为
	fx.Invoke(func(*http.Server) {}),
)
