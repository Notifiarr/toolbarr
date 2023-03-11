  import {toasts}  from "svelte-toasts"
  import {SaveConfigItem} from "../wailsjs/go/app/App"

  export const toast = (type, title, msg, dark, seconds=7) => {
    const toast = toasts.add({
      title: title,
      description: (type=="error"?"Error: ":"") +msg,
      duration: seconds*1000,
      theme: dark ? "dark" : "light",
      type: type,
      onClick: () => {toast.remove()},
      showProgress: true,
    })
  }

  export function saveValue(name, val, dark, reload) {
    if (val == "") { return true }
  
    SaveConfigItem(name, val, reload).then((msg) => {
      toast("success", "", msg, dark)
      return true
    }, (error) => {
      toast("error", "", error, dark)
      return false
    })

    return true
  }