<script>
  export let instance
  export let info
  
  import { Alert, Badge, Button, Card, FormGroup, Input, InputGroup, Modal, Progress, Spinner, Table } from "sveltestrap"
  import { PickFolder, UpdateDBRootFolder, DeleteDBRootFolder } from "../../../../wailsjs/go/app/App.js"
  import T, { _ } from "../../../libs/Translate.svelte"
  import { faFolderOpen } from "@fortawesome/free-solid-svg-icons"
  import Fa from "svelte-fa"
  import { app } from "../../../libs/config.js"
  import { EventsOff, EventsOn } from "../../../../wailsjs/runtime/runtime.js"
  import { onOnce } from "../../../libs/funcs.js"

  let changeMe = ""
  let changing = ""
  let updating = false
  let progress
  let totals 
  let msg = ""
  let msgType

  function pickFolder() {
    PickFolder($app.Home).then(
      // PickFolder does not appaned a slash, and we need one here, so add it.
      path => { if (path != "") changeMe = path + ($app.IsWindows?"\\":"/") },
    )
  }

  function updateRoot() {
    msg = ""
    updating = true
    progress = {}
    totals = {}

    EventsOn("DBitemTotals", (data) => {
      totals = data
      progress[info.Table.Items] = 0
      progress[info.Table.Files] = 0
    })
    EventsOn("DBfileCount", (data) => {Object.keys(data).map((key) => {progress[key] = data[key]})})

    UpdateDBRootFolder(instance, changing, changeMe).then(
      resp => {
        msgType = "success"
        msg = resp.Msg
        info = resp.Info
        EventsOff("DBitemTotals", "DBfileCount")
        onOnce(() => {
          progress = undefined
          updating = false
          changeMe = ""
        }, 1)
      },
      err => {
        EventsOff("DBitemTotals", "DBfileCount")
        msgType = "danger"
        msg = err
        updating = false
        changeMe = ""
        progress = undefined
      },
    )
  }

  function deleteFolder(rf) {
    DeleteDBRootFolder(instance, rf).then(
      resp => {
        if (resp.Msg != "") {
          msgType = "success"
          msg = resp.Msg
          info = resp.Info
        }
      },
      err => {
        msgType = "danger"
        msg = err
      },
    )
  }
</script>

<li>{info.Table.Items}: {Object.keys(info.Items).length}</li>
<li>{$_("words.Files")}: {Object.keys(info.Files).length}</li>
<!-- li>
  Recycle Bin: {info.Recycle}
  <Badge color="primary"><a href="#top" on:click={() => {}}>{$_("words.Change")}</a></Badge>
</li -->
<li>
  Invalid paths: {Object.keys(info.InvalidI).length} {info.Table.Items}, 
  {Object.keys(info.InvalidF).length} {info.Table.Files}
</li>
<Card body color="primary">
  <Alert fade isOpen={msg!=""} color={msgType}>{msg}</Alert>
  <Table responsive>
    <tr><th>{$_("words.Actions")}</th><th>{$_("instances.RootFolder")}</th><th>{$_("words.Files")}</th><th>{info.Table.Items}</th></tr>
    {#each info.RootFolders as rf}
      <tr>
        <td>
          <Badge color="primary"><a href="#top" on:click={() => {changeMe=rf;changing=rf}}>{$_("words.Change")}</a></Badge>
          <Badge color="primary"><a href="#top" on:click={() => {deleteFolder(rf)}}>{$_("words.Delete")}</a></Badge>
        </td>
        <td><code>{rf}</code></td>
        <td>{info.FoldersF[rf]?info.FoldersF[rf]:0}</td>
        <td>{info.FoldersI[rf]?info.FoldersI[rf]:0}</td>
      </tr>
    {/each}
  </Table>
</Card>

<Modal centered body isOpen={changeMe!=""}>
  <T id="instances.EnterNewRootFolder" files={info.Table.Files} items={info.Table.Items}/>
  <FormGroup floating>
    <InputGroup>
      <Button disabled={updating} class="setting-name" color="secondary" on:click={pickFolder}>
        <Fa icon={faFolderOpen} /> {$_("words.Browse")}
      </Button>
      <Input type="text" disabled={updating} id="Path" bind:value={changeMe}/>
    </InputGroup>
  </FormGroup>
  {#if !updating}
    <Button 
      disabled={changeMe == changing || (changeMe.slice(-1) != "/" && changeMe.slice(-1) != "\\")} 
      color="success" on:click={updateRoot}>{$_("words.Save")}
    </Button>
    <Button color="danger" on:click={() => {changeMe=""}}>{$_("words.Cancel")}</Button>
  {:else if progress != undefined}
    <Progress striped animated color="success" value={progress[info.Table.Items]/totals[info.Table.Items]*100}>
      {info.Table.Items}: {progress[info.Table.Items]}/{totals[info.Table.Items]}
    </Progress>
    <Progress striped animated color="success" value={progress[info.Table.Files]/totals[info.Table.Files]*100}>
      {info.Table.Files}: {progress[info.Table.Files]}/{totals[info.Table.Files]}
    </Progress>
  {:else}
    <Spinner size="sm" color="info" /> {$_("words.Updating")} ...
  {/if}
</Modal>