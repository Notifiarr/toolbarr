<script lang="ts">
  import { Badge, Input, InputGroupText, Tooltip } from "sveltestrap"
  import { toast, onOnce } from "/src/libs/funcs"
  import { SaveConfigItem } from "/wailsjs/go/app/App"
  import { conf } from "/src/libs/config"
  import { _ } from "/src/libs/Translate.svelte"
  import { createEventDispatcher } from "svelte"

  // Like `text` or `select`.
  export let type
  // Used to update the valie in Go. Needs full struct path for gorilla/schema.
  export let id
  // Used if you do not want this value changed directly.
  export let readonly = false
  // Similar to readonly, but the input dims/greys out.
  export let disabled = false
  // Similar to readonly, but the update method will not work.
  export let locked = false
  // Avoid reloading running config when saving? Most things need to.
  export let noreload = false
  // Avoid sending a toast when not in dev mode?
  export let notoast = false
  // Optional tooptip to bind to input. Derived automatically from language file.
  export let tooltip = $_("configtooltip."+id)
  // Where the tooltip goes.
  export let placement = undefined
  // Optional value. Should only be used for binding.
  export let value = undefined

  // Set this here to avoid a type-warning.
  placement = placement ? placement : "top"

  let valid // Controls the green/red on success/error of a change.
  let timer // Allows clearing the green/red marks after an interval.
  let input // Allows importers to call the exported functions.
  let last  // This stays simple to trigger the reactive if block.

  // Messages get dispatched on update, change and error.
	const dispatch = createEventDispatcher();

  // This reactive if block updates the config when the 'value' changes.
  $: if (value == undefined) {
    value = last = $conf[id] // Set the initial value from the config.
  } else if (value != last) {
    clearInterval(timer) // clears previous timer (if exists)
    last = value // prevent infinite loops.
    dispatch("update", {val: value});
    SaveConfigItem(id, value, !noreload).then(saved, failed)
  }

  // This runs on successful save to config file.
  function saved(resp){
    valid = true         // set green check mark
    $conf[id] = value  // updates running config in javascript store
    timer = onOnce(() => {valid=undefined}, 5) // clears green check mark
    dispatch("change", {val: value});
    // Send only 1 toast.
    if (!notoast) toast("success", resp.Msg, $_("words.CONFIG"))
    if (notoast&&$conf.DevMode)
      toast("warning", resp.Msg, $_("words.CONFIG") + " ("+$_("words.debug")+")")
  }

  // This runs when the save to config file fails.
  function failed(err) {
    valid = false // set red X mark, does not clear
    toast("error", err, $_("configvalues.CONFIGERROR"), 9)
    dispatch("error", {val: value});
  }

  // Allows updating/saving inputs (readonly or not) from other inputs.
  export function update(val) { value = val }
</script>

{#if tooltip != "configtooltip."+id && tooltip != ""}
  <Tooltip target={id.replace(/([\.:])/g, '\\$1')} {placement}>{tooltip}</Tooltip>
{/if}
{#if value!=undefined}
  {#if $_("configinput."+id) != "configinput."+id}
    <InputGroupText class="setting-name">{$_("configinput."+id)}</InputGroupText>
  {/if}
  <Input class="setting-input" bind:this={input} {type} {id} {disabled} invalid={valid==false} {valid}
    readonly={readonly||locked} bind:checked={value} bind:value={value}><slot/>
  </Input>
{:else}
  <Badge color="danger">ID '{id}' has no value in config.</Badge>
{/if}
