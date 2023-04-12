<script>
  export let info
  export let instance

  import { Card, Button } from "sveltestrap"
  import T, { _ } from "../../../libs/Translate.svelte"
  import { createEventDispatcher } from "svelte"
  import Loading from "../loading.svelte"

  const dispatch = createEventDispatcher()
  function update() { info = undefined; dispatch('update') }
</script>

<Button on:click={update}>Update {instance.Name} Data</Button>

<Card body color="danger">{$_("incompletePage")}</Card>
<Loading isOpen={!info}/>
{#if info}
<code><pre class="code">{JSON.stringify(info, null, 3)}</pre></code>
{/if}
