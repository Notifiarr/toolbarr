package logs

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/wailsapp/wails/v2/pkg/logger"
	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
	"golift.io/rotatorr"
	"golift.io/rotatorr/timerotator"
)

const defaultName = "app"

// Log Levels.
const (
	LogLevelNormal Level = iota
	LogLevelDebug
	LogLevelTrace
)

// Level represents the log level.
type Level int

// LogConfig allows sending logs to rotating files.
type LogConfig struct {
	Name  string
	Path  string
	Size  uint
	Mode  FileMode
	Level Level
	Files uint
}

// String returns the human-readable log level.
func (l Level) String() string {
	switch {
	case l < LogLevelNormal:
		return "Disabled"
	case l == LogLevelNormal:
		return "Normal"
	case l == LogLevelDebug:
		return "Debug"
	case l == LogLevelTrace:
		return "Trace"
	default:
		return "Unknown"
	}
}

// Logger provides some methods with baked in assumptions.
type Logger struct {
	logger  *log.Logger // Shares a Writer with InfoLog.
	rotator *rotatorr.Logger
	config  *LogConfig
	mu      sync.RWMutex
	File    string
}

// New returns an uninitialized logger.
// Config is optional, but must be provided here or with Setup().
func New() *Logger {
	return &Logger{
		logger: log.New(os.Stdout, "", log.LstdFlags),
		config: &LogConfig{},
	}
}

// SetupLogging splits log writers into a file and/or stdout.
// Config is optional, but must be provided here or with New().
func (l *Logger) Setup(ctx context.Context, config LogConfig) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.config = &config

	wailsRuntime.LogSetLogLevel(ctx, map[Level]logger.LogLevel{
		LogLevelDebug:  logger.DEBUG,
		LogLevelNormal: logger.INFO,
		LogLevelTrace:  logger.TRACE,
	}[config.Level])

	l.setLogPaths()
	l.openLogFile()
	l.Print("Opened log file: " + l.File)
}

// setLogPaths sets the log paths for app and http logs.
func (l *Logger) setLogPaths() {
	if l.config.Name == "" {
		l.config.Name = defaultName
	}

	// Regular log file.
	if l.config.Path == "" {
		l.config.Path = filepath.Join("~", "."+l.config.Name)
	}

	if f, err := homedir.Expand(l.config.Path); err == nil {
		l.config.Path = f
	} else if runtime.GOOS == "windows" {
		l.config.Path = `C:\`
	} else {
		l.config.Path = "/tmp"
	}

	if f, err := filepath.Abs(l.config.Path); err == nil {
		l.config.Path = f
	}
}

func (l *Logger) openLogFile() {
	l.File = filepath.Join(l.config.Path, l.config.Name+".log")
	l.rotator = rotatorr.NewMust(&rotatorr.Config{
		Filepath: l.File,                         // log file name.
		FileSize: int64(l.config.Size) * 1048576, //nolint:gomnd
		FileMode: l.config.Mode.Mode(),           // set file mode.
		Rotatorr: &timerotator.Layout{
			FileCount:  int(l.config.Files), // number of files to keep.
			PostRotate: l.postLogRotate,     // method to run after rotating.
		},
	})
	l.logger.SetOutput(l.rotator)
	redirectStderr(l.rotator.File)
}

// This is only for the main log. To deal with stderr.
func (l *Logger) postLogRotate(fileName, newFile string) {
	if l.rotator != nil && l.rotator.File != nil {
		redirectStderr(l.rotator.File) // Log panics.
	}

	if newFile != "" && l != nil {
		go l.Printf("Rotated log file to: %s", newFile)
	}
}

func (l *Logger) Close() error {
	l.Print("Closing Logger")
	l.logger.SetOutput(io.Discard)

	if err := l.rotator.Close(); err != nil {
		return fmt.Errorf("closing log rotator: %w", err)
	}

	return nil
}
