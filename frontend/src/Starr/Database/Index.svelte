<script>
  export let starrApp

  import { Input, InputGroup, InputGroupText } from "sveltestrap"
  import { conf } from "../../libs/config.js"
  import { _ } from "../../libs/Translate.svelte"
  import Inspector from "./Inspector.svelte"
  import Migrate from "./Migrate.svelte"

  let instance
  let activeTab=Migrate

  $: if ($conf.Instances[starrApp] != undefined && instance == undefined) {
    instance = $conf.Instances[starrApp][0]
  }
</script>

<h3>{starrApp} Database Tools</h3>
<p>
  Select a tool and an instance to get started. 
  <span class="text-danger">The database file must not be in use. </span>
  Either work on a copy, or exit {starrApp} before proceeding.
</p>

<InputGroup>
  <InputGroupText class="setting-name">Tool</InputGroupText>
  <Input type="select" bind:value={activeTab}>
    <option value={Migrate}>Migrate Filesystem Paths</option>
    <option value={Inspector}>SQLite3 Database Inspector</option>
  </Input>
</InputGroup>
<InputGroup>
  <InputGroupText class="setting-name">Instance</InputGroupText>
  <Input type="select" id="instance" bind:value={instance}>
    {#if $conf.Instances[starrApp] != null}
      {#each $conf.Instances[starrApp] as instance}
        <option value={instance}>{instance.Name}</option>
      {/each}
      {#if $conf.Instances[starrApp].length == 0}
        <option disabled>- no instances configured -</option>
      {/if}
    {/if}
  </Input>
</InputGroup>

<hr>
{#if instance}
  <svelte:component this={activeTab} {instance}/>
{/if}