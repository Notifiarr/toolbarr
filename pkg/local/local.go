// Package local provides procedures to run commands and do things on the local system.
package local

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"time"
)

var ErrUnsupported = fmt.Errorf("unsupported OS")

const timeout = 15 * time.Second

// StartCmd starts a command.
func StartCmd(ctx context.Context, c string, v ...string) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, c, v...)
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard
	modifyCmd(cmd)

	return cmd.Run() //nolint:wrapcheck
}
