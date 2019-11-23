package airvisual

import "time"

// Go structs generated from JSON using this
// super cool service https://mholt.github.io/json-to-go/

type CityData struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

type Location struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Weather struct {
	Timestamp           time.Time `json:"ts"`
	Temperature         int       `json:"tp"` // in Celsius
	AtmosphericPressure int       `json:"pr"` // in hPa
	Humidity            int       `json:"hu"` // in %
	WindSpeed           float64   `json:"ws"` // in m/s
	WindDirection       int       `json:"wd"` // as an angle of 360° (N=0, E=90, S=180, W=270)
	IconCode            string    `json:"ic"`
}

type Pollution struct {
	Timestamp          time.Time `json:"ts"`
	AQIUS              int       `json:"aqius"` // main pollutant for US AQI
	MainPollutantUS    string    `json:"mainus"`
	AQIChina           int       `json:"aqicn"`
	MainPollutantChina string    `json:"maincn"` // main pollutant for Chinese AQI
}

type Current struct {
	Weather   Weather   `json:"weather"`
	Pollution Pollution `json:"pollution"`
}

type Data struct {
	City     string   `json:"city"`
	State    string   `json:"state"`
	Country  string   `json:"country"`
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}
