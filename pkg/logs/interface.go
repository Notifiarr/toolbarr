package logs

import (
	"errors"
	"fmt"
	"runtime/debug"

	"golift.io/rotatorr"
	"golift.io/version"
)

func (l *Logger) Printf(msg string, v ...any) {
	l.writeMsg(fmt.Sprintf("[INFO] "+msg, v...))
}

func (l *Logger) Print(msg string) {
	l.writeMsg("[INFO] " + msg)
}

func (l *Logger) Trace(msg string) {
	l.Tracef(msg)
}

func (l *Logger) Tracef(msg string, v ...any) {
	if l.config.Level >= LogLevelTrace {
		l.writeMsg(fmt.Sprintf("[TRACE] "+msg, v...))
	}
}

func (l *Logger) Debugf(msg string, v ...any) {
	if l.config.Level >= LogLevelDebug {
		l.writeMsg(fmt.Sprintf("[DEBUG] "+msg, v...))
	}
}

func (l *Logger) Debug(msg string) {
	l.Debugf(msg)
}

func (l *Logger) Info(msg string) {
	l.writeMsg("[INFO] " + msg)
}

func (l *Logger) Warning(msg string) {
	l.writeMsg("[WARN] " + msg)
}

func (l *Logger) Error(msg string) {
	l.writeMsg("[ERROR] " + msg)
}

func (l *Logger) Errorf(msg string, v ...any) {
	l.writeMsg(fmt.Sprintf("[ERROR] "+msg, v...))
}

func (l *Logger) Fatal(msg string) {
	l.writeMsg("[FATAL] " + msg)
}

const callDepth = 2

func (l *Logger) writeMsg(msg string) {
	if err := l.logger.Output(callDepth, msg); err != nil {
		if errors.Is(err, rotatorr.ErrWriteTooLarge) {
			l.writeSplitMsg(msg)
			return
		}

		fmt.Println("Logger Error:", err) //nolint:forbidigo
	}
}

// writeSplitMsg splits the message in half and attempts to write each half.
// If the message is still too large, it'll be split again, and the process continues until it works.
func (l *Logger) writeSplitMsg(msg string) {
	half := len(msg) / 2 //nolint:gomnd // split messages in half, recursively as needed.
	l.writeMsg(msg[:half])
	l.writeMsg("...continuing: " + msg[half:])
}

// CapturePanic can be deferred in any go routine to log any panic that occurs.
func (l *Logger) CapturePanic() {
	if r := recover(); r != nil {
		l.logger.Output(callDepth, //nolint:errcheck
			fmt.Sprintf("Go Panic! This is a bug!\n%s-%s %s %v\n%s",
				version.Version, version.Revision, version.Branch, r, string(debug.Stack())))
		panic(r)
	}
}
