<script>
	export let starrApp
  import { Alert } from "sveltestrap"
  import { _ } from "../../libs/Translate.svelte"
</script>

<h3>{starrApp} Source Information</h3>
<Alert color="danger">{$_("incompletePage")}</Alert>
<p>This page contains information about the latest release. Perhaps recent commit info. Whatever else we want to pack in that is "info" about {starrApp}'s source.</p>
