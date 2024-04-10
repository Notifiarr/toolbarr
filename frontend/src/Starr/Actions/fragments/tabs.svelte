<script context="module" lang="ts">
  import type { StarrApp } from "/src/libs/config"
  import type { ComponentType } from "svelte"
  import appProfiles from "../AppProfiles.svelte"
  import indexers from "../Indexers.svelte"
  import downloadClients from "../DownloadClients.svelte"
  import importLists from "../ImportLists.svelte"
  import customFilters from "../CustomFilters.svelte"
  import blockList from "../BlockLists.svelte"
  import qualityProfiles from "../QualityProfiles.svelte"
  import {
    AppProfiles,
    Indexers,
    Downloaders,
    ImportLists,
    CustomFilters,
    BlockList,
    QualityProfiles,
  } from "/wailsjs/go/starrs/Starrs"

  export type Tab = {
    data?: (instance: any) => Promise<any>
    pageData?: (instance: any, pageSize: number, page: number, sortKey: string, sortDir: string) => Promise<any>
    id: string
    component?: ComponentType
  }

  // All apps have these tabs.
  const commonTabs: Tab[] = [
    { data: Indexers, id: "Indexers", component: indexers },
    { data: Downloaders, id: "DownloadClients", component: downloadClients }
  ]

  // Everything but Prowlarr.
  const starrTabs = commonTabs.concat([
    { pageData: BlockList, id: "BlockLists", component: blockList },
    { data: QualityProfiles, id: "QualityProfiles", component: qualityProfiles },
    { data: ImportLists, id: "ImportLists", component: importLists },
  ])

  const tabs = {
    "Lidarr": starrTabs,
    "Prowlarr": commonTabs.concat([
      { data: AppProfiles, id: "AppProfiles", component: appProfiles },
      { data: CustomFilters, id: "CustomFilters", component: customFilters },
    ]),
    "Radarr": starrTabs,
    "Readarr": starrTabs,
    "Sonarr": starrTabs,
    "Whisparr": starrTabs,
  }

  // So consumers can find the start page.
  export const startTab = commonTabs[0]
</script>

<script lang="ts">
  export let starrApp: StarrApp
  export let showTitle = false
  export let updating: boolean
  export let tab: Tab // bind this and pass in startTab.

  import T, { _ } from "../../../libs/Translate.svelte"
  import { Fade, Nav, NavItem, NavLink } from "@sveltestrap/sveltestrap"

  function changeTab(e: MouseEvent, newTab: Tab) {
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
          {@html $_("instances."+thisTab.id)}
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
