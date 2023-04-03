<script>
  export let info

  import { Button, Spinner } from "sveltestrap"
  import T, { _ } from "../../../libs/Translate.svelte"
  import { toast } from "../../../libs/funcs"
  import { createEventDispatcher } from "svelte"

  const dispatch = createEventDispatcher()
  function update() { info = undefined; dispatch('update') }
</script>

<Button on:click={update}>Update</Button>

{#if !info} <Spinner/> {$_("words.Loading")} ... {:else}
<code><pre class="code">{JSON.stringify(info, null, 3)}</pre></code>
{/if}
