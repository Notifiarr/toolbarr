<script>
  export let starrApp
  import { Alert } from "sveltestrap"
  import { _ } from "../../libs/Translate.svelte"
</script>

<h3>{starrApp} Activity</h3>
<Alert color="danger">{$_("incompletePage")}</Alert>
<p>{starrApp} activity and history shows up here.</p>
