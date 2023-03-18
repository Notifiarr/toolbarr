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
	a.log.Trace("Call:CheckUpdate()")

	if release := a.checkUpdateChecked(); release != nil {
		return release, nil
	}

	var (
		release *update.Update
		err     error
	)

	updates := a.config.Settings().Updates
	if updates == "unstable" {
		release, err = update.CheckUnstable(a.ctx, "Toolbarr", version.Revision)
	} else {
		release, err = update.CheckGitHub(a.ctx, "Notifiarr/toolbarr", version.Version)
	}

	if err != nil {
		a.log.Errorf("Checking for current %s release: %w", updates, err)
		return nil, fmt.Errorf("checking for current %s release: %w", updates, err)
	}

	a.updates.Lock()
	defer a.updates.Unlock()

	a.updates.release = &Release{
		Update: release,
		Size:   mnd.FormatBytes(release.RelSize),
	}
	a.updates.date = time.Now()
	a.log.Printf("Checked Current %s release: %v (%s)",
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
	a.log.Trace("Call:DownloadUpdate()")

	a.updates.RLock()
	defer a.updates.RUnlock()

	a.log.Printf("Downloading File")

	err := fmt.Errorf("%w: missing release, check first?", ErrInvalidInput)
	if a.updates.release == nil {
		return nil, err
	}

	updates := a.config.Settings().Updates

	a.updates.progress, err = update.DownloadURL(
		a.ctx, a.updates.release.CurrURL, path.Base(a.updates.release.CurrURL), nil)
	if err != nil {
		a.log.Errorf("Downloading %s update: %w", updates, err)
		return nil, fmt.Errorf("downloading failed: %w", err)
	}

	size := mnd.FormatBytes(a.updates.progress.Size())
	a.log.Printf("Downloading %s release from %s to %s (%s)",
		updates, a.updates.release.CurrURL, a.updates.progress.Path(), size)

	return &UpdateInfo{
		Msg:  fmt.Sprintln("Downloading", size, "to", a.updates.progress.Path()),
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
			a.log.Errorf("Opening Folder: %s: %w", path, err)
		}
	}()

	defer func() {
		time.Sleep(5 * time.Second) //nolint:gomnd
		a.Quit()
	}()

	return "Launching Installer: " + path, nil
}

func (a *App) OpenFolder(path string) string {
	a.log.Tracef("Call:OpenFolder(%s)", path)

	go func() {
		err := ui.OpenFolder(a.ctx, path)
		if err != nil {
			a.log.Errorf("Opening Folder: %s: %w", path, err)
		}
	}()

	return "Opening Path: " + path
}
