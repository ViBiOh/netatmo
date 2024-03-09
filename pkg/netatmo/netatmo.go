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

	absto "github.com/ViBiOh/absto/pkg/model"
	"github.com/ViBiOh/flags"
	"github.com/ViBiOh/httputils/v4/pkg/cron"
	"go.opentelemetry.io/otel/metric"
)

type Service struct {
	storage      absto.Storage
	token        Token
	clientID     string
	clientSecret string
	scopes       string
	devices      []Device
	mutex        sync.RWMutex
}

type Config struct {
	ClientID     string
	ClientSecret string
	Scopes       string
}

func Flags(fs *flag.FlagSet, prefix string) *Config {
	var config Config

	flags.New("ClientID", "Client ID").Prefix(prefix).DocPrefix("netatmo").StringVar(fs, &config.ClientID, "", nil)
	flags.New("ClientSecret", "Client Secret").Prefix(prefix).DocPrefix("netatmo").StringVar(fs, &config.ClientSecret, "", nil)
	flags.New("Scopes", "Scopes, comma separated").Prefix(prefix).DocPrefix("netatmo").StringVar(fs, &config.Scopes, "", nil)

	return &config
}

func New(config *Config, storage absto.Storage, meterProvider metric.MeterProvider) (*Service, error) {
	app := &Service{
		clientID:     config.ClientID,
		clientSecret: config.ClientSecret,
		scopes:       config.Scopes,
		storage:      storage,
	}

	if err := app.createMetrics(meterProvider, "temperature", "humidity", "noise", "co2", "pressure"); err != nil {
		return nil, fmt.Errorf("create metrics: %w", err)
	}

	return app, nil
}

func (s *Service) Start(ctx context.Context) {
	if err := s.loadToken(ctx); err != nil {
		slog.LogAttrs(ctx, slog.LevelError, "load token", slog.Any("error", err))
	}

	cron.New().Each(time.Minute*5).OnSignal(syscall.SIGUSR1).Now().OnError(func(ctx context.Context, err error) {
		slog.LogAttrs(ctx, slog.LevelError, "refresh cron", slog.Any("error", err))
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

func (s *Service) HasScope(scope string) bool {
	return strings.Contains(s.scopes, scope)
}
