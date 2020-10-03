package main

import (
	"flag"
	"net/http"
	"os"
	"strings"

	"github.com/ViBiOh/goweb/pkg/netatmo"
	"github.com/ViBiOh/httputils/v3/pkg/alcotest"
	"github.com/ViBiOh/httputils/v3/pkg/cors"
	"github.com/ViBiOh/httputils/v3/pkg/httputils"
	"github.com/ViBiOh/httputils/v3/pkg/logger"
	"github.com/ViBiOh/httputils/v3/pkg/owasp"
	"github.com/ViBiOh/httputils/v3/pkg/prometheus"
)

const (
	devicesPath = "/devices"
)

func main() {
	fs := flag.NewFlagSet("netatmo", flag.ExitOnError)

	serverConfig := httputils.Flags(fs, "api")
	alcotestConfig := alcotest.Flags(fs, "")
	loggerConfig := logger.Flags(fs, "logger")
	prometheusConfig := prometheus.Flags(fs, "prometheus")
	owaspConfig := owasp.Flags(fs, "")
	corsConfig := cors.Flags(fs, "cors")

	netatmoConfig := netatmo.Flags(fs, "")

	logger.Fatal(fs.Parse(os.Args[1:]))

	alcotest.DoAndExit(alcotestConfig)
	logger.Global(logger.New(loggerConfig))
	defer logger.Close()

	server := httputils.New(serverConfig)
	prometheusApp := prometheus.New(prometheusConfig)
	netatmoApp := netatmo.New(netatmoConfig, prometheusApp)

	netatmoHandler := http.StripPrefix(devicesPath, netatmoApp.Handler())

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, devicesPath) {
			netatmoHandler.ServeHTTP(w, r)
			return
		}

		w.WriteHeader(http.StatusNotFound)
	})

	go netatmoApp.Start()

	server.Middleware(prometheusApp.Middleware)
	server.Middleware(owasp.New(owaspConfig).Middleware)
	server.Middleware(cors.New(corsConfig).Middleware)
	server.ListenServeWait(handler)
}
