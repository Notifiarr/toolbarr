<script lang="ts">
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
      Tooltip
  } from "sveltestrap" 
  import { EventsOn } from "/wailsjs/runtime"
  import Fa from "svelte-fa"
  import { faGear, faBookBible, faLink } from "@fortawesome/free-solid-svg-icons"
  import About from "/src/About.svelte" 
  import Landing from "/src/Landing.svelte"
  import Starr from "/src/Starr/Index.svelte"
  import Toolbox from "/src/Toolbox.svelte"
  import Links from "/src/Links.svelte"
  import Settings from "/src/Settings/Index.svelte"
  import Applogo from "/src/libs/Applogo.svelte"
  import { ToastContainer, FlatToast }  from "svelte-toasts"
  import { app, conf, AllStarrs } from "./libs/config"
  import ConfigInput from "/src/libs/Input.svelte"
  import bgVint from "/src/assets/images/vintage-background.png"
  import bgDark from "/src/assets/images/dark-background.png"
  import { _, isReady } from "/src/libs/Translate.svelte"
  import { toast } from "/src/libs/funcs"
  import darkCSS from "/src/assets/bootstrap.darkly.css?url"
  import liteCSS from "/src/assets/bootstrap.flatly.css?url"
  import icons from "/src/assets/bootstrap-icons.css?url"

  let isOpen = false // nav open/closer tracker (mobile)
  $: pageName = $app.Title
  // Keep dark-mode class up to date with dark config setting.
  $: $conf.Dark ? window.document.body.classList.add("dark-mode") : window.document.body.classList.remove("dark-mode")

  function nav(newPage: string) {
    isOpen = false
    pageName = newPage
  }

  // Sometimes the config changes outside the GUI.
  EventsOn("configChanged", data => {
    $conf = data.Settings
    if ($conf.DevMode) toast("warning", data.Msg, $_("words.CONFIG") + " ("+$_("words.debug")+")")
  })

  /* Prevent right-click when dev mode is disabled. */
  function blockRightClick(e) { if (!$conf.DevMode) e.preventDefault() }
  document.removeEventListener("contextmenu", blockRightClick)
  document.addEventListener("contextmenu", blockRightClick)
</script> 

<svelte:head>
  <!-- Preload these to prevent a white page for a moment when switching dark on/off. -->
  <link rel="preload" as="image" href={bgVint}>
  <link rel="preload" as="image" href={bgDark}>
  <link rel="stylesheet" href={$conf.Dark?darkCSS:liteCSS} />
  <link rel="stylesheet" href={icons} />
</svelte:head>

<!-- This gets used by any toast from any page. -->
<ToastContainer placement="bottom-right" let:data={data}><FlatToast {data} /></ToastContainer>

<!-- This if statement prevents the app from loading until the config and locale is retrieved from the backend. -->
{#if Object.keys($conf).length > 0 && $isReady == true}
<Navbar color="secondary" dark={$conf.Dark} expand="md py-0">
  <NavbarBrand on:click={(e) => (pageName = $app.Title,e.preventDefault())}>
    <Applogo size="25px" app={pageName} /> 
    {$_("words."+pageName) == "words."+pageName ? pageName : $_("words."+pageName)}
  </NavbarBrand>
  {#if $conf.Hide.Dark != true}
    <ConfigInput type="switch" id="Dark" notoast noreload></ConfigInput>
  {/if}
  <NavbarToggler on:click={() => (isOpen = !isOpen)} />
  <Collapse {isOpen} navbar expand="md">
    <Nav class="ms-auto" navbar>
      {#each AllStarrs as starrApp}
        {#if $conf.Hide[starrApp] != true}
          <Tooltip target={starrApp} class="d-none d-md-block" placement="bottom">{starrApp}</Tooltip>
          <NavLink id={starrApp} on:click={()=>nav(starrApp)}>
            <Applogo size="20px" app={starrApp} /> <span class="d-md-none">{starrApp}</span>
          </NavLink>
        {/if}
      {/each}
      {#if $conf.Hide.Settings != true}
      <Dropdown nav inNavbar autoClose>
        <DropdownToggle nav>
          <Applogo size="20px" app="Settings" /> <span class="d-md-none">{$_("words.Configuration")}</span>
        </DropdownToggle>
        <DropdownMenu dark={$conf.Dark} end>
          <DropdownItem active={pageName=="Settings"} on:click={()=>nav("Settings")}>
            <Fa primaryColor="sienna" icon={faGear} /> {$_("words.Settings")}
          </DropdownItem>
          <DropdownItem active={pageName=="Toolbox"} on:click={()=>nav("Toolbox")}>
            <Applogo size="19px" app="Toolbox" /> {$_("words.Toolbox")}
          </DropdownItem>
          <DropdownItem active={pageName=="Links"} on:click={()=>nav("Links")}>
            <Fa primaryColor="dodgerblue" icon={faLink} /> {$_("words.Links")}
          </DropdownItem>
          <DropdownItem active={pageName=="About"} on:click={()=>nav("About")}>
            <Fa primaryColor="mediumpurple" icon={faBookBible} /> {$_("words.About")}
          </DropdownItem>
        </DropdownMenu>
      </Dropdown>
      {/if}
    </Nav>
  </Collapse>
</Navbar>
<br />

<main>
  <!-- if blocks cause the page to destroy when navigating away, using hidden={false} does not -->
  {#if pageName == "About"}
    <About />
  {/if}
  <Landing hidden={pageName != $app.Title} />
  <Toolbox hidden={pageName != "Toolbox"} />
  <Links  hidden={pageName != "Links"} />
  {#if pageName == "Settings"}
    <Settings />
  {/if}
  {#each AllStarrs as starrApp}
  <!-- the key block forces the page to refresh when the hide method is toggled.-->
    {#key $conf.Hide[starrApp]}
      <Starr hidden={pageName != starrApp} {starrApp} />
    {/key}
  {/each}
</main>
{/if}<!-- $isReady -->

<style>
  main {
    /* Set main height to 100% minus nav bar height. */
    height: calc(100% - 75px);
  }
</style>
