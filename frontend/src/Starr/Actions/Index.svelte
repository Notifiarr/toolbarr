<script>
  export let starrApp
  import { Alert } from "sveltestrap"
  import { _ } from "../../libs/Translate.svelte"
</script>

<h3>{starrApp} Actions</h3>
<Alert color="danger">{$_("incompletePage")}</Alert>
<p>Actions to do things to {starrApp} are found here.</p>
