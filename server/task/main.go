package task

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/harshilsharma63/mattermost-plugin-aqi/server/config"
	"github.com/harshilsharma63/mattermost-plugin-aqi/server/service"
	"github.com/harshilsharma63/mattermost-plugin-aqi/server/util"
)

func PublishPollutionData() error {
	citiesData, err := service.GetPollutionData()
	if err != nil {
		return err
	}

	dataToPublish := map[string]interface{}{
		"id":   uuid.New().String(),
		"data": citiesData,
	}

	data, err := json.Marshal(dataToPublish)
	if err != nil {
		return err
	}

	_ = config.Mattermost.KVDelete(config.CacheKeyPollutionData)
	if appErr := config.Mattermost.KVSetWithExpiry(config.CacheKeyPollutionData, data, config.PollutionDataCacheExpirySeconds); appErr != nil {
		return errors.New(appErr.Error())
	}

	if err := util.PublishToAllTeams(map[string]interface{}{
		"data": string(data),
	}, config.Mattermost); err != nil {
		return err
	}

	return nil
}
