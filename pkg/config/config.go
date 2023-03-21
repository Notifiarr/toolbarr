//nolint:gomnd
package config

import (
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Notifiarr/toolbarr/pkg/logs"
	"github.com/Notifiarr/toolbarr/pkg/mnd"
	"github.com/Notifiarr/toolbarr/pkg/translations"
	"golang.org/x/text/language"
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
	if settings == nil {
		settings = &Settings{
			File: i.File,
			LogConfig: logs.LogConfig{
				Name:  i.Name,
				Path:  filepath.Join("~", "."+i.Name),
				Size:  4,
				Mode:  "0640",
				Files: 10,
				Lang:  language.English.String(),
			},
			Instances: make(Instances),
			Hide:      make(map[string]bool),
			Updates:   "production",
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

	if s.Updates == "" {
		s.Updates = "production"
	}

	if s.Instances == nil {
		s.Instances = make(Instances)
	}

	if s.Hide == nil {
		s.Hide = make(map[string]bool)
	}

	if translations.Languages(language.English.String())[s.Lang] == "" {
		s.Lang = language.English.String()
	}

	return s
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
