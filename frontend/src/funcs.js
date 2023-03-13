  import { toasts }  from "svelte-toasts"
  import { SaveConfigItem } from "../wailsjs/go/app/App"
  import { dark } from './Settings/store.js'

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

  export function saveValue(name, val, reload) {
    if (val == "") { return true }
  
    SaveConfigItem(name, val, reload).then((msg) => {
      toast("success", msg)
      return true
    }, (error) => {
      toast("error", error)
      return false
    })

    return true
  }