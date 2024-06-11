<script lang="ts">
  import { Badge, Input, Tooltip } from "@sveltestrap/sveltestrap"

  export let updating: boolean
  export let id: string|number
  export let selected: {[key: string]: boolean}
  export let item: any = undefined

  let link: HTMLElement
</script>

<tr class={selected[id]?"bg-secondary":""}>
  <td class="p-0">
    <div class="switch">
      <Input disabled={updating} invalid={selected[id]} type="switch" bind:checked={selected[id]}/>
    </div>
  </td>
  <td class="d-none d-md-table-cell">{id}</td>

  {#if item && (item.listType || (item.protocol && item.implementation))}
    <td class="d-none d-sm-table-cell">
      <Tooltip target={link}>
        {item.listType?item.listType:item.protocol} {item.infoLink?item.infoLink:''}
      </Tooltip>
      <span bind:this={link}>
        <Badge color="info">
          {#if item.infoLink}
            <open-browser href={item.infoLink}>{item.implementation}</open-browser>
          {:else}
            {item.implementation}
          {/if}
        </Badge>
      </span>
    </td>
  {/if}

  <slot/>
</tr>

<style>
  .switch {
    padding: 0 5px;
  }
</style>
