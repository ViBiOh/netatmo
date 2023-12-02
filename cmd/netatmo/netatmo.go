package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	_ "net/http/pprof"

	"github.com/ViBiOh/flags"
	"github.com/ViBiOh/httputils/v4/pkg/health"
	"github.com/ViBiOh/httputils/v4/pkg/logger"
	"github.com/ViBiOh/httputils/v4/pkg/request"
	"github.com/ViBiOh/httputils/v4/pkg/telemetry"
	"github.com/ViBiOh/netatmo/pkg/netatmo"
)

func main() {
	fs := flag.NewFlagSet("netatmo", flag.ExitOnError)
	fs.Usage = flags.Usage(fs)

	healthConfig := health.Flags(fs, "")

	loggerConfig := logger.Flags(fs, "logger")
	telemetryConfig := telemetry.Flags(fs, "telemetry")

	netatmoConfig := netatmo.Flags(fs, "")

	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	logger.Init(loggerConfig)

	ctx := context.Background()

	go func() {
		fmt.Println(http.ListenAndServe("localhost:9999", http.DefaultServeMux))
	}()

	telemetryApp, err := telemetry.New(ctx, telemetryConfig)
	if err != nil {
		slog.ErrorContext(ctx, "create telemetry", "err", err)
		os.Exit(1)
	}

	defer telemetryApp.Close(ctx)

	logger.AddOpenTelemetryToDefaultLogger(telemetryApp)
	request.AddOpenTelemetryToDefaultClient(telemetryApp.MeterProvider(), telemetryApp.TracerProvider())

	healthApp := health.New(ctx, healthConfig)

	netatmoApp, err := netatmo.New(netatmoConfig, telemetryApp.MeterProvider())
	if err != nil {
		slog.ErrorContext(ctx, "create netatmo", "err", err)
		os.Exit(1)
	}

	netatmoApp.Start(healthApp.DoneCtx())
}
