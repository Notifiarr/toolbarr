package app

import (
	"context"

	"github.com/Notifiarr/toolbarr/pkg/config"
	"github.com/Notifiarr/toolbarr/pkg/logs"
	"github.com/Notifiarr/toolbarr/pkg/mnd"
	"github.com/wailsapp/wails/v2/pkg/menu"
	wr "github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct.
type App struct {
	ctx        context.Context
	log        *logs.Logger
	config     *config.Config
	configFile string // empty unless passed in from cli
	updates    updates
	menu       *menu.Menu
}

// New creates a new App application struct.
func New(logger *logs.Logger, configFile string, appMenu *menu.Menu) *App {
	return &App{log: logger, configFile: configFile, menu: appMenu}
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
		a.ErrorDialog(a.log.Translate("Config Problem"), err.Error())
		a.Quit()

		return
	}

	a.config = conf
	a.log.Setup(ctx, conf.Settings().LogConfig)
	a.setupMenu()
}

func (a *App) setupMenu() {
	hideMenu := a.menu.AddSubmenu(a.log.Translate("Hide"))
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

	wr.MenuSetApplicationMenu(a.ctx, a.menu)
	a.menu = hideMenu
}

// toggleMenuItem powers the app 'Hide' menu.
func (a *App) toggleMenuItem(item string, checked bool) {
	settings := a.config.Settings()
	settings.Hide[item] = checked
	name := "Hide." + item
	msg := a.log.Translate("Saved: '%s' Value: %v", name, checked)

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

func (a *App) ErrorDialog(title, msg string) {
	_, _ = wr.MessageDialog(a.ctx, wr.MessageDialogOptions{
		Type:    wr.ErrorDialog,
		Title:   title,
		Message: msg,
	})
}
