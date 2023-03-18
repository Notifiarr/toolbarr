package app

import (
	"context"

	"github.com/Notifiarr/toolbarr/pkg/config"
	"github.com/Notifiarr/toolbarr/pkg/logs"
	"github.com/Notifiarr/toolbarr/pkg/mnd"
	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct.
type App struct {
	ctx        context.Context
	log        *logs.Logger
	config     *config.Config
	configFile string // empty unless passed in from cli
	updates    updates
}

// New creates a new App application struct.
func New(logger *logs.Logger, configFile string) *App {
	return &App{log: logger, configFile: configFile}
}

// Startup is called when the app starts.
// The context is saved so runtime methods may be called.
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx

	defer a.log.CapturePanic()

	conf, err := config.Get(&config.Input{
		File: a.configFile,
		Name: mnd.Name,
		Dir:  "com.notifiarr." + mnd.Name,
	})
	if err != nil {
		_, _ = wailsRuntime.MessageDialog(a.ctx, wailsRuntime.MessageDialogOptions{
			Type:    "error",
			Title:   "Config Problem",
			Message: err.Error(),
		})

		a.Quit()

		return
	}

	a.config = conf
	a.log.Setup(ctx, conf.Settings().LogConfig)
}

// Quit shuts the app down.
func (a *App) Quit() {
	wailsRuntime.Quit(a.ctx)
}
