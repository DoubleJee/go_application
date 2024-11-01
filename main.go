package main

import (
	"context"
	"go_application/tool"
	"go_application/web"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewStartPrinter(lx fx.Lifecycle, log *zap.Logger) *tool.StartPrinter {
	printer := &tool.StartPrinter{}

	lx.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			printer.SetContent("向你们全世界宣布！新应用启动成功")
			printer.SetLog(log)
			printer.Print()
			return nil
		},
	})

	return printer
}

func main() {

	fx.New(
		// 构造器，提供对象生成逻辑
		fx.Provide(NewStartPrinter, zap.NewExample),
		// 执行器，发起现在要依赖某对象行为
		fx.Invoke(func(*tool.StartPrinter) {}),
		// easy.Model,

		web.Model,
	).Run()
}
