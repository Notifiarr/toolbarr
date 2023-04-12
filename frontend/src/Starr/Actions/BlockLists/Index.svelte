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

<Card body color="danger">{$_("incompletePage")}</Card>
<br>
{#if info}
  {instance.Name} Block Lists: {info.records.length}
{/if}
<Button disabled={!info} on:click={update}>Update</Button>
<Loading isOpen={!info}/>
