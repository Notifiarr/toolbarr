<script>
  export let isOpen = true

  import { Card, Fade, Spinner } from "sveltestrap"
  import { _ } from "../../libs/Translate.svelte"
</script>

<Fade {isOpen}>
  <Card body color="secondary">
    <span>
      <Spinner size="sm" color="info" />
      <h5 style="display:inline-block">{$_("words.Loading")} ...</h5>
    </span>
  </Card>
</Fade>