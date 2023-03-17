<script>
  import { 
    Button,
    Container,
    Form,
    InputGroup,
    InputGroupText,
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

  let activeTab = Logs
  let confHelp = false
  let confSpin = false
  const toggleConfHelp = () => (confHelp = !confHelp);

  function createWindowsShortcut(e) {
    e.preventDefault()
    CreateShortcut().then(
      msg => (toast("primary", msg)),
      error => (toast("error", error))
    )
  }
</script>

<BGLogo url="golift">
  <Container>
    <Row>
      <p>This is where the application settings are found.</p>
      <Form class="Settings">
        <div on:mouseenter={() => {confSpin=true}} on:mouseleave={() => {confSpin=false}}>
          <InputGroup>
            <InputGroupText class="setting-name">Config File</InputGroupText>
            <ConfigInput locked type="text" id="File" name="File" 
            tooltip="Can only be changed on application launch" placement="bottom" />
            <Button on:click={(e) => (e.preventDefault(),toggleConfHelp())}>
              <Fa primaryColor="cyan" spin={confSpin} icon="{faQuestion}" />
            </Button>
          </InputGroup>
        </div>
      <br />
        <Nav tabs fill>
          <NavLink href="#" on:click={() => (activeTab = Logs)} active={activeTab == Logs}>Logging</NavLink>
          <NavLink href="#" on:click={() => (activeTab = Advanced)} active={activeTab == Advanced}>Advanced</NavLink>
        </Nav>
        <br />
        <svelte:component this={activeTab} />
      </Form>
    </Row>
  </Container>

  <Offcanvas style="width:50%;min-width:390px;max-width:550px" class="{$conf.Dark ? "bg-secondary" : "bg-light"}" isOpen={confHelp} toggle={toggleConfHelp} header="Custom Config Path" placement="end">
    {#if $app.IsLinux}
      <p>
        Toolbarr will look for <code>toolbarr.conf</code> in the same folder as the <code>toolbarr</code> binary.
        If it is not found, then a location inside your home folder is used for the config file.
        If you want the config file to live next to the binary: copy it there, restart this app, and it will be used.
        To use a custom config file path on Linux, provide it as a cli argument when you launch the executable.
      </p>
      <h5>Example</h5>
      <p>
      <code>toolbarr -c /path/to/toolbarr.conf</code><br>
      With a full path:<br>
      <code>{$app.Exe} -c {$app.Home}/.toolbarr/toolbarr.conf</code><br>
      Make a bash alias or script to do this for you.</p>
    {:else if $app.IsMac}
    <p>
      Toolbarr will look for <code>toolbarr.conf</code> in the same folder as <code>Toolbarr.app</code>.
      If it is not found, then a location inside your home folder is used for the config file.
      If you want the config file to live next to the app: copy it there, restart this app, and it will be used.
      It's difficult to use a custom config location on a mac and is not recommended.
    </p>
    {:else}
      <p>
        Toolbarr will look for <code>toolbarr.conf</code> in the same folder as <code>Toolbarr.exe</code>.
        If it is not found, then a location inside your home folder is used for the config file.
        If you want the config file to live next to the exe file: copy it there, restart this app, and it will be used.
        If you want a custom config file location, follow the instructions below.
      </p>
      <p>
      To use a custom config file on Windows, you must create and edit a shortcut. 
      When you use the installer a shortcut is placed in your start menu, 
      and optionally on your desktop. You may use either of these files, or create a new one. 
      Click the Create Shortcut button below to create a new shortcut on your desktop now.<br>
      <Button size="sm" color="info" on:click={createWindowsShortcut}>Create Shortcut</Button>
      </p>
      <h5>Directions</h5>
      <ol>
        <li>Right-click the short cut and click Properties.</li>
        <li>Click the Shortcut tab if not already there.</li>
        <li>In Target, ADD this:</li>
        <li><code>-c "C:\path\to\toolbarr.conf"</code></li>
        <li>Replace the provided path with your own.</li>
      </ol>
      <h5>Example</h5>
      <img width="100%" alt="screenshot of shortcut window" src={windowsConf}>
    {/if}
  </Offcanvas>
</BGLogo>

<style>
  :global(.Settings .setting-name) {
    min-width:120px;
    max-width:120px;
  }
</style>
