<script context="module" type="ts">
  import appProfiles from "./AppProfiles.svelte"
  import indexers from "./Indexers.svelte"
  import downloadClients from "./DownloadClients.svelte"
  import importLists from "./ImportLists.svelte"
  import customFilters from "./CustomFilters.svelte"
  import blockList from "./BlockLists.svelte"
  import qualityProfiles from "./QualityProfiles.svelte"
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
  const commonTabs = []
  commonTabs.push({getData: Indexers, link: "Indexers", component: indexers})
  commonTabs.push({getData: Downloaders, link: "DownloadClients", component: downloadClients})

  // Everything but Prowlarr.
  const starrTabs = commonTabs.concat([
    {getData: BlockList, link: "BlockList", component: blockList},
    {getData: QualityProfiles, link: "QualityProfiles", component: qualityProfiles},
    {getData: ImportLists, link: "ImportLists", component: importLists},
  ])

  const tabs = {
    "Lidarr": starrTabs,
    "Prowlarr": commonTabs.concat([
      {getData: AppProfiles, link: "AppProfiles", component: appProfiles},
      {getData: CustomFilters, link: "CustomFilters", component: customFilters},
    ]),
    "Radarr": starrTabs,
    "Readarr": starrTabs,
    "Sonarr": starrTabs,
    "Whisparr": starrTabs,
  }

  // So consumers can find the start page.
  export const startTab = commonTabs[0]
</script>

<script>
  export let starrApp
  export let showTitle = false
  export let updating
  export let tab // bind this and pass in startTab.
 
  import T, { _ } from "../../libs/Translate.svelte"
  import { Fade, Nav, NavItem, NavLink } from "sveltestrap"

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

    {#each tabs[starrApp] as thisTab}
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
