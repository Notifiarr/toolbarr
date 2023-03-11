package app

import (
	"fmt"
	"sync"
	"time"

	"github.com/Notifiarr/toolbarr/pkg/update"
	"golift.io/version"
)

var ErrInvalidInput = fmt.Errorf("invalid input provided")

type updates struct {
	sync.RWMutex
	release *update.Update
	date    time.Time
}

func (a *App) CheckUpdate() (*update.Update, error) {
	if release := a.checkUpdateChecked(); release != nil {
		return release, nil
	}

	var (
		release *update.Update
		err     error
	)

	if a.config.Updates == "unstable" {
		release, err = update.CheckUnstable(a.ctx, "Toolbarr", version.Version+"-"+version.Revision)
	} else {
		release, err = update.CheckGitHub(a.ctx, "Notifiarr/toolbarr", version.Version+"-"+version.Revision)
	}

	if err != nil {
		return nil, fmt.Errorf("checking for current %s release: %w", a.config.Updates, err)
	}

	a.updates.Lock()
	defer a.updates.Unlock()

	a.updates.release = release
	a.updates.date = time.Now()

	return release, nil
}

func (a *App) checkUpdateChecked() *update.Update {
	a.updates.RLock()
	defer a.updates.RUnlock()

	if a.updates.release != nil && time.Since(a.updates.date) < 10*time.Minute {
		return a.release
	}

	return nil
}

func (a *App) DownloadUpdate() (string, error) {
	a.updates.RLock()
	defer a.updates.RUnlock()

	if a.updates.release == nil {
		return "", fmt.Errorf("%w: missing release, check first?", ErrInvalidInput)
	}

	filePath, err := update.DownloadURL(a.ctx, a.updates.release.CurrURL, "temp.file")
	if err != nil {
		return "", fmt.Errorf("downloading update: %w", err)
	}

	return "Downloaded: " + filePath, nil
}

func (a *App) LaunchInstaller() (string, error) {
	return "", nil
}
