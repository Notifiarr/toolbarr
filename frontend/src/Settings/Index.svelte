<script>
  export let hidden = false

  import { 
    Button,
    Container,
    Form,
    InputGroup,
    Nav,
    NavLink,
    Offcanvas,
    Row,
  } from "sveltestrap"
  import Logs from "./Logs.svelte"
  import Advanced from "./Advanced.svelte"
  import { CreateShortcut } from "../../wailsjs/go/app/App.js"
  import BGLogo from "../libs/BackgroundLogo.svelte"
  import Fa from "svelte-fa"
  import { faQuestion } from "@fortawesome/free-solid-svg-icons"
  import { conf, app } from "../libs/config.js"
  import windowsConf from "../assets/images/windows-conf-file.png"
  import { toast } from "../libs/funcs"
  import ConfigInput from "../libs/Input.svelte"
  import General from "./General.svelte"
  import T, { _ } from "../libs/Translate.svelte"

  let activeTab = Logs
  let confHelp = false
  let confSpin = false

  function createWindowsShortcut(e) {
    e.preventDefault()
    CreateShortcut().then(
      msg => (toast("info", msg)),
      error => (toast("error", error))
    )
  }
</script>

<BGLogo {hidden} url="golift">
  <div class="settings">
  <Container>
    <Row>
      <p>{$_("main_application_settings")}</p>
      <Form class="Settings">
        <div on:mouseenter={() => {confSpin=true}} on:mouseleave={() => {confSpin=false}}>
          <InputGroup>
            <ConfigInput locked type="text" id="File" placement="bottom" />
            <Button on:click={(e) => {e.preventDefault();confHelp = !confHelp}}>
              <Fa primaryColor="cyan" spin={confSpin} icon="{faQuestion}" />
            </Button>
          </InputGroup>
        </div>
      <br />
        <Nav tabs fill>
          <NavLink href="#" on:click={() => (activeTab = General)} active={activeTab == General}>{$_("words.General")}</NavLink>
          <NavLink href="#" on:click={() => (activeTab = Logs)} active={activeTab == Logs}>{$_("words.Logs")}</NavLink>
          <NavLink href="#" on:click={() => (activeTab = Advanced)} active={activeTab == Advanced}>{$_("words.Advanced")}</NavLink>
        </Nav>
        <br />
        <svelte:component this={activeTab} />
      </Form>
    </Row>
  </Container>
  </div>

  <Offcanvas 
    style="width:50%;min-width:390px;max-width:550px"
    class="{$conf.Dark ? "bg-secondary" : "bg-light"}"
    isOpen={confHelp}
    toggle={() => {confHelp = !confHelp}}
    header={$_("customconfig.CustomConfigPath")} placement="end">
    {#if $app.IsLinux}
      <p><T id="customconfig.linuxHelp" title={$app.Title} name={$app.Name}/></p>
      <h5>{$_("words.Example")}</h5>
      <p>
        <code>{$app.Name} -c /path/to/{$app.Name}.conf</code><br>
        {$_("customconfig.Withafullpath")}:<br>
        <code>{$app.Exe} -c {$app.Home}/.{$app.Name}/{$app.Name}.conf</code><br>
        {$_("customconfig.Makeabashscript")}
      </p>
    {:else if $app.IsMac}
      <p><T id="customconfig.macHelp" title={$app.Title} name={$app.Name}/></p>
    {:else}
      <p><T id="customconfig.winHelp" title={$app.Title} name={$app.Name}/></p>
      <p>
        <T id="customconfig.winCustom" title={$app.Title} name={$app.Name}/>
        <br>
        <Button size="sm" color="info" on:click={createWindowsShortcut}>{$_("customconfig.CreateShortcut")}</Button>
      </p>
      <h5>{$_("words.Directions")}</h5>
      <ol><T id="customconfig.winDirectionList" name={$app.Name}/></ol>
      <h5>{$_("words.Example")}</h5>
      <img width="100%" alt="{$_("screenshotofshortcutwindow")}" src={windowsConf}>
    {/if}
  </Offcanvas>
</BGLogo>

<style>
  .settings :global(.setting-name) {
    min-width: max-content;
    width: 25%;
    max-width:200px;
  }
</style>
