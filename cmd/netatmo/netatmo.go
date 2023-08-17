package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"

	_ "net/http/pprof"

	"github.com/ViBiOh/flags"
	"github.com/ViBiOh/httputils/v4/pkg/alcotest"
	"github.com/ViBiOh/httputils/v4/pkg/cors"
	"github.com/ViBiOh/httputils/v4/pkg/health"
	"github.com/ViBiOh/httputils/v4/pkg/httputils"
	"github.com/ViBiOh/httputils/v4/pkg/logger"
	"github.com/ViBiOh/httputils/v4/pkg/owasp"
	"github.com/ViBiOh/httputils/v4/pkg/recoverer"
	"github.com/ViBiOh/httputils/v4/pkg/server"
	"github.com/ViBiOh/httputils/v4/pkg/telemetry"
	"github.com/ViBiOh/netatmo/pkg/netatmo"
)

const (
	devicesPath = "/devices"
)

func main() {
	fs := flag.NewFlagSet("netatmo", flag.ExitOnError)
	fs.Usage = flags.Usage(fs)

	appServerConfig := server.Flags(fs, "")
	healthConfig := health.Flags(fs, "")

	alcotestConfig := alcotest.Flags(fs, "")
	loggerConfig := logger.Flags(fs, "logger")
	telemetryConfig := telemetry.Flags(fs, "telemetry")
	owaspConfig := owasp.Flags(fs, "")
	corsConfig := cors.Flags(fs, "cors")

	netatmoConfig := netatmo.Flags(fs, "")

	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	alcotest.DoAndExit(alcotestConfig)

	logger.Init(loggerConfig)

	ctx := context.Background()

	go func() {
		fmt.Println(http.ListenAndServe("localhost:9999", http.DefaultServeMux))
	}()

	telemetryApp, err := telemetry.New(ctx, telemetryConfig)
	if err != nil {
		slog.Error("create telemetry", "err", err)
		os.Exit(1)
	}

	appServer := server.New(appServerConfig)
	healthApp := health.New(healthConfig)

	netatmoApp, err := netatmo.New(netatmoConfig, telemetryApp.GetMeterProvider())
	if err != nil {
		slog.Error("create netatmo", "err", err)
		os.Exit(1)
	}

	netatmoHandler := http.StripPrefix(devicesPath, netatmoApp.Handler())

	appHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, devicesPath) {
			netatmoHandler.ServeHTTP(w, r)
			return
		}

		w.WriteHeader(http.StatusNotFound)
	})

	go netatmoApp.Start(healthApp.Done(ctx))

	endCtx := healthApp.End(ctx)

	go appServer.Start(endCtx, "http", httputils.Handler(appHandler, healthApp, recoverer.Middleware, telemetryApp.Middleware("http"), owasp.New(owaspConfig).Middleware, cors.New(corsConfig).Middleware))

	healthApp.WaitForTermination(appServer.Done())
	server.GracefulWait(appServer.Done())
}
