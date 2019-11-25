package service

import (
	"github.com/harshilsharma63/mattermost-plugin-aqi/server/config"
	"github.com/harshilsharma63/mattermost-plugin-aqi/server/providers/airvisual"
	"time"
)

func GetPollutionData() ([]*airvisual.CityData, error) {
	citiesData := []*airvisual.CityData{}

	for i, location := range config.GetConfig().DerivedLocations {
		if i > 0 {
			time.Sleep(5 * time.Second)
		}

		cityData, err := airvisual.GetCityData(config.GetConfig().AirVisualAPIKey, location.Country, location.State, location.City)
		if err != nil {
			return nil, err
		}

		citiesData = append(citiesData, cityData)
	}

	return citiesData, nil
}
