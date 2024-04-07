<script lang="ts">
  export let starrApp: StarrApp
  export let instance: Instance
  export let showTitle: boolean | undefined

  import type { Instance, StarrApp } from "/src/libs/config"
  import { Alert, Card, CardBody, CardFooter, CardHeader, CardTitle } from "@sveltestrap/sveltestrap"
  import T, { _ } from "/src/libs/Translate.svelte"
</script>

<Card outline class="mt-2">
  <span class={showTitle?"":"d-none"}>
    <CardHeader>
      <CardTitle class="mb-0">{$_("instances.SQLite3DatabaseInspector")}</CardTitle>
    </CardHeader>
  </span>

  <CardBody>
    <Alert fade={false} color="danger">{$_("incompletePage")}</Alert>
    <p>This tool allows you to browse your {starrApp} database.</p>
    {#if instance}
      {#if instance && instance.DBPath == ""}
        <Card body color="danger"><T id="instances.NoDatabasePathConfigured" starrApp={instance.App} name={instance.Name}/></Card>
      {:else}
        <Card body color="info">DB explorer for {instance.Name} goes here.</Card>
      {/if}
    {/if}
  </CardBody>

  {#if instance && instance.DBPath}
    <CardFooter><code>{instance.DBPath}</code></CardFooter>
  {/if}
</Card>