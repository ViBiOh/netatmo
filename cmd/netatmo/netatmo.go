package main

import (
	"context"

	"github.com/ViBiOh/httputils/v4/pkg/logger"
)

func main() {
	config := newConfig()

	ctx := context.Background()

	clients, err := newClients(ctx, config)
	logger.FatalfOnErr(ctx, err, "client")

	go clients.Start()
	defer clients.Close(ctx)

	adapters, err := newAdapters(config, clients)
	logger.FatalfOnErr(ctx, err, "adapters")

	services, err := newServices(config, clients, adapters)
	logger.FatalfOnErr(ctx, err, "services")

	services.Start(clients.health.DoneCtx())
}
