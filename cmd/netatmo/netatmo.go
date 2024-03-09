package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "net/http/pprof"

	"github.com/ViBiOh/absto/pkg/absto"
	"github.com/ViBiOh/flags"
	"github.com/ViBiOh/httputils/v4/pkg/e2e"
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
	abstoConfig := absto.Flags(fs, "storage")

	cipherSecret := flags.New("CipherSecret", "Secret for ciphering token, 32 characters").DocPrefix("secret").String(fs, "", nil)

	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	logger.Init(loggerConfig)

	ctx := context.Background()

	go func() {
		fmt.Println(http.ListenAndServe("localhost:9999", http.DefaultServeMux))
	}()

	telemetryApp, err := telemetry.New(ctx, telemetryConfig)
	logger.FatalfOnErr(ctx, err, "create telemetry")

	defer telemetryApp.Close(ctx)

	logger.AddOpenTelemetryToDefaultLogger(telemetryApp)
	request.AddOpenTelemetryToDefaultClient(telemetryApp.MeterProvider(), telemetryApp.TracerProvider())

	healthApp := health.New(ctx, healthConfig)

	e2eService := e2e.New(*cipherSecret)

	storageService, err := absto.New(abstoConfig, telemetryApp.TracerProvider())
	logger.FatalfOnErr(ctx, err, "create storage")

	netatmoApp, err := netatmo.New(netatmoConfig, storageService, e2eService, telemetryApp.MeterProvider())
	logger.FatalfOnErr(ctx, err, "create netatmo")

	netatmoApp.Start(healthApp.DoneCtx())
}
