import { derived, readable, writable } from 'svelte/store'
import { GetConfig } from "../../wailsjs/go/app/App"

/* Not exported because it's not kept up to date. */
const config = readable({
    IsWindows: false,
    IsMac: false,
    IsLinux: false,
    DevMode: false,
    Dark: false,
  },
  (set) => { GetConfig().then(conf => set(conf)) },
)

/* writable */
export const devMode = writable(false, (set) => { derived(config, $config => set($config.DevMode)) })
export const dark = writable(false, (set) => { derived(config, $config => set($config.Dark)) })
/* here for convenience. */
export const isWindows = derived(config, $config => $config.IsWindows)
export const isLinux = derived(config, $config => $config.IsLinux)
export const isMac = derived(config, $config => $config.IsMac)
