package netatmo

import (
	"fmt"
	"regexp"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	sanitizeRegexp = regexp.MustCompile(`(?mi)(\S+).*`)
)

func createMetrics(prometheusRegisterer prometheus.Registerer, names ...string) (map[string]*prometheus.GaugeVec, error) {
	if prometheusRegisterer == nil {
		return nil, nil
	}

	metrics := make(map[string]*prometheus.GaugeVec)
	for _, name := range names {
		metric, err := createMetric(prometheusRegisterer, name)
		if err != nil {
			return nil, err
		}

		metrics[name] = metric
	}

	return metrics, nil
}

func createMetric(prometheusRegisterer prometheus.Registerer, name string) (*prometheus.GaugeVec, error) {
	if prometheusRegisterer == nil {
		return nil, nil
	}

	gauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: Source,
		Name:      name,
	}, []string{"station", "module"})

	if err := prometheusRegisterer.Register(gauge); err != nil {
		return nil, fmt.Errorf("unable to registrer %s: %s", name, err)
	}

	return gauge, nil
}

func (a *App) setMetric(name, station, module string, value float64) {
	metric, ok := a.metrics[name]
	if !ok {
		return
	}

	labels := prometheus.Labels{
		"station": station,
		"module":  module,
	}

	metric.With(labels).Set(value)
}

func (a *App) updatePrometheus() {
	for _, device := range a.devices {
		stationName := sanitizeName(device.StationName)

		a.setMetric("temperature", stationName, device.ModuleName, device.DashboardData.Temperature)
		a.setMetric("humidity", stationName, device.ModuleName, device.DashboardData.Humidity)
		a.setMetric("noise", stationName, device.ModuleName, device.DashboardData.Noise)
		a.setMetric("co2", stationName, device.ModuleName, device.DashboardData.CO2)
		a.setMetric("pressure", stationName, device.ModuleName, device.DashboardData.Pressure)

		for _, module := range device.Modules {
			a.setMetric("temperature", stationName, module.ModuleName, module.DashboardData.Temperature)
			a.setMetric("humidity", stationName, module.ModuleName, module.DashboardData.Humidity)
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
