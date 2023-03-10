<script>
  import {Input,InputGroup,InputGroupText,Tooltip} from "sveltestrap"
  import {SaveConfigItem} from "../../wailsjs/go/app/App.js"
  import {toasts}  from "svelte-toasts"
	import {devMode,dark} from './store.js'

  let validProps = {}
  let invalidProps = {}
  $: dm = $devMode + ""

  const showToast = (type, msg) => {
    const toast = toasts.add({
      title: "",
      description: (type=="error"?"Error: ":"") +msg,
      duration: 7000, // 0 or negative to avoid auto-remove
      theme: $dark ? "dark" : "light",
      type: type,
      onClick: () => {toast.remove()},
      showProgress: true,
    })
  }

  function saveValue(name, id, val) {
    if (val == "") { return }
    SaveConfigItem(name, val, false).then((msg) => {
      showToast("success", msg)
      validProps[id] = true
      setInterval(() => {validProps[id]=false}, 5000)
    }, (error) => {
      showToast("error", error)
      invalidProps[id] = false
    })
  }

  function saveInput(event) {
    if (event.target.id == 'DevMode') {
      $devMode = event.target.value == "false" ? false : true
    }
    saveValue(event.target.name, event.target.id, event.target.value)
  }
</script>

<InputGroup>
  <InputGroupText class="setting-name">Dev Mode</InputGroupText>
  <Input valid={validProps.DevMode} bind:invalid={invalidProps.DevMode} on:change={saveInput} value={dm} type="select" name="Advanced.DevMode" id="DevMode">
    <option value="true">Enabled</option>
    <option value="false">Disabled</option>
  </Input>
  <Tooltip target="DevMode" placement="top">Enable this when a developer instructs you to do so.</Tooltip>
</InputGroup>
