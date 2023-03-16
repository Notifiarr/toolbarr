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
  import { conf } from "./libs/config.js"
  import ConfigInput from "./libs/Input.svelte"
  import bgVint from "./assets/images/vintage-background.png"
  import bgDark from "./assets/images/dark-background.png"

  let isOpen = false // nav open/closer tracker (mobile)
  let app = "Toolbarr" // start page (landing)
  // Keep dark-mode class up to date with dark config setting.
  $: $conf.Dark ? window.document.body.classList.add("dark-mode") : window.document.body.classList.remove("dark-mode")

  function nav(a) {
    isOpen = false
    app = a
  }

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

<!-- This if statement prevents the app from loading until the config is retrieved from the backend. -->
{#if Object.keys($conf).length > 0}
<Navbar color="secondary"  expand="md py-0">
  <NavbarBrand on:click={(e) => (app = "Toolbarr",e.preventDefault())}><Applogo size="25px" {app} /> {app}</NavbarBrand>
  <ConfigInput type="switch" id="Dark" name="Dark" notoast noreload></ConfigInput>
  <NavbarToggler on:click={() => (isOpen = !isOpen)} />
  <Collapse {isOpen} navbar expand="md">
    <Nav class="ms-auto" navbar>
      {#each ["Lidarr", "Prowlarr", "Radarr", "Readarr", "Sonarr", "Whisparr"] as appLink}
        <Tooltip target={appLink} class="d-none d-md-block" placement="bottom">{appLink}</Tooltip>
        <NavLink id={appLink} on:click={()=>nav(appLink)}>
          <Applogo size="20px" app={appLink} /><span class="d-md-none">{appLink}</span>
        </NavLink>
      {/each}
      <Dropdown nav inNavbar>
        <DropdownToggle nav>
          <Applogo size="20px" app="Settings" /> <span class="d-md-none">Configuration</span>
        </DropdownToggle>
        <DropdownMenu dark={$conf.Dark} end>
          <DropdownItem on:click={()=>nav("Settings")}><Fa primaryColor="sienna" icon="{faGear}" /> Settings</DropdownItem>
          <DropdownItem on:click={()=>nav("Toolbox")}><Applogo size="19px" app="Toolbox" /> Toolbox</DropdownItem>
          <DropdownItem on:click={()=>nav("Links")}><Fa primaryColor="dodgerblue" icon="{faLink}" /> Links</DropdownItem>
          <DropdownItem on:click={()=>nav("About")}><Fa primaryColor="mediumpurple" icon="{faBookBible}" /> About</DropdownItem>
        </DropdownMenu>
      </Dropdown>
    </Nav>
  </Collapse>
</Navbar>
<br />

<main>
  {#if app == "Toolbarr" }
  <Landing />
  {:else if app == "About"}
  <About />
  {:else if app == "Toolbox" }
  <Toolbox />
  {:else if app == "Links" }
  <Links />
  {:else if app == "Settings" }
  <Settings />
  {:else}
  <Starr {app} />
  {/if}
</main>
{/if}

<style>
  main {
    /* Set main height to 100% minus nav bar height. */
    height: calc(100% - 75px);
  }
</style>
