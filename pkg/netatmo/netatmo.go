package netatmo

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/ViBiOh/httputils/v4/pkg/cron"
	"github.com/ViBiOh/httputils/v4/pkg/flags"
	"github.com/ViBiOh/httputils/v4/pkg/httpjson"
	"github.com/ViBiOh/httputils/v4/pkg/logger"
	"github.com/prometheus/client_golang/prometheus"
)

// App of package
type App struct {
	prometheusCollectors map[string]prometheus.Gauge
	registerer           prometheus.Registerer

	clientID     string
	clientSecret string
	accessToken  string
	refreshToken string
	scopes       string

	devices []Device

	mutex sync.RWMutex
}

// Config of package
type Config struct {
	accessToken  *string
	refreshToken *string
	clientID     *string
	clientSecret *string
	scopes       *string
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
func New(config Config, registerer prometheus.Registerer) *App {
	return &App{
		clientID:             strings.TrimSpace(*config.clientID),
		clientSecret:         strings.TrimSpace(*config.clientSecret),
		accessToken:          strings.TrimSpace(*config.accessToken),
		refreshToken:         strings.TrimSpace(*config.refreshToken),
		scopes:               strings.TrimSpace(*config.scopes),
		prometheusCollectors: make(map[string]prometheus.Gauge),
		registerer:           registerer,
	}
}

// Handler for request. Should be use with net/http
func (a *App) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}

		a.mutex.RLock()
		defer a.mutex.RUnlock()

		httpjson.WriteArray(w, http.StatusOK, a.devices, httpjson.IsPretty(r))
	})
}

// Start periodic fetch of data from netatmo API
func (a *App) Start(done <-chan struct{}) {
	if !a.Enabled() {
		logger.Warn("app is disabled")
		return
	}

	cron.New().Each(time.Minute*5).Now().OnError(func(err error) {
		logger.Error("%s", err)
	}).Start(func(ctx context.Context) error {
		devices, err := a.getDevices(ctx)
		if err != nil {
			return fmt.Errorf("unable to fetch devices: %s", err)
		}

		a.mutex.Lock()
		defer a.mutex.Unlock()

		a.devices = devices
		a.updatePrometheus()

		return nil
	}, done)
}

// Enabled check if app is enabled
func (a *App) Enabled() bool {
	return a.accessToken != ""
}

// HasScope check if given scope is configured
func (a *App) HasScope(scope string) bool {
	return strings.Contains(a.scopes, scope)
}
