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
)

const confExt = ".conf"

var ErrEmptyInput = fmt.Errorf("input must have at least name or path")

// Input data to open a config file.
// If Dir!="" then config is placed in a sub directory.
type Input struct {
	Path   string
	Name   string
	Dir    string
	Logger *logs.Logger
}

// Config is the data read and written to/from the config file.
type Config struct {
	*logs.LogConfig
	Advanced
	General
	*logs.Logger
	File string
	Dark bool
}

// Advanced settings. Mostly affects front end only.
type Advanced struct {
	DevMode bool
	Updates string
}

// General application settings.
type General struct {
	Language string
	TimeZone string
}

// New returns a config with defaults.
func New(appName string, logger *logs.Logger) *Config {
	return &Config{
		LogConfig: &logs.LogConfig{
			Name:  appName,
			Path:  filepath.Join("~", "."+appName),
			Size:  4,
			Mode:  "0640",
			Files: 10,
		},
		Advanced: Advanced{
			Updates: "production",
		},
		General: General{
			Language: "en-us",
			TimeZone: time.Local.String(),
		},
		Logger: logger,
	}
}

// Get opens/reads or creates/writes a config file.
func Get(input *Input) (*Config, error) {
	input.findConfigFileFolder()

	if input.Path != "" {
		return input.getCustomPath()
	} else if input.Name == "" {
		return nil, ErrEmptyInput
	}

	configDir, err := os.UserConfigDir()
	if err != nil {
		configDir, err = os.Executable()
		if configDir = filepath.Dir(configDir); err != nil {
			return nil, fmt.Errorf("could not find a suitable config directory:  %s: %w", configDir, err)
		}
	}

	if input.Dir != "" {
		configDir = filepath.Join(configDir, input.Dir)
	}

	if err = os.MkdirAll(configDir, 0o755); err != nil {
		return nil, fmt.Errorf("creating config dir:  %s: %w", configDir, err)
	}

	configFile := filepath.Join(configDir, input.Name+confExt)
	if _, err = os.Stat(configFile); err == nil {
		return input.openConfig(configFile)
	}

	config := New(input.Name, input.Logger)
	config.File = configFile

	return config, config.Write()
}

func (i *Input) getCustomPath() (*Config, error) {
	var err error
	if i.Path, err = filepath.Abs(i.Path); err != nil {
		return nil, fmt.Errorf("invalid config path provided: %w", err)
	}

	config := New(i.Name, i.Logger)
	config.File = i.Path

	if _, err := os.Stat(config.File); err == nil {
		return i.openConfig(config.File)
	}

	configDir := filepath.Dir(config.File)
	if err := os.MkdirAll(configDir, 0o755); err != nil {
		return nil, fmt.Errorf("creating config path dir:  %s: %w", configDir, err)
	}

	return config, config.Write()
}

func (i *Input) openConfig(configFile string) (*Config, error) {
	cnfOpen, err := os.Open(configFile)
	if err != nil {
		return nil, fmt.Errorf("opening config file: %s: %w", configFile, err)
	}
	defer cnfOpen.Close()

	var config Config
	if err = gob.NewDecoder(cnfOpen).Decode(&config); err != nil {
		return nil, fmt.Errorf("decoding config file:  %s: %w", configFile, err)
	}

	config.File = configFile
	config.Logger = i.Logger
	config.Name = i.Name

	if config.Updates == "" {
		config.Updates = "production"
	}

	if config.General.Language == "" {
		config.General.Language = "en-us"
	}

	if config.General.TimeZone == "" {
		config.General.TimeZone, _ = time.Now().Zone()
	}

	return &config, nil
}

// Write writes the config file.
func (c *Config) Write() error {
	cnfOpen, err := os.Create(c.File)
	if err != nil {
		return fmt.Errorf("creating config file:  %s: %w", c.File, err)
	}
	defer cnfOpen.Close()

	if err = gob.NewEncoder(cnfOpen).Encode(c); err != nil {
		return fmt.Errorf("encoding config file:  %s: %w", c.File, err)
	}

	return nil
}

// Copy returns a copy of the config data. Useful for updating the config file.
func (c *Config) Copy() *Config {
	newConfig := *c
	newConfig.LogConfig = &logs.LogConfig{
		Name:  c.Name,
		Path:  c.Path,
		Size:  c.Size,
		Mode:  c.Mode,
		Level: c.Level,
		Files: c.Files,
	}

	return &newConfig
}

// Update merges in a new config.
func (c *Config) Update(merge *Config) {
	// All config structs must be added here.
	c.LogConfig = merge.LogConfig
	c.Advanced = merge.Advanced
	c.Dark = merge.Dark
	c.File = merge.File
}

// findConfigFileFolder checks the app/exe directory for a config file.
// Returns the directory if the file is found.
func (i *Input) findConfigFileFolder() {
	exe, err := os.Executable()
	if i.Path != "" || i.Name == "" || err != nil {
		return
	}

	path := filepath.Dir(exe)
	if mnd.IsMac && filepath.Base(path) == "MacOS" {
		path = filepath.Dir(filepath.Dir(filepath.Dir(path)))
	}

	path = filepath.Join(path, i.Name+confExt)
	if _, err := os.Stat(path); err == nil {
		i.Path = path // file exists next to executable.
	}
}
