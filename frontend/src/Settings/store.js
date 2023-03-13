import { writable } from 'svelte/store'

/* We have to save devMode in a store because it's "used" in Navbar,
 * but it's "changed" in a lower/included page: settings
 */
export const devMode = writable(false)
/* dark is here for convenience. */
export const dark = writable(false)
/* I just put these here tomake them easy to deal with.
 * I'd like to figure out how to set them in this file.
 */
export const isWindows = writable(false)
export const isLinux = writable(false)
export const isMac = writable(false)