<script>
  export let info
  export let instance

  import { Card, Button } from "sveltestrap"
  import T, { _ } from "../../../libs/Translate.svelte"
  import Loading from "../../loading.svelte"


</script>

<Card body color="danger">{$_("incompletePage")}</Card>
<br>

{instance.Name} Block Lists: {info.records.length}
