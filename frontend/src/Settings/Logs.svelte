<script>
  export let dark
  import {
    Button,
    Input,
    InputGroup,
    InputGroupText,
    Tooltip,
  } from "sveltestrap"
  import Fa from "svelte-fa"
  import {faFolderOpen} from "@fortawesome/free-solid-svg-icons"
  import {GetConfig, IsWindows, PickFolder, SaveConfigItem} from "../../wailsjs/go/app/App.js"
  import {toasts}  from "svelte-toasts"

  const showToast = (type, msg) => {
    const toast = toasts.add({
      title: "",
      description: (type=="error"?"Error: ":"") +msg,
      duration: 7000, // 0 or negative to avoid auto-remove
      theme: dark ? "dark" : "light",
      type: type,
      onClick: () => {toast.remove()},
      showProgress: true,
    })
  }

  let validProps = {}
  let invalidProps = {}
  let conf = {}
  let isWindows = false

  function init() {
    GetConfig().then((result) => (conf = result))
    IsWindows().then((result) => (isWindows = result))
  }
  init()

  function saveInput(event) {
    if (event.target.value == "") {
      return
    }

    conf[event.target.id] = event.target.value
    SaveConfigItem(event.target.name, event.target.value, true).then((msg) => {
      showToast("success", msg)
      validProps[event.target.name] = true
      setInterval(() => {validProps[event.target.name]=false}, 5000)
    }, (error) => {
      showToast("error", error)
      invalidProps[event.target.name] = false
    })
  }
  
  function getLogFolder(event) {
    event.preventDefault()
    PickFolder("Log File Path").then(path => (conf.Path = path != "" ? path : conf.Path), e => showToast("error", e))
  }
</script>

<InputGroup >
  <Tooltip target="LogFilePath" placement="top">Must be a directory</Tooltip>
  <InputGroupText>Log File Path</InputGroupText>
  <Input readonly bind:valid={validProps.LogFilePath} bind:invalid={invalidProps.LogFilePath} on:change={saveInput} bind:value={conf.Path} placeholder="click browse to locate a log folder ->" id="Path" name="LogConfig.Path" />
  <Button color="success" on:click={getLogFolder}>
    <Fa icon="{faFolderOpen}" />
    <span class="d-none d-md-inline-block">Browse</span>
  </Button>
</InputGroup>
<InputGroup>
  <InputGroupText>Log Level</InputGroupText>
  <Input bind:valid={validProps.LogLevel} bind:invalid={invalidProps.LogFileMode} on:change={saveInput} type="select" name="LogConfig.Level" id="Level">
    <option value="-1" selected={conf.Level < 0}>Logging Disabled</option>
    <option value="0" selected={conf.Level == 0}>Normal Logging</option>
    <option value="1" selected={conf.Level == 1}>Debug Logging</option>
    <option value="2" selected={conf.Level == 2}>Trace Logging</option>
  </Input>
</InputGroup>
{#if !isWindows}
<InputGroup>
  <InputGroupText>Log File Mode</InputGroupText>
  <Input bind:valid={validProps.LogFileMode} bind:invalid={invalidProps.LogFileMode} on:change={saveInput} type="select" name="LogConfig.Mode" id="Mode">
    <option value="-1">Default (UMask)</option>
    <option value="0600" selected={conf.Mode=="0600"}>0600 (rw-------)</option>
    <option value="0640" selected={conf.Mode=="0640"}>0640 (rw-r-----)</option>
    <option value="0644" selected={conf.Mode=="0644"}>0644 (rw-r--r--)</option>
    <option value="0660" selected={conf.Mode=="0660"}>0660 (rw-rw----)</option>
    <option value="0664" selected={conf.Mode=="0664"}>0664 (rw-rw-r--)</option>
  </Input>
</InputGroup>
{/if}
<InputGroup>
  <Tooltip target="LogFileSize" placement="top">Size in megabytes</Tooltip>
  <InputGroupText>Log File Size</InputGroupText>
  <Input bind:valid={validProps.LogFileSize} bind:invalid={invalidProps.LogFileSize} on:change={saveInput} value={conf.Size} type="number" name="LogConfig.Size" id="Size" min="1" max="100" />
</InputGroup>
<InputGroup>
  <Tooltip target="LogFiles" placement="top">How many files to keep when rotating</Tooltip>
  <InputGroupText>Log Files</InputGroupText>
  <Input bind:valid={validProps.LogFiles} bind:invalid={invalidProps.LogFiles} on:change={saveInput} value={conf.Files} type="number" name="LogConfig.Files" id="Files" min="1" max="100" />
</InputGroup>

