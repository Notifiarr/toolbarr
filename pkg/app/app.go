package app

import (
	"fmt"
	"runtime"

	"github.com/Notifiarr/toolbarr/pkg/config"
	"github.com/gorilla/schema"
	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

//nolint:gochecknoglobals // this is supposed to be global.
var decoder = schema.NewDecoder()

func (a *App) GetConfig() *config.Config {
	return a.config
}

func (a *App) IsWindows() bool {
	return runtime.GOOS == "windows"
}

// SaveConfigItem saves a single item to the running config and writes the config file.
func (a *App) SaveConfigItem(name string, value string, reload bool) (string, error) {
	config := a.config.Copy()

	err := decoder.Decode(config, map[string][]string{name: {value}})
	if err != nil {
		return "", fmt.Errorf("decoding '%s' value '%s' failed: %w", name, value, err)
	}

	if err = config.Write(); err == nil && reload {
		if err = a.config.Logger.Close(); err == nil {
			a.config.Logger.Setup(a.ctx, config.LogConfig)
		}

		a.config.Update(config)
	}

	return fmt.Sprintf("Config Item '%s' saved! Value: %s", name, value), err
}

// PickFolder opens the folder selector.
func (a *App) PickFolder(id string) (string, error) {
	dir, err := wailsRuntime.OpenDirectoryDialog(a.ctx, wailsRuntime.OpenDialogOptions{
		DefaultDirectory:     a.config.Path,
		Title:                id,
		CanCreateDirectories: true,
		ShowHiddenFiles:      true,
	})
	if err != nil {
		wailsRuntime.LogError(a.ctx, err.Error())
		return "", fmt.Errorf("opening directory browser: %w", err)
	}

	return dir, nil
}
