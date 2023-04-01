package starrs

import (
	"context"

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
	config *StarrConfig
	*starr.Config
}

// Instances is the configured list of instances. The map key is the instance type, e.g. Lidarr.
type Instances map[string][]StarrConfig

// StarrConfig is the configuration for an instance.
type StarrConfig struct {
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
		instances[k] = make([]StarrConfig, len(i[k]))
		copy(instances[k], i[k])
	}

	return instances
}
