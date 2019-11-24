package controller

import (
	"encoding/json"
	"github.com/harshilsharma63/mattermost-plugin-aqi/server/config"
	"github.com/harshilsharma63/mattermost-plugin-aqi/server/providers/airvisual"
	"github.com/harshilsharma63/mattermost-plugin-aqi/server/task"
	"github.com/harshilsharma63/mattermost-plugin-aqi/server/util"
	"net/http"
)

var refreshData = &Endpoint{
	Path:         "/refresh",
	Execute:      executeRefreshData,
	RequiresAuth: true,
}

func executeRefreshData(w http.ResponseWriter, r *http.Request) {
	// trying to fetch pollution data from cache
	data, appErr := config.Mattermost.KVGet(config.CacheKeyPollutionData)
	if appErr != nil {
		http.Error(w, appErr.Error(), appErr.StatusCode)
		return
	}

	// if no cached data found, fetch from provider servers
	if data == nil {
		if err := task.PublishPollutionData(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	// if cached data was found, use it
	citiesData := []*airvisual.CityData{}
	if err := json.Unmarshal(data, &citiesData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dataToPublish := map[string]interface{}{
		"pollutionData": string(data),
	}

	if err := util.PublishToAllTeams(dataToPublish, config.Mattermost); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
