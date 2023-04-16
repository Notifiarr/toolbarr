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
    {getData: Indexers, link: "Indexers", component: indexers},
    {getData: Downloaders, link: "DownloadClients", component: downloadClients},
  ]
  // Add some more tabs depending on the app.
  const prowlarrTabs = [
    {getData: AppProfiles, link: "AppProfiles", component: appProfiles},
    {getData: CustomFilters, link: "CustomFilters", component: customFilters},
  ]
  // Everything but Prowlarr.
  const otherTabs = [
    {getData: BlockList, link: "BlockList", component: blockList},
    {getData: QualityProfiles, link: "QualityProfiles", component: qualityProfiles},
    {getData: ImportLists, link: "ImportLists", component: importLists},
  ]

  // So consumers can find the start page.
  export let startTab = commonTabs[0]
</script>

<script>
  export let starrApp
  export let showTitle = false
  export let updating
  export let tab // bind and pass in startTab.
 
  import T, { _ } from "../../libs/Translate.svelte"
  import { Fade, Nav, NavItem, NavLink } from "sveltestrap"

  let tabs = commonTabs.concat(starrApp == "Prowlarr" ? prowlarrTabs : otherTabs)

  function changeTab(e, newTab) {
    e.preventDefault()
    if (!updating) tab = newTab
  }
</script>

<div id="container">
  <Nav {...$$props}>
    <Fade isOpen={showTitle}>
      <NavItem>
        <NavLink class="title" disabled>
          <h6 class="text-center text-primary">
            <strong> &nbsp; <T id="instances.AppMenu" {starrApp}/> &nbsp; </strong>
          </h6>
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
    padding: 5px 18px 5px 18px;
    font-weight: bold;
  }

  #container :global(.title) {
    text-decoration: underline;
    text-decoration-style: wavy;
  }
</style>