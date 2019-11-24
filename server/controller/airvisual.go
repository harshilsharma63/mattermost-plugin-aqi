package controller

import (
	"github.com/harshilsharma63/mattermost-plugin-aqi/server/task"
	"net/http"
)

var refreshData = &Endpoint{
	Path:         "/pollution",
	Execute:      executeRefreshData,
	RequiresAuth: true,
}

func executeRefreshData(w http.ResponseWriter, r *http.Request) {
	if err := task.PublishPollutionData(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
