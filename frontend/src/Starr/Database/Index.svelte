<script lang="ts">
  export let starrApp: StarrApp
  export let hidden: boolean = false

  import Applogo from "/src/libs/Applogo.svelte"
  import { Accordion, AccordionItem, Badge, Input, InputGroup, InputGroupText } from "@sveltestrap/sveltestrap"
  import type { Instance, StarrApp } from "/src/libs/config"
  import { conf } from "/src/libs/config"
  import T, { _ } from "/src/libs/Translate.svelte"
  import Inspector from "./Inspector.svelte"
  import Migrator from "./Migrator/Index.svelte"
  import type { ComponentType } from "svelte"

  
  const tabs: {title: string, target: ComponentType}[] = []
  if (starrApp != "Prowlarr") tabs.push({title: "instances.FilesystemPathsMigrator", target: Migrator})
  tabs.push({title: "instances.SQLite3DatabaseInspector", target: Inspector})

  let activeTab = tabs[0]
  let showTitle = true
  let idx = $conf.Instance[starrApp] // Start with default instance.
  let instance: Instance|undefined = $conf.Instances[starrApp]?$conf.Instances[starrApp][idx]:undefined
  $: if (!instance || !$conf.Instances[starrApp].includes(instance)) {
    instance = $conf.Instances[starrApp]?$conf.Instances[starrApp][$conf.Instance[starrApp]]:undefined
  }
</script>

<Accordion>
  <!-- This is the page title. It's collapsible. -->
  <AccordionItem active on:toggle={()=>showTitle=!showTitle}>
    <span slot="header" style="width:95%;">
      <Applogo style="float:right" size="25px" app={starrApp} />
      <h4 class="d-inline-block accordian-header">{@html $_("instances.DBTools")}</h4>
      {#if activeTab}<Badge color="primary">{$_(activeTab.title)}</Badge>{/if}
    </span>
    <p><T id="instances.DBToolsSelector" {starrApp}/></p>

    <!-- DB Tool selector menu. -->
    <InputGroup>
      <InputGroupText class="setting-name">{$_("instances.DBTool")}</InputGroupText>
      <Input type="select" bind:value={activeTab}>
        {#each tabs as tab}
          <option value={tab}>{$_(tab.title)}</option>
        {/each}
      </Input>
    </InputGroup>

    <!-- Instance selector menu. -->
    <InputGroup>
      <InputGroupText class="setting-name">{$_("words.Instance")}</InputGroupText>
      <Input type="select" id="instance" invalid={!instance||!instance.DBPath} bind:value={instance}>
        {#if $conf.Instances != undefined && $conf.Instances[starrApp]}
          {#each $conf.Instances[starrApp] as i}
            <option value={i}>{i.Name}: {i.URL}</option>
          {/each}
        {:else}
          <option disabled>- {$_("instances.noInstancesConfigured")} -</option>
        {/if}
      </Input>
    </InputGroup>
  </AccordionItem>
</Accordion>

{#if activeTab && !hidden}
  <!-- Display the selected tool, pass in selected instance. -->
  <svelte:component this={activeTab.target} {instance} {starrApp} {showTitle}/>
{/if}

<style>
  .accordian-header {
    margin-bottom: 0 !important;
  }
</style>