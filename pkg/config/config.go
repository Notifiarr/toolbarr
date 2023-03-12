//nolint:gomnd
package config

import (
	"encoding/gob"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"

	"github.com/Notifiarr/toolbarr/pkg/logs"
)

// TODO: load config file from same folder as exe or app.
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
	App
	*logs.Logger
	File string
	Dark bool
}

type Advanced struct {
	DevMode bool
	Updates string
}

type App struct {
	IsWindows bool
	IsLinux   bool
	IsMac     bool
	Exe       string
	Home      string
}

// New returns a config with defaults.
func New(appName string, logger *logs.Logger) *Config {
	exec, _ := os.Executable()
	user, _ := user.Current()

	return &Config{
		LogConfig: &logs.LogConfig{
			Name:  appName,
			Path:  filepath.Join("~", "."+appName),
			Size:  4,
			Mode:  "0640",
			Files: 10,
		},
		App: App{
			IsWindows: runtime.GOOS == "windows",
			IsLinux:   runtime.GOOS == "linux",
			IsMac:     runtime.GOOS == "darwin",
			Exe:       exec,
			Home:      user.HomeDir,
		},
		Advanced: Advanced{
			Updates: "production",
		},
		Logger: logger,
	}
}

// Get opens/reads or creates/writes a config file.
func Get(input Input) (*Config, error) {
	if input.Path != "" {
		return getCustomPath(&input)
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

	configFile := filepath.Join(configDir, input.Name+".conf")
	if _, err = os.Stat(configFile); err == nil {
		return input.openConfig(configFile)
	}

	config := New(input.Name, input.Logger)
	config.File = configFile

	return config, config.Write()
}

func getCustomPath(input *Input) (*Config, error) {
	var err error
	if input.Path, err = filepath.Abs(input.Path); err != nil {
		return nil, fmt.Errorf("invalid config path provided: %w", err)
	}

	config := New(input.Name, input.Logger)
	config.File = input.Path

	if _, err := os.Stat(config.File); err == nil {
		return input.openConfig(config.File)
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

	user, _ := user.Current()
	exec, _ := os.Executable()
	config.File = configFile
	config.Logger = i.Logger
	config.Name = i.Name
	config.App = App{
		IsWindows: runtime.GOOS == "windows",
		IsLinux:   runtime.GOOS == "linux",
		IsMac:     runtime.GOOS == "darwin",
		Exe:       exec,
		Home:      user.HomeDir,
	}

	if config.Updates == "" {
		config.Updates = "production"
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
	c.App = merge.App
}
