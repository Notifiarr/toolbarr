<script>
  export let starrApp
  export let hidden = true

  import Applogo from "../../libs/Applogo.svelte"
  import { Accordion, AccordionItem, Badge, Input, InputGroup, InputGroupText } from "sveltestrap"
  import { conf } from "../../libs/config.js"
  import T, { _ } from "../../libs/Translate.svelte"
  import Inspector from "./Inspector.svelte"
  import Migrator from "./Migrator/Index.svelte"

  const tabs = []
  if (starrApp != "Prowlarr") tabs.push({title: $_("instances.FilesystemPathsMigrator"), target: Migrator})
  tabs.push({title: $_("instances.SQLite3DatabaseInspector"), target: Inspector})

  let activeTab = tabs[0]
  let showTitle = true

  $: instance = $conf.Instances[starrApp] ? $conf.Instances[starrApp][0] : undefined
</script>

<Accordion>
  <!-- This is the page title. It's collapsible. -->
  <AccordionItem active on:toggle={()=>showTitle=!showTitle}>
    <span slot="header" style="width:95%;">
      <Applogo style="float:right" size="25px" app={starrApp} />
      <h4 class="d-inline-block accordian-header">{@html $_("instances.DBTools")}</h4>
      {#if activeTab}<Badge color="primary">{activeTab.title}</Badge>{/if}
    </span>
    <p><T id="instances.DBToolsSelector" {starrApp}/></p>

    <!-- DB Tool selector menu. -->
    <InputGroup>
      <InputGroupText class="setting-name">{$_("instances.DBTool")}</InputGroupText>
      <Input type="select" bind:value={activeTab}>
        {#each tabs as tab}
          <option value={tab}>{tab.title}</option>
        {/each}
      </Input>
    </InputGroup>

    <!-- Instance selector menu. -->
    <InputGroup>
      <InputGroupText class="setting-name">{$_("word.Instance")}</InputGroupText>
      <Input type="select" id="instance" bind:value={instance}>
        {#if $conf.Instances[starrApp] != null}
          {#each $conf.Instances[starrApp] as instance}
            <option value={instance}>{instance.Name}: {instance.URL}</option>
          {/each}
          {#if $conf.Instances[starrApp].length == 0}
            <option disabled>- {$_("instances.noInstancesConfigured")} -</option>
          {/if}
        {:else}
          <option disabled>- {$_("instances.noInstancesConfigured")} -</option>
        {/if}
      </Input>
    </InputGroup>
  </AccordionItem>
</Accordion>

<!-- Display the selected tool, pass in selected instance. -->
{#if !hidden && activeTab}
  <svelte:component this={activeTab.target} {instance} {starrApp} {showTitle}/>
{/if}

<style>
  .accordian-header {
    margin-bottom: 0 !important;
  }
</style>