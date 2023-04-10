<script>
  export let starrApp
  export let instance
  export let tab
  export let showTitle
  
  import { Card, CardBody, CardFooter, CardHeader, CardTitle, Spinner } from "sveltestrap"
  import T, { _ } from "../../libs/Translate.svelte"
  import { toast } from "../../libs/funcs"

  let info
  $: activeTab = (tab != activeTab || info == undefined) ? undefined : tab

  $: if (activeTab == undefined) {
    info = undefined
    activeTab = tab
    if (instance && instance.URL) update()
  }

  const update = () => tab.fn(instance).then(
    rep => info = rep,
    err => toast("error", err),
  )
</script>

<Card outline color="dark" class="mt-2">
  {#if showTitle}
  <CardHeader>
    <CardTitle class="mb-0">{$_("instances."+tab.link)}</CardTitle>
  </CardHeader>
  {/if}

  <CardBody>
    {#if !instance} 
      <Card body color="danger">
        <T id="instances.NoURLConfigured" starrApp={starrApp} name="***"/>
      </Card>
    {:else if instance.URL == ""}
      <Card body color="danger">
        <T id="instances.NoURLConfigured" starrApp={starrApp} name={instance.Name}/>
      </Card>
    {:else if info}
      <!-- We have all the pieces we need. Load the form component. -->
      <svelte:component this={tab.lib} {instance} {info} on:update={update}/>
    {:else}
      <Card body color="secondary">
        <span>
          <Spinner size="sm" color="info" />
          <h5 style="display:inline-block">{$_("words.Loading")} ...</h5>
        </span>
      </Card>
    {/if}
  </CardBody>

  {#if instance && instance.URL}
    <CardFooter><code>{instance.URL}</code></CardFooter>
  {/if}
</Card>
