//nolint:lll
//go:generate go run golang.org/x/text/cmd/gotext@latest -srclang=en update -out=catalog.go -lang=en,es,da,nl,fr,it,pt,sv github.com/Notifiarr/toolbarr/pkg/...
//
// https://www.alexedwards.net/blog/i18n-managing-translations

package translations

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
	"golang.org/x/text/message"
)

// Languages should match the languages in the generate command above.
func Languages(current string) map[string]string {
	langs := make(map[string]string)
	curTag := language.MustParse(current)

	for _, t := range message.DefaultCatalog.Languages() {
		cur := display.Languages(curTag)
		title := cases.Title(curTag)
		selfTitle := cases.Title(t)
		langs[t.String()] = fmt.Sprintf("%s (%s)", title.String(cur.Name(t)), selfTitle.String(display.Self.Name(t)))
	}

	return langs
}
