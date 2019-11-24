package task

import (
	"encoding/json"
	"errors"
	"github.com/harshilsharma63/mattermost-plugin-aqi/server/config"
	"github.com/harshilsharma63/mattermost-plugin-aqi/server/providers/airvisual"
	"github.com/harshilsharma63/mattermost-plugin-aqi/server/util"
	"time"
)

func PublishPollutionData() error {
	citiesData := []*airvisual.CityData{}

	for i, location := range config.GetConfig().DerivedLocations {
		if i > 0 {
			time.Sleep(5 * time.Second)
		}

		cityData, err := airvisual.GetCityData(config.GetConfig().AirVisualAPIKey, location.Country, location.State, location.City)
		if err != nil {
			return err
		}

		citiesData = append(citiesData, cityData)
	}

	data, err := json.Marshal(citiesData)
	if err != nil {
		return err
	}

	if appErr := config.Mattermost.KVSetWithExpiry(config.CacheKeyPollutionData, data, config.PollutionDataCacheExpirySeconds); appErr != nil {
		return errors.New(appErr.Error())
	}

	dataToPublish := map[string]interface{}{
		"pollutionData": string(data),
	}

	if err := util.PublishToAllTeams(dataToPublish, config.Mattermost); err != nil {
		return err
	}

	return nil
}
