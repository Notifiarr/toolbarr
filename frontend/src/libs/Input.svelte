<script lang="ts">
  import { Input, Tooltip } from "sveltestrap"
  import { toast, onOnce } from "./funcs.js"
  import { SaveConfigItem } from "../../wailsjs/go/app/App.js"
  import { conf } from "./config"

  // Like `text` or `select`.
  export let type
  // Used to update the valie in Go. Needs full struct path for gorilla/schema.
  export let id
  // Used as the config item name, only struct member name goes here.
  export let name
  // Used if you do not want this value changed directly.
  export let readonly = false
  // Similar to readonly, but the input dims/greys out.
  export let disabled = false
  // Similar to readonly, but the update method will not work.
  export let locked = false
  // Avoid relaoding running config when saving? Most things need to.
  export let noreload = false
  // Avoid sending a toast when not in dev mode?
  export let notoast = false
  // Optional tooptip to bind to input.
  export let tooltip = ""
  // Where the tooltip goes.
  export let placement = "top"
  // Optional value. Should only be used for binding.
  export let value = undefined
  // Set the value from the config.
  value = $conf[name]

  // These control the green/red on success/error of a change.
  let valid = false
  let invalid = false
  // Controls if we use .checked or .value. Might need adjustments?
  const checkbox = (type=="switch" || type=="checkbox")

  async function saveValue(e) {
    if (locked) return
    // We have to use e.target because value does not update before this runs.
    const val = checkbox?e.target.checked:e.target.value
    await SaveConfigItem(id, val, !noreload).then(saved, failed)
  }

  function saved(resp) {
    // Convert the strinfified value back into a type.. by guessing?
    $conf[name] = resp.Val == "true" ? true : resp.Val == "false" ? false :
      /^[1-9]\d*(\.\d+)?$/.test(resp.Val) ? Number(resp.Val) : resp.Val
    invalid = false
    valid = true
    onOnce(() => {valid=false}, 5)
    // Send 1 toast.
    if (!notoast) toast("success", resp.Msg, "CONFIG")
    if (notoast&&$conf.DevMode) toast("warning", resp.Msg, "CONFIG (debug)")
  }

  function failed(err) {
    toast("error", err)
    invalid = true
    valid = false
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
      readonly={readonly||locked}
      bind:this={input}
      on:change={saveValue}
      bind:checked={value}
    ><slot/></Input>
  {:else}
    <Input {type} {id} {name} 
      {disabled} {valid} {invalid}
      readonly={readonly||locked}
      bind:this={input}
      bind:value={value}
      on:change={saveValue}
    ><slot/></Input>
  {/if}

  {#if tooltip != ""}
    <Tooltip target={id.replace(/([\.:])/g, '\\$1')} {placement}>{tooltip}</Tooltip>
  {/if}
{:else}
  Provided name '{name}' has no value in config.
{/if}
