package main

import (
	"github.com/harshilsharma63/mattermost-plugin-aqi/server/task"
	"net/http"
	"time"

	"github.com/mattermost/mattermost-server/plugin"

	"github.com/harshilsharma63/mattermost-plugin-aqi/server/config"
	"github.com/harshilsharma63/mattermost-plugin-aqi/server/controller"
)

type Plugin struct {
	plugin.MattermostPlugin

	handler http.Handler
	running bool
}

func (p *Plugin) OnActivate() error {
	config.Mattermost = p.API

	if err := p.OnConfigurationChange(); err != nil {
		return err
	}

	p.Run()

	return nil
}

func (p *Plugin) OnConfigurationChange() error {
	if config.Mattermost != nil {
		var configuration config.Configuration

		if err := config.Mattermost.LoadPluginConfiguration(&configuration); err != nil {
			config.Mattermost.LogError("Error in LoadPluginConfiguration: " + err.Error())
			return err
		}

		if err := configuration.ProcessConfiguration(); err != nil {
			config.Mattermost.LogError("Error in ProcessConfiguration: " + err.Error())
			return err
		}

		if err := configuration.IsValid(); err != nil {
			config.Mattermost.LogError("Error in Validating Configuration: " + err.Error())
			return err
		}

		config.SetConfig(&configuration)

		_ = config.Mattermost.KVDelete(config.CacheKeyPollutionData)

		go func() {
			if err := task.PublishPollutionData(); err != nil {
				config.Mattermost.LogError("", err, nil)
			}
		}()
	}
	return nil
}

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	conf := config.GetConfig()

	if err := conf.IsValid(); err != nil {
		p.API.LogError("This plugin is not configured: " + err.Error())
		http.Error(w, "This plugin is not configured.", http.StatusNotImplemented)
		return
	}

	path := r.URL.Path
	endpoint := controller.Endpoints[path]

	if endpoint == nil {
		p.handler.ServeHTTP(w, r)
	} else if !endpoint.RequiresAuth || controller.Authenticated(w, r) {
		endpoint.Execute(w, r)
	}
}

func (p *Plugin) Run() {
	if !p.running {
		p.running = true
		p.runner()
	}
}

func (p *Plugin) runner() {
	go func() {
		<-time.NewTimer(config.RunnerInterval).C
		if err := task.PublishPollutionData(); err != nil {
			config.Mattermost.LogError("", err, nil)
		}
		if !p.running {
			return
		}
		p.runner()
	}()
}

func main() {
	plugin.ClientMain(&Plugin{})
}
