package config

import (
	"encoding/gob"
	"fmt"
	"os"

	"github.com/Notifiarr/toolbarr/pkg/logs"
)

// Settings is the data read and written to/from the Settings file.
// Avoid pointers and complex types.
type Settings struct {
	logs.LogConfig
	Language string
	TimeZone string
	Dark     bool
	DevMode  bool
	Updates  string
	File     string // should not be changed.
}

func (c *Config) Stop() {
	if c.ask != nil {
		close(c.ask)
	}
}

func (c *Config) Start() {
	if c.ask != nil {
		return
	}

	c.ask = make(chan *Settings)
	c.rep = make(chan *Settings)

	go c.watch()
}

func (c *Config) Settings() *Settings {
	c.ask <- nil
	return <-c.rep
}

// Update merges in a new config.
func (c *Config) Update(settings *Settings) *Settings {
	c.ask <- settings
	return <-c.rep
}

// Write writes the config file and updates the running settings.
func (c *Config) Write(settings *Settings) error {
	if settings == nil {
		settings = c.Settings()
	}

	settings.File = c.file

	cnfOpen, err := os.Create(c.file)
	if err != nil {
		return fmt.Errorf("creating config file:  %s: %w", c.file, err)
	}
	defer cnfOpen.Close()

	if err = gob.NewEncoder(cnfOpen).Encode(settings); err != nil {
		return fmt.Errorf("encoding config file:  %s: %w", c.file, err)
	}

	c.Update(settings)

	return nil
}

// watch for updates and settings requests.
// runs in a go routine and holds the running config.
func (c *Config) watch() {
	defer func() {
		close(c.rep)
		c.ask = nil
	}()

	for q := range c.ask {
		if q != nil {
			q.File, c.settings = c.file, q
		}

		settings := *c.settings
		c.rep <- &settings
	}
}
