//nolint:goerr113
package app

import (
	"fmt"
	"path"
	"sync"
	"time"

	"github.com/Notifiarr/toolbarr/pkg/mnd"
	"github.com/Notifiarr/toolbarr/pkg/ui"
	"github.com/Notifiarr/toolbarr/pkg/update"
	"golift.io/version"
)

var ErrInvalidInput = fmt.Errorf("invalid input provided")

type updates struct {
	sync.RWMutex
	release  *Release
	progress *update.Progress
	date     time.Time
}

type Release struct {
	*update.Update
	Size string
}

func (a *App) CheckUpdate() (*Release, error) {
	a.log.Tracef("Call:CheckUpdate()")

	if release := a.checkUpdateChecked(); release != nil {
		return release, nil
	}

	var (
		release *update.Update
		err     error
	)

	updates := a.config.Settings().Updates
	if updates == "unstable" {
		release, err = update.CheckUnstable(a.ctx, mnd.Title, version.Revision)
	} else {
		release, err = update.CheckGitHub(a.ctx, mnd.UserRepo, version.Version)
	}

	if err != nil {
		a.log.Errorf("Checking for current %s release: %v", updates, err)
		return nil, fmt.Errorf(a.log.Translate("Checking for current %s release: %s", updates, err.Error()))
	}

	a.updates.Lock()
	defer a.updates.Unlock()

	a.updates.release = &Release{
		Update: release,
		Size:   mnd.FormatBytes(release.RelSize),
	}
	a.updates.date = time.Now()
	a.log.Infof("Checked Current %s release: %v (%s)",
		updates, release.Version, mnd.FormatBytes(release.RelSize))

	return a.updates.release, nil
}

func (a *App) checkUpdateChecked() *Release {
	a.updates.RLock()
	defer a.updates.RUnlock()

	if a.updates.release != nil && time.Since(a.updates.date) < 10*time.Second {
		return a.updates.release
	}

	return nil
}

type UpdateInfo struct {
	Msg  string
	Path string
	Size string
}

func (a *App) DownloadUpdate() (*UpdateInfo, error) {
	a.log.Tracef("Call:DownloadUpdate()")

	a.updates.RLock()
	defer a.updates.RUnlock()

	a.log.Infof("Downloading File")

	err := fmt.Errorf(a.log.Translate("Missing release, check first?"))
	if a.updates.release == nil {
		return nil, err
	}

	updates := a.config.Settings().Updates

	a.updates.progress, err = update.DownloadURL(
		a.ctx, a.updates.release.CurrURL, path.Base(a.updates.release.CurrURL), nil)
	if err != nil {
		a.log.Errorf("Downloading %s update: %v", updates, err)
		return nil, fmt.Errorf(a.log.Translate("Downloading failed: %v", err))
	}

	size := mnd.FormatBytes(a.updates.progress.Size())
	a.log.Infof("Downloading %s release from %s to %s (%s)",
		updates, a.updates.release.CurrURL, a.updates.progress.Path(), size)

	return &UpdateInfo{
		Msg:  a.log.Translate("Downloading %s to %s", size, a.updates.progress.Path()),
		Path: a.updates.progress.Path(),
		Size: size,
	}, nil
}

func (a *App) LaunchInstaller(path string) (string, error) {
	a.log.Tracef("Call:LaunchInstaller(%s)", path)

	var err error

	go func() {
		if mnd.IsMac {
			err = ui.OpenFolder(a.ctx, path)
		} else {
			err = ui.OpenCmd(a.ctx, path)
		}

		if err != nil {
			a.log.Errorf("Launching installer: %s: %v", path, err)
		}
	}()

	defer func() {
		time.Sleep(5 * time.Second) //nolint:gomnd
		a.Quit()
	}()

	return a.log.Translate("Launching Installer: %s", path), nil
}

func (a *App) OpenFolder(path string) string {
	a.log.Tracef("Call:OpenFolder(%s)", path)

	go func() {
		err := ui.OpenFolder(a.ctx, path)
		if err != nil {
			a.log.Errorf("Opening Folder: %s: %v", path, err)
		}
	}()

	return a.log.Translate("Opening Path: %s", path)
}
