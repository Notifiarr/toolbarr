// Package mnd provides re-usable constants for the Notifiarr application packages.
package mnd

import (
	"runtime"
	"time"
)

// Application Constants.
const (
	Mode0755  = 0o755
	Mode0750  = 0o750
	Mode0600  = 0o600
	Kilobyte  = 1024
	Megabyte  = Kilobyte * Kilobyte
	KB100     = Kilobyte * 100
	OneDay    = 24 * time.Hour
	Base10    = 10
	Base8     = 8
	Bits64    = 64
	Bits32    = 32
	Disabled  = "disabled"
	Success   = "success"
	HelpLink  = "Notifiarr Discord: https://notifiarr.com/discord"
	UserRepo  = "Notifiarr/toolbarr"
	BugIssue  = "This is a bug please report it on github: https://github.com/" + UserRepo + "/issues/new"
	IsLinux   = runtime.GOOS == "linux"
	IsWindows = runtime.GOOS == "windows"
	IsMac     = runtime.GOOS == "darwin"
)

// Application Defaults.
const (
	Title       = "Toolbarr"
	DefaultName = "toolbarr"
)
