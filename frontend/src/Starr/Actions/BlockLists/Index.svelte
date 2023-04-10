<script>
  export let info
  export let instance

  import { Card, Button, Spinner } from "sveltestrap"
  import T, { _ } from "../../../libs/Translate.svelte"
  import { createEventDispatcher } from "svelte"

  const dispatch = createEventDispatcher()
  function update() { info = undefined; dispatch('update') }
</script>

<Button on:click={update}>Update</Button>

<Card body color="danger">{$_("incompletePage")}</Card>

{#if !info} <Spinner/> {$_("words.Loading")} {instance.Name} ... {:else}
<code><pre class="code">{JSON.stringify(info, null, 3)}</pre></code>
{/if}
