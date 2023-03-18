package app

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
	"time"

	"github.com/Notifiarr/toolbarr/pkg/config"
	"github.com/Notifiarr/toolbarr/pkg/mnd"
	"github.com/gorilla/schema"
	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
	"golift.io/version"
)

//nolint:gochecknoglobals // this is supposed to be global.
var decoder = schema.NewDecoder()

func (a *App) GetConfig() *config.Settings {
	a.log.Trace("Call:GetConfig()")
	return a.config.Settings()
}

type ConfigSaved struct {
	Msg string
	Val any
}

// SaveConfigItem saves a single item to the running config and writes the config file.
func (a *App) SaveConfigItem(name string, value any, reload bool) (*ConfigSaved, error) {
	a.log.Tracef("Call:SaveConfigItem(%s,%v,%v)", name, value, reload)
	config := a.config.Settings()

	err := decoder.Decode(config, map[string][]string{name: {fmt.Sprint(value)}})
	if err != nil {
		a.log.Errorf("Writing config: decoding '%s' value '%v' failed: %w", name, value, err)
		return nil, fmt.Errorf("decoding '%s' value '%v' failed: %w", name, value, err)
	}

	if err = a.config.Write(config); err != nil {
		a.log.Error("Error writing config: " + err.Error())
		return nil, fmt.Errorf("writing config: %w", err)
	}

	if reload {
		_ = a.log.Close()
		a.log.Setup(a.ctx, a.config.Settings().LogConfig)
	}

	msg := fmt.Sprintf("Saved: '%s' Value: %v", name, value)
	a.log.Print("Config " + msg)

	return &ConfigSaved{Msg: msg, Val: value}, nil
}

// PickFolder opens the folder selector.
func (a *App) PickFolder(path string) (string, error) {
	a.log.Tracef("Call:PickFolder(%s)", path)

	if path == "" {
		path = a.config.Settings().Path
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
	StartTime int64
	Version   string
	Revision  string
	Branch    string
	BuildUser string
	BuildDate string
	GoVersion string
	Started   string
	Info
}

type Info struct {
	IsWindows bool
	IsLinux   bool
	IsMac     bool
	Name      string
	Title     string
	Exe       string
	Home      string
	Username  string
}

func (a *App) Version() Version {
	a.log.Trace("Call:Version()")

	user, _ := user.Current()
	exec, _ := os.Executable()

	return Version{
		StartTime: version.Started.Unix(),
		Version:   version.Version,
		Revision:  version.Revision,
		Branch:    version.Branch,
		BuildUser: version.BuildUser,
		BuildDate: version.BuildDate,
		GoVersion: runtime.Version(),
		Started:   version.Started.Round(time.Second).String(),
		Info: Info{
			Title:     mnd.Title,
			Name:      mnd.Name,
			IsWindows: mnd.IsWindows,
			IsLinux:   mnd.IsLinux,
			IsMac:     mnd.IsMac,
			Exe:       exec,
			Home:      user.HomeDir,
			Username:  user.Name,
		},
	}
}
