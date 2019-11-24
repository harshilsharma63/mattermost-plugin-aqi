package controller

import (
	"encoding/json"
	"github.com/harshilsharma63/mattermost-plugin-aqi/server/config"
	"github.com/harshilsharma63/mattermost-plugin-aqi/server/providers/airvisual"
	"net/http"
)

var getPollutionData = &Endpoint{
	Path:         "/pollution",
	Execute:      executeGetPollutionData,
	RequiresAuth: true,
}

func executeGetPollutionData(w http.ResponseWriter, r *http.Request) {
	citiesData := []*airvisual.CityData{}

	for _, location := range config.GetConfig().DerivedLocations {
		cityData, err := airvisual.GetCityData(config.GetConfig().AirVisualAPIKey, location.Country, location.State, location.City)
		if err != nil {
			config.Mattermost.LogError(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		citiesData = append(citiesData,  cityData)
	}

	data, err := json.Marshal(citiesData)
	if err != nil {
		config.Mattermost.LogError(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(data); err != nil {
		config.Mattermost.LogError("Error occurred in writing data to HTTP response", err, nil)
	}
}
