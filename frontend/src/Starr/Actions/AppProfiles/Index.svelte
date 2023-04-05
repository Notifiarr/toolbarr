<script>
  export let info

  import { Spinner } from "sveltestrap"
  import T, { _ } from "../../../libs/Translate.svelte"
  import { toast } from "../../../libs/funcs"
  import { createEventDispatcher } from "svelte"

  const dispatch = createEventDispatcher()
</script>

{#if !info} <Spinner/> {$_("words.Loading")} ... {:else}
<code><pre class="code">{JSON.stringify(info, null, 3)}</pre></code>
{/if}