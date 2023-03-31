<script>
  export let starrApp

  import { Accordion, AccordionItem, Badge} from "sveltestrap"
  import { conf } from "../../libs/config.js"
  import Fa from "svelte-fa"
  import { faPlus } from "@fortawesome/free-solid-svg-icons"
  import Instance from "./Instance.svelte"
  import T from "../../libs/Translate.svelte"

  // Keep UI up to date with existing count of instances.
  $: count = $conf.Instances[starrApp] ? $conf.Instances[starrApp].length : 0
  $: id = count<2 ? 0 : count-1
  let open = true
</script>

<Accordion stayOpen>
  {#if $conf.Instances != undefined && $conf.Instances[starrApp]}
    <!-- Create a list of configured starrApp instances. -->
    {#each $conf.Instances[starrApp] as instance, i }
      <AccordionItem active={id==i} on:toggle={(e) => {id = i;open = e.detail}}>
        <span slot="header">
          <Badge class="superbadge" color={id==i&&open?"success":"secondary"}>{i}</Badge>
          <h4 class="d-inline-block accordian-header">{instance.Name}</h4>
          <Badge color="primary">{instance.URL}</Badge>
        </span> 
        <Instance index={i} {starrApp} {instance} />
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
    <Instance index={count} {starrApp} />
  </AccordionItem>
</Accordion>


<style>
  .accordian-header {
    margin-bottom: 0 !important;
  }
</style>