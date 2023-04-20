<script lang="ts">
  export let starrApp: StarrApp
  export let instance: Instance
  export let tab: Tab
  export let showTitle: boolean
  export let updating: boolean
  export let hidden: boolean

  import type { Tab } from "./fragments/tabs.svelte"
  import { Button, Card, CardBody, CardFooter, CardHeader, CardTitle, Collapse } from "sveltestrap"
  import T, { _ } from "../../libs/Translate.svelte"
  import { toast } from "../../libs/funcs"
  import Loading  from "../loading.svelte"
  import { conf, type Instance, type StarrApp } from "../../libs/config"
  import Fa from "svelte-fa"
  import { faCaretDown, faCaretUp } from "@fortawesome/free-solid-svg-icons"

  let rawOpen = false
  let info = undefined
  let prevTab = tab
  let prevURL = ""
  // update info when tab or instance changes.
  $: if (tab&&instance&&!hidden) update()

  async function update() {
    if (prevURL === instance.URL && prevTab === tab && info) return

    prevTab = tab
    updating = true
    info = undefined

    if (instance.URL=="") return
    await tab.data(instance).then(
      rep => { info = rep; prevURL = instance.URL },
      err => toast("error", err),
    )
    updating = false
  }
</script>

<Card outline color="dark" class="mt-1">
  {#if showTitle}
    <CardHeader>
      <CardTitle class="mb-0">{$_("instances."+tab.id)}</CardTitle>
    </CardHeader>
  {/if}

  <CardBody>
    {#if info}
    <div id="container">
      <!-- We have all the pieces we need. Load the selected tab's component. -->
      <svelte:component this={tab.component} {instance} bind:info={info} {tab}/>
    </div>
    {/if}
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
    {#if !instance || instance.URL == ""}
      <Card body color="danger">
        <T id="instances.NoURLConfigured" {starrApp} name={instance?instance.Name:"***"}/>
      </Card>
    {:else if !info}
      <Loading/>
    {/if}
  </CardBody>

  {#if instance && instance.URL}
    <CardFooter><code>{instance.URL}</code></CardFooter>
  {/if}
</Card>

<style>
  #container :global(.link) {
    cursor: pointer;
    text-decoration: underline;
  }
</style>