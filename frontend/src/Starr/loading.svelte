<script>
  export let isOpen = true

  import { Card, Collapse, Spinner } from "sveltestrap"
  import { _ } from "../libs/Translate.svelte"
</script>

<Collapse expand {isOpen}>
  <Card body color="secondary">
    <span>
      <Spinner type="border" size="sm" color="info" />
      <h5 style="display:inline-block">{$_("words.Loading")} ...</h5>
    </span>
  </Card>
</Collapse>