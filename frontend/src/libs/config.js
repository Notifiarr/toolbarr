import { writable, readable } from "svelte/store"
import { GetConfig, Version } from "../../wailsjs/go/app/App"

// Setting a constant with an empty object prevents typecript errors downstream.
// As opposed to just passing an empty object into writable() or readable().
const eo = {}
export const conf = writable(eo, set => { GetConfig().then(v => set(v)) })
export const app = readable(eo, set => { Version().then(v => set(v)) })
