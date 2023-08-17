package netatmo

import (
	"context"
	"regexp"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var sanitizeRegexp = regexp.MustCompile(`(?mi)(\S+).*`)

func (a *App) createMetrics(meterProvider metric.MeterProvider, names ...string) error {
	meter := meterProvider.Meter("github.com/ViBiOh/netatmo/pkg/netatmo")

	for _, name := range names {
		if err := a.createMetric(meter, name); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) createMetric(meter metric.Meter, name string) error {
	callback := func(ctx context.Context, fo metric.Float64Observer) error {
		a.mutex.RLock()
		defer a.mutex.RUnlock()

		switch name {
		case "temperature":
			for _, device := range a.devices {
				stationName := sanitizeName(device.StationName)

				fo.Observe(device.DashboardData.Temperature, metric.WithAttributes(
					attribute.String("station", stationName),
					attribute.String("module", device.ModuleName),
				))

				for _, module := range device.Modules {
					fo.Observe(module.DashboardData.Temperature, metric.WithAttributes(
						attribute.String("station", stationName),
						attribute.String("module", module.ModuleName),
					))
				}
			}
		case "humidity":
			for _, device := range a.devices {
				stationName := sanitizeName(device.StationName)

				fo.Observe(device.DashboardData.Humidity, metric.WithAttributes(
					attribute.String("station", stationName),
					attribute.String("module", device.ModuleName),
				))

				for _, module := range device.Modules {
					fo.Observe(module.DashboardData.Humidity, metric.WithAttributes(
						attribute.String("station", stationName),
						attribute.String("module", module.ModuleName),
					))
				}
			}
		case "noise":
			for _, device := range a.devices {
				stationName := sanitizeName(device.StationName)

				fo.Observe(device.DashboardData.Noise, metric.WithAttributes(
					attribute.String("station", stationName),
					attribute.String("module", device.ModuleName),
				))
			}
		case "co2":
			for _, device := range a.devices {
				stationName := sanitizeName(device.StationName)

				fo.Observe(device.DashboardData.CO2, metric.WithAttributes(
					attribute.String("station", stationName),
					attribute.String("module", device.ModuleName),
				))
			}
		case "pressure":
			for _, device := range a.devices {
				stationName := sanitizeName(device.StationName)

				fo.Observe(device.DashboardData.Pressure, metric.WithAttributes(
					attribute.String("station", stationName),
					attribute.String("module", device.ModuleName),
				))
			}
		}

		return nil
	}

	_, err := meter.Float64ObservableGauge("netatmo_"+name, metric.WithFloat64Callback(callback))

	return err
}

func sanitizeName(name string) string {
	matches := sanitizeRegexp.FindAllStringSubmatch(name, -1)
	if len(matches) == 0 {
		return name
	}

	return matches[0][1]
}
