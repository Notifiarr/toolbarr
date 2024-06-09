<script lang="ts">
  export let info: any
  export let form: any
  export let idx: number
  export let item: any
  export let itemIdx: number
  export let disabled = false

  import { _ } from "/src/libs/Translate.svelte"
  import { Input, InputGroup, InputGroupText, Tooltip } from "@sveltestrap/sveltestrap"

  let input: HTMLElement
  // changed.
  $: invalid = info[idx].fields[itemIdx].value != form[idx].fields[itemIdx].value
</script>

  {#if item.helpText}
    <Tooltip target={input}>{item.helpText}</Tooltip>
  {/if}

  <div id="input" bind:this={input}>
    <InputGroup>
      <InputGroupText class="setting-name">{item.label?item.label:item.name}</InputGroupText>

      {#if item.type == "select" || item.type == "tagSelect"}
        {#if item.selectOptions && typeof item.value != "object"}
          <Input type="select" {invalid} bind:value={form[idx].fields[itemIdx].value} {disabled}>
              {#each info[idx].fields[itemIdx].selectOptions as val}
                <option value={val.name}>{val.name}</option>
              {/each}
          </Input>
        {:else}<!-- /else (item.selectOptions) -->
          <select multiple disabled class="multi" bind:value={form[idx].fields[itemIdx].value}>
            {#if typeof item.value == "object"}
              {#each info[idx].fields[itemIdx].value as val}
                <option value={val}>{val}</option>
              {/each}
            {:else}
              <option value={info[idx].fields[itemIdx].value}>{info[idx].fields[itemIdx].value}</option>
            {/if}
          </select>
        {/if}<!-- /if (item.selectOptions) -->
      {:else if item.type == "checkbox"}
      <Input type="select" {invalid} bind:value={form[idx].fields[itemIdx].value} {disabled}>
        <option value={true}>{$_("configvalues.Enabled")}</option>
        <option value={false}>{$_("configvalues.Disabled")}</option>
      </Input>
      {:else}
        <Input type={item.type} {invalid} bind:value={form[idx].fields[itemIdx].value} {disabled}/>
      {/if}<!-- /if (item.type) -->
    </InputGroup>
  </div>

  <style>
    .multi {
      width: calc(100% - 160px)
    }

    #input :global(.setting-name) {
      min-width: max-content !important;
      width: 160px !important;
    }
  </style>
