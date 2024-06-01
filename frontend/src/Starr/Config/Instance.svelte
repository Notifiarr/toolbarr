<script lang="ts">
  export let starrApp: StarrApp
  export let index: number
  export let instance: Instance = { // It's new.
    Key: "",
    User: "",
    Pass: "",
    DBPath: "",
    App: starrApp,
    Name: starrApp+(index+1),
    URL: `http://127.0.0.1:${port[starrApp]}/${starrApp.toLowerCase()}`,
    SSL: false,
    Form: false,
    Timeout: 120000000000,
  }
  export let defaultInstance: boolean = false
  export let newInstance: boolean = false

  import type { StarrApp, Instance } from "/src/libs/config"
  import { Input, InputGroup, InputGroupText, Button, Form, Alert, FormGroup, Tooltip } from "@sveltestrap/sveltestrap"
  import { app, conf } from "/src/libs/config"
  import { port } from "/src/libs/info"
  import Fa from "svelte-fa"
  import { faFolderOpen, faLock, faUnlock, faEye, faEyeSlash } from "@fortawesome/free-solid-svg-icons"
  import { PickFile, SaveInstance, RemoveInstance } from "/wailsjs/go/app/App"
  import { TestInstance } from "/wailsjs/go/starrs/Starrs"
  import { toast } from "/src/libs/funcs"
  import T, { _ } from "/src/libs/Translate.svelte"

  let isDefault = defaultInstance
  let info: string = ""
  let testOK = false
  let passLocked = true
  let keyLocked = !newInstance

  let reset: Instance = {...instance}

  function pickDbFile(start?: string) {
    PickFile(start?start:$app.Home?$app.Home:"/", "sqlite3 (*.db)", "*.db").then(
      path => { if (path != "" && instance) instance.DBPath = path },
    )
  }

  function testInstance(e: MouseEvent) {
    e.preventDefault()
    info = ""
    const interval = setInterval(() => {
      clearInterval(interval)
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
    }, 200)
  }

  function removeInstance(e: MouseEvent) {
    e.preventDefault()
    RemoveInstance(index, starrApp).then(
        resp => {
          if (resp.List) $conf.Instances[starrApp] = resp.List
          if (resp.Msg) toast("info", resp.Msg)
        },
        err => {toast("error", err)},
      )
  }

  function resetInstance(e: MouseEvent) {
    e.preventDefault()
    instance = {...reset}
  }

  function saveInstance(e: MouseEvent) {
    e.preventDefault()
    SaveInstance(index, instance, defaultInstance).then(
      resp => {
        if (resp.List) $conf.Instances[starrApp] = resp.List
        if (resp.Msg) toast("success", resp.Msg)
        if (instance) reset = {...instance}
        if (defaultInstance) $conf.Instance[starrApp] = index
        isDefault = defaultInstance
      },
      err => {toast("error", err)},
    )
  }
</script>

{$_("instances.instanceConfig")}
<div id="container">
  {#if instance}
  <Form>
    <Input type="text" hidden bind:value={instance.App} />
    <!-- Name -->
    <FormGroup floating>
      <div class="text-danger input-label" slot="label" style="display:{reset.Name != instance.Name?"block":"none"}">{$_("words.Unsaved")}</div>
      <InputGroup>
        <InputGroupText class="setting-name">{$_("words.Name")}</InputGroupText>
        <Input invalid={reset.Name != instance.Name} type="text" placeholder={$_("20charactermax")} maxlength={20}  bind:value={instance.Name} />
      </InputGroup>
    </FormGroup>
    <!-- URL -->
    <FormGroup floating>
      <div class="text-danger input-label" slot="label" style="display:{reset.URL != instance.URL?"block":"none"}">{$_("words.Unsaved")}</div>
      <InputGroup>
        <InputGroupText class="setting-name">{$_("words.URL")}</InputGroupText>
        <Input invalid={reset.URL != instance.URL} type="text" bind:value={instance.URL} />
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
        <Input invalid={reset.Key != instance.Key} type={keyLocked?"password":"text"} readonly={keyLocked}  placeholder={$_("32characterhash")} bind:value={instance.Key}/>
      </InputGroup>
    </FormGroup>
    <!-- Username -->
    <FormGroup floating>
      <div class="text-danger input-label" slot="label" style="display:{reset.User != instance.User?"block":"none"}">{$_("words.Unsaved")}</div>
      <InputGroup>
        <InputGroupText class="setting-name">{$_("words.Username")}</InputGroupText>
        <Input invalid={reset.User != instance.User} type="text" placeholder="admin" bind:value={instance.User}/>
        <InputGroupText id="form{index}{starrApp}">{$_("words.Form")}</InputGroupText>
        <Tooltip target="form{index}{starrApp}">{$_("configtooltip.Form")}</Tooltip>
        <InputGroupText color="success">
          <Input type="switch" id="form{index}{starrApp}switch" invalid={instance.Form!=reset.Form} bind:checked={instance.Form}/>
        </InputGroupText>
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
        <Input invalid={reset.Pass != instance.Pass} type={passLocked?"password":"text"} placeholder="******" bind:value={instance.Pass}/>
      </InputGroup>
    </FormGroup>
    <!-- Timeout -->
    <FormGroup floating>
      <div class="text-danger input-label" slot="label" style="display:{reset.Timeout != instance.Timeout?"block":"none"}">{$_("words.Unsaved")}</div>
      <InputGroup>
        <InputGroupText class="setting-name">{$_("words.Timeout")}</InputGroupText>
        <Input invalid={reset.Timeout != instance.Timeout} type="select" bind:value={instance.Timeout}>
          <option value={30000000000}><T id="configvalues.CountSeconds" count={30}/></option>
          <option value={45000000000}><T id="configvalues.CountSeconds" count={45}/></option>
          <option value={60000000000}><T id="configvalues.CountMinutes" count={1}/></option>
          <option value={120000000000}><T id="configvalues.CountMinutes" count={2}/></option>
          <option value={180000000000}><T id="configvalues.CountMinutes" count={3}/></option>
          <option value={240000000000}><T id="configvalues.CountMinutes" count={4}/></option>
          <option value={300000000000}><T id="configvalues.CountMinutes" count={5}/></option>
        </Input>
        <InputGroupText id="default{index}{starrApp}">{$_("words.Default")}</InputGroupText>
        <Tooltip target="default{index}{starrApp}">{$_("configtooltip.AfterRestarting")}</Tooltip>
        <InputGroupText color="success"><Input type="switch" id="default{index}{starrApp}switch" disabled={defaultInstance&&isDefault} invalid={defaultInstance!=isDefault} bind:checked={defaultInstance}/></InputGroupText>
      </InputGroup>
    </FormGroup>
    <!-- DB file path -->
    <FormGroup floating>
      <InputGroup>
        <Button style="text-align:left" class="setting-name" color="secondary" on:click={(e) => {e.preventDefault(); pickDbFile(instance?.DBPath)}}>
          <Fa icon="{faFolderOpen}" /> {$_("words.DBPath")}
        </Button>
        <Input feedback={$_("words.Unsaved")} invalid={reset.DBPath != instance.DBPath}
          type="text" bind:value={instance.DBPath}
          placeholder="{($app.IsWindows?"C:\\some\\path\\":"/some/path/")+instance.App.toLowerCase()+".db"}" />
      </InputGroup>
    </FormGroup>
    <!-- action buttons -->
    <Button class="actions" color="primary" on:click={saveInstance}>{$_("words.Save")}</Button>
    <Button class="actions" color="success" on:click={testInstance}>{$_("words.Test")}</Button>
    {#if !newInstance}
    <Button class="actions ml" color="warning" on:click={resetInstance}>{$_("words.Reset")}</Button>
    <Button class="actions" disabled={newInstance} color="danger" on:click={removeInstance}>{$_("words.Delete")}</Button>
    {/if}
    <Alert fade={false} isOpen={info!=""} color={testOK?"success":"danger"} dismissible>{@html info}</Alert>
  </Form>
  {/if}
</div>

<style>
  /* style the buttons a little different */
  #container :global(button.actions) {
    padding:0px 6px 0px 6px;
    margin:4px 1px;
  }

  /* move a button over a few pixels */
  #container :global(button.ml) {
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
