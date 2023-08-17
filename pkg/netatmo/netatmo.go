package netatmo

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/ViBiOh/flags"
	"github.com/ViBiOh/httputils/v4/pkg/cron"
	"github.com/ViBiOh/httputils/v4/pkg/httpjson"
	"go.opentelemetry.io/otel/metric"
)

// App of package
type App struct {
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
		accessToken:  flags.New("AccessToken", "Access Token").Prefix(prefix).DocPrefix("netatmo").String(fs, "", nil),
		refreshToken: flags.New("RefreshToken", "Refresh Token").Prefix(prefix).DocPrefix("netatmo").String(fs, "", nil),
		clientID:     flags.New("ClientID", "Client ID").Prefix(prefix).DocPrefix("netatmo").String(fs, "", nil),
		clientSecret: flags.New("ClientSecret", "Client Secret").Prefix(prefix).DocPrefix("netatmo").String(fs, "", nil),
		scopes:       flags.New("Scopes", "Scopes, comma separated").Prefix(prefix).DocPrefix("netatmo").String(fs, "", nil),
	}
}

// New creates new App from Config
func New(config Config, meterProvider metric.MeterProvider) (*App, error) {
	app := &App{
		clientID:     strings.TrimSpace(*config.clientID),
		clientSecret: strings.TrimSpace(*config.clientSecret),
		accessToken:  strings.TrimSpace(*config.accessToken),
		refreshToken: strings.TrimSpace(*config.refreshToken),
		scopes:       strings.TrimSpace(*config.scopes),
	}

	if err := app.createMetrics(meterProvider, "temperature", "humidity", "noise", "co2", "pressure"); err != nil {
		return nil, fmt.Errorf("create metrics: %w", err)
	}

	return app, nil
}

// Handler for request. Should be use with net/http
func (a *App) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}

		a.mutex.RLock()
		defer a.mutex.RUnlock()

		httpjson.WriteArray(w, http.StatusOK, a.devices)
	})
}

// Start periodic fetch of data from netatmo API
func (a *App) Start(ctx context.Context) {
	if !a.Enabled() {
		slog.Warn("app is disabled")
		return
	}

	cron.New().Each(time.Minute*5).OnSignal(syscall.SIGUSR1).Now().OnError(func(err error) {
		slog.Error("refresh cron", "err", err)
	}).Start(ctx, func(ctx context.Context) error {
		devices, err := a.getDevices(ctx)
		if err != nil {
			return fmt.Errorf("fetch devices: %w", err)
		}

		a.mutex.Lock()
		defer a.mutex.Unlock()

		a.devices = devices

		return nil
	})
}

// Enabled check if app is enabled
func (a *App) Enabled() bool {
	return a.accessToken != ""
}

// HasScope check if given scope is configured
func (a *App) HasScope(scope string) bool {
	return strings.Contains(a.scopes, scope)
}
