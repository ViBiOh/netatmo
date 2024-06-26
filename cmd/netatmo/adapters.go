package main

import (
	"fmt"

	"github.com/ViBiOh/absto/pkg/absto"
	abstoModel "github.com/ViBiOh/absto/pkg/model"
)

type adapters struct {
	storage abstoModel.Storage
}

func newAdapters(config configuration, clients clients) (adapters, error) {
	storageService, err := absto.New(config.absto, clients.telemetry.TracerProvider())
	if err != nil {
		return adapters{}, fmt.Errorf("absto: %w", err)
	}

	return adapters{
		storage: storageService,
	}, nil
}
