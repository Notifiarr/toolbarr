<script>
  export let starrApp
  export let hidden = true

  import { Container, TabContent, TabPane } from "sveltestrap"
  import Activity from "./Actvity/Index.svelte"
  import Library from "./Library/Index.svelte"
  import Actions from "./Actions/Index.svelte"
  import Config from "./Config/Index.svelte"
  import Database from "./Database/Index.svelte"
  import Source from "./Source/Index.svelte"
  import { _ } from "../libs/Translate.svelte"
  import BackgroundLogo from "../libs/BackgroundLogo.svelte"

  // Keep track of the active tab.
  let tab
</script>


  <BackgroundLogo {hidden} url="starr">
    <div class="container">
      <TabContent on:tab={(e) => (tab = e.detail)}>
      <TabPane tabId="Config" tab={$_("words.Config")} active><Config {starrApp}/></TabPane>
      <!-- <TabPane tabId="Source" tab={$_("words.Source")}><Source {starrApp}/></TabPane>
      <TabPane tabId="Activity" tab={$_("words.Activity")}><Activity {starrApp}/></TabPane>
      <TabPane tabId="Library" tab={$_("words.Library")}><Library {starrApp}/></TabPane> -->
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
