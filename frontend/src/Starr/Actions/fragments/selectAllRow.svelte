<script lang="ts">
  import { Badge, Input, Tooltip } from "sveltestrap"

  export let updating: boolean
  export let id: string
  export let selected: {[key: string]: boolean}
  export let item = undefined

  let link
</script>

<tr class={selected[id]?"bg-secondary":""}>
  <td style="padding:0">
    <div class="switch">
      <Input disabled={updating} invalid={selected[id]} type="switch" bind:checked={selected[id]}/>
    </div>
  </td>
  <td class="d-none d-md-table-cell">{id}</td>
  {#if item && item.infoLink}
    <td class="d-none d-sm-table-cell">
      <Tooltip target={link}>
        {item.listType?item.listType:item.protocol} {item.infoLink}
      </Tooltip>
      <span bind:this={link}>
        <Badge color="info">
          <open-browser href={item.infoLink}>{item.implementation}</open-browser>
        </Badge>
      </span>
    </td>
  {/if}
  <slot/>
</tr>

<style>
  .switch {
    padding:0 5px;
  }
</style>
