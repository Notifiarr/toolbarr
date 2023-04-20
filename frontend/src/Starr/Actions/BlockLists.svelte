<script lang="ts">
  export let info
  export let instance

  import { Card } from "sveltestrap"
  import T, { _ } from "../../libs/Translate.svelte"


</script>

<Card body color="danger">{$_("incompletePage")}</Card>
<br>

{instance.Name} Block Lists: {info.records.length}
