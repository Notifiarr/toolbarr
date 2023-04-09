<script>
  export let info

  import { Card, Spinner } from "sveltestrap"
  import T, { _ } from "../../../libs/Translate.svelte"
  import { toast } from "../../../libs/funcs"
</script>

<Card body color="danger">{$_("incompletePage")}</Card>

{#if !info} <Spinner/> {$_("words.Loading")} ... {:else}
<code><pre class="code">{JSON.stringify(info, null, 3)}</pre></code>
{/if}