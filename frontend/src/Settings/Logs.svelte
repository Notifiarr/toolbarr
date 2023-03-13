<script>
  import {
    Button,
    Input,
    InputGroup,
    InputGroupText,
    Tooltip,
  } from "sveltestrap"
  import Fa from "svelte-fa"
  import {faFolderOpen} from "@fortawesome/free-solid-svg-icons"
  import {GetConfig, PickFolder} from "../../wailsjs/go/app/App.js"
  import { saveValue } from "../funcs";
  import { isWindows } from './store.js';

  let validProps = {}
  let invalidProps = {}
  let conf = {}

  GetConfig().then((result) => (conf = result))

  function saveInput(e) {
    validProps[e.target.id] = saveValue(e.target.name, e.target.value, true)
    invalidProps[e.target.id] = !validProps[e.target.id]
    setInterval(() => {validProps[e.target.id]=false}, 5000)
  }

  // This func opens a 'pick a folder' dialog and populates the Log File Path with the selected value.
  function getLogFolder(event) {
    event.preventDefault()
    PickFolder("Log File Path").then((path) => {
      if (path == "") { return }
      validProps['Path'] = saveValue('LogConfig.Path', path, true)
      invalidProps['Path'] = !validProps['Path']
      conf.Path = path
    })
  }
</script>

<p>This app writes logs file, and rotates them itself. Use the settings below to control that behavior.
Settings save automatically when changed.</p>
<InputGroup>
  <Tooltip target="Path" placement="top">Must be a directory</Tooltip>
  <InputGroupText class="setting-name">Log File Path</InputGroupText>
  <Input valid={validProps.Path} invalid={invalidProps.Path} bind:value={conf.Path} id="Path" name="LogConfig.Path" />
  <Button color="success" on:click={getLogFolder}>
    <Fa icon="{faFolderOpen}" />
    <span class="d-none d-md-inline-block">Browse</span>
  </Button>
</InputGroup>
<InputGroup>
  <InputGroupText class="setting-name">Log Level</InputGroupText>
  <Input valid={validProps.Level} invalid={invalidProps.Level} on:change={saveInput} type="select" name="LogConfig.Level" id="Level">
    <!-- option value="-1" selected={conf.Level < 0}>Logging Disabled</option -->
    <option value="0" selected={conf.Level == 0}>Normal Logging</option>
    <option value="1" selected={conf.Level == 1}>Debug Logging</option>
    <option value="2" selected={conf.Level == 2}>Trace Logging</option>
  </Input>
</InputGroup>
{#if !$isWindows}
<InputGroup>
  <InputGroupText class="setting-name">Log File Mode</InputGroupText>
  <Input valid={validProps.Mode} invalid={invalidProps.Mode} on:change={saveInput} type="select" name="LogConfig.Mode" id="Mode">
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
  <Tooltip target="Size" placement="top">Rotate files when they reach this size</Tooltip>
  <InputGroupText class="setting-name">Log File Size</InputGroupText>
  <Input valid={validProps.Size} invalid={invalidProps.Size} on:change={saveInput} type="select" name="LogConfig.Size" id="Size">
    {#each Array(20) as _, i}
      <option value={i+1} selected={conf.Size==i+1}>{i+1} Megabyte{i==0?'':'s'}</option>
    {/each}
  </Input>
</InputGroup>
<InputGroup>
  <Tooltip target="Files" placement="top">How many files to keep when rotating</Tooltip>
  <InputGroupText class="setting-name">Log Files</InputGroupText>
  <Input valid={validProps.Files} invalid={invalidProps.Files} on:change={saveInput} type="select" name="LogConfig.Files" id="Files">
    <option value="0" selected={conf.Size==0}>Disable Rotation</option>
    {#each Array(50) as _, i}
      <option value={i+1} selected={conf.Size==i+1}>{i+1} File{i==0?'':'s'}</option>
    {/each}
  </Input>
</InputGroup>

