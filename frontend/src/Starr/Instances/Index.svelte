<script>
  export let starrApp

  import { TabContent, TabPane } from "sveltestrap"
  import Activity from "./Actvity/Index.svelte"
  import Library from "./Library/Index.svelte"
  import Actions from "./Actions/Index.svelte"
  import Config from "./Config/Index.svelte"
  import { _ } from "../../libs/Translate.svelte"
  // Move things around when screen width changes.
  let width
</script>

<svelte:window bind:innerWidth={width}/>

<div class="container">
  <TabContent vertical={width>767} pills>
    <TabPane tabId="Config" tab="{$_("words.Config")}" active><Config {starrApp}/></TabPane>
    <TabPane tabId="Activity" tab="{$_("words.Activity")}"><Activity {starrApp}/></TabPane>
    <TabPane tabId="Library" tab="{$_("words.Library")}"><Library {starrApp}/></TabPane>
    <TabPane tabId="Actions" tab="{$_("words.Actions")}"><Actions {starrApp}/></TabPane>
  </TabContent>
</div>

<style>
  .container :global(.superbadge) {
    vertical-align: top;
    font-size:9px;
  }

  .container {
    margin-top: 2px;
  }

  .container :global(.setting-name) {
    min-width:120px;
    max-width:120px;
  }

  .container :global(.tab-pane) {
    margin-top: 3px;
    width:100%;
  }
</style>