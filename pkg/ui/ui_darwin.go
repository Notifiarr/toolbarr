package ui

import (
	"context"
	"os"
	"os/exec"
)

func modifyCmd(_ *exec.Cmd) {}

// OpenCmd opens anything.
func OpenCmd(ctx context.Context, cmd ...string) error {
	return StartCmd(ctx, "open", cmd...)
}

// OpenLog opens Log Files.
func OpenLog(ctx context.Context, logFile string) error {
	return OpenCmd(ctx, "-b", "com.apple.Console", logFile)
}

// OpenFile open Config Files (with TextEdit).
func OpenFile(ctx context.Context, filePath string) error {
	return OpenCmd(ctx, "-t", filePath)
}

// OpenFolder opens a folder in Finder.
func OpenFolder(ctx context.Context, filePath string) error {
	if stat, err := os.Stat(filePath); err != nil || !stat.IsDir() {
		return OpenCmd(ctx, "-R", filePath)
	}

	return OpenCmd(ctx, filePath)
}
