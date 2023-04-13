<script>
  export let starrApp
  export let instance
  export let tab
  export let showTitle
  
  import { Button, Card, CardBody, CardFooter, CardHeader, CardTitle, Collapse } from "sveltestrap"
  import T, { _ } from "../../libs/Translate.svelte"
  import { toast } from "../../libs/funcs"
  import Loading  from "./loading.svelte"
  import { conf } from "../../libs/config"
  import Fa from "svelte-fa"
  import { faCaretDown, faCaretUp } from "@fortawesome/free-solid-svg-icons"

  let updating = true
  let rawOpen = false
  let info
  $: activeTab = (tab != activeTab || info == undefined) ? undefined : tab

  $: if (activeTab == undefined) {
    info = undefined
    activeTab = tab
    rawOpen = false
    if (instance && instance.URL) update()
  }

  async function update() {
    updating = true
    await tab.fn(instance).then(
      rep => info = rep,
      err => toast("error", err),
    )
    updating = false
  }
</script>

<Card outline color="dark" class="mt-2">
  {#if showTitle}
  <CardHeader>
    <CardTitle class="mb-0">{$_("instances."+tab.link)}</CardTitle>
  </CardHeader>
  {/if}

  <CardBody>
    {#if !instance || instance.URL == ""}
      <Card body color="danger">
        <T id="instances.NoURLConfigured" {starrApp} name={instance?instance.Name:"***"}/>
      </Card>
    {:else if !info}
      <Loading/>
    {:else}
      <!-- We have all the pieces we need. Load the selected tab's component. -->
      <svelte:component this={tab.lib} {instance} {info} on:update={update} {updating}/>
      <!-- show raw data button for dev mode -->
      <Collapse isOpen={$conf.DevMode}>
        <hr>
        <Button size="sm" on:click={() => (rawOpen = !rawOpen)} class="mb-1">
          Raw Data <Fa icon={rawOpen?faCaretDown:faCaretUp}/>
        </Button>
        <Card color="secondary">
          <Collapse isOpen={rawOpen}>
            <code><pre>{JSON.stringify(info, null, 3)}</pre></code>
          </Collapse>
        </Card>
      </Collapse>
    {/if}
  </CardBody>

  {#if instance && instance.URL}
    <CardFooter><code>{instance.URL}</code></CardFooter>
  {/if}
</Card>
