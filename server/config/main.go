package config

import (
	"encoding/json"
	"github.com/mattermost/mattermost-server/plugin"
	"github.com/pkg/errors"
	"go.uber.org/atomic"
	"strings"
	"time"
)

const (
	CommandPrefix             = PluginName
	URLMappingKeyPrefix       = "url_"
	ServerExeToWebappRootPath = "/../webapp"

	URLPluginBase = "/plugins/" + PluginName
	URLStaticBase = URLPluginBase + "/static"

	HeaderMattermostUserID = "Mattermost-User-Id"

	RunnerInterval = 30 * time.Second
)

var (
	config     atomic.Value
	Mattermost plugin.API
)

type Location struct {
	Country string
	State string
	City string
}

type Configuration struct {
	AirVisualAPIKey string
	Locations string

	// derived attributes
	DerivedLocations []Location
}

func GetConfig() *Configuration {
	return config.Load().(*Configuration)
}

func SetConfig(c *Configuration) {
	config.Store(c)
}

func (c *Configuration) ProcessConfiguration() error {
	// any post-processing on configurations goes here
	c.AirVisualAPIKey = strings.Trim(c.AirVisualAPIKey, " ")
	c.Locations = strings.Trim(c.Locations, " ")

	l := []string{}
	if err := json.Unmarshal([]byte(c.Locations), &l); err != nil {
		return err
	}

	for _, location := range l {
		components := strings.Split(location, "_")
		c.DerivedLocations = append(c.DerivedLocations, Location{
			Country: components[0],
			State:   components[1],
			City:    components[2],
		})
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
