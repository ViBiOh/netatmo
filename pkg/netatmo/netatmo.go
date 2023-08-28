package netatmo

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/ViBiOh/flags"
	"github.com/ViBiOh/httputils/v4/pkg/cron"
	"go.opentelemetry.io/otel/metric"
)

type Service struct {
	clientID     string
	clientSecret string
	accessToken  string
	refreshToken string
	scopes       string

	devices []Device

	mutex sync.RWMutex
}

type Config struct {
	AccessToken  string
	RefreshToken string
	ClientID     string
	ClientSecret string
	Scopes       string
}

func Flags(fs *flag.FlagSet, prefix string) *Config {
	var config Config

	flags.New("AccessToken", "Access Token").Prefix(prefix).DocPrefix("netatmo").StringVar(fs, &config.AccessToken, "", nil)
	flags.New("RefreshToken", "Refresh Token").Prefix(prefix).DocPrefix("netatmo").StringVar(fs, &config.RefreshToken, "", nil)
	flags.New("ClientID", "Client ID").Prefix(prefix).DocPrefix("netatmo").StringVar(fs, &config.ClientID, "", nil)
	flags.New("ClientSecret", "Client Secret").Prefix(prefix).DocPrefix("netatmo").StringVar(fs, &config.ClientSecret, "", nil)
	flags.New("Scopes", "Scopes, comma separated").Prefix(prefix).DocPrefix("netatmo").StringVar(fs, &config.Scopes, "", nil)

	return &config
}

func New(config *Config, meterProvider metric.MeterProvider) (*Service, error) {
	app := &Service{
		clientID:     config.ClientID,
		clientSecret: config.ClientSecret,
		accessToken:  config.AccessToken,
		refreshToken: config.RefreshToken,
		scopes:       config.Scopes,
	}

	if err := app.createMetrics(meterProvider, "temperature", "humidity", "noise", "co2", "pressure"); err != nil {
		return nil, fmt.Errorf("create metrics: %w", err)
	}

	return app, nil
}

func (s *Service) Start(ctx context.Context) {
	if !s.Enabled() {
		slog.Warn("app is disabled")
		return
	}

	cron.New().Each(time.Minute*5).OnSignal(syscall.SIGUSR1).Now().OnError(func(err error) {
		slog.Error("refresh cron", "err", err)
	}).Start(ctx, func(ctx context.Context) error {
		devices, err := s.getDevices(ctx)
		if err != nil {
			return fmt.Errorf("fetch devices: %w", err)
		}

		s.mutex.Lock()
		defer s.mutex.Unlock()

		s.devices = devices

		return nil
	})
}

func (s *Service) Enabled() bool {
	return s.accessToken != ""
}

func (s *Service) HasScope(scope string) bool {
	return strings.Contains(s.scopes, scope)
}
