package main

import (
	"github.com/MurashovVen/outsider-sdk/app/configuration"
)

type config struct {
	configuration.Default
	configuration.GRPCServer
	configuration.TelegramClient
	configuration.Mongo
}
