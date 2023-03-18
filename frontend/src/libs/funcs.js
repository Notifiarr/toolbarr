import { toasts }  from "svelte-toasts"
import { conf } from "./config.js"
import { onDestroy } from "svelte"

// Keep track of toasts so their theme may be kept up to date.
const sentToasts = []

let isDark = false;
conf.subscribe(value => {
  isDark = value.Dark
  // Updates the active toasts themes if dark mode changes.
  sentToasts.forEach(t => t.theme = isDark ? "dark" : "light")
})

export const toast = (type, msg, title="", seconds=7) => {
  const thisToast = toasts.add({
    title: title != "" ? title : type == "error" ? "ERROR" : "",
    description: msg,
    duration: seconds*1000,
    theme: isDark ? "dark" : "light",
    type: type,
    onClick: () => {thisToast.remove()},
    showProgress: true,
    onRemove: () => { // Remove itself from the running list.
      const index = sentToasts.indexOf(thisToast)
      if (index > -1) sentToasts.splice(index, 1)
    }
  })
  sentToasts.push(thisToast)
}

// onInterval sets an interval and destroys it when it when the page changes.
export function onInterval(callback, seconds) {
  const interval = setInterval(callback, seconds*1000)
  onDestroy(() => clearInterval(interval))
  return interval
}

// onOnce sets a timer and expires it after one invocation.
export function onOnce(callback, seconds) {
  const interval = setInterval(input => {
    clearInterval(interval)
    callback(input)
  }, seconds*1000)
  return interval
}
