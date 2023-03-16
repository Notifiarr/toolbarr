import { readable, writable } from 'svelte/store'
import { GetConfig } from "../../wailsjs/go/app/App"

/* This gets all the values and only runs GetConfig a single time. */
let derivedConf = undefined
function getConfig(set) {
  if (derivedConf != undefined) {
    set(derivedConf)
    return
  }
  GetConfig().then(v => {
    derivedConf = v
    set(derivedConf)
  })
}

// All the values shown below come from one endpoint, but
// we split them here to make them more meaningful in svelte.

const defaultApp = {
  IsWindows: false,
  IsMac:     false,
  IsLinux:   false,
  Exe:       "",
}

const defaultConfig = {
  DevMode: false, // use $devMode instead.
  Dark:    false, // use $dark instead.
}

const defaultUser = {
  Home: "",
  Username: "",
}

// Use this for read-only app items.
export const app = readable(defaultApp, (set) => { getConfig(set) })
// Read only user-specific items.
export const user = readable(defaultUser, (set) => { getConfig(set) })

/* writable */
const conf = readable(defaultConfig, (set) => { getConfig(set) })
export const devMode = writable(true)
export const dark = writable(false)
/* this only sets an initial value. */
conf.subscribe(value => {
  devMode.set(value.DevMode)
  dark.set(value.Dark)
})