<script lang="ts">
  import { Badge, Input, InputGroupText, Tooltip } from "@sveltestrap/sveltestrap"
  import type { TooltipPlacement } from "@sveltestrap/sveltestrap/dist/Tooltip/Tooltip";
  import type { InputType } from "@sveltestrap/sveltestrap/src/shared";
  import { toast, onOnce } from "/src/libs/funcs"
  import { SaveConfigItem } from "/wailsjs/go/app/App"
  import { conf, type AppConfig } from "/src/libs/config"
  import { _ } from "/src/libs/Translate.svelte"
  import { SvelteComponent, createEventDispatcher } from "svelte"

  // Like `text` or `select`.
  export let type: InputType|undefined
  // Used to update the value in Go. Needs full struct path for gorilla/schema.
  export let id: string
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
  export let placement: TooltipPlacement|undefined = undefined
  // Optional value. Should only be used for binding.
  export let value: any = undefined

  // Set this here to avoid a type-warning.
  placement = placement ? placement : "top"

  let valid: boolean|undefined // Controls the green/red on success/error of a change.
  let timer: number|undefined // Allows clearing the green/red marks after an interval.
  let input: SvelteComponent // Allows importers to call the exported functions.
  let last: any  // This stays simple to trigger the reactive if block.

  // Messages get dispatched on update, change and error.
	const dispatch = createEventDispatcher();

  // This reactive if block updates the config when the 'value' changes.
  $: if (value == undefined) {
    value = last = $conf[id as keyof AppConfig] // Set the initial value from the config.
  } else if (value != last) {
    clearInterval(timer) // clears previous timer (if exists)
    last = value // prevent infinite loops.
    dispatch("update", {val: value});
    SaveConfigItem(id, value, !noreload).then(saved, failed)
  }

  // This runs on successful save to config file.
  function saved(resp: any){
    valid = true      // set green check mark
    if ($conf[id as keyof AppConfig] !== undefined) {
      // updates running config in javascript store
      $conf[id as keyof AppConfig] = value
    }

    timer = onOnce(() => {valid=undefined}, 5) // clears green check mark
    dispatch("change", {val: value});
    // Send only 1 toast.
    if (!notoast) toast("success", resp.Msg, $_("words.CONFIG"))
    if (notoast&&$conf.DevMode)
      toast("warning", resp.Msg, $_("words.CONFIG") + " ("+$_("words.debug")+")")
  }

  // This runs when the save to config file fails.
  function failed(err: string) {
    valid = false // set red X mark, does not clear
    toast("error", err, $_("configvalues.CONFIGERROR"), 9)
    dispatch("error", {val: value});
  }

  // Allows updating/saving inputs (readonly or not) from other inputs.
  export function update(val: any) { value = val }
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
