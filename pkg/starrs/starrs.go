package starrs

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/Notifiarr/toolbarr/pkg/logs"
	"github.com/Notifiarr/toolbarr/pkg/mnd"
	"golift.io/starr"
)

const waitTime = 666 * time.Millisecond

var ErrInvalidApp = errors.New("an invalid app was provided; this may be a bug")

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
	SSL     bool          // verify ssl cert?
	Form    bool          // Use form login instead of basic auth?
	Timeout time.Duration // How long to wait for the app's API.
	App     string        // Radarr, Sonarr, etc
	Name    string        // Custom name: Radarr2, Radarr4k, etc.
	URL     string        // url to app.
	User    string        // username for app.
	Pass    string        // password for app.
	Key     string        // api key for app.
	DBPath  string        // path to database for app.
}

// Startup runs after wails initializes so we can save the context.
func Startup(ctx context.Context, starrs *Starrs, log *logs.Logger, app mnd.App) {
	const pruneInterval = 10 * time.Minute

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
	starrConfig := &starr.Config{
		APIKey: config.Key,
		URL:    strings.TrimSuffix(config.URL, "/") + "/",
		Client: starr.Client(config.Timeout, config.SSL),
	}

	if config.Form {
		starrConfig.Username = config.User
		starrConfig.Password = config.Pass
	} else {
		starrConfig.HTTPUser = config.User
		starrConfig.HTTPPass = config.Pass
	}

	return &instance{
		Starrs: s,
		config: config,
		Config: starrConfig,
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

type Selected map[int64]bool

func (s Selected) Count() (count int) { //nolint:nonamedreturns
	for _, b := range s {
		if b {
			count++
		}
	}

	return count
}
