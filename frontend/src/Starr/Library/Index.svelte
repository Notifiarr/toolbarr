<script>
  export let starrApp
  import { Alert } from "sveltestrap"
  import { _ } from "../../libs/Translate.svelte"
</script>

<h3>{starrApp} Library</h3>
<Alert color="danger">{$_("incompletePage")}</Alert>
<p>Information about your {starrApp} library is found here.</p>
