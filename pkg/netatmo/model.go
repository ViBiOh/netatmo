package netatmo

import "time"

const (
	DevicesAction = "devices"
	Source        = "netatmo"
)

type StationsData struct {
	Body struct {
		Devices []Device `json:"devices"`
	} `json:"body"`
}

type Device struct {
	StationName string `json:"station_name"`
	ModuleName  string `json:"module_name"`
	Modules     []struct {
		ModuleName    string        `json:"module_name"`
		BatterPercent int           `json:"battery_percent"`
		DashboardData DashboardData `json:"dashboard_data"`
	} `json:"modules"`
	DashboardData DashboardData `json:"dashboard_data"`
}

type DashboardData struct {
	Temperature float64 `json:"Temperature"`
	Humidity    float64 `json:"Humidity"`
	Noise       float64 `json:"Noise"`
	CO2         float64 `json:"CO2"`
	Pressure    float64 `json:"Pressure"`
}

type Error struct {
	Error struct {
		Message string
		Code    int
	}
}

type Token struct {
	ExpiresAt    time.Time `json:"expires_at"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresIn    int64     `json:"expires_in"`
}

func (t Token) IsExpired() bool {
	return t.ExpiresAt.Before(time.Now())
}

func (t Token) ComputeExpire() Token {
	t.ExpiresAt = time.Now().Add(time.Second * time.Duration(t.ExpiresIn))

	return t
}
