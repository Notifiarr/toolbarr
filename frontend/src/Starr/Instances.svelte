<script>
  export let starrApp

  import { Accordion, AccordionItem, Badge, TabContent, TabPane } from "sveltestrap"
  import { conf } from "../libs/config.js"
  import Fa from "svelte-fa"
  import { faPlus } from "@fortawesome/free-solid-svg-icons"
  import Instance from "./InstanceConfig.svelte"

  // Keep UI up to date with existing count of instances.
  $: count = $conf.Instances[starrApp] ? $conf.Instances[starrApp].length : 0
  $: id = count<2 ? 0 : count-1
  let open = true
  // Move things around when screen width changes.
  let width
</script>

<svelte:window bind:innerWidth={width}/>

<div class="container">
  <TabContent vertical={width>767} pills>
    <TabPane tabId="Config" tab="Config" active>
      <Accordion stayOpen>
        {#if $conf.Instances != undefined && $conf.Instances[starrApp]}
          {#each $conf.Instances[starrApp] as instance, i }
            <AccordionItem active={id==i} on:toggle={(e) => {id = i;open = e.detail}}>
              <span slot="header">
                <Badge class="superbadge" color={id==i&&open?"success":"secondary"}>{i}</Badge>
                <h4 class="d-inline-block">{instance.Name}</h4>
                <Badge color="primary">{instance.URL}</Badge>
              </span> 
              <Instance index={i} {starrApp} instance={instance} />
            </AccordionItem>
          {/each}
        {/if}
        <AccordionItem active={id==count} on:toggle={(e) => {id = count;open = e.detail}}>
          <h5 slot="header">
            <Fa primaryColor={id==count&&open?"green":"orange"} icon={faPlus} />
            Add New {starrApp} Instance
          </h5>
          <Instance index={count} {starrApp} instance={undefined} />
        </AccordionItem>
      </Accordion>
    </TabPane>
    <TabPane tabId="Activity" tab="Activity">
      <p>Application activity and history shows up here.</p>
    </TabPane>
    <TabPane tabId="Library" tab="Library">
      <p>Information about your {starrApp} library is found here.</p>
    </TabPane>
    <TabPane tabId="Actions" tab="Actions">
      <p>Actions to do things to your app are found here.</p>
    </TabPane>
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