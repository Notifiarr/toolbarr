package app

import (
	"context"

	"github.com/Notifiarr/toolbarr/pkg/config"
	"github.com/Notifiarr/toolbarr/pkg/logs"
	"github.com/Notifiarr/toolbarr/pkg/mnd"
	"github.com/Notifiarr/toolbarr/pkg/starrs"
	"github.com/wailsapp/wails/v2/pkg/menu"
	wr "github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct.
type App struct {
	ctx     context.Context
	log     *logs.Logger
	config  *config.Config
	updates updates
	*Config
}

// Config is the input data passed in from main.go.
type Config struct {
	Logger     *logs.Logger
	ConfigFile string // empty unless passed in from cli
	AppMenu    *menu.Menu
	Starrs     *starrs.Starrs
}

// New creates a new App application struct.
func New(logger *logs.Logger, config *Config) *App {
	return &App{log: logger, Config: config}
}

// Startup is called when the app starts.
// The context is saved so runtime methods may be called.
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx

	defer a.log.CapturePanic()

	conf, err := config.Get(&config.Input{
		File: a.ConfigFile,
		Name: mnd.Name,
		Dir:  "com.notifiarr." + mnd.Name,
	})
	if err != nil {
		a.ErrorDialog(a.log.Translate("Config Problem"), err.Error())
		a.Quit()

		return
	}

	a.config = conf
	a.log.Setup(ctx, conf.Settings().LogConfig)
	starrs.Startup(ctx, a.Starrs, a.log, a)
	a.setupMenu()
}

// setupMenu configures the menu bar at the top of the application.
func (a *App) setupMenu() {
	hideMenu := a.AppMenu.AddSubmenu(a.log.Translate("Hide"))
	settings := a.config.Settings()
	hideMenuItems := []string{
		"Dark",
		"Lidarr",
		"Prowlarr",
		"Radarr",
		"Readarr",
		"Sonarr",
		"Whisparr",
		"Settings",
	}

	for _, arr := range hideMenuItems {
		hideMenu.AddCheckbox(arr, settings.Hide[arr], nil, func(data *menu.CallbackData) {
			a.toggleMenuItem(data.MenuItem.Label, data.MenuItem.Checked)
		})
	}

	wr.MenuSetApplicationMenu(a.ctx, a.AppMenu)
	a.AppMenu = hideMenu
}

// toggleMenuItem powers the app 'Hide' menu.
func (a *App) toggleMenuItem(item string, value bool) {
	settings := a.config.Settings()
	settings.Hide[item] = value
	name := "Hide." + item
	msg := a.log.Translate("Saved: '%s' Value: %v", name, value)

	settings, err := a.config.Write(settings)
	if err != nil {
		a.log.Errorf("Error writing config: %v", err.Error())
		a.ErrorDialog(a.log.Translate("Config Problem"), err.Error())
	} else {
		a.log.Infof("Config %s", msg)
		wr.EventsEmit(a.ctx, "configChanged", map[string]any{
			"Settings": settings,
			"Msg":      msg,
		})
	}
}

// Quit shuts the app down.
func (a *App) Quit() {
	wr.Quit(a.ctx)
}

// ErrorDialog displays an error on screen.
func (a *App) ErrorDialog(title, msg string) {
	_, _ = wr.MessageDialog(a.ctx, wr.MessageDialogOptions{
		Type:    wr.ErrorDialog,
		Title:   title,
		Message: msg,
	})
}
