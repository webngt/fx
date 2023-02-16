package main

import (
	"net/http"

	ports "github.com/webngt/fx/internal/http"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Provide(
			ports.NewHTTPServer,
			ports.NewServeMux,
			ports.NewHandler,
			zap.NewProduction,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
