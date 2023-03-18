//go:generate touch foo.go
//go:generate gotext -srclang=en-US update -out=catalog.go -lang=en-US,de-DE,fr-CH github.com/Notifiarr/toolbarr/pkg/...
//
// https://www.alexedwards.net/blog/i18n-managing-translations

package translations

// Languages should match the languages in the generate command above.
var Languages = []string{ //nolint:gochecknoglobals
	"en-US",
	"de-DE",
	"fr-CH",
}
