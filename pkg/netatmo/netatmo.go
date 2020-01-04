package netatmo

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/ViBiOh/httputils/v3/pkg/cron"
	"github.com/ViBiOh/httputils/v3/pkg/flags"
	"github.com/ViBiOh/httputils/v3/pkg/httpjson"
	"github.com/ViBiOh/httputils/v3/pkg/logger"
	httpprom "github.com/ViBiOh/httputils/v3/pkg/prometheus"
	"github.com/ViBiOh/httputils/v3/pkg/swagger"
	"github.com/prometheus/client_golang/prometheus"
)

// App of package
type App interface {
	Handler() http.Handler
	Swagger() (swagger.Configuration, error)
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

	registerer prometheus.Registerer
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
func New(config Config, prometheusApp httpprom.App) App {
	var registerer prometheus.Registerer
	if prometheusApp != nil {
		registerer = prometheusApp.Registerer()
	}

	return &app{
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
func (a *app) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}

		a.mutex.RLock()
		defer a.mutex.RUnlock()

		httpjson.ResponseArrayJSON(w, http.StatusOK, a.devices, httpjson.IsPretty(r))
	})
}

func (a *app) Swagger() (swagger.Configuration, error) {
	return swagger.Configuration{
		Paths: `/devices:
  get:
    description: Netatmo devices
    responses:
      '200':
        description: Devices data
        content:
          application/json:
            schema:
              type: object
              properties:
                results:
                  $ref: '#/components/schemas/Devices'`,
		Components: `DashboardData:
  type: object
  properties:
    temperature:
      type: number
      format: float
      description: Temperature in celcius
    humidity:
      type: number
      format: float
      description: Humidity in percentage
    noise:
      type: number
      format: float
      description: Noise in dB
    co2:
      type: number
      format: float
      description: CO2 in ppm

Device:
  type: object
  properties:
    station_name:
      type: string
      description: Station name
    dashboard_data:
      description: Station data
      $ref: '#/components/schemas/DashboardData'
    modules:
      type: array
      description: Additional modules
      items:
        type: object
        properties:
          module_name:
            type: string
            description: Module name
          dashboard_data:
            description: Station data
            $ref: '#/components/schemas/DashboardData'

Devices:
  type: array
  items:
    $ref: '#/components/schemas/Device'`,
	}, nil
}

// Start periodic fetch of data from netatmo API
func (a *app) Start() {
	if !a.Enabled() {
		logger.Warn("app is disabled")
		return
	}

	cron.New().Each(time.Minute*5).Now().Start(func(_ time.Time) error {
		devices, err := a.GetDevices(context.Background())
		if err != nil {
			return fmt.Errorf("unable to fetch devices: %s", err)
		}

		a.mutex.Lock()
		defer a.mutex.Unlock()

		a.devices = devices
		a.updatePrometheus()

		return nil
	}, func(err error) {
		logger.Error("%s", err)
	})
}

// Enabled check if app is enabled
func (a *app) Enabled() bool {
	return a.accessToken != ""
}

// HasScope check if given scope is configured
func (a *app) HasScope(scope string) bool {
	return strings.Contains(a.scopes, scope)
}
