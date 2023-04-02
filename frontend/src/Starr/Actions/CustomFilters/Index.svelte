<script>
  export let info

  import { Spinner } from "sveltestrap"
  import T, { _ } from "../../../libs/Translate.svelte"
  import { toast } from "../../../libs/funcs"
</script>

<code><pre class="code">{JSON.stringify(info, null, 3)}</pre></code>
