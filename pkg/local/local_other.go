//go:build !windows && !darwin

package local

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func modifyCmd(_ *exec.Cmd) {}

// OpenCmd opens anything.
func OpenCmd(ctx context.Context, cmd ...string) error {
	return StartCmd(ctx, "xdg-open", cmd...)
}

// OpenLog opens Log Files.
func OpenLog(_ context.Context, _ string) error {
	return fmt.Errorf("%w: %s", ErrUnsupported, runtime.GOOS)
}

// OpenFile open Config Files.
func OpenFile(ctx context.Context, filePath string) error {
	return OpenCmd(ctx, "file://"+filePath)
}

// OpenFolder opens a folder in Finder.
func OpenFolder(ctx context.Context, filePath string) error {
	if stat, err := os.Stat(filePath); err != nil || !stat.IsDir() {
		filePath = filepath.Dir(filePath)
	}

	return OpenCmd(ctx, "file://"+filePath)
}

func CreateShortcut() (string, error) {
	return "Shortcuts are not supported on Linux!", nil
}
