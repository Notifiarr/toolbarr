<script>
  export let starrApp
  export let index
  export let instance = undefined

  import { Input, InputGroup, InputGroupText, Button, Form, Alert } from "sveltestrap"
  import { app, conf } from "../../../libs/config.js"
  import { port } from "../../../libs/info.js"
  import Fa from "svelte-fa"
  import { faFolderOpen, faLock, faUnlock, faEye, faEyeSlash } from "@fortawesome/free-solid-svg-icons"
  import { PickFolder, SaveInstance, TestInstance, RemoveInstance } from "../../../../wailsjs/go/app/App.js"
  import { toast } from "../../../libs/funcs.js"
  import { onDestroy } from "svelte"
  import { _ } from "../../../libs/Translate.svelte"

  let newInstance = false
  if (instance == undefined || Object.keys(instance).length == 0) {
    newInstance = true
    instance = { // It's new.
      App: starrApp,
      Name: starrApp+(index+1),
      URL: `http://127.0.0.1:${port[starrApp]}/${starrApp.toLowerCase()}`
    }
  }

  let info
  let testOK = false
  let passLocked = true
  let keyLocked = !newInstance

  const reset = {...instance}
  onDestroy(() => Object.keys(reset).map((k) => {
    instance[k] = reset[k]
  }))

  // This func opens a "pick a folder" dialog and populates the Log File Path with the current config file folder.
  function openFolder(event) {
    event.preventDefault()
    PickFolder($app.Home).then(path => {if (path != "") instance.DBPath = path})
  }

  function saveInstance(e) {
    e.preventDefault()
    const val = new FormData(e.target)

    if (val.get("Reset")) {
      instance = {...reset}
    } else if (val.get("Test")) {
      info = undefined
      TestInstance(instance).then(
        msg => {
          info = msg
          testOK = true
        },
        err => {
          info = err
          testOK = false
        },
      )
    } else if (val.get("Delete")) {
      RemoveInstance(index, starrApp).then(
        resp => {
          if (resp.List) $conf.Instances[starrApp] = resp.List
          if (resp.Msg) toast("info", resp.Msg)
        },
        err => {toast("error", err)},
      )
    } else {
      SaveInstance(index, instance).then(
        resp => {
          if (resp.List) $conf.Instances[starrApp] = resp.List
          if (resp.Msg) toast("success", resp.Msg)
          Object.keys(instance).map((k) =>{ reset[k] = instance[k] })
        },
        err => {toast("error", err)},
      )
    }
  }
</script>

{$_("instances.instanceConfig")}

<div id="container">
  {#if instance}
  <Form on:submit={saveInstance}>
    <Input type="text" hidden id="App" bind:value={instance.App} />
    <InputGroup>
      <InputGroupText class="setting-name">{$_("words.Name")}</InputGroupText>
      <Input type="text" id="Name" bind:value={instance.Name} />
    </InputGroup>
    <InputGroup>
      <InputGroupText class="setting-name">{$_("words.URL")}</InputGroupText>
      <Input type="text" id="URL" bind:value={instance.URL} />
    </InputGroup>
    <InputGroup>
      <InputGroupText class="setting-name">{$_("words.APIKey")}</InputGroupText>
      <Input type={keyLocked?"password":"text"} readonly={keyLocked} id="Key" placeholder="32 character hash" bind:value={instance.Key}/>
      <Button color={keyLocked?"success":"danger"} on:click={(e) => {e.preventDefault();keyLocked=!keyLocked}}>
        <Fa icon={keyLocked?faLock:faUnlock} />
      </Button>
    </InputGroup>
    <InputGroup>
      <InputGroupText class="setting-name">{$_("words.Username")}</InputGroupText>
      <Input type="text" id="User" placeholder="admin" bind:value={instance.User}/>
    </InputGroup>
    <InputGroup>
      <InputGroupText class="setting-name">{$_("words.Password")}</InputGroupText>
      <Input type={passLocked?"password":"text"} id="Pass" placeholder="******" bind:value={instance.Pass}/>
      <Button color={passLocked?"danger":"success"} outline on:click={(e) => {e.preventDefault();passLocked=!passLocked}}>
        <Fa primaryColor={passLocked?"green":"red"} icon={passLocked?faEyeSlash:faEye} />
      </Button>
    </InputGroup>
    <InputGroup>
      <InputGroupText class="setting-name">{$_("words.DBPath")}</InputGroupText>
      <Input type="text" id="DBPath" bind:value={instance.DBPath} placeholder="{$app.IsWindows?"C:\\some\\path\\"+instance.App+"\\db":"/some/path/"+instance.App+"/db"}" />
      <Button color="success" on:click={openFolder}>
        <Fa icon="{faFolderOpen}" />
        <span class="d-none d-md-inline-block">{$_("words.Browse")}</span>
      </Button>
    </InputGroup>

    <Button class="actions" color="primary">{$_("words.Save")}</Button>
    <Button class="actions" color="success" name="Test" value={true}>{$_("words.Test")}</Button>
    {#if !newInstance}
      <Button class="actions delete" disabled={newInstance} color="danger" name="Delete" value={true}>{$_("words.Delete")}</Button>
      <Button class="actions" disabled={newInstance} color="warning" name="Reset" value={true}>{$_("words.Reset")}</Button>
    {/if}
    <Alert isOpen={info!=undefined} color={testOK?"success":"danger"} dismissible>{info}</Alert>
  </Form>
  {/if}
</div>

<style>
  /* style the buttons a little different */
  #container :global(button.actions) {
    padding:0px 6px 0px 6px;
    margin-top:2px;
    margin-left:2px;
  }

  /* move delete button over a few pixels */
  #container :global(button.delete) {
    margin-left:20px;
  }

  /* Give the info alert a small body */
  #container :global(.alert) {
    margin: 2px;
    min-height:50px;
    padding:1px 20px 1px 5px;
  }
</style>