<script>
  export let starrApp

  import Applogo from "../../libs/Applogo.svelte"
  import { Accordion, AccordionItem, Badge, Fade, Input, InputGroup, InputGroupText } from "sveltestrap"
  import { conf } from "../../libs/config.js"
  import T, { _ } from "../../libs/Translate.svelte"
  import Inspector from "./Inspector.svelte"
  import Migrator from "./Migrator.svelte"

  const tabs = [
    {title: $_("instances.FilesystemPathsMigrator"), target: Migrator},
    {title: $_("instances.SQLite3DatabaseInspector"), target: Inspector},
  ]

  let activeTab = tabs[0]
  let instance
  let showTitle = true

  $: if ($conf.Instances[starrApp] != undefined && instance == undefined) {
    instance = $conf.Instances[starrApp][0]
  }
</script>

<Accordion>
  <!-- This is the page title. It's collapsable. -->
  <AccordionItem active on:toggle={()=>showTitle=!showTitle}>
    <span slot="header" style="width:95%;">
      <Applogo style="float:right" size="25px" app={starrApp} />
      <h4 class="d-inline-block accordian-header">{@html $_("instances.DBTools")}</h4>
      <Badge color="primary">{activeTab.title}</Badge>
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
      <InputGroupText class="setting-name">Instance</InputGroupText>
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
{#if instance}
  <svelte:component this={activeTab.target} {instance} {showTitle}/>
{/if}

<style>
  .accordian-header {
    margin-bottom: 0 !important;
  }
</style>