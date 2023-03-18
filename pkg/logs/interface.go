package logs

import (
	"errors"
	"fmt"
	"log"
	"runtime/debug"

	"golang.org/x/text/message"
	"golift.io/rotatorr"
	"golift.io/version"
)

// Translate is like Sprintf.
func (l *Logger) Translate(msg message.Reference, v ...any) string {
	// Locking allows changing the translation language on the fly.
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.printer.Sprintf(msg, v...)
}

func (l *Logger) Errorf(msg message.Reference, v ...any) {
	l.writeMsg("["+l.Translate("ERROR")+"] "+l.Translate(msg, v...), l.logger)
}

func (l *Logger) Warnf(msg message.Reference, v ...any) {
	l.writeMsg("["+l.Translate("WARN")+"] "+l.Translate(msg, v...), l.logger)
}

func (l *Logger) Infof(msg message.Reference, v ...any) {
	l.writeMsg("["+l.Translate("INFO")+"] "+l.Translate(msg, v...), l.logger)
}

func (l *Logger) Debugf(msg message.Reference, v ...any) {
	l.writeMsg("["+l.Translate("DEBUG")+"] "+l.Translate(msg, v...), l.debug)
}

func (l *Logger) Tracef(msg string, v ...any) {
	l.writeMsg("["+l.Translate("TRACE")+"] "+fmt.Sprintf(msg, v...), l.trace)
}

func (l *wailsInterface) Trace(msg string) {
	l.log.writeMsg("["+l.log.Translate("TRACE")+"] "+msg, l.log.trace)
}

func (l *wailsInterface) Debug(msg string) {
	l.log.writeMsg("["+l.log.Translate("DEBUG")+"] "+msg, l.log.debug)
}

func (l *wailsInterface) Print(msg string) {
	l.log.writeMsg("["+l.log.Translate("INFO")+"] "+msg, l.log.logger)
}

func (l *wailsInterface) Info(msg string) {
	l.Print(msg)
}

func (l *wailsInterface) Warning(msg string) {
	l.log.writeMsg("["+l.log.Translate("WARN")+"] "+msg, l.log.logger)
}

func (l *wailsInterface) Error(msg string) {
	l.log.writeMsg("["+l.log.Translate("ERROR")+"] "+msg, l.log.logger)
}

func (l *wailsInterface) Fatal(msg string) {
	l.log.writeMsg("["+l.log.Translate("FATAL")+"] "+msg, l.log.logger)
}

const callDepth = 2

func (l *Logger) writeMsg(msg string, logger *log.Logger) {
	if err := logger.Output(callDepth, msg); err != nil {
		if errors.Is(err, rotatorr.ErrWriteTooLarge) {
			l.writeSplitMsg(msg, logger)
			return
		}

		l.printer.Printf("Logger Error: %v", err)
	}
}

// writeSplitMsg splits the message in half and attempts to write each half.
// If the message is still too large, it'll be split again, and the process continues until it works.
func (l *Logger) writeSplitMsg(msg string, logger *log.Logger) {
	half := len(msg) / 2 //nolint:gomnd // split messages in half, recursively as needed.
	l.writeMsg(msg[:half], logger)
	l.writeMsg(l.Translate("...continuing: ")+msg[half:], logger)
}

// CapturePanic can be deferred in any go routine to log any panic that occurs.
func (l *Logger) CapturePanic() {
	if r := recover(); r != nil {
		l.logger.Output(callDepth, //nolint:errcheck
			l.Translate("Go Panic! This is a bug!\n%s-%s %s %v\n%s",
				version.Version, version.Revision, version.Branch, r, string(debug.Stack())))
		panic(r)
	}
}
