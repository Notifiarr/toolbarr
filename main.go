package main

import (
	"embed"
	"flag"

	"github.com/Notifiarr/toolbarr/pkg/app"
	logs "github.com/Notifiarr/toolbarr/pkg/logs"
	"github.com/Notifiarr/toolbarr/pkg/mnd"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

//nolint:gomnd
func main() {
	// Create an instance of the app structure
	log := logs.New()
	defer log.CapturePanic()

	var configFile string

	flag.StringVar(&configFile, "c", "", "Config file path. Determined automatically if not provided.")
	flag.Parse()

	appMenu := menu.NewMenu()
	fileMenu := appMenu.AddSubmenu("File")
	appMenu.Append(menu.EditMenu())

	app := app.New(log, configFile)

	if mnd.IsWindows {
		fileMenu.AddText("Exit", keys.OptionOrAlt("f4"), func(a *menu.CallbackData) { app.Quit() })
	} else {
		fileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(a *menu.CallbackData) { app.Quit() })
	}

	// Create application with options
	err := wails.Run(&options.App{
		Title:              mnd.Title,
		Width:              1024,
		Height:             720,
		MinWidth:           480,
		MinHeight:          600,
		MaxWidth:           1500,
		Fullscreen:         false,
		MaxHeight:          1200,
		AssetServer:        &assetserver.Options{Assets: assets},
		OnStartup:          app.Startup,
		Menu:               appMenu,
		Bind:               []interface{}{app},
		Logger:             log.Wails,
		LogLevel:           logger.DEBUG,
		LogLevelProduction: logger.INFO,
	})
	if err != nil {
		panic("Error: " + err.Error())
	}
}
