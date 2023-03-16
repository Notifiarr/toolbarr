import { toasts }  from "svelte-toasts"
import { conf } from "./config.js"
import { onDestroy } from "svelte"

let isDark = false;
conf.subscribe(value => (isDark = value.Dark))

export const toast = (type, msg, title="", seconds=7) => {
  const toast = toasts.add({
    title: title != "" ? title : type == "error" ? "ERROR" : "",
    description: msg,
    duration: seconds*1000,
    theme: isDark ? "dark" : "light",
    type: type,
    onClick: () => {toast.remove()},
    showProgress: true,
  })
}

// onInterval sets an interval and destroys it when it when the page changes.
export function onInterval(callback, seconds) {
	const interval = setInterval(callback, seconds*1000)
	onDestroy(() => clearInterval(interval))
}

// onOnce sets a timer and expires it after one invocation.
export function onOnce(callback, seconds) {
	const interval = setInterval(input => {
    clearInterval(interval)
    callback(input)
  }, seconds*1000);
}
