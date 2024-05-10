package main

import (
	"context"
	"flag"
	"os"

	"github.com/ViBiOh/absto/pkg/absto"
	"github.com/ViBiOh/flags"
	"github.com/ViBiOh/httputils/v4/pkg/e2e"
	"github.com/ViBiOh/httputils/v4/pkg/health"
	"github.com/ViBiOh/httputils/v4/pkg/logger"
	"github.com/ViBiOh/httputils/v4/pkg/pprof"
	"github.com/ViBiOh/httputils/v4/pkg/request"
	"github.com/ViBiOh/httputils/v4/pkg/telemetry"
	"github.com/ViBiOh/netatmo/v2/pkg/netatmo"
)

func main() {
	fs := flag.NewFlagSet("netatmo", flag.ExitOnError)
	fs.Usage = flags.Usage(fs)

	healthConfig := health.Flags(fs, "")

	loggerConfig := logger.Flags(fs, "logger")
	telemetryConfig := telemetry.Flags(fs, "telemetry")
	pprofConfig := pprof.Flags(fs, "pprof")

	netatmoConfig := netatmo.Flags(fs, "")
	abstoConfig := absto.Flags(fs, "storage")

	cipherSecret := flags.New("CipherSecret", "Secret for ciphering token, 32 characters").DocPrefix("secret").String(fs, "", nil)

	_ = fs.Parse(os.Args[1:])

	ctx := context.Background()

	logger.Init(ctx, loggerConfig)

	healthApp := health.New(ctx, healthConfig)

	telemetryApp, err := telemetry.New(ctx, telemetryConfig)
	logger.FatalfOnErr(ctx, err, "create telemetry")

	defer telemetryApp.Close(ctx)

	logger.AddOpenTelemetryToDefaultLogger(telemetryApp)
	request.AddOpenTelemetryToDefaultClient(telemetryApp.MeterProvider(), telemetryApp.TracerProvider())

	service, version, env := telemetryApp.GetServiceVersionAndEnv()
	pprofService := pprof.New(pprofConfig, service, version, env)

	go pprofService.Start(healthApp.DoneCtx())

	e2eService := e2e.New(*cipherSecret)

	storageService, err := absto.New(abstoConfig, telemetryApp.TracerProvider())
	logger.FatalfOnErr(ctx, err, "create storage")

	netatmoApp, err := netatmo.New(netatmoConfig, storageService, e2eService, telemetryApp.MeterProvider())
	logger.FatalfOnErr(ctx, err, "create netatmo")

	netatmoApp.Start(healthApp.DoneCtx())
}
