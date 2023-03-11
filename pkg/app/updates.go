package app

import (
	"fmt"
	"path"
	"sync"
	"time"

	"github.com/Notifiarr/toolbarr/pkg/update"
	"golift.io/version"
)

var ErrInvalidInput = fmt.Errorf("invalid input provided")

type updates struct {
	sync.RWMutex
	release  *update.Update
	download string
	date     time.Time
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
		release, err = update.CheckGitHub(a.ctx, "Notifiarr/toolbarr", version.Version)
	}

	if err != nil {
		a.config.Errorf("Checking for current %s release: %w", a.config.Updates, err)
		return nil, fmt.Errorf("checking for current %s release: %w", a.config.Updates, err)
	}

	a.updates.Lock()
	defer a.updates.Unlock()

	a.updates.release = release
	a.updates.date = time.Now()
	a.config.Printf("Checked Current %s release: %s", a.config.Updates, release.Current)

	return release, nil
}

func (a *App) checkUpdateChecked() *update.Update {
	a.updates.RLock()
	defer a.updates.RUnlock()

	if a.updates.release != nil && time.Since(a.updates.date) < 10*time.Second {
		return a.release
	}

	return nil
}

func (a *App) DownloadUpdate() (string, error) {
	a.updates.RLock()
	defer a.updates.RUnlock()

	a.config.Printf("Downloading File")
	err := fmt.Errorf("%w: missing release, check first?", ErrInvalidInput)
	if a.updates.release == nil {
		return "", err
	}

	a.updates.download, err = update.DownloadURL(a.ctx, a.updates.release.CurrURL, path.Base(a.updates.release.CurrURL), nil)
	if err != nil {
		a.config.Errorf("Downloading %s update: %w", a.config.Updates, err)
		return "", fmt.Errorf("downloading update: %w", err)
	}

	a.config.Printf("Downloaded %s release from %s to %s", a.config.Updates, a.updates.release.CurrURL, a.updates.download)

	return "Downloaded: " + a.updates.download, nil
}

func (a *App) LaunchInstaller() (string, error) {
	return "", fmt.Errorf("cannot launch an installer yet!")
}
