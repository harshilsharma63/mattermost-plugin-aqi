package config

import (
	"encoding/json"
	"github.com/harshilsharma63/mattermost-plugin-aqi/server/util"
	"github.com/mattermost/mattermost-server/model"
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

	CacheKeyPollutionData           = "pollution_data"
	PollutionDataCacheExpirySeconds = 3600

	CacheKeyLocationConfigHash = "location_config_hash"
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

func PurgePollutionDataIfRequired() *model.AppError {
	should, appErr := shouldPurgeLocationData()
	if appErr != nil {
		return appErr
	}

	if should {
		if appErr := Mattermost.KVDelete(CacheKeyPollutionData); appErr != nil {
			return appErr
		}

		newLocationConfigHash := util.GetMD5Hash(GetConfig().Locations)
		if appErr := Mattermost.KVSet(CacheKeyPollutionData, []byte(newLocationConfigHash)); appErr != nil {
			return appErr
		}
	}

	return nil
}

func shouldPurgeLocationData() (bool, *model.AppError) {
	data, appErr := Mattermost.KVGet(CacheKeyLocationConfigHash)
	if appErr != nil {
		return false, appErr
	}

	if data == nil {
		return true, nil
	}

	oldLocationConfigHash := string(data)
	newLocationConfigHash := util.GetMD5Hash(GetConfig().Locations)

	return oldLocationConfigHash == newLocationConfigHash, nil
}
