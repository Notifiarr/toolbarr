package logs

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/Notifiarr/toolbarr/pkg/mnd"
	"github.com/wailsapp/wails/v2/pkg/logger"
	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golift.io/rotatorr"
	"golift.io/rotatorr/timerotator"
)

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
	Lang  string
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
	Wails   logger.Logger // special interface for wails.
	logger  *log.Logger
	debug   *log.Logger // Shares a Writer with logger.
	trace   *log.Logger // Shares a Writer with logger.
	rotator *rotatorr.Logger
	config  *LogConfig
	mu      sync.RWMutex
	File    string
	printer *message.Printer
}

type wailsInterface struct {
	log *Logger
}

// New returns an uninitialized logger.
// Config is optional, but must be provided here or with Setup().
func New() *Logger {
	logger := &Logger{
		logger:  log.New(os.Stdout, "", log.LstdFlags),
		debug:   log.New(os.Stdout, "", log.LstdFlags),
		trace:   log.New(os.Stdout, "", log.LstdFlags),
		config:  &LogConfig{},
		printer: message.NewPrinter(language.English),
	}
	logger.Wails = &wailsInterface{log: logger}

	return logger
}

// SetupLogging splits log writers into a file and/or stdout.
// Config is optional, but must be provided here or with New().
func (l *Logger) Setup(ctx context.Context, config LogConfig) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.config = &config
	l.printer = message.NewPrinter(message.MatchLanguage(language.English.String(), config.Lang))

	wailsRuntime.LogSetLogLevel(ctx, map[Level]logger.LogLevel{
		LogLevelDebug:  logger.DEBUG,
		LogLevelNormal: logger.INFO,
		LogLevelTrace:  logger.TRACE,
	}[config.Level])
	l.setLogPaths()
	l.openLogFile()

	go l.Infof("Opened log file: %v", l.File)
}

// setLogPaths sets the log paths for app and http logs.
func (l *Logger) setLogPaths() {
	if l.config.Name == "" {
		l.config.Name = mnd.Name
	}

	// Regular log file.
	if l.config.Path == "" {
		l.config.Path = "/tmp"
		if mnd.IsWindows {
			l.config.Path = `C:\`
		}
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
	redirectStderr(l.rotator.File)
	l.logger.SetOutput(l.rotator)

	if l.config.Level >= LogLevelDebug {
		l.debug.SetOutput(l.rotator)
	}

	if l.config.Level >= LogLevelTrace {
		l.trace.SetOutput(l.rotator)
	}
}

// This is only for the main log. To deal with stderr.
func (l *Logger) postLogRotate(_, newFile string) {
	if l.rotator != nil && l.rotator.File != nil {
		redirectStderr(l.rotator.File) // Log panics.
	}

	if newFile != "" && l != nil {
		go l.Infof("Rotated log file to: %s", newFile)
	}
}

func (l *Logger) Close() error {
	l.Infof("Closing Logger")
	l.logger.SetOutput(os.Stdout)
	l.debug.SetOutput(os.Stdout)
	l.trace.SetOutput(os.Stdout)

	if err := l.rotator.Close(); err != nil {
		return fmt.Errorf("%s %w", l.Translate("closing log rotator:"), err)
	}

	return nil
}
