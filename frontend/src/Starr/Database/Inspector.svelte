<script>
    export let instance
    export let showTitle

    import { Alert, Card, CardBody, CardFooter, CardHeader, CardTitle } from "sveltestrap"
    import T, { _ } from "../../libs/Translate.svelte"
</script>

<Card outline class="mt-2">
  <span class={showTitle?"":"d-none"}>
    <CardHeader>
      <CardTitle class="mb-0">{$_("instances.SQLite3DatabaseInspector")}</CardTitle>
    </CardHeader>
  </span>
  <CardBody>
    <Alert color="danger">{$_("incompletePage")}</Alert>
    <p>
      This tool allows you to browse your {instance.App} database. 
    </p>

    {#if instance.DBPath == ""}
      <Card body color="danger"><T id="instances.NoDatabasePathConfigured" starrApp={instance.App} name={instance.Name}/></Card>
    {:else}
      <Card body color="info">DB explorer for {instance.Name} goes here.</Card>
    {/if}
  </CardBody>
  <CardFooter><code>{instance.DBPath}</code></CardFooter>
</Card>