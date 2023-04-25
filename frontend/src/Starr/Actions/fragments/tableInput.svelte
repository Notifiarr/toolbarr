<script lang="ts">
  export let info: any
  export let form: any
  export let idx: number
  export let field: string
  export let updating: boolean
  export let type: InputType
  export let min = 1
  export let max = 100

  import { _ } from "/src/libs/Translate.svelte"
  import { Button, Input, InputGroup, Tooltip } from "sveltestrap"
  import type { InputType } from "sveltestrap/src/Input"

  function edit(editable: boolean) {
    plaintext = !editable
    style = editable?"":"zoom-in"
  }

  let plaintext = true
  let style = "cursor:zoom-in"
  let cell
  $: changed = form[idx][field]!=info[idx][field]
</script>

{#if form[idx][field] !== undefined}
  <td bind:this={cell} class={changed?"bg-warning":""}>
    <div class="{type} p-0">
      {#if type == "switch"}
        <Input disabled={updating} type="switch" bind:checked={form[idx][field]} />
      {:else if type == "text"}
      <div class="link"
        on:focusout={() => edit(false)}
        on:focusin={() => edit(true)}
        on:keyup={() => edit(true)}
        on:click={() => edit(true)} >
        <Input {style} {plaintext} disabled={updating} invalid={changed} {type} bind:value={form[idx][field]}/>
      </div>
      {:else if type == "select"}
        <InputGroup size="sm">
          <Input disabled={updating} invalid={changed} {type} bind:value={form[idx][field]}>
            <slot/>
          </Input>
        </InputGroup>
      {:else if type == "range"}
        <Tooltip target={cell}>{form[idx][field]}</Tooltip>
        <Input disabled={updating} bsSize="sm" {type} {min} {max} bind:value={form[idx][field]}/>
      {:else}
        <Button disabled>'{type}' not supported; add it in formInput.svelte</Button>
      {/if}
    </div>
  </td>
{/if}

<style>
  .range {
    height:20px;
    margin-top: -12px !important;
  }

  .switch {
    height:20px;
    margin-top: -5px !important;
  }

  .select, .text { /* Smash that input box into the table row. */
    margin: -3px 3px -3px 0px;
  }
</style>
