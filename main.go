package main

import (
	"embed"
	"flag"
	"runtime"

	"github.com/Notifiarr/toolbarr/pkg/app"
	logs "github.com/Notifiarr/toolbarr/pkg/logs"
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

	app := app.New(log, configFile)
	appMenu := menu.NewMenu()
	FileMenu := appMenu.AddSubmenu("File")
	FileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(a *menu.CallbackData) {
		app.Quit()
	})

	if runtime.GOOS == "darwin" {
		// on macos, append EditMenu to enable Cmd+C,Cmd+V,Cmd+Z.
		appMenu.Append(menu.EditMenu())
	}

	// Create application with options
	err := wails.Run(&options.App{
		Title:              "Toolbarr",
		Width:              1024,
		Height:             720,
		MinWidth:           480,
		MinHeight:          600,
		MaxWidth:           1500,
		MaxHeight:          1200,
		AssetServer:        &assetserver.Options{Assets: assets},
		OnStartup:          app.Startup,
		Menu:               appMenu,
		Bind:               []interface{}{app},
		Logger:             log,
		LogLevel:           logger.DEBUG,
		LogLevelProduction: logger.INFO,
	})
	if err != nil {
		panic("Error: " + err.Error())
	}
}
