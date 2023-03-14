<script context="module">
  import { Input, Tooltip } from "sveltestrap"
  import { toast } from "./funcs.js"
  import { GetConfig, SaveConfigItem } from "../../wailsjs/go/app/App.js"
  import { devMode } from '../Settings/settings.js'
  import { onMount } from 'svelte';
  let conf = undefined
</script>

<script>
  // Like `text` or `select`.
  export let type
  // Used to update the valie in Go. Needs full struct path for gorilla/schema.
  export let id
  // Used as the config item name, only struct member name goes here.
  export let name
  // Used if you don't want this value changed.
  export let readonly = false
  // Similar to readonly, but the input dims/greys out.
  export let disabled = false
  // Similar to readonly, but the update method won't work.
  export let locked = false
  // Avoid relaoding running config when saving? Most things need to.
  export let noreload = false
  // Avoid sending a toast when not in dev mode?
  export let notoast = false
  // Optional tooptip to bind to input.
  export let tooltip = ""
  // Optional value. Should only be used for binding.
  export let value = undefined

  let sync = false
  let valid = false
  let invalid = false

  onMount(async () => {
    if (conf == undefined) {
      await GetConfig().then(result => {
        conf = result
        value = conf[name]
        sync = true
      })
    } else {
      value = conf[name]
      sync = true
    }
	});

  // Keep conf[name] up to date, in case it gets used across component pages.
  $: if (sync&&value!=undefined) conf[name] = value

  const checkbox = (type=="switch" || type=="checkbox")

  async function saveValue(e) {
    if (locked) return
    // We have to use e.target because value does not update before this runs.
    await SaveConfigItem(id, checkbox?e.target.checked:e.target.value, !noreload).then((msg) => {
      if (!notoast) {toast("success", msg, "CONFIG")}
      if (notoast&&$devMode) {toast("warning", msg, "DEBUG")}
      invalid = false
      valid = true
      setInterval(() => {valid=false}, 5000)
    }, (error) => {
      toast("error", error)
      invalid = true
      valid = false
    })
  }

  // Allows importers to call the exported functions below.
  let input 

  // Allows updating/saving inputs (readonly or not) from other inputs.
  export function update(newVal) {
    value = newVal
    saveValue({target:{value:newVal,checked:newVal}})
  }
</script>

{#if value!=undefined}
  {#if checkbox}
    <Input {type} {id} {name} 
      {disabled} {valid} {invalid}
      readonly={(readonly||locked)}
      bind:this={input}
      on:change={saveValue}
      bind:checked={value}
    ><slot/></Input>
  {:else}
    <Input {type} {id} {name} 
      {disabled} {valid} {invalid}
      readonly={(readonly||locked)}
      bind:this={input}
      bind:value={value}
      on:change={saveValue}
    ><slot/></Input>
  {/if}

  {#if tooltip != ""}
    <Tooltip target={id.replace(/\./g, '\\.')} placement="top">{tooltip}</Tooltip>
  {/if}
{:else}
Provided id '{id}' has no value in config.
{/if}
