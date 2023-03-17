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

  let valid // Controls the green/red on success/error of a change.
  let timer // Allows clearing the green/red marks after an interval.
  let input // Allows importers to call the exported functions.
  let last  // This stays simple to trigger the reactive if block.

  // This reactive if block updates the config when the 'value' changes.
  $: if (value == undefined) {
    value = last = $conf[name] // Set the initial value from the config.
  } else if (value != last) {
    clearInterval(timer) // clears previous timer (if exists)
    last = value // prevent infinite loops.
    SaveConfigItem(id, value, !noreload).then(saved, failed)
  }

  // This runs on successful save to config file.
  function saved(resp){
    valid = true         // set green check mark
    $conf[name] = value  // updates running config in javascript store
    timer = onOnce(() => {valid=undefined}, 5) // clears green check mark
    // Send only 1 toast.
    if (!notoast) toast("success", resp.Msg, "CONFIG")
    if (notoast&&$conf.DevMode) toast("warning", resp.Msg, "CONFIG (debug)")
  }

  // This runs when the save to config file fails.
  function failed(err) {
    valid = false // set red X mark, does not clear
    toast("error", err, "CONFIG ERROR", 9)
  }

  // Allows updating/saving inputs (readonly or not) from other inputs.
  export function update(val) { value = val }
</script>

{#if value!=undefined}
  {#if tooltip != ""}
    <Tooltip target={id.replace(/([\.:])/g, '\\$1')} {placement}>{tooltip}</Tooltip>
  {/if}
  <Input {type} {id} {name} {disabled}
    valid={valid==true}
    invalid={valid==false}
    readonly={readonly||locked}
    bind:this={input}
    bind:checked={value}
    bind:value={value}
    ><slot/>
  </Input>
{:else}
  Name '{name}' has no value in config.
{/if}
