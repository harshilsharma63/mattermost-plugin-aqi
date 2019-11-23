package config

import (
	"github.com/mattermost/mattermost-server/plugin"
	"github.com/pkg/errors"
	"go.uber.org/atomic"
	"strings"
)

const (
	CommandPrefix             = PluginName
	URLMappingKeyPrefix       = "url_"
	ServerExeToWebappRootPath = "/../webapp"

	URLPluginBase = "/plugins/" + PluginName
	URLStaticBase = URLPluginBase + "/static"

	HeaderMattermostUserID = "Mattermost-User-Id"
)

var (
	config     atomic.Value
	Mattermost plugin.API
)

type Configuration struct {
	AirVisualAPIKey string
}

func GetConfig() *Configuration {
	return config.Load().(*Configuration)
}

func SetConfig(c *Configuration) {
	config.Store(c)
}

func (c *Configuration) ProcessConfiguration() error {
	// any post-processing on configurations goes here
	if strings.Trim(c.AirVisualAPIKey, " ") == "" {
		return errors.New("AirVisualAPIKey cannot be empty")
	}

	return nil
}

func (c *Configuration) IsValid() error {
	// Add config validations here.
	// Check for required fields, formats, etc.

	if strings.Trim(c.AirVisualAPIKey, " ") == "" {
		Mattermost.LogError("AirVisualAPIKey cannot be empty")
		return errors.New("AirVisualAPIKey cannot be empty")
	}

	return nil
}
