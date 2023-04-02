<script context="module">
  const commonTabs = [
    {fn: Indexers, link: "Indexers", lib: indexers},
    {fn: Downloaders, link: "DownloadClients", lib: downloadClients},
  ]
  // Add some more tabs depending on the app.
  const prowlarrTabs = [
    {fn: CustomFilters, link: "AppProfiles", lib: appProfiles},
    {fn: AppProfiles, link: "CustomFilters", lib: customFilters},
    ...commonTabs,
  ]
  const otherTabs = [
    // method: BlockList, title: "instances.BlockList", lib: blockList},
    {fn: QualityProfiles, link: "QualityProfiles", lib: qualityProfiles},
    {fn: ImportLists, link: "ImportLists", lib: importLists},
    ...commonTabs,
  ]

  // This is the active tab; exported, so consumers can find the start page.
  export let activeTab = commonTabs[0]
</script>

<script>
  export let starrApp

  import { _ } from "../../libs/Translate.svelte"
  import importLists from "./ImportLists/Index.svelte"
  import indexers from "./Indexers/Index.svelte"
  import appProfiles from "./AppProfiles/Index.svelte"
  import downloadClients from "./DownloadClients/Index.svelte"
  import qualityProfiles from "./QualityProfiles/Index.svelte"
  import blockList from "./BlockLists/Index.svelte"
  import customFilters from "./CustomFilters/Index.svelte"
  import { Nav, NavItem, NavLink} from "sveltestrap"
  import { createEventDispatcher } from 'svelte';
  import { 
    AppProfiles, 
    Indexers, 
    Downloaders, 
    ImportLists, 
    CustomFilters,
    BlockLists,
    QualityProfiles,
  } from "../../../wailsjs/go/starrs/Starrs.js"

  const dispatch = createEventDispatcher();
  let tab = activeTab
  let tabs = starrApp == "Prowlarr" ? prowlarrTabs : otherTabs

  function nav(e, newTab) { 
    e.preventDefault()
    tab = newTab 
    activeTab = tab
    dispatch('tab', tab)
  }
</script>

<Nav {...$$props}>
  {#each tabs as thisTab}
  <NavItem>
    <NavLink class="f" active={tab==thisTab} on:click={(e) => {nav(e,thisTab)}} href="/">
      {@html $_("instances."+thisTab.link)}
    </NavLink>
  </NavItem>
  {/each}
</Nav>
