package netatmo

import (
	"context"
	"fmt"
	"regexp"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var sanitizeRegexp = regexp.MustCompile(`(?mi)(\S+).*`)

func createMetrics(meterProvider metric.MeterProvider, names ...string) (map[string]metric.Float64Gauge, error) {
	meter := meterProvider.Meter("github.com/ViBiOh/netatmo/v2/pkg/netatmo")
	metrics := make(map[string]metric.Float64Gauge, len(names))

	for _, name := range names {
		gauge, err := meter.Float64Gauge("netatmo." + name)
		if err != nil {
			return nil, fmt.Errorf("create gauge for `%s`: %w", name, err)
		}

		metrics[name] = gauge
	}

	return metrics, nil
}

func (s *Service) observeMetrics(ctx context.Context, devices []Device) {
	for name, meter := range s.metrics {
		switch name {
		case "temperature":
			for _, device := range devices {
				stationName := sanitizeName(device.StationName)

				meter.Record(ctx, device.DashboardData.Temperature, metric.WithAttributes(
					attribute.String("station", stationName),
					attribute.String("module", device.ModuleName),
				))

				for _, module := range device.Modules {
					meter.Record(ctx, module.DashboardData.Temperature, metric.WithAttributes(
						attribute.String("station", stationName),
						attribute.String("module", module.ModuleName),
					))
				}
			}

		case "humidity":
			for _, device := range devices {
				stationName := sanitizeName(device.StationName)

				meter.Record(ctx, device.DashboardData.Humidity, metric.WithAttributes(
					attribute.String("station", stationName),
					attribute.String("module", device.ModuleName),
				))

				for _, module := range device.Modules {
					meter.Record(ctx, module.DashboardData.Humidity, metric.WithAttributes(
						attribute.String("station", stationName),
						attribute.String("module", module.ModuleName),
					))
				}
			}

		case "noise":
			for _, device := range devices {
				stationName := sanitizeName(device.StationName)

				meter.Record(ctx, device.DashboardData.Noise, metric.WithAttributes(
					attribute.String("station", stationName),
					attribute.String("module", device.ModuleName),
				))
			}

		case "co2":
			for _, device := range devices {
				stationName := sanitizeName(device.StationName)

				meter.Record(ctx, device.DashboardData.CO2, metric.WithAttributes(
					attribute.String("station", stationName),
					attribute.String("module", device.ModuleName),
				))
			}

		case "pressure":
			for _, device := range devices {
				stationName := sanitizeName(device.StationName)

				meter.Record(ctx, device.DashboardData.Pressure, metric.WithAttributes(
					attribute.String("station", stationName),
					attribute.String("module", device.ModuleName),
				))
			}
		}
	}
}

func sanitizeName(name string) string {
	matches := sanitizeRegexp.FindAllStringSubmatch(name, -1)
	if len(matches) == 0 {
		return name
	}

	return matches[0][1]
}
