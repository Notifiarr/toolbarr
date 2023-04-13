<script>
  export let info
  export let instance
  export let updating

  import { Card, Button } from "sveltestrap"
  import T, { _ } from "../../../libs/Translate.svelte"
  import { createEventDispatcher } from "svelte"
  import Loading from "../loading.svelte"

  const dispatch = createEventDispatcher()
  function update() { dispatch('update') }
</script>

<Card body color="danger">{$_("incompletePage")}</Card>
<br>

{instance.Name} Block Lists: {updating?0:info.records.length}

<Button disabled={updating} on:click={update}>Update</Button>
<Loading isOpen={updating}/>
