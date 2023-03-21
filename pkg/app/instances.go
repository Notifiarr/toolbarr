package app

import (
	"fmt"

	"github.com/Notifiarr/toolbarr/pkg/config"
	"golift.io/starr"
)

type SavedInstance struct {
	Msg  string
	List []config.Instance
}

func (a *App) SaveInstance(idx int, instance config.Instance) (*SavedInstance, error) {
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
		return nil, fmt.Errorf("%s %w", a.log.Translate("writing config:"), err)
	}

	return &SavedInstance{
		Msg:  msg,
		List: settings.Instances[instance.App],
	}, nil
}

func (a *App) RemoveInstance(idx int, starrApp string) (*SavedInstance, error) {
	a.log.Tracef("Call:RemoveInstance(%d,%s)", idx, starrApp)

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
		return nil, fmt.Errorf("%s %w", a.log.Translate("writing config:"), err)
	}

	return &SavedInstance{Msg: "Removed instance", List: settings.Instances[starrApp]}, nil
}

func (a *App) TestInstance(instance *config.Instance) (string, error) {
	a.log.Tracef("Call:TestInstance(%s,%s)", instance.App, instance.Name)

	switch starr.App(instance.App) {
	case starr.Lidarr:
		return testLidarr(a.ctx, starrConfig(instance))
	case starr.Prowlarr:
		return testProwlarr(a.ctx, starrConfig(instance))
	case starr.Radarr:
		return testRadarr(a.ctx, starrConfig(instance))
	case starr.Readarr:
		return testReadarr(a.ctx, starrConfig(instance))
	case starr.Sonarr:
		return testSonarr(a.ctx, starrConfig(instance))
	case "Whisparr":
		return testWhisparr(a.ctx, starrConfig(instance))
	}

	return "tested! " + instance.Name, nil
}