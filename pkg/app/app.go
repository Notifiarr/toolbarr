package app

import (
	"fmt"
	"runtime"
	"time"

	"github.com/Notifiarr/toolbarr/pkg/config"
	"github.com/gorilla/schema"
	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
	"golift.io/version"
)

//nolint:gochecknoglobals // this is supposed to be global.
var decoder = schema.NewDecoder()

func (a *App) GetConfig() *config.Config {
	a.config.Trace("Call:GetConfig()")
	return a.config
}

type ConfigSaved struct {
	Msg string
	Val any
}

// SaveConfigItem saves a single item to the running config and writes the config file.
func (a *App) SaveConfigItem(name string, value any, reload bool) (*ConfigSaved, error) {
	a.config.Tracef("Call:SaveConfigItem(%s,%v,%v)", name, value, reload)
	config := a.config.Copy()

	err := decoder.Decode(config, map[string][]string{name: {fmt.Sprint(value)}})
	if err != nil {
		a.config.Errorf("Writing config: decoding '%s' value '%v' failed: %w", name, value, err)
		return nil, fmt.Errorf("decoding '%s' value '%v' failed: %w", name, value, err)
	}

	if err = config.Write(); err != nil {
		a.config.Error("Error writing config: " + err.Error())
		return nil, fmt.Errorf("writing config: %w", err)
	}

	a.config.Update(config)

	if reload {
		_ = a.config.Logger.Close()
		a.config.Logger.Setup(a.ctx, config.LogConfig)
	}

	msg := fmt.Sprintf("Saved: '%s' Value: %v", name, value)
	a.config.Print("Config " + msg)

	return &ConfigSaved{Msg: msg, Val: value}, nil
}

// PickFolder opens the folder selector.
func (a *App) PickFolder(path string) (string, error) {
	if path == "" {
		path = a.config.Path
	}

	dir, err := wailsRuntime.OpenDirectoryDialog(a.ctx, wailsRuntime.OpenDialogOptions{
		DefaultDirectory:     path,
		Title:                "Choose Folder",
		CanCreateDirectories: true,
		ShowHiddenFiles:      true,
	})
	if err != nil {
		wailsRuntime.LogError(a.ctx, err.Error())
		return "", fmt.Errorf("opening directory browser: %w", err)
	}

	return dir, nil
}

type Version struct {
	Version   string
	Revision  string
	Branch    string
	BuildUser string
	BuildDate string
	GoVersion string
	Started   string
	Running   int
}

func (a *App) Version() Version {
	return Version{
		Version:   version.Version,
		Revision:  version.Revision,
		Branch:    version.Branch,
		BuildUser: version.BuildUser,
		BuildDate: version.BuildDate,
		GoVersion: runtime.Version(),
		Started:   version.Started.Round(time.Second).String(),
		Running:   int(time.Since(version.Started).Seconds()),
	}
}
