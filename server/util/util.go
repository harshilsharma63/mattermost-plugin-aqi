package util

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
	"github.com/pkg/errors"
)

func PublishToAllTeams(data map[string]interface{}, mattermost plugin.API) error {
	teams, appErr := mattermost.GetTeams()
	if appErr != nil {
		return errors.New(appErr.Error())
	}

	for _, team := range teams {
		mattermost.PublishWebSocketEvent(
			"receive_pollution_data",
			data,
			&model.WebsocketBroadcast{
				TeamId: team.Id,
			},
		)
	}

	return nil
}
