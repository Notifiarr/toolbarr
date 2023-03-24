//nolint:goerr113
package app

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
	"time"

	"github.com/Notifiarr/toolbarr/pkg/app/cmds"
	"github.com/Notifiarr/toolbarr/pkg/config"
	"github.com/Notifiarr/toolbarr/pkg/mnd"
	"github.com/Notifiarr/toolbarr/pkg/translations"
	"github.com/gorilla/schema"
	wr "github.com/wailsapp/wails/v2/pkg/runtime"
	"golift.io/version"
)

//nolint:gochecknoglobals // this is supposed to be global.
var decoder = schema.NewDecoder()

// Ask the user a Yes/No question.
func (a *App) Ask(title, msg string) bool {
	resp, err := wr.MessageDialog(a.ctx, wr.MessageDialogOptions{
		Type:          wr.QuestionDialog,
		Title:         title,
		Message:       msg,
		Buttons:       []string{"Yes", "No"},
		DefaultButton: "Yes",
		CancelButton:  "No",
	})
	if err != nil {
		a.log.Errorf("Dialog Failed: %v", err)
	}

	return resp == "Yes"
}

// GetConfig returns a copy of the app settings.
func (a *App) GetConfig() *config.Settings {
	a.log.Tracef("Call:GetConfig()")
	return a.config.Settings()
}

// ConfigSaved sends back the value and a message when a config item changes.
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

	if _, err = a.config.Write(config); err != nil {
		a.log.Errorf("Error writing config: %v", err.Error())
		return nil, fmt.Errorf(a.log.Translate("Writing config: %v", err))
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

	dir, err := wr.OpenDirectoryDialog(a.ctx, wr.OpenDialogOptions{
		DefaultDirectory:     path,
		Title:                a.log.Translate("Choose Folder"),
		CanCreateDirectories: true,
		ShowHiddenFiles:      true,
	})
	if err != nil {
		wr.LogError(a.ctx, err.Error())
		return "", fmt.Errorf(a.log.Translate("Opening directory browser: %v", err))
	}

	return dir, nil
}

// Version lets the frontend know who it is.
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

// Info provides additional immutable data to the front end.
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

// Version returns the app version and other immutable info.
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

// Languages returns a list of languages the backend supports.
func (a *App) Languages() map[string]string {
	a.log.Tracef("Call:Languages()")
	return translations.Languages(a.config.Settings().Lang)
}

// CreateShortcut makes a shortcut to the exe on Windows desktop.
func (a *App) CreateShortcut() (string, error) {
	a.log.Tracef("Call:CreateShortcut()")
	return cmds.CreateShortcut() //nolint:wrapcheck
}
