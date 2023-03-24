package starrs

import (
	"github.com/Notifiarr/toolbarr/pkg/logs"
	"golift.io/starr"
)

// StarrConfig allows interacting with the instances using a standard interface.
type StarrConfig struct {
	*logs.Logger
	*starr.Config
	instance *Instance
}

// Instances is the configured list of instances.
type Instances map[string][]Instance

// Instance config.
type Instance struct {
	App    string // Radarr, Sonarr, etc
	Name   string // Custom name: Radarr2, Radarr4k, etc.
	URL    string // url to app.
	User   string // username for app.
	Pass   string // password for app.
	Key    string // api key for app.
	DBPath string // path to database for app.
	SSL    bool   // verify ssl cert?
}

// Copy the instances list. Using copies provides thread safety at the risk of deviergence.
func (i Instances) Copy() Instances {
	instances := make(Instances)
	for k := range i {
		instances[k] = make([]Instance, len(i[k]))
		copy(instances[k], i[k])
	}

	return instances
}
