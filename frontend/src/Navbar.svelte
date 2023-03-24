<script type="ts">
  import { 
      Collapse, 
      Dropdown, 
      DropdownItem,
      DropdownMenu, 
      DropdownToggle,
      Nav,
      Navbar,
      NavbarBrand, 
      NavbarToggler,
      NavLink, 
      Styles, 
      Tooltip
  } from "sveltestrap"
  import { EventsOn } from "../wailsjs/runtime"
  import Fa from "svelte-fa"
  import { faGear, faBookBible, faLink } from "@fortawesome/free-solid-svg-icons"
  import About from "./About.svelte"
  import Landing from "./Landing.svelte"
  import Starr from "./Starr/Starr.svelte"
  import Toolbox from "./Toolbox.svelte"
  import Links from "./Links.svelte"
  import Settings from "./Settings/Settings.svelte"
  import Applogo from "./libs/Applogo.svelte"
  import { ToastContainer, FlatToast }  from "svelte-toasts"
  import { app as ver, conf } from "./libs/config.js"
  import ConfigInput from "./libs/Input.svelte"
  import bgVint from "./assets/images/vintage-background.png"
  import bgDark from "./assets/images/dark-background.png"
  import { _, isReady } from "./libs/Translate.svelte"
  import { toast } from "./libs/funcs"

  let isOpen = false // nav open/closer tracker (mobile)
  let app
  const starrs = ["Lidarr", "Prowlarr", "Radarr", "Readarr", "Sonarr", "Whisparr"]

  $: if (app == undefined) app = $ver.Title // start page (landing)

  // Keep dark-mode class up to date with dark config setting.
  $: $conf.Dark ? window.document.body.classList.add("dark-mode") : window.document.body.classList.remove("dark-mode")

  function nav(a) {
    isOpen = false
    app = a
  }

  // Sometimes the config changes outside the GUI.
  EventsOn("configChanged", data => {
    $conf = data
    if ($conf.DevMode) toast("warning", "Config Updated", "EVENT (debug)")
  })

  /* Prevent right-click when dev mode is disabled. */
  function blockRightClick(e) {if (!$conf.DevMode) e.preventDefault() }
  document.removeEventListener("contextmenu", blockRightClick)
  document.addEventListener("contextmenu", blockRightClick)
</script> 

<!-- Preload these to prevent a white page for a moment when switching dark on/off. -->
<link rel="preload" as="image" href={bgVint}>
<link rel="preload" as="image" href={bgDark}>
<Styles />

<!-- This gets used by any toast from any page. -->
<ToastContainer placement="bottom-right" let:data={data}><FlatToast {data} /></ToastContainer>

<!-- This if statement prevents the app from loading until the config and locale is retrieved from the backend. -->
{#if Object.keys($conf).length > 0 && $isReady == true}
<Navbar color="secondary" dark={$conf.Dark} expand="md py-0">
  <NavbarBrand on:click={(e) => (app = $ver.Title,e.preventDefault())}>
    <Applogo size="25px" {app} /> {$_("words."+app) == "words."+app ? app : $_("words."+app)}
  </NavbarBrand>
  {#if $conf.Hide.Dark != true}
    <ConfigInput type="switch" id="Dark" notoast noreload></ConfigInput>
  {/if}
  <NavbarToggler on:click={() => (isOpen = !isOpen)} />
  <Collapse {isOpen} navbar expand="md">
    <Nav class="ms-auto" navbar>
      {#each starrs as appLink}
        {#if $conf.Hide[appLink] != true}
          <Tooltip target={appLink} class="d-none d-md-block" placement="bottom">{appLink}</Tooltip>
          <NavLink id={appLink} on:click={()=>nav(appLink)}>
            <Applogo size="20px" app={appLink} /> <span class="d-md-none">{appLink}</span>
          </NavLink>
        {/if}
      {/each}
      {#if $conf.Hide.Settings != true}
      <Dropdown nav inNavbar>
        <DropdownToggle nav>
          <Applogo size="20px" app="Settings" /> <span class="d-md-none">{$_("words.Configuration")}</span>
        </DropdownToggle>
        <DropdownMenu dark={$conf.Dark} end>
          <DropdownItem on:click={()=>nav("Settings")}><Fa primaryColor="sienna" icon={faGear} /> {$_("words.Settings")}</DropdownItem>
          <DropdownItem on:click={()=>nav("Toolbox")}><Applogo size="19px" app="Toolbox" /> Toolbox</DropdownItem>
          <DropdownItem on:click={()=>nav("Links")}><Fa primaryColor="dodgerblue" icon={faLink} /> {$_("words.Links")}</DropdownItem>
          <DropdownItem on:click={()=>nav("About")}><Fa primaryColor="mediumpurple" icon={faBookBible} /> {$_("words.About")}</DropdownItem>
        </DropdownMenu>
      </Dropdown>
      {/if}
    </Nav>
  </Collapse>
</Navbar>
<br />

<main>
  <!-- if blocks cause the page to destroy when navigating away, using hidden={false} does not -->
  {#if app == "About"}
    <About />
  {/if}
  <Landing hidden={app != $ver.Title} />
  <Toolbox hidden={app != "Toolbox"} />
  <Links  hidden={app != "Links"} />
  {#if app == "Settings"}
    <Settings />
  {/if}
  {#each starrs as starrApp}
  <!-- the key block forces the page to refresh when the hide method is toggled.-->
    {#key $conf.Hide[starrApp]}
      <Starr hidden={app != starrApp} {starrApp} />
    {/key}
  {/each}
</main>
{/if}

<style>
  main {
    /* Set main height to 100% minus nav bar height. */
    height: calc(100% - 75px);
  }
</style>
