<script lang="ts">
  export let starrApp: StarrApp
  export let hidden = true

  import { TabContent, TabPane } from "sveltestrap"
  import Actions from "./Actions/Index.svelte"
  import Config from "./Config/Index.svelte"
  import Database from "./Database/Index.svelte"
  import { _ } from "/src/libs/Translate.svelte"
  import BackgroundLogo from "/src/libs/BackgroundLogo.svelte"
  import type { StarrApp } from "/src/libs/config"

  // Keep track of the active tab.
  let tab: string | number
</script>

<BackgroundLogo {hidden} url="starr">
  <div class="container">
    <TabContent on:tab={(e) => (tab = e.detail)}>
      <TabPane tabId="Config" tab={$_("words.Config")} active><Config {starrApp}/></TabPane>
      <TabPane tabId="Actions" tab={$_("words.Actions")}><Actions hidden={tab != "Actions"} {starrApp}/></TabPane>
      <TabPane tabId="Database">
        <span slot="tab">{@html $_("instances.DBTools")}</span>
        <Database hidden={tab != "Database"} {starrApp}/>
      </TabPane>
    </TabContent>
  </div>
</BackgroundLogo>

<style>
  .container :global(.superbadge) {
    vertical-align: top;
    font-size:9px;
  }

  .container :global(.setting-name) {
    min-width: max-content;
    width: 25%;
    max-width:160px;
  }

  .container :global(.tab-pane) {
    margin-top: 3px;
    width:100%;
  }
</style>
