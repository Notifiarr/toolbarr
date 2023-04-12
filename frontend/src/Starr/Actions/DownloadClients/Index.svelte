<script>
  export let info
  export let instance

  import { Card } from "sveltestrap"
  import T, { _ } from "../../../libs/Translate.svelte"
  import { toast } from "../../../libs/funcs"
</script>

<Card body color="danger">{$_("incompletePage")}</Card>
<br>
{instance.Name} Download Clients: {info.length}
