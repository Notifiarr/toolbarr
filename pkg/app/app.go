package app

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
	"time"

	"github.com/Notifiarr/toolbarr/pkg/config"
	"github.com/Notifiarr/toolbarr/pkg/mnd"
	"github.com/Notifiarr/toolbarr/pkg/translations"
	"github.com/gorilla/schema"
	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
	"golift.io/version"
)

//nolint:gochecknoglobals // this is supposed to be global.
var decoder = schema.NewDecoder()

func (a *App) GetConfig() *config.Settings {
	a.log.Tracef("Call:GetConfig()")
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
		a.log.Errorf("Writing config: decoding '%s' value '%v' error: %v", name, value, err)
		return nil, fmt.Errorf("%s %w", a.log.Translate("decoding '%s' value '%v' error:", name, value), err)
	}

	if err = a.config.Write(config); err != nil {
		a.log.Errorf("Error writing config: %v", err.Error())
		return nil, fmt.Errorf("%s %w", a.log.Translate("writing config:"), err)
	}

	if reload {
		_ = a.log.Close()
		a.log.Setup(a.ctx, a.config.Settings().LogConfig)
	}

	msg := a.log.Translate("Saved: '%s' Value: %v", name, value)
	a.log.Infof("Config %s", msg)

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
		Title:                a.log.Translate("Choose Folder"),
		CanCreateDirectories: true,
		ShowHiddenFiles:      true,
	})
	if err != nil {
		wailsRuntime.LogError(a.ctx, err.Error())
		return "", fmt.Errorf("%s %w", a.log.Translate("opening directory browser:"), err)
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
	a.log.Tracef("Call:Version()")

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

func (a *App) Languages() map[string]string {
	a.log.Tracef("Call:Languages()")
	return translations.Languages(a.config.Settings().Lang)
}
