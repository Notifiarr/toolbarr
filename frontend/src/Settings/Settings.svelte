<script>
  import { 
    Button,
    Container,
    Form,
    Input,
    InputGroup,
    InputGroupText,
    Nav,
    NavLink,
    Offcanvas,
    Row,
  } from "sveltestrap"
  import Logs from "./Logs.svelte"
  import Advanced from "./Advanced.svelte"
  import { GetConfig, CreateShortcut } from "../../wailsjs/go/app/App.js"
  import BGLogo from "../BackgroundLogo.svelte"
  import Fa from "svelte-fa"
  import { faQuestion } from "@fortawesome/free-solid-svg-icons"
  import { isLinux, isMac, dark } from './store.js';
  import windowsConf from "../assets/images/windows-conf-file.png"
  import { toast } from "../funcs";

  let activeTab = "Logs"
  let confHelp = false
  const toggleConfHelp = () => (confHelp = !confHelp);

  let conf = {}
  GetConfig().then(result => conf = result)

  function createWindowsShortcut(e) {
    e.preventDefault()
    CreateShortcut().then(
      msg => (toast("success", msg)),
      error => (toast("error", error))
    )
  }
</script>

<BGLogo url="golift">
  <Container>
    <Row>
      <p>
        This is where the application settings are found.
      </p>
      <Form class="Settings">
        <InputGroup>
          <InputGroupText class="setting-name">Config File</InputGroupText>
          <Input disabled value={conf.File} />
          <Button on:click={(e) => (e.preventDefault(),toggleConfHelp())}><Fa primaryColor="cyan" icon="{faQuestion}" /></Button>
        </InputGroup>
        <br />
        <Nav tabs fill>
          <NavLink href="#" on:click={() => (activeTab = "Logs")} active={activeTab == "Logs"}>Logging</NavLink>
          <NavLink href="#" on:click={() => (activeTab = "Advanced")} active={activeTab == "Advanced"}>Advanced</NavLink>
        </Nav>
        <br />
        {#if activeTab == "Advanced"}<Advanced />{/if}
        {#if activeTab == "Logs"}<Logs />{/if}
      </Form>
    </Row>
  </Container>
  <Offcanvas style="width:50%;min-width:390px;max-width:550px" class="{$dark ? 'bg-dark' : ''}" isOpen={confHelp} toggle={toggleConfHelp} header="Custom Config Path" placement="end">
    {#if $isLinux}
      <p>To use a custom config file path on Linux, simply provide it as a cli argument when you launch the executable.</p>
      <h5>Example</h5>
      <p>
      <code>toolbarr -c /path/to/toolbarr.conf</code><br>
      With a full path:<br>
      <code>{conf.Exe} -c {conf.Home}/.toolbarr/toolbarr.conf</code><br>
      Make a bash alias or script to do this for you.</p>
    {:else if $isMac}
      <p>It's difficult to use a custom config location on a mac and is not recommended.</p>
    {:else}
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
        <li>In Target, you're going to ADD this:</li>
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