<script>
  import { 
      Collapse, 
      Dropdown, 
      DropdownItem,
      DropdownMenu, 
      DropdownToggle,
      Input,
      Nav,
      Navbar,
      NavbarBrand, 
      NavbarToggler,
      NavLink, 
      Styles, 
      Tooltip
  } from "sveltestrap"
  import Fa from "svelte-fa"
  import {faGear,faBookBible,faLink} from "@fortawesome/free-solid-svg-icons"
  import About from "./About.svelte"
  import Landing from "./Landing.svelte"
  import Starr from "./Starr/Starr.svelte"
  import Toolbox from "./Toolbox.svelte"
  import Links from "./Links.svelte"
  import Settings from "./Settings/Settings.svelte"
  import Applogo from "./Applogo.svelte"
  import {ToastContainer, FlatToast}  from "svelte-toasts"
  import bgVint from "./assets/images/vintage-background.png"
  import bgDark from "./assets/images/dark-background.png"
  import {GetConfig, SaveConfigItem} from "../wailsjs/go/app/App"
  import {toasts}  from "svelte-toasts"

  let navIsOpen = false
  function toggleNavOpen(event) {
    navIsOpen = event.detail.isOpen
  }

  let app = "Toolbarr"

  const classes = window.document.body.classList
  
  let conf = {}
  GetConfig().then(result => (
    conf = result,
    conf.Dark ? classes.add("dark-mode") : classes.remove("dark-mode")
  ))

  function saveDark(event) {
    conf.Dark = event.target.checked
    conf.Dark ? classes.add("dark-mode") : classes.remove("dark-mode")

    conf[event.target.id] = event.target.checked
    SaveConfigItem(event.target.name, event.target.checked+"", false).then((msg) => {
      /* nothing shown */
    }, (error) => {
      const toast = toasts.add({
        title: "Error Saving Config",
        description: error,
        duration: 7000,
        theme: conf.Dark ? "dark" : "light",
        type: "error",
        onClick: () => {toast.remove()},
        showProgress: true,
      })
    })
  }

  /* Prevent right-click. */
  // document.addEventListener("contextmenu", event => event.preventDefault())
</script> 
<link rel="preload" as="image" href={bgVint}>
<link rel="preload" as="image" href={bgDark}>
<Styles />
<Navbar color="secondary"  expand="md py-0">
  <NavbarBrand on:click={(e) => (e.preventDefault(), app = "Toolbarr")}><Applogo size="25px" {app} /> {app}</NavbarBrand>
  <Input type="switch" on:change={saveDark} id="Dark" name="Dark" checked={conf.Dark} />
  <NavbarToggler on:click={() => (navIsOpen = !navIsOpen)} />
  <Collapse isOpen={navIsOpen} navbar expand="md" on:update={toggleNavOpen}>
    <Nav class="ms-auto" navbar>
      {#each ["Lidarr", "Prowlarr", "Radarr", "Readarr", "Sonarr", "Whisparr"] as appLink}
        <Tooltip target="{appLink}-navbar" class="d-none d-md-block" placement="bottom">{appLink}</Tooltip>
        <NavLink id="{appLink}-navbar" on:click={() => (app = appLink)}>
          <Applogo size="20px" app={appLink} /><span class="d-md-none">{appLink}</span>
        </NavLink>
      {/each}
      <Dropdown nav inNavbar>
        <DropdownToggle nav><Applogo size="20px" app="Settings" /> <span class="d-sm-inline-block d-md-none">Configuration</DropdownToggle>
          <DropdownMenu dark={conf.Dark} end>
          <DropdownItem on:click={() => (app = "Settings")}><Fa primaryColor="sienna" icon="{faGear}" /> Settings</DropdownItem>
          <DropdownItem on:click={() => (app = "Toolbox")}><Applogo size="19px" app="Toolbox" /> Toolbox</DropdownItem>
          <DropdownItem on:click={() => (app = "Links")}><Fa primaryColor="dodgerblue" icon="{faLink}" /> Links</DropdownItem>
          <DropdownItem on:click={() => (app = "About")}><Fa primaryColor="mediumpurple" icon="{faBookBible}" /> About</DropdownItem>
        </DropdownMenu>
      </Dropdown>
    </Nav>
  </Collapse>
</Navbar>
<br />

<!-- This gets used by any toast from any page. -->
<ToastContainer placement="bottom-right" theme={conf.Dark ? "dark" : "light"} let:data={data}>
  <FlatToast {data} /> <!-- default slot as toast component -->
</ToastContainer>

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
  <Settings dark={conf.Dark} />
  {:else}
  <Starr {app} dark={conf.Dark} />
  {/if}
</main>