package app

import (
	"fmt"

	"github.com/Notifiarr/toolbarr/pkg/starrs"
)

// SavedInstance is the response to the frontend after instances changes.
type SavedInstance struct {
	Msg  string
	List []starrs.AppConfig
}

// SaveInstance saves the configuration for an instance.
func (a *App) SaveInstance(idx int, instance starrs.AppConfig) (*SavedInstance, error) {
	a.log.Tracef("Call:SaveInstance(%d,%s,%s)", idx, instance.App, instance.Name)

	msg := a.log.Translate("Saved %s instance configuration!", instance.App)
	settings := a.config.Settings()

	if idx >= len(settings.Instances[instance.App]) { // new instance.
		settings.Instances[instance.App] = append(settings.Instances[instance.App], instance)
		msg = a.log.Translate("Added new %s instance!", instance.App)
	} else { // updating an instance.
		settings.Instances[instance.App][idx] = instance
	}

	settings, err := a.config.Write(settings)
	if err != nil {
		return nil, fmt.Errorf(a.log.Translate("Error writing config: %v", err.Error()))
	}

	return &SavedInstance{
		Msg:  msg,
		List: settings.Instances[instance.App],
	}, nil
}

// RemoveInstance deletes an instance from the configuration.
func (a *App) RemoveInstance(idx int, starrApp string) (*SavedInstance, error) {
	a.log.Tracef("Call:RemoveInstance(%d, %s)", idx, starrApp)

	settings := a.config.Settings()
	app := settings.Instances[starrApp]

	if idx > len(app) {
		return &SavedInstance{
			Msg:  a.log.Translate("Provided %s instance %d does not exist", starrApp, idx),
			List: app,
		}, nil
	}

	question := a.log.Translate("Really delete %s instance? Name: %s", starrApp, app[idx].Name)
	if !a.Ask(a.log.Translate("Remove Instance"), question) {
		return &SavedInstance{}, nil
	}

	settings.Instances[starrApp] = append(app[:idx], app[idx+1:]...)

	settings, err := a.config.Write(settings)
	if err != nil {
		return nil, fmt.Errorf(a.log.Translate("Error writing config: %v", err.Error()))
	}

	return &SavedInstance{
		Msg:  a.log.Translate("Removed %s Instance", starrApp),
		List: settings.Instances[starrApp],
	}, nil
}
