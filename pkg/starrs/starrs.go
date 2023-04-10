package starrs

import (
	"context"
	"strings"

	"github.com/Notifiarr/toolbarr/pkg/logs"
	"github.com/Notifiarr/toolbarr/pkg/mnd"
	"golift.io/starr"
)

// Starrs holds the running data and provides the frontend a place
// to interact with starr instances and their databases.
type Starrs struct {
	ctx context.Context
	app mnd.App
	log *logs.Logger
}

// instance allows interacting with the instances via HTTP API using a standard interface.
type instance struct {
	*Starrs
	config *AppConfig
	*starr.Config
}

// Instances is the configured list of instances. The map key is the instance type, e.g. Lidarr.
type Instances map[string][]AppConfig

// AppConfig is the configuration for an instance.
type AppConfig struct {
	App    string // Radarr, Sonarr, etc
	Name   string // Custom name: Radarr2, Radarr4k, etc.
	URL    string // url to app.
	User   string // username for app.
	Pass   string // password for app.
	Key    string // api key for app.
	DBPath string // path to database for app.
	SSL    bool   // verify ssl cert?
}

// Startup runs after wails inializes so we can save the context.
func Startup(ctx context.Context, starrs *Starrs, log *logs.Logger, app mnd.App) {
	starrs.ctx = ctx
	starrs.app = app
	starrs.log = log
}

// Copy the instances list. Using copies provides thread safety at the risk of inconsistency.
func (i Instances) Copy() Instances {
	instances := make(Instances)
	for k := range i {
		instances[k] = make([]AppConfig, len(i[k]))
		copy(instances[k], i[k])
	}

	return instances
}

func (s *Starrs) newInstance(config *AppConfig) *instance {
	return &instance{
		Starrs: s,
		config: config,
		Config: &starr.Config{
			APIKey: config.Key,
			URL:    strings.TrimSuffix(config.URL, "/") + "/",
			// HTTPPass: instance.Pass,
			// HTTPUser: instance.User,
			Username: config.User,
			Password: config.Pass,
			Client:   starr.Client(timeout, config.SSL),
		},
	}
}

func (s *Starrs) newAPIinstance(config *AppConfig) (*instance, error) {
	instance := s.newInstance(config)
	if instance.APIKey == "" {
		data, err := instance.testWithoutKey()
		if err != nil {
			return nil, err
		}

		instance.Config.APIKey = data.Key
	}

	return instance, nil
}
