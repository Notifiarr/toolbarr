package cmds

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"syscall"

	"github.com/Notifiarr/toolbarr/pkg/mnd"
)

const createShortcut = `$WshShell = New-Object -comObject WScript.Shell
$Shortcut = $WshShell.CreateShortcut("%s")
$Shortcut.TargetPath = "%s"
$Shortcut.Save()`

type PowerShell struct {
	powerShell string
}

// newPowerShell creates a new session for powershell use.
func newPowerShell() (*PowerShell, error) {
	shell, err := exec.LookPath("powershell.exe")
	if err != nil {
		return nil, fmt.Errorf("cannot locate powershell.exe: %w", err)
	}

	return &PowerShell{powerShell: shell}, nil
}

func (p *PowerShell) execute(args ...string) (string, error) {
	args = append([]string{"-NoProfile", "-NonInteractive"}, args...)
	cmd := exec.Command(p.powerShell, args...) //nolint:gosec
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output

	return output.String(), cmd.Run()
}

func CreateShortcut() (string, error) {
	shell, err := newPowerShell()
	if err != nil {
		return "", err
	}

	user, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("cannot locate Desktop: %w", err)
	}

	exePath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("cannot locate %s: %w", mnd.Title, err)
	}

	linkPath := filepath.Join(user.HomeDir, "Desktop", mnd.Title+".lnk")
	if _, err := os.Stat(linkPath); err == nil {
		return "Shortcut already exists!", nil
	}

	shortcut := fmt.Sprintf(createShortcut, linkPath, exePath)
	if output, err := shell.execute(shortcut); err != nil {
		return "", fmt.Errorf("shortcut creation failed: %w: %s", err, output)
	}

	return "Created shortcut!", nil
}
