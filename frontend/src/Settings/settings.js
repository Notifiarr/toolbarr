import { writable } from 'svelte/store'
import { GetConfig } from "../../wailsjs/go/app/App"

/* How do we make these readable instead of writable? */
export const isWindows = writable(false)
export const isLinux = writable(false)
export const isMac = writable(false)
/* here for convenience. */
export const devMode = writable(false)
export const dark = writable(false)

let ready

if (!ready) {
    GetConfig().then(result => {
        devMode.set(result.DevMode)
        isLinux.set(result.IsLinux)
        isMac.set(result.IsMac)
        isWindows.set(result.IsWindows)
        dark.set(result.Dark)
        ready = true
    })
}

