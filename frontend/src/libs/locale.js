import { conf } from "./config"
import { derived} from "svelte/store"
import { init, register, locale, _ } from "svelte-i18n"

/*
https://phrase.com/blog/posts/a-step-by-step-guide-to-svelte-localization-with-svelte-i18n-v3/
https://phrase.com/blog/posts/how-to-localize-a-svelte-app-with-svelte-i18n/
https://lokalise.com/blog/svelte-i18n/
*/

export { _ }
export const isReady = derived(locale, $locale => typeof $locale === 'string');

// Start with English.
const initialLocale = "en"
register(initialLocale, () => import(`./locale/${initialLocale}.json`))
init({fallbackLocale: initialLocale, initialLocale})

// Keep it up to date in case the user changes the conf.
conf.subscribe(val => {
  register(val.Lang, () => import(`./locale/${val.Lang}.json`))
  locale.set(val.Lang)
})
