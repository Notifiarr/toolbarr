package mnd

// App provides a simple interface to do things in the wails app.
type App interface {
	Ask(title string, msg string) bool
}
