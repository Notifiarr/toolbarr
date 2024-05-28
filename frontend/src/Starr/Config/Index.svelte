<script lang="ts">
  export let starrApp: StarrApp

  import type { StarrApp } from "/src/libs/config"
  import { Accordion, AccordionItem, Badge} from "@sveltestrap/sveltestrap"
  import { conf } from "/src/libs/config"
  import Fa from "svelte-fa"
  import { faPlus } from "@fortawesome/free-solid-svg-icons"
  import Instance from "./Instance.svelte"
  import T from "/src/libs/Translate.svelte"

  let count: number
  let oldCount: number = 0
  let id: number = 0
  // Keep UI up to date with existing count of instances.
  $: count = $conf.Instances[starrApp] ? $conf.Instances[starrApp].length : 0
  $: if (count > oldCount) {
    // we added a new instance
    id = count-1
    oldCount = count
  }
  $: if (count < oldCount) {
    // we removed an instance
    id = count
    oldCount = count
  }
  let open = true
</script>

<Accordion stayOpen>
  {#if $conf.Instances != undefined && $conf.Instances[starrApp]}
    <!-- Create a list of configured starrApp instances. -->
    {#each $conf.Instances[starrApp] as instance, idx }
      <AccordionItem active={id==idx} on:toggle={(e) => {id = idx;open = e.detail}}>
        <span slot="header">
          <Badge class="superbadge" color={id==idx&&open?"success":"secondary"}>{idx}</Badge>
          <h4 class="d-inline-block accordian-header">{instance.Name}</h4>
          <Badge color="primary">{instance.URL}</Badge>
        </span> 
        <Instance index={idx} {starrApp} {instance} defaultInstance={$conf.Instance[starrApp] == idx}/>
      </AccordionItem>
    {/each}
  {/if}

  <!-- add new instance option -->
  <AccordionItem active={id==count} on:toggle={(e) => {id = count;open = e.detail}}>
    <span slot="header">
      <h5 class="accordian-header">
        <Fa primaryColor={id==count&&open?"green":"orange"} icon={faPlus} />
        <T id="instances.addNewInstance" {starrApp}/>
      </h5>
    </span>
    <Instance index={count} {starrApp} newInstance/>
  </AccordionItem>
</Accordion>

<style>
  .accordian-header {
    margin-bottom: 0 !important;
  }
</style>