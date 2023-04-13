<script context="module">
  import appProfiles from "./AppProfiles/Index.svelte"
  import indexers from "./Indexers/Index.svelte"
  import downloadClients from "./DownloadClients/Index.svelte"
  import importLists from "./ImportLists/Index.svelte"
  import customFilters from "./CustomFilters/Index.svelte"
  import blockList from "./BlockLists/Index.svelte"
  import qualityProfiles from "./QualityProfiles/Index.svelte"
  import {
    AppProfiles,
    Indexers,
    Downloaders,
    ImportLists,
    CustomFilters,
    BlockList,
    QualityProfiles,
  } from "../../../wailsjs/go/starrs/Starrs.js"

  // All apps have these tabs.
  const commonTabs = [
    {fn: Indexers, link: "Indexers", lib: indexers},
    {fn: Downloaders, link: "DownloadClients", lib: downloadClients},
  ]
  // Add some more tabs depending on the app.
  const prowlarrTabs = [
    {fn: AppProfiles, link: "AppProfiles", lib: appProfiles},
    {fn: CustomFilters, link: "CustomFilters", lib: customFilters},
  ]
  // Everything but Prowlarr.
  const otherTabs = [
    {fn: BlockList, link: "BlockList", lib: blockList},
    {fn: QualityProfiles, link: "QualityProfiles", lib: qualityProfiles},
    {fn: ImportLists, link: "ImportLists", lib: importLists},
  ]

  // This is the active tab; exported, so consumers can find the start page.
  export let startTab = commonTabs[0]
</script>

<script>
  export let starrApp
  export let showTitle = false

  import T, { _ } from "../../libs/Translate.svelte"
  import { Fade, Nav, NavItem, NavLink } from "sveltestrap"
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();
  let tab = startTab
  let tabs = commonTabs.concat(starrApp == "Prowlarr" ? prowlarrTabs : otherTabs)

  function changeTab(e, newTab) {
    e.preventDefault()
    startTab = tab = newTab
    dispatch('tab', tab)
  }
</script>

<div id="container">
  <Nav {...$$props}>
    <Fade isOpen={showTitle}>
      <NavItem>
        <NavLink class="nav-link title" disabled>
          <h6 class="text-center text-primary"><T id="instances.AppMenu" {starrApp}/></h6>
        </NavLink>
      </NavItem>
    </Fade>

    {#each tabs as thisTab}
    <NavItem>
      <NavLink class="nav-link" active={tab == thisTab} on:click={(e) => {changeTab(e, thisTab)}} href="/">
        {@html $_("instances."+thisTab.link)}
      </NavLink>
    </NavItem>
    {/each}
  </Nav>
</div>

<style>
  #container :global(.nav-link) {
    white-space: nowrap;
    /* padding: 4px 15px 4px 15px; */
  }

  #container :global(.title) {
    text-decoration: underline;
  }
</style>