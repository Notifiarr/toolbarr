<script>
  export let starrApp
  export let hidden = true

  import { TabContent, TabPane } from "sveltestrap"
  import Activity from "./Actvity/Index.svelte"
  import Library from "./Library/Index.svelte"
  import Actions from "./Actions/Index.svelte"
  import Config from "./Config/Index.svelte"
  import Database from "./Database/Index.svelte"
  import Source from "./Source/Index.svelte"
  import { _ } from "../libs/Translate.svelte"
  import BackgroundLogo from "../libs/BackgroundLogo.svelte"
  // Move things around when screen width changes.
  let width 
</script>

<svelte:window bind:innerWidth={width}/>

<BackgroundLogo {hidden} url="starr">
  <div class="container">
    <TabContent vertical={width>767} pills={width>767}>
      <!-- {#if width>767}
        <TabPane disabled><span slot="tab" style="text-decoration:underline;color:gray">MENU</span></TabPane>
      {/if} -->
      <TabPane tabId="Config" tab={$_("words.Config")} active><Config {starrApp}/></TabPane>
      <TabPane tabId="Source" tab={$_("words.Source")}><Source {starrApp}/></TabPane>
      <TabPane tabId="Activity" tab={$_("words.Activity")}><Activity {starrApp}/></TabPane>
      <TabPane tabId="Library" tab={$_("words.Library")}><Library {starrApp}/></TabPane>
      <TabPane tabId="Actions" tab={$_("words.Actions")}><Actions {starrApp}/></TabPane>
      <TabPane tabId="DBTools"><span slot="tab">{@html $_("instances.DBTools")}</span><Database {starrApp}/></TabPane>
    </TabContent>
  </div>
</BackgroundLogo>

<style>
  .container :global(.superbadge) {
    vertical-align: top;
    font-size:9px;
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
