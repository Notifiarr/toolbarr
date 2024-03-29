package app

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"time"

	"github.com/Notifiarr/toolbarr/pkg/config"
	"github.com/Notifiarr/toolbarr/pkg/local"
	"github.com/Notifiarr/toolbarr/pkg/mnd"
	"github.com/Notifiarr/toolbarr/pkg/translations"
	"github.com/gorilla/schema"
	wr "github.com/wailsapp/wails/v2/pkg/runtime"
	"golift.io/version"
)

//nolint:gochecknoglobals // this is supposed to be global.
var decoder = schema.NewDecoder()

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

	user, err := user.Current()
	if err != nil {
		a.log.Warnf("Getting user data: %v", err.Error())
	}

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

// Ask the user a Yes/No question.
func (a *App) Ask(title, msg string) bool {
	a.log.Tracef("Call:Ask(%s,%s)", title, msg)

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
		return nil, fmt.Errorf(a.log.Translate("Writing config: decoding '%s' value '%v' error: %v", name, value, err))
	}

	if _, err = a.config.Write(config); err != nil {
		a.log.Errorf("Error writing config: %v", err.Error())
		return nil, fmt.Errorf(a.log.Translate("Error writing config: %v", err.Error()))
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

	dir, err := wr.OpenDirectoryDialog(a.ctx, wr.OpenDialogOptions{
		DefaultDirectory:     a.fixFolderPath(path),
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

// PickFile opens the file selector.
func (a *App) PickFile(path, extname, extensions string) (string, error) {
	a.log.Tracef("Call:PickFile(%s,%s,%s)", path, extname, extensions)

	dir, err := wr.OpenFileDialog(a.ctx, wr.OpenDialogOptions{
		DefaultDirectory: a.fixFolderPath(path),
		Title:            a.log.Translate("Choose File"),
		ShowHiddenFiles:  true,
		Filters: []wr.FileFilter{{
			DisplayName: extname,
			Pattern:     extensions,
		}},
	})
	if err != nil {
		wr.LogError(a.ctx, err.Error())
		return "", fmt.Errorf(a.log.Translate("Opening file browser: %v", err))
	}

	return dir, nil
}

// Languages returns a list of languages the backend supports.
func (a *App) Languages() map[string]string {
	a.log.Tracef("Call:Languages()")
	return translations.Languages(a.config.Settings().Lang)
}

// CreateShortcut makes a shortcut to the exe on Windows desktop.
func (a *App) CreateShortcut() (string, error) {
	a.log.Tracef("Call:CreateShortcut()")
	return local.CreateShortcut() //nolint:wrapcheck
}

// fixFolderPath returns an existing folder path.
func (a *App) fixFolderPath(path string) string {
	if path == "" {
		path = a.config.Settings().Path
	}

	stat, err := os.Stat(path)
	if err == nil && stat.IsDir() {
		return path
	}

	// Work up the tree until we find a folder to start in.
	for {
		path = filepath.Dir(path)

		_, err = os.Stat(path)
		if err == nil || path == "/" || (mnd.IsWindows && len(path) < 4) {
			break
		}
	}

	return path
}
