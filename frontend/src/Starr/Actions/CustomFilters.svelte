<script lang="ts">
  export let info
  export let instance

  import { Card } from "sveltestrap"
  import T, { _ } from "/src/libs/Translate.svelte"
  import { toast } from "/src/libs/funcs"
</script>

<Card body color="danger">{$_("incompletePage")}</Card>
<br>
{instance.Name} Custom Filters: {info.length}
