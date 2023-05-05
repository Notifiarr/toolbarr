<script lang="ts">
  export let info: any
  export let form: any
  export let idx: number
  export let field: string
  export let applyAll: string = "" // Apply this setting to all import lists in {instance.Name}.
  export let desc: string = "" // translation key
  export let name: string = "" // translation key
  export let starrApp: undefined | string = undefined
  export let type: InputType

  import { _ } from "/src/libs/Translate.svelte"
  import Fa from "svelte-fa"
  import { faGroupArrowsRotate } from "@fortawesome/free-solid-svg-icons"
  import { Button, Input, InputGroup, InputGroupText, Tooltip } from "sveltestrap"
  import type { InputType } from "sveltestrap/src/Input"

  const vals = { values:{"starrApp": starrApp} }
  $: tooltip = desc ? $_(desc, vals) : $_(`instances.${field}Desc`, vals)
  $: title = name ? $_(name) : $_(`instances.${field}Title`)

  function applyToAll(field, idx) {
    for (const listID in form) {
      form[listID][field] = form[idx][field]
    }
  }

  let input
  let button
</script>

{#if form[idx][field] !== undefined}
  {#if `instances.${field}Desc` != tooltip}
    <Tooltip target={input}>{tooltip}</Tooltip>
  {/if}

  <div id="input" bind:this={input}>
    <InputGroup>
      <InputGroupText class="setting-name">{title}</InputGroupText>
      <Input invalid={form[idx][field] != info[idx][field]} bind:value={form[idx][field]} type={type=="checkbox"?"select":type}>
        {#if type == "checkbox"}
          <option value={true}>{$_("configvalues.Enabled")}</option>
          <option value={false}>{$_("configvalues.Disabled")}</option>
        {/if}
        <slot></slot>
      </Input>

      {#if applyAll}
        <Tooltip target={button}>{applyAll}</Tooltip>
        <span bind:this={button}>
          <Button on:click={()=>{applyToAll(field, idx)}}>
            <Fa primaryColor="darkCyan" icon={faGroupArrowsRotate}/>
          </Button>
        </span>
      {/if}
    </InputGroup>
  </div>
{/if}

<style>
  #input :global(.setting-name) {
    min-width: max-content !important;
    width: 160px !important;
  }
</style>
