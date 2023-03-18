//nolint:lll
//go:generate go run golang.org/x/text/cmd/gotext@latest -srclang=en-US update -out=catalog.go -lang=en-US,es-MX,da-DK,nl-AN,fr-FR,it-IT,pt-PT,sv-SE github.com/Notifiarr/toolbarr/pkg/...
//
// https://www.alexedwards.net/blog/i18n-managing-translations

package translations

import "golang.org/x/text/message"

// Languages should match the languages in the generate command above.
func Languages() []string {
	tags := message.DefaultCatalog.Languages()
	langs := make([]string, len(tags))

	for i := range tags {
		langs[i] = tags[i].String()
	}

	return langs
}
