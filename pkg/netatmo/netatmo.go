package netatmo

import (
	"context"
	"flag"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/ViBiOh/httputils/v3/pkg/flags"
	"github.com/ViBiOh/httputils/v3/pkg/httperror"
	"github.com/ViBiOh/httputils/v3/pkg/httpjson"
	"github.com/ViBiOh/httputils/v3/pkg/logger"
	"github.com/prometheus/client_golang/prometheus"
)

// App of package
type App interface {
	Handler() http.Handler
	Start()
	Enabled() bool
}

// Config of package
type Config struct {
	accessToken  *string
	refreshToken *string
	clientID     *string
	clientSecret *string
	scopes       *string
}

type app struct {
	clientID     string
	clientSecret string
	accessToken  string
	refreshToken string
	scopes       string

	mutex                sync.RWMutex
	devices              []Device
	prometheusCollectors map[string]prometheus.Gauge
}

// Flags adds flags for configuring package
func Flags(fs *flag.FlagSet, prefix string) Config {
	return Config{
		accessToken:  flags.New(prefix, "netatmo").Name("AccessToken").Default("").Label("Access Token").ToString(fs),
		refreshToken: flags.New(prefix, "netatmo").Name("RefreshToken").Default("").Label("Refresh Token").ToString(fs),
		clientID:     flags.New(prefix, "netatmo").Name("ClientID").Default("").Label("Client ID").ToString(fs),
		clientSecret: flags.New(prefix, "netatmo").Name("ClientSecret").Default("").Label("Client Secret").ToString(fs),
		scopes:       flags.New(prefix, "netatmo").Name("Scopes").Default("").Label("Scopes, comma separated").ToString(fs),
	}
}

// New creates new App from Config
func New(config Config) App {
	return &app{
		clientID:             strings.TrimSpace(*config.clientID),
		clientSecret:         strings.TrimSpace(*config.clientSecret),
		accessToken:          strings.TrimSpace(*config.accessToken),
		refreshToken:         strings.TrimSpace(*config.refreshToken),
		scopes:               strings.TrimSpace(*config.scopes),
		prometheusCollectors: make(map[string]prometheus.Gauge),
	}
}

// Handler for request. Should be use with net/http
func (a *app) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}

		a.mutex.RLock()
		defer a.mutex.RUnlock()

		if err := httpjson.ResponseArrayJSON(w, http.StatusOK, a.devices, httpjson.IsPretty(r)); err != nil {
			httperror.InternalServerError(w, err)
		}
	})
}

// Start periodic fetch of data from netatmo API
func (a *app) Start() {
	if !a.Enabled() {
		logger.Warn("app is disabled")
		return
	}

	timer := time.NewTimer(time.Minute)
	for {
		<-timer.C
		devices, err := a.GetDevices(context.Background())
		if err != nil {
			logger.Error("unable to fetch devices: %+v", err)
		} else {
			a.mutex.Lock()
			a.devices = devices
			a.mutex.Unlock()

			a.updatePrometheus()
		}

		timer.Reset(time.Minute)
	}
}

// Enabled check if app is enabled
func (a *app) Enabled() bool {
	return a.accessToken != ""
}

// HasScope check if given scope is configured
func (a *app) HasScope(scope string) bool {
	return strings.Contains(a.scopes, scope)
}
