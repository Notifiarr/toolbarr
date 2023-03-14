import { toasts }  from "svelte-toasts"
import { dark } from '../Settings/settings.js'

let isDark = false;
dark.subscribe(value => (isDark = value))

export const toast = (type, msg, title="", seconds=7) => {
  const toast = toasts.add({
    title: title != "" ? title : type == "error" ? 'ERROR' : '',
    description: msg,
    duration: seconds*1000,
    theme: isDark ? "dark" : "light",
    type: type,
    onClick: () => {toast.remove()},
    showProgress: true,
  })
}

