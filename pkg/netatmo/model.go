package netatmo

const (
	// DevicesAction action for listing devices
	DevicesAction = "devices"

	// Source constant for worker message
	Source = "netatmo"
)

var (
	noneStationsData StationsData
)

// StationsData contains data retrieved when getting stations datas
type StationsData struct {
	Body struct {
		Devices []Device `json:"devices"`
	} `json:"body"`
}

// Device contains a device data
type Device struct {
	StationName string `json:"station_name"`
	ModuleName  string `json:"module_name"`
	Modules     []struct {
		ModuleName    string        `json:"module_name"`
		DashboardData DashboardData `json:"dashboard_data"`
	} `json:"modules"`
	DashboardData DashboardData `json:"dashboard_data"`
}

// DashboardData contains dashboard data
type DashboardData struct {
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
	Noise       float32 `json:"noise"`
	CO2         float32 `json:"co2"`
}

// Error describes error
type Error struct {
	Error struct {
		Message string
		Code    int
	}
}

// Token describes refresh token response
type Token struct {
	AccessToken string `json:"access_token"`
}
