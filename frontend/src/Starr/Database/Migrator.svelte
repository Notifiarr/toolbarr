<script>
  export let instance
  export let showTitle
  import { Card, CardBody, CardFooter, CardHeader, CardTitle } from "sveltestrap"
  import T, { _ } from "../../libs/Translate.svelte"
</script>

<Card outline class="mt-2">
  <span class={showTitle?"":"d-none"}>
    <CardHeader>
      <CardTitle class="mb-0">{$_("instances.FilesystemPathsMigrator")}</CardTitle>
    </CardHeader>
  </span>
  <CardBody>
    <p><T id ="instances.MigratorDescription" starrApp={instance.App}/></p>
    {#if instance.DBPath == ""}
      <Card body color="danger">
        <T id="instances.NoDatabasePathConfigured" starrApp={instance.App} name={instance.Name}/>
      </Card>
    {:else}
      <Card body color="info">Migration form for {instance.Name} goes here.</Card>
    {/if}
  </CardBody>
  <CardFooter><code>{instance.DBPath}</code></CardFooter>
</Card>