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
	var output services
	var err error

	e2eService := e2e.New(*config.cipherSecret)

	output.netatmo, err = netatmo.New(config.netatmo, adapters.storage, e2eService, clients.telemetry.TracerProvider(), clients.telemetry.MeterProvider())
	if err != nil {
		return output, fmt.Errorf("netatmo: %w", err)
	}

	return output, nil
}

func (s services) Start(ctx context.Context) {
	s.netatmo.Start(ctx)
}
