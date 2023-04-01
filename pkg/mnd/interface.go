package mnd

// App provides a simple interface to do things in the wails app.
type App interface {
	Ask(string, string) bool
}
