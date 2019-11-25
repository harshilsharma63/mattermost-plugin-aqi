package controller

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/harshilsharma63/mattermost-plugin-aqi/server/config"
	"github.com/harshilsharma63/mattermost-plugin-aqi/server/providers/airvisual"
	"github.com/harshilsharma63/mattermost-plugin-aqi/server/service"
	"net/http"
)

var refreshData = &Endpoint{
	Path:         "/refresh",
	Execute:      executeRefreshData,
	RequiresAuth: true,
}

func executeRefreshData(w http.ResponseWriter, r *http.Request) {
	citiesData := []*airvisual.CityData{}
	dataToPublish := map[string]interface{}{}

	// trying to fetch pollution data from cache
	data, appErr := config.Mattermost.KVGet(config.CacheKeyPollutionData)
	if appErr != nil {
		http.Error(w, appErr.Error(), appErr.StatusCode)
		return
	}

	if data != nil {
		if err := json.Unmarshal(data, &dataToPublish); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		var err error
		citiesData, err = service.GetPollutionData()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		dataToPublish = map[string]interface{}{
			"id":   uuid.New().String(),
			"data": citiesData,
		}

		// now cache the data
		data, _ := json.Marshal(dataToPublish)
		if appErr := config.Mattermost.KVSetWithExpiry(config.CacheKeyPollutionData, data, config.PollutionDataCacheExpirySeconds); appErr != nil {
			http.Error(w, appErr.Error(), appErr.StatusCode)
			return
		}
	}

	data, err := json.Marshal(dataToPublish)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
