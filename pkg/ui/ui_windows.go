package ui

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

func modifyCmd(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
}

// OpenCmd opens anything.
func OpenCmd(ctx context.Context, cmd ...string) error {
	return StartCmd(ctx, "cmd", append([]string{"/c", "start"}, cmd...)...)
}

// OpenLog opens Log Files.
func OpenLog(ctx context.Context, logFile string) error {
	return OpenCmd(
		ctx, "PowerShell", "Get-Content", "-Tail", "1000", "-Wait", "-Encoding", "utf8", "-Path", "'"+logFile+"'")
}

// OpenFile open files and folders.
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
