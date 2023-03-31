package app

import (
	"fmt"

	"github.com/Notifiarr/toolbarr/pkg/starrs"
)

type SavedInstance struct {
	Msg  string
	List []starrs.Instance
}

func (a *App) SaveInstance(idx int, instance starrs.Instance) (*SavedInstance, error) {
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

func (a *App) TestInstance(instance *starrs.Instance) (string, error) {
	a.log.Tracef("Call:TestInstance(%s, %s)", instance.App, instance.Name)

	test, err := starrs.TestInstance(a.ctx, a.log, instance)
	if err != nil {
		return "", err
	}

	msg := a.log.Translate("<li>Connection test successful! Found %s (%s) with version %s.</li>",
		test.App, test.Name, test.Version)

	if test.App != instance.App {
		msg = a.log.Translate("Connection test failed! Wrong app found. Expected %s but found %s. App Version: %s",
			instance.App, test.App, test.Version)
		return "", fmt.Errorf(msg)
	}

	if instance.DBPath == "" {
		return msg, nil
	}

	if test, err = starrs.TestDBPath(a.ctx, a.log, instance); err != nil {
		return "", err
	}

	msg += " " + a.log.Translate("<li>Database file test successful! SQLite3 version: %s, tables: %d</li>",
		test.Version, test.Count)

	return msg, nil
}
