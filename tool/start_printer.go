package tool

import (
	"go.uber.org/zap"
)

type StartPrinter struct {
	content string
	log     *zap.Logger
}

func (p *StartPrinter) SetContent(content string) {
	p.content = content
}
func (p *StartPrinter) SetLog(log *zap.Logger) {
	p.log = log
}

func (p *StartPrinter) Print() {
	banner := []string{
		" ______  ______  ______ ",
		"/  ____||___  / |___  / ",
		"| |  __    / /     / /  ",
		"| | |_ |  / /     / /   ",
		"| |__| | / /__   / /__  ",
		"\\______//_____| /_____| ",
	}

	for _, line := range banner {
		p.log.Info(line)
	}

	p.log.Info("启动完毕！=> :", zap.String("content", p.content))
}
