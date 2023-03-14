import { derived, readable, writable } from 'svelte/store'
import { GetConfig } from "../../wailsjs/go/app/App"

/* Not exported because it's not kept up to date. */
const config = readable({
    IsWindows: false,
    IsMac: false,
    IsLinux: false,
    DevMode: true,
    Dark: false,
  },
  (set) => { GetConfig().then(conf => set(conf)) },
)

/* here for convenience. */
export const isWindows = derived(config, $config => $config.IsWindows)
export const isLinux = derived(config, $config => $config.IsLinux)
export const isMac = derived(config, $config => $config.IsMac)
/* writable */
export const devMode = writable(true)
export const dark = writable(false)
/* this only sets an initial value. */
config.subscribe(conf => {
  devMode.set(conf.DevMode)
  dark.set(conf.Dark)
})