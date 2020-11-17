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

func sanitizeSource(source string) string {
	matches := sanitizeRegexp.FindAllStringSubmatch(source, -1)
	if len(matches) == 0 {
		return source
	}

	return matches[0][1]
}

func (a *app) getMetrics(prefix, suffix string) prometheus.Gauge {
	gauge, ok := a.prometheusCollectors[fmt.Sprintf("%s_%s", prefix, suffix)]
	if !ok {
		gauge = prometheus.NewGauge(prometheus.GaugeOpts{
			Name: fmt.Sprintf("%s_%s_%s", sanitizeSource(strings.ToLower(Source)), prefix, suffix),
		})

		a.prometheusCollectors[fmt.Sprintf("%s_%s", prefix, suffix)] = gauge

		if a.registerer != nil {
			a.registerer.MustRegister(gauge)
		}
	}

	return gauge
}

func (a *app) updatePrometheus() {
	for _, device := range a.devices {
		a.getMetrics(strings.ToLower(device.StationName), "temperature").Set(float64(device.DashboardData.Temperature))
		a.getMetrics(strings.ToLower(device.StationName), "humidity").Set(float64(device.DashboardData.Humidity))
		a.getMetrics(strings.ToLower(device.StationName), "noise").Set(float64(device.DashboardData.Noise))
		a.getMetrics(strings.ToLower(device.StationName), "co2").Set(float64(device.DashboardData.CO2))

		for _, module := range device.Modules {
			a.getMetrics(strings.ToLower(fmt.Sprintf("%s_%s", device.StationName, module.ModuleName)), "temperature").Set(float64(module.DashboardData.Temperature))
			a.getMetrics(strings.ToLower(fmt.Sprintf("%s_%s", device.StationName, module.ModuleName)), "humidity").Set(float64(module.DashboardData.Humidity))

		}
	}
}
