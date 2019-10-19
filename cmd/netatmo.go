package main

import (
	"flag"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/ViBiOh/goweb/pkg/netatmo"
	httputils "github.com/ViBiOh/httputils/v2/pkg"
	"github.com/ViBiOh/httputils/v2/pkg/alcotest"
	"github.com/ViBiOh/httputils/v2/pkg/cors"
	"github.com/ViBiOh/httputils/v2/pkg/logger"
	"github.com/ViBiOh/httputils/v2/pkg/owasp"
	"github.com/ViBiOh/httputils/v2/pkg/prometheus"
)

const (
	devicesPath = "/devices"
	docPath     = "doc/"
)

func main() {
	fs := flag.NewFlagSet("netatmo", flag.ExitOnError)

	serverConfig := httputils.Flags(fs, "api")
	alcotestConfig := alcotest.Flags(fs, "")
	prometheusConfig := prometheus.Flags(fs, "prometheus")
	owaspConfig := owasp.Flags(fs, "")
	corsConfig := cors.Flags(fs, "cors")

	netatmoConfig := netatmo.Flags(fs, "")

	logger.Fatal(fs.Parse(os.Args[1:]))

	alcotest.DoAndExit(alcotestConfig)

	prometheusApp := prometheus.New(prometheusConfig)
	owaspApp := owasp.New(owaspConfig)
	corsApp := cors.New(corsConfig)

	netatmoApp := netatmo.New(netatmoConfig)
	netatmoHandler := http.StripPrefix(devicesPath, netatmoApp.Handler())

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, devicesPath) {
			netatmoHandler.ServeHTTP(w, r)
			return
		}

		w.Header().Set("Cache-Control", "no-cache")
		http.ServeFile(w, r, path.Join(docPath, r.URL.Path))
	})

	httpHandler := httputils.ChainMiddlewares(handler, prometheusApp, owaspApp, corsApp)
	go netatmoApp.Start()

	httputils.New(serverConfig).ListenAndServe(httpHandler, httputils.HealthHandler(nil), nil)
}
