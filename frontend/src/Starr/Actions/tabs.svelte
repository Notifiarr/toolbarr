<script context="module">
  import importLists from "./ImportLists/Index.svelte"
  import indexers from "./Indexers/Index.svelte"
  import appProfiles from "./AppProfiles/Index.svelte"
  import downloadClients from "./DownloadClients/Index.svelte"
  import qualityProfiles from "./QualityProfiles/Index.svelte"
  import blockList from "./BlockLists/Index.svelte"
  import customFilters from "./CustomFilters/Index.svelte"
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
    {fn: CustomFilters, link: "AppProfiles", lib: appProfiles},
    {fn: AppProfiles, link: "CustomFilters", lib: customFilters},
  ]

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

  import { _ } from "../../libs/Translate.svelte"
  import { Nav, NavItem, NavLink} from "sveltestrap"
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

<div class="container">
  <Nav {...$$props}>
    {#each tabs as thisTab}
    <NavItem>
      <NavLink class="nav-link" active={tab==thisTab} on:click={(e) => {changeTab(e,thisTab)}} href="/">
        {@html $_("instances."+thisTab.link)}
      </NavLink>
    </NavItem>
    {/each}
  </Nav>
</div>

<style>
  .container :global(.nav-link) {
    white-space: nowrap;
    padding: 4px 15px 4px 15px;
  }
</style>