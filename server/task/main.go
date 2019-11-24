package task

import (
	"encoding/json"
	"errors"
	"github.com/harshilsharma63/mattermost-plugin-aqi/server/config"
	"github.com/harshilsharma63/mattermost-plugin-aqi/server/providers/airvisual"
	"github.com/mattermost/mattermost-server/model"
)

func PublishPollutionData() error {
	citiesData := []*airvisual.CityData{}

	for _, location := range config.GetConfig().DerivedLocations {
		cityData, err := airvisual.GetCityData(config.GetConfig().AirVisualAPIKey, location.Country, location.State, location.City)
		if err != nil {
			return err
		}

		citiesData = append(citiesData, cityData)
	}

	teams, appErr := config.Mattermost.GetTeams()
	if appErr != nil {
		return errors.New(appErr.Error())
	}

	data, err := json.Marshal(citiesData)
	if err != nil {
		return err
	}

	for _, team := range teams {
		config.Mattermost.PublishWebSocketEvent(
			"receive_pollution_data",
			map[string]interface{}{
				"pollutionData": string(data),
			},
			&model.WebsocketBroadcast{
				TeamId: team.Id,
			},
		)
	}

	return nil
}
