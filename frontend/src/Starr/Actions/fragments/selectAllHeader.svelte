<script lang="ts">
  import type { Color } from "sveltestrap/src/shared";
  import { _ } from "../../../libs/Translate.svelte"
  import { Icon } from "sveltestrap"

  export let updating: boolean
  export let all: boolean
  export let selected: {[key: string]: boolean}
  export let icon = "trash"
  export let color: Color = "danger"


  function selectAll() {
    if (updating) return
    all = !all
    Object.keys(selected).forEach((key) => {
       selected[key] = all
    })
  }
</script>

<th>
  <span>
    <Icon class="small text-{color}" name={icon}/>
    <span class={updating?"":"link"} on:keypress={selectAll} on:click={selectAll}>
      {$_("words.All")}
    </span>
  </span>
</th>
<th class="d-none d-md-table-cell">ID</th>
