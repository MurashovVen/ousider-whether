package main

import (
	"context"
	"errors"

	"github.com/MurashovVen/outsider-sdk/app"
	"github.com/MurashovVen/outsider-sdk/app/configuration"
	"github.com/MurashovVen/outsider-sdk/app/logger"
	"github.com/MurashovVen/outsider-sdk/app/termination"
	sdkgrpc "github.com/MurashovVen/outsider-sdk/grpc"
	"github.com/MurashovVen/outsider-sdk/tg"
	"go.uber.org/zap"

	"outsider-whether/internal/controller/grpc"
	"outsider-whether/internal/service"
)

func main() {
	cfg := new(config)
	configuration.MustProcessConfig(cfg)

	log := logger.MustCreateLogger(cfg.Env)

	tgClient := tg.MustCreateAndConnect(cfg.TelegramBotToken)

	application := app.New(
		log,
		app.AppendWorks(
			sdkgrpc.NewServer(
				cfg.GRPCServerAddress,
				[]sdkgrpc.ServerRegisterer{
					grpc.New(
						service.New(tgClient),
					),
				},
				sdkgrpc.DefaultServerOptions(log)...,
			),
		),
	)

	if err := application.Run(context.Background()); err != nil && !errors.Is(err, termination.ErrStopped) {
		log.Error("running error: %s", zap.Error(err))
	}
}
