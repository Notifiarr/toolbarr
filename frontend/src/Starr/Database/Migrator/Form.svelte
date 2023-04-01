<script>
  export let instance
  export let info
  
  import {
    Alert,
    Badge,
    Button,
    Card,
    FormGroup,
    Input,
    InputGroup,
    Modal,
    Progress,
    Spinner,
    Table,
  } from "sveltestrap"
  import {
    UpdateRootFolder,
    DeleteRootFolder,
    UpdateRecycleBin,
    UpdateInvalidItems,
  } from "../../../../wailsjs/go/starrs/Starrs.js"
  import { PickFolder } from "../../../../wailsjs/go/app/App.js"
  import T, { _ } from "../../../libs/Translate.svelte"
  import { faClose, faFolderOpen } from "@fortawesome/free-solid-svg-icons"
  import Fa from "svelte-fa"
  import { app } from "../../../libs/config.js"
  import { EventsOff, EventsOn } from "../../../../wailsjs/runtime/runtime.js"
  import { onOnce } from "../../../libs/funcs.js"

  // Used when updating root folder paths.
  let newPath
  let oldPath
  // Progress for root folder updates.
  let progress = undefined
  let totals
   // Used for root folder and recycle bin changes.
  let updating = false
  // Controls the Modal for updating root folder paths.
  let action
  // Message shows up after an action completes.
  let msg = ""
  let msgType
  // Controls the Modals for invalid paths display.
  let invalidModals = {}
  let invalidIDs = {}

  // Fill in the invalidIDs array.
  // This allows collecting check boxes for items to fix.
  function ewPa() {
    Object.keys(info.Invalid).map((table) => {
      invalidIDs[table] = {}
      info.Invalid[table].map((invalidItem) => {
        invalidIDs[table][invalidItem.ID] = false
      })
    })
  }
  ewPa()

  function success(resp) {
    disableEvents()      // turn off events
    onOnce(() => {       // wait ~1s for these
      action = undefined // close action modal
      updating = false   // re-enable buttons
    }, 1.2)
    if (resp.Msg) {
      msg = resp.Msg      // update displayed message
      msgType = "success" // update message color
    }
    if (resp.Info) {info = resp.Info; ewPa()} // update db info
  }

  function error(err) {
    disableEvents()      // disable events
    action = undefined   // close action modal
    msgType = "danger"   // update message color
    msg = err            // update displayed message
    updating = false     // re-enable buttons.
    progress = undefined // disable progress bars
  }

  // disableEvents turns events off. Used for progress bars.
  function disableEvents() {
    EventsOff("DBitemTotals", "DBfileCount")  // stop listeners
    onOnce(() => { progress = undefined }, 2) // destroy progress bars
  }

  // startEvents turns events on. Used for progress bars.
  function startEvents() {
    updating = true // disable buttons
    totals = {}     // empty totals

    EventsOn("DBitemTotals", v => {
      progress = {} // enable progress bar(s)
      totals = v    // update totals, this only comes in once.
      Object.keys(v).map(key => { progress[key] = 0 })
    })
    // update progress, this comes in a lot.
    EventsOn("DBfileCount", v => progress = { ...progress, ...v })
  }

  function updateInvalid(table) {
    startEvents()
    UpdateInvalidItems(instance, table, newPath, invalidIDs[table]).then(
      resp => {
        success(resp) // handle response
        onOnce(() => { invalidModals[table] = false }, 1) // close modal after 1s
      },
      err => {
        error(err) // handle error.
        invalidModals[table] = false // close modal
      }
    )
  }

  function updateRoot() { // changes a root folder path
    startEvents()
    UpdateRootFolder(instance, oldPath, newPath).then(success, error)
  }

  function pickFolder(e) {
    e.preventDefault()
    PickFolder($app.Home).then((path) => {
      // PickFolder does not append a slash, and we need one here, so add it.
      if (path != "") newPath = path + ($app.IsWindows?"\\":"/")
    })
  }

  function updateRecycleBin(e) {
    e.preventDefault()
    updating = true
    UpdateRecycleBin(instance, newPath).then(success, error)
  }

  function unsetReycleBin(e) {
    e.preventDefault()
    updating = true
    UpdateRecycleBin(instance, "").then(success, error)
  }

  function openRecycleChanger(e) {
    e.preventDefault()
    oldPath = newPath = info.Recycle
    action = updateRecycleBin
  }

  function deleteFolder(rf) {
    updating = true
    DeleteRootFolder(instance, rf).then(success, error)
  }

  function selectAll(table, all) {
    Object.keys(invalidIDs[table]).map(id => invalidIDs[table][id] = all)
  }

  function openInvalidModal(e, table) {
    e.preventDefault()
    invalidModals[table] = true
    newPath = info.RootFolders[0] ? info.RootFolders[0] : ""
  }
</script>

<!-- This card shows up if there are any invalid paths found. -->
{#if Object.keys(info.Invalid).length > 0}
<Card outline body color="danger">
  <h5 class="text-danger">{$_("instances.InvalidPaths")}</h5>
  {#each Object.keys(info.Invalid) as table}
    <li><a href="/" on:click={(e)=>{openInvalidModal(e, table)}}>{table}</a>: {info.Invalid[table].length}</li>
    <!-- Each table with invalid paths includes a modal to view/fix them. -->
    <Modal size="lg" body isOpen={invalidModals[table]}>
      <h5>
        <T id="instances.Invalidpathsintable" {table}/>
        <span style="float:right;"><!-- extra close button -->
          <a style="cursor:pointer;" href="/" on:click={(e)=>{e.preventDefault();invalidModals[table]=false}}>
            <Fa icon={faClose} color="red"/>
          </a>
        </span>
      </h5>
      {#if info.Table[table].Column == "Path" && info.Table[table].Name != '""'}
        <p>{$_("instances.unsupportedColumn")}</p>
      {/if}
      <Table bordered responsive>
        <tr><th></th><th>ID</th><th>{$_("words.Path")}</th><th>{$_("words.Name")}</th></tr>
        {#each info.Invalid[table] as invalidItem}
          <tr>
            <td><Input disabled={updating} class="input-selector" bind:checked={invalidIDs[table][invalidItem.ID]} type="switch"/></td>
            <td>{invalidItem.ID}</td>
            <td>{invalidItem.Path}</td>
            <td>{invalidItem.Name}</td>
          </tr>
        {/each}
      </Table>
      {#if info.Table[table].Column != "Path" || info.Table[table].Name == '""'}
        <InputGroup>
          <Button color="info" disabled>{$_("instances.NewRoot")}</Button>
          <Input disabled={updating} bind:value={newPath} type="select">
            {#each info.RootFolders as rf} <option value={rf}>{rf}</option> {/each}
          </Input>
        </InputGroup>
      {/if}
      {#if progress != undefined}
        {#each Object.keys(progress) as table}
          <Progress animated color="success"
            value={totals[table]>0?(progress[table]/totals[table]*100):100}>
            {table}: {progress[table]}/{totals[table]}
          </Progress>
        {/each}
      {:else if updating}
        <Spinner size="sm" color="info" /> {$_("words.Updating")} ...
      {:else}
        <InputGroup>
          <Button color="danger" on:click={() => invalidModals[table]=false}>{$_("words.Cancel")}</Button>
          {#if info.Table[table].Column != "Path" || info.Table[table].Name == '""'}
          <Button color="success" on:click={()=>{updateInvalid(table)}}>{$_("instances.SaveSelected")}</Button>
          <Button color="primary" on:click={()=>{selectAll(table, true)}}>{$_("instances.SelectAll")}</Button>
          <Button color="secondary" on:click={()=>{selectAll(table, false)}}>{$_("instances.SelectNone")}</Button>
          {/if}
        </InputGroup>
      {/if}
    </Modal>
  {/each}
</Card>
{/if}

<!-- Main page card that shows all the counters and action buttons. -->
<Card body outline color="primary">
  <p>
    {$_("instances.RecycleBin")}: <code>{info.Recycle?info.Recycle:"("+$_("instances.NoPathSet")+")"}</code>
    <Badge color="primary"><a href="/" on:click={openRecycleChanger}>{$_("words.Change")}</a></Badge>
    {#if info.Recycle != ""}
      <Badge color="primary"><a href="/" on:click={unsetReycleBin}>{$_("words.Unset")}</a></Badge>
    {/if}
  </p>
  <Alert fade isOpen={msg!=""} color={msgType}>{msg}</Alert>
  <h5>{$_("instances.RootFolders")}</h5>
  <Table bordered responsive>
    <tr>
      <th>{$_("words.Actions")}</th><th>{$_("words.Path")}</th>
      {#each Object.keys(info.Folders) as table}
        <th>{table}</th>
      {/each}
    </tr>
    {#each info.RootFolders as rf}
      <tr>
        <td>
          <Badge color="primary">
            <a href="/" on:click={(e) => {e.preventDefault();oldPath=newPath=rf;action=updateRoot}}>
              {$_("words.Change")}
            </a>
          </Badge>
          <Badge color="primary">
            <a href="/" on:click={(e) => {e.preventDefault();deleteFolder(rf)}}>{$_("words.Delete")}</a>
          </Badge>
        </td>
        <td><code>{rf}</code></td>
        {#each Object.keys(info.Folders) as table }
          <td>{info.Folders[table][rf]?info.Folders[table][rf]:0}</td>
        {/each}
      </tr>
    {/each}
  </Table>
</Card>

<!-- This Modal is used to change root folder or recycle bin paths.-->
<Modal centered body isOpen={action!=undefined}>
  {action==updateRoot?$_("instances.EnterNewRootFolder"):$_("instances.EnterNewRecycleBin")}
  <FormGroup floating>
    <InputGroup>
      <Button disabled={updating} class="setting-name" color="secondary" on:click={pickFolder}>
        <Fa icon={faFolderOpen} /> {$_("words.Browse")}
      </Button>
      <Input type="text" disabled={updating} bind:value={newPath}/>
    </InputGroup>
  </FormGroup>
  {#if progress != undefined}
    {#each Object.keys(progress) as table}
      <Progress animated color="success"
        value={totals[table]>0?(progress[table]/totals[table]*100):100}>
        {table}: {progress[table]}/{totals[table]}
      </Progress>
    {/each}
  {:else if updating}
    <Spinner size="sm" color="info" /> {$_("words.Updating")} ...
  {:else}
    <Button 
      disabled={newPath == oldPath || (newPath.slice(-1) != "/" && newPath.slice(-1) != "\\")} 
      color="success" on:click={action}>{$_("words.Save")}
    </Button>
    <Button color="danger" on:click={() => {action=undefined}}>{$_("words.Cancel")}</Button>
  {/if}
</Modal>

<style>
  :global(.input-selector) {
    margin: -15px -15px 0 32px;
  }
</style>