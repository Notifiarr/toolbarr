<script lang="ts">
  import type { Color } from "sveltestrap/src/shared"
  import { _ } from "../../../libs/Translate.svelte"
  import { Icon } from "sveltestrap"
  import { onOnce } from "../../../libs/funcs"

  export let updating: boolean
  export let all: boolean
  export let selected: {[key: string]: boolean}
  export let icon = "trash"
  export let color: Color = "danger"


  function selectAll() {
    if (updating) return
    all = !all
    var idx = 0.05 // initial delay to click.

    Object.keys(selected).forEach((key) => { // gets faster and faster.
      onOnce(() => {selected[key] = all}, idx += 0.077 - (idx/24))
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
