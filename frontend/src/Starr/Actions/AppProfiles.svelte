<script lang="ts">
  export let info
  export let instance: Instance
  import type { Instance } from "/src/libs/config"
  import { Card } from "@sveltestrap/sveltestrap"
  import T, { _ } from "/src/libs/Translate.svelte"
  import { toast } from "/src/libs/funcs"
</script>

<Card body color="danger">{$_("incompletePage")}</Card>
<br>
{instance.Name} App Profiles: {info.length}
