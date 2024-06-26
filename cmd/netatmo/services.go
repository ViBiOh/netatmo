package main

import (
	"context"
	"fmt"

	"github.com/ViBiOh/httputils/v4/pkg/e2e"
	"github.com/ViBiOh/netatmo/v2/pkg/netatmo"
)

type services struct {
	netatmo *netatmo.Service
}

func newServices(config configuration, clients clients, adapters adapters) (services, error) {
	e2eService := e2e.New(*config.cipherSecret)

	netatmoApp, err := netatmo.New(config.netatmo, adapters.storage, e2eService, clients.telemetry.MeterProvider())
	if err != nil {
		return services{}, fmt.Errorf("netatmo: %w", err)
	}

	return services{
		netatmo: netatmoApp,
	}, nil
}

func (s services) Start(ctx context.Context) {
	s.netatmo.Start(ctx)
}
