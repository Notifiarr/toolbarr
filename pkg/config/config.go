//nolint:gomnd
package config

import (
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/Notifiarr/toolbarr/pkg/logs"
	"github.com/Notifiarr/toolbarr/pkg/mnd"
	"github.com/Notifiarr/toolbarr/pkg/starrs"
	"github.com/Notifiarr/toolbarr/pkg/translations"
	"github.com/mitchellh/go-homedir"
	"golang.org/x/text/language"
	"golift.io/starr"
)

const confExt = ".conf"

var ErrEmptyInput = fmt.Errorf("input must have at least name or path")

// Input data to open a config file.
// If Dir!="" then config is placed in a sub directory.
type Input struct {
	// File path to config file. Optional.
	// Name and Dir are not used for config file path if this is provided.
	File string
	// Application name. Used to construct a config file path.
	// Also used for log file names. Overrides what's in config file.
	Name string
	// Sub-directory to put the config file into. Optional.
	Dir string
}

// Config provides an interface to read and write config values, thread safe.
type Config struct {
	file     string // file read from and written to.
	settings *Settings
	ask      chan *Settings
	rep      chan *Settings
}

// Get opens/reads or creates/writes a config file.
func Get(i *Input) (*Config, error) {
	return i.load()
}

// load the config from the provided input data.
func (i *Input) load() (*Config, error) {
	i.findConfigFileFolder()

	if i.File != "" {
		return i.openCustomPath()
	} else if i.Name == "" {
		return nil, ErrEmptyInput
	}

	configDir, err := i.getConfigDir()
	if err != nil {
		return nil, err
	}

	if err = os.MkdirAll(configDir, 0o755); err != nil {
		return nil, fmt.Errorf("creating config dir:  %s: %w", configDir, err)
	}

	i.File = filepath.Join(configDir, i.Name+confExt)
	if _, err = os.Stat(i.File); err == nil {
		return i.openConfigFile()
	}

	return i.defaultConfig()
}

func (i *Input) defaultConfig() (*Config, error) {
	c := i.newConfig(nil)
	_, err := c.Write(nil)

	return c, err
}

// newConfig returns a config with defaults.
func (i *Input) newConfig(settings *Settings) *Config {
	home, _ := os.UserHomeDir()
	if home == "" {
		home = "/tmp"
		if mnd.IsWindows {
			home = `C:\`
		}
	}

	if settings == nil {
		settings = &Settings{
			File: i.File,
			LogConfig: logs.LogConfig{
				Name:  i.Name,
				Path:  filepath.Join(home, "."+i.Name),
				Size:  4,
				Mode:  "0640",
				Files: 10,
				Lang:  language.English.String(),
			},
			Default: Default{
				Instance: map[string]int{
					starr.Lidarr.String():   0,
					starr.Prowlarr.String(): 0,
					starr.Radarr.String():   0,
					starr.Readarr.String():  0,
					starr.Sonarr.String():   0,
					starr.Whisparr.String(): 0,
				},
			},
			Instances: starrs.Instances{
				starr.Lidarr.String():   []starrs.AppConfig{},
				starr.Prowlarr.String(): []starrs.AppConfig{},
				starr.Radarr.String():   []starrs.AppConfig{},
				starr.Readarr.String():  []starrs.AppConfig{},
				starr.Sonarr.String():   []starrs.AppConfig{},
				starr.Whisparr.String(): []starrs.AppConfig{},
			},
			Hide:    make(map[string]bool),
			Updates: "production",
		}
	}

	config := &Config{
		file:     i.File,
		settings: settings,
	}
	config.Start()

	return config
}

// getConfigDir finds a good place in the home folder for a config file.
// If a good place in home folder is not found, then a place next to the executable is used.
func (i *Input) getConfigDir() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		configDir, err = os.Executable()
		if configDir = filepath.Dir(configDir); err != nil {
			return "", fmt.Errorf("could not find a suitable config directory: %s: %w", configDir, err)
		}
	}

	if i.Dir != "" {
		configDir = filepath.Join(configDir, i.Dir)
	}

	return configDir, nil
}

func (i *Input) openCustomPath() (*Config, error) {
	var err error
	if i.File, err = filepath.Abs(i.File); err != nil {
		return nil, fmt.Errorf("invalid config path provided: %w", err)
	}

	if _, err := os.Stat(i.File); err == nil {
		return i.openConfigFile()
	}

	// Custom config path does not exist, so create it.
	config := i.newConfig(nil)

	configDir := filepath.Dir(config.file)
	if err := os.MkdirAll(configDir, 0o755); err != nil {
		return nil, fmt.Errorf("creating config path dir:  %s: %w", configDir, err)
	}

	_, err = config.Write(nil)

	return config, err
}

func (i *Input) openConfigFile() (*Config, error) {
	cnfOpen, err := os.Open(i.File)
	if err != nil {
		return nil, fmt.Errorf("opening config file: %s: %w", i.File, err)
	}
	defer cnfOpen.Close()

	var settings Settings
	if err = gob.NewDecoder(cnfOpen).Decode(&settings); err != nil {
		return nil, fmt.Errorf("decoding config file:  %s: %w", i.File, err)
	}

	return i.newConfig(i.setDefaults(&settings)), nil
}

// setDefaults only needs to handle "new" things that get added later.
func (i *Input) setDefaults(s *Settings) *Settings { //nolint:varnamelen
	s.Name = i.Name
	s.File = i.File
	s.LogConfig.Path, _ = homedir.Expand(s.LogConfig.Path)

	if s.Updates == "" {
		s.Updates = "production"
	}

	if s.Instances == nil {
		s.Instances = starrs.Instances{}
	}

	for _, app := range []starr.App{
		starr.Lidarr, starr.Prowlarr, starr.Radarr, starr.Readarr, starr.Sonarr, starr.Whisparr,
	} {
		if s.Instances[app.String()] == nil {
			s.Instances[app.String()] = []starrs.AppConfig{}
		}
	}

	if s.Hide == nil {
		s.Hide = make(map[string]bool)
	}

	if translations.Languages(language.English.String())[s.Lang] == "" {
		s.Lang = language.English.String()
	}

	if s.Default.Instance == nil {
		s.Default.Instance = make(map[string]int)
	}

	setInstanceDefaults(s)

	return s
}

func setInstanceDefaults(s *Settings) { //nolint:varnamelen
	for app := range s.Instances {
		for idx := range s.Instances[app] {
			if s.Instance[app] == 0 || s.Instance[app] >= len(s.Instances[app]) {
				s.Instance[app] = 0
			}

			if s.Instances[app][idx].Timeout < 1 {
				s.Instances[app][idx].Timeout = time.Minute * 2
			}
		}
	}
}

// findConfigFileFolder checks the app/exe directory for a config file.
// Returns the directory if the file is found.
func (i *Input) findConfigFileFolder() {
	exe, err := os.Executable()
	if i.File != "" || i.Name == "" || err != nil {
		return
	}

	path := filepath.Dir(exe)
	if mnd.IsMac && filepath.Base(path) == "MacOS" {
		path = filepath.Dir(filepath.Dir(filepath.Dir(path)))
	}

	path = filepath.Join(path, i.Name+confExt)
	if _, err := os.Stat(path); err == nil {
		i.File = path // file exists next to executable.
	}
}
