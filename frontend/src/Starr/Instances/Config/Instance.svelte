<script>
  export let starrApp
  export let index
  export let instance = undefined

  import { Input, InputGroup, InputGroupText, Button, Form, Alert, FormGroup, Badge } from "sveltestrap"
  import { app, conf } from "../../../libs/config.js"
  import { port } from "../../../libs/info.js"
  import Fa from "svelte-fa"
  import { faFolderOpen, faLock, faUnlock, faEye, faEyeSlash } from "@fortawesome/free-solid-svg-icons"
  import { PickFile, SaveInstance, TestInstance, RemoveInstance } from "../../../../wailsjs/go/app/App.js"
  import { toast } from "../../../libs/funcs.js"
  import { onDestroy } from "svelte"
  import { _ } from "../../../libs/Translate.svelte"

  let newInstance = false
  if (instance == undefined || Object.keys(instance).length == 0) {
    newInstance = true
    instance = { // It's new.
      Key: "",
      User: "",
      Pass: "",
      DBPath: "",
      App: starrApp,
      Name: starrApp+(index+1),
      URL: `http://127.0.0.1:${port[starrApp]}/${starrApp.toLowerCase()}`,
    }
  }

  let info
  let testOK = false
  let passLocked = true
  let keyLocked = !newInstance

  const reset = {}
  Object.keys(instance).map((k) =>{ reset[k] = instance[k] })
  onDestroy(() => Object.keys(reset).map((k) => {
    instance[k] = reset[k]
  }))

  function pickDbFile(start) {
    PickFile(start?start:$app.Home, "sqlite3 (*.db)", "*.db").then(
      path => { if (path != "") instance.DBPath = path },
    )
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
    <!-- Name -->
    <FormGroup floating>
      <div class="text-danger input-label" slot="label" style="display:{reset.Name != instance.Name?"block":"none"}">{$_("words.Unsaved")}</div>
      <InputGroup>
        <InputGroupText class="setting-name">{$_("words.Name")}</InputGroupText>
        <Input invalid={reset.Name != instance.Name} type="text" placeholder={$_("17charactermax")} maxlength={17} id="Name" bind:value={instance.Name} />
      </InputGroup>
    </FormGroup>
    <!-- URL -->
    <FormGroup floating>
      <div class="text-danger input-label" slot="label" style="display:{reset.URL != instance.URL?"block":"none"}">{$_("words.Unsaved")}</div>
      <InputGroup>
        <InputGroupText class="setting-name">{$_("words.URL")}</InputGroupText>
        <Input invalid={reset.URL != instance.URL} type="text" id="URL" bind:value={instance.URL} />
      </InputGroup>
    </FormGroup>
    <!-- API Key -->
    <FormGroup floating>
      <div class="text-danger input-label" slot="label" style="display:{reset.Key != instance.Key?"block":"none"}">{$_("words.Unsaved")}</div>
      <InputGroup>
        <InputGroupText class="setting-name">{$_("words.APIKey")}</InputGroupText>
        <Button color={keyLocked?"success":"danger"} on:click={(e) => {e.preventDefault();keyLocked=!keyLocked}}>
          <Fa style="width:20px" icon={keyLocked?faLock:faUnlock} />
        </Button>
        <Input invalid={reset.Key != instance.Key} type={keyLocked?"password":"text"} readonly={keyLocked} id="Key" placeholder={$_("32characterhash")} bind:value={instance.Key}/>
      </InputGroup>
    </FormGroup>
    <!-- Username -->
    <FormGroup floating>
      <div class="text-danger input-label" slot="label" style="display:{reset.User != instance.User?"block":"none"}">{$_("words.Unsaved")}</div>
      <InputGroup>
        <InputGroupText class="setting-name">{$_("words.Username")}</InputGroupText>
        <Input invalid={reset.User != instance.User}  type="text" id="User" placeholder="admin" bind:value={instance.User}/>
      </InputGroup>
    </FormGroup>
    <!-- Password -->
    <FormGroup floating>
      <div class="text-danger input-label" slot="label" style="display:{reset.Pass != instance.Pass?"block":"none"}">{$_("words.Unsaved")}</div>
      <InputGroup>
        <InputGroupText class="setting-name">{$_("words.Password")}</InputGroupText>
        <Button color={passLocked?"danger":"success"} outline on:click={(e) => {e.preventDefault();passLocked=!passLocked}}>
          <Fa style="width:20px" primaryColor={passLocked?"green":"red"} icon={passLocked?faEyeSlash:faEye} />
        </Button>
        <Input invalid={reset.Pass != instance.Pass} type={passLocked?"password":"text"} id="Pass" placeholder="******" bind:value={instance.Pass}/>
      </InputGroup>
    </FormGroup>
    <!-- DB file path -->
    <FormGroup floating>
      <InputGroup>
        <Button class="setting-name" color="secondary" on:click={(e) => {e.preventDefault(); pickDbFile(instance.DBPath)}}>
          <Fa icon="{faFolderOpen}" /> {$_("words.DBPath")}
        </Button>
        <Input feedback={$_("words.Unsaved")} invalid={reset.DBPath != instance.DBPath} type="text" id="DBPath" bind:value={instance.DBPath} placeholder="{$app.IsWindows?"C:\\some\\path\\"+instance.App+"\\db":"/some/path/"+instance.App+".db"}" />
      </InputGroup>
    </FormGroup>
    <!-- action buttons -->
    <Button class="actions" color="primary">{$_("words.Save")}</Button>
    <Button class="actions" color="success" name="Test" value={true}>{$_("words.Test")}</Button>
    {#if !newInstance}
      <Button class="actions delete" disabled={newInstance} color="danger" name="Delete" value={true}>{$_("words.Delete")}</Button>
      <Button class="actions" disabled={newInstance} color="warning" name="Reset" value={true}>{$_("words.Reset")}</Button>
    {/if}
    <Alert isOpen={info!=undefined} color={testOK?"success":"danger"} dismissible>{@html info}</Alert>
  </Form>
  {/if}
</div>

<style>
  /* style the buttons a little different */
  #container :global(button.actions) {
    padding:0px 6px 0px 6px;
    margin:4px 1px;
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

  /* make input labels land in the right spot */
  #container :global(.input-label) {
    margin-top: 5px;
    font-size: 12px;
  }

  #container :global(.form-floating) {
    margin-bottom: 0.3rem !important;
  }
</style>
