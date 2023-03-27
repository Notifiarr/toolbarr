<script>
  export let starrApp
  export let instance
  export let showTitle

  import { Card, CardBody, CardFooter, CardHeader, CardTitle, Spinner } from "sveltestrap"
  import T, { _ } from "../../../libs/Translate.svelte"
  import { GetMigratorInfo } from "../../../../wailsjs/go/app/App.js"
  import { toast } from "../../../libs/funcs"
  import Form from "./Form.svelte"

  let info
  if (instance && instance.DBPath) {
    GetMigratorInfo(instance).then(
      rep => info = rep,
      err => toast("error", err),
    )
  }
</script>

<Card outline class="mt-2">
  <span class={showTitle?"":"d-none"}>
    <CardHeader>
      <CardTitle class="mb-0">{$_("instances.FilesystemPathsMigrator")}</CardTitle>
    </CardHeader>
  </span>

  <CardBody>
    <p><T id ="instances.MigratorDescription" starrApp={starrApp}/></p>
    {#if instance}
      {#if instance.DBPath == ""}
        <Card body color="danger"><T id="instances.NoDatabasePathConfigured" starrApp={starrApp} name={instance.Name}/></Card>
      {:else if info}
        <!-- We have all the pieces we need. Load the form component. -->
        <Form {instance} {info}/>
      {:else}
        <Card body color="secondary">
          <span>
            <Spinner size="sm" color="info" />
            <h5 style="display:inline-block">{$_("words.Loading")} ...</h5>
          </span>
        </Card>
      {/if}
    {:else}
      <Card body color="danger"><T id="instances.NoDatabasePathConfigured" starrApp={starrApp} name="***"/></Card>
    {/if}
  </CardBody>

  {#if instance && instance.DBPath}
    <CardFooter><code>{instance.DBPath}</code></CardFooter>
  {/if}
</Card>