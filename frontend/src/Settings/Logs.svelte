<script lang="ts">
  import { Button, InputGroup } from "sveltestrap"
  import Fa from "svelte-fa"
  import { faFolderOpen } from "@fortawesome/free-solid-svg-icons"
  import { PickFolder } from "/wailsjs/go/app/App"
  import { app, conf } from "/src/libs/config"
  import ConfigInput from "/src/libs/Input.svelte"
  import T, { _ } from "/src/libs/Translate.svelte"

  let confPath
  // This func opens a "pick a folder" dialog and populates the Log File Path with the current config file folder.
  function getLogFolder(event) {
    event.preventDefault()
    PickFolder("").then(path => {if (path != "") confPath.update(path)})
  }
</script>

<p>{$_("logs_application_settings")}</p>
<InputGroup>
  <ConfigInput readonly bind:this={confPath} type="text" id="Path" />
  <Button color="success" on:click={getLogFolder}>
    <Fa icon="{faFolderOpen}" />
    <span class="d-none d-md-inline-block">Browse</span>
  </Button>
</InputGroup>
<InputGroup>
  <ConfigInput type="select" id="Level">
    <option value={0}>{$_("configvalues.NormalLogging")}</option>
    <option value={1}>{$_("configvalues.DebugLogging")}</option>
    {#if $conf.DevMode}
      <option value="dev mode is cool">make an error</option>
    {/if}
    <option value={2}>{$_("configvalues.TraceLogging")}</option>
  </ConfigInput>
</InputGroup>
{#if !$app.IsWindows}
<InputGroup>
  <ConfigInput type="select" id="Mode">
    <option value="0600">0600 (rw-------)</option>
    <option value="0640">0640 (rw-r-----)</option>
    <option value="0644">0644 (rw-r--r--)</option>
    <option value="0660">0660 (rw-rw----)</option>
    <option value="0664">0664 (rw-rw-r--)</option>
  </ConfigInput>
</InputGroup>
{/if}
<InputGroup>
  <ConfigInput type="select" id="Size">
    {#each Array(20) as u, i}
      <option value={i+1}><T id="configvalues.Megabyte" count={i+1}/></option>
    {/each}
  </ConfigInput>
</InputGroup>
<InputGroup>
  <ConfigInput type="select" id="Files">
    <option value={0}>{$_("configvalues.DisableRotation")}</option>
    {#each Array(50) as u, i}
      <option value={i+1}><T id="configvalues.File" count={i+1}/></option>
    {/each}
  </ConfigInput>
</InputGroup>
