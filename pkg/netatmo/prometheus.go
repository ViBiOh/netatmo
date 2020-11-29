package netatmo

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	sanitizeRegexp = regexp.MustCompile(`(?mi)(\S+).*`)
)

func sanitizeName(name string) string {
	matches := sanitizeRegexp.FindAllStringSubmatch(name, -1)
	if len(matches) == 0 {
		return name
	}

	return matches[0][1]
}

func (a *app) getMetrics(device, module, suffix string) prometheus.Gauge {
	var name string
	if len(module) == 0 {
		name = strings.ToLower(fmt.Sprintf("%s_%s", device, suffix))
	} else {
		name = strings.ToLower(fmt.Sprintf("%s_%s_%s", device, module, suffix))
	}

	gauge, ok := a.prometheusCollectors[name]
	if !ok {
		gauge = prometheus.NewGauge(prometheus.GaugeOpts{
			Name: fmt.Sprintf("%s_%s", Source, name),
		})

		a.prometheusCollectors[name] = gauge

		if a.registerer != nil {
			a.registerer.MustRegister(gauge)
		}
	}

	return gauge
}

func (a *app) updatePrometheus() {
	for _, device := range a.devices {
		stationName := sanitizeName(device.StationName)

		a.getMetrics(stationName, device.ModuleName, "temperature").Set(float64(device.DashboardData.Temperature))
		a.getMetrics(stationName, device.ModuleName, "humidity").Set(float64(device.DashboardData.Humidity))
		a.getMetrics(stationName, device.ModuleName, "noise").Set(float64(device.DashboardData.Noise))
		a.getMetrics(stationName, device.ModuleName, "co2").Set(float64(device.DashboardData.CO2))

		for _, module := range device.Modules {
			a.getMetrics(stationName, module.ModuleName, "temperature").Set(float64(module.DashboardData.Temperature))
			a.getMetrics(stationName, module.ModuleName, "humidity").Set(float64(module.DashboardData.Humidity))

		}
	}
}
