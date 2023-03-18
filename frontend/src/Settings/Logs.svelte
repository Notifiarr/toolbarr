<script>
  import { Button, InputGroup, InputGroupText } from "sveltestrap"
  import Fa from "svelte-fa"
  import { faFolderOpen } from "@fortawesome/free-solid-svg-icons"
  import { PickFolder } from "../../wailsjs/go/app/App.js"
  import { conf } from "../libs/config.js"
  import ConfigInput from "../libs/Input.svelte"

  let confPath
  // This func opens a "pick a folder" dialog and populates the Log File Path with the current config file folder.
  function getLogFolder(event) {
    event.preventDefault()
    PickFolder("").then(path => {if (path != "") confPath.update(path)})
  }
</script>

<p>This app writes logs file, and rotates them itself. Use the settings below to control that behavior.
Settings save automatically when changed.</p>
<InputGroup>
  <InputGroupText class="setting-name">Log File Path</InputGroupText>
  <ConfigInput readonly bind:this={confPath} type="text" id="Path" tooltip="Must be a directory" />
  <Button color="success" on:click={getLogFolder}>
    <Fa icon="{faFolderOpen}" />
    <span class="d-none d-md-inline-block">Browse</span>
  </Button>
</InputGroup>
<InputGroup>
  <InputGroupText class="setting-name">Log Level</InputGroupText>
  <ConfigInput type="select" id="Level">
    <option value={0}>Normal Logging</option>
    <option value={1}>Debug Logging</option>
    {#if $conf.DevMode}
      <option value="dev mode is cool">make an error</option>
    {/if}
    <option value={2}>Trace Logging</option>
  </ConfigInput>
</InputGroup>
{#if !$conf.IsWindows}
<InputGroup>
  <InputGroupText class="setting-name">Log File Mode</InputGroupText>
  <ConfigInput type="select" id="Mode">
    <option value="-1">Default (UMask)</option>
    <option value="0600">0600 (rw-------)</option>
    <option value="0640">0640 (rw-r-----)</option>
    <option value="0644">0644 (rw-r--r--)</option>
    <option value="0660">0660 (rw-rw----)</option>
    <option value="0664">0664 (rw-rw-r--)</option>
  </ConfigInput>
</InputGroup>
{/if}
<InputGroup>
  <InputGroupText class="setting-name">Log File Size</InputGroupText>
  <ConfigInput type="select" id="Size" tooltip="Rotate log file when it reaches this size">
    {#each Array(20) as _, i}
      <option value={i+1}>{i+1} Megabyte{i==0?"":"s"}</option>
    {/each}
  </ConfigInput>
</InputGroup>
<InputGroup>
  <InputGroupText class="setting-name">Log Files</InputGroupText>
  <ConfigInput type="select" id="Files" tooltip="How many backup files to keep when rotating">
    <option value={0}>Disable Rotation</option>
    {#each Array(50) as _, i}
      <option value={i+1}>{i+1} File{i==0?"":"s"}</option>
    {/each}
  </ConfigInput>
</InputGroup>
