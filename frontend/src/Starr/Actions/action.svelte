<script lang="ts">
  export let starrApp: StarrApp
  export let instance: Instance
  export let tab: Tab
  export let showTitle: boolean
  export let updating: boolean
  export let hidden: boolean

  import type { Tab } from "./fragments/tabs.svelte"
  import {
    Button,
    Card,
    CardBody,
    CardFooter,
    CardHeader,
    CardTitle,
    Collapse,
  } from "@sveltestrap/sveltestrap"
  import T, { _ } from "/src/libs/Translate.svelte"
  import { toast } from "/src/libs/funcs"
  import Loading from "/src/Starr/loading.svelte"
  import Paginate  from "./fragments/paginate.svelte"
  import { conf, type Instance, type StarrApp } from "/src/libs/config"
  import Fa from "svelte-fa"
  import { faCaretDown, faCaretUp } from "@fortawesome/free-solid-svg-icons"

  let rawOpen = false
  let info = undefined
  let prevTab = tab
  let prevURL = ""
  // These are only used for pageable content.
  let page = 1
  let pages = 1
  let pageSize = 10
  let sortKey = "date" 
  let sortDir: boolean = false // descending

  // update info when tab or instance changes.
  $: if (tab&&instance&&!hidden) update({})

  async function update(e) {
    if (prevURL === instance.URL && prevTab === tab && info && !e.detail) return

    prevTab = tab
    updating = true
    info = undefined

    if (instance.URL=="") return
    if (tab.page) {
      await tab.data(instance, pageSize, page, sortKey, sortDir?"ascending":"descending").then(
        rep => {
          info = rep
          prevURL = instance.URL
          pages = Math.ceil(info.totalRecords / info.pageSize)
        },
        err => toast("error", err),
      )
    } else {
      await tab.data(instance).then(
        rep => { info = rep; prevURL = instance.URL },
        err => toast("error", err),
      )
    }

    updating = false
  }
</script>

<Card color={$conf.Dark?"secondary":"light"} class="mt-1">
  {#if showTitle}
    <CardHeader>
      <CardTitle class="mb-0">{$_("instances."+tab.id)}</CardTitle>
    </CardHeader>
  {/if}

  <CardBody>
    {#if info}
    <div id="container" class={$conf.Dark?"dark-mode":""}>
      <!-- We have all the pieces we need. Load the selected tab's component. -->
      {#if !tab.page}
        <svelte:component this={tab.component} {instance} {tab} bind:info bind:updating />
      {:else}<!-- tab is pagable-->
        <svelte:component {instance} {tab}
          bind:info bind:updating bind:sortKey bind:sortDir 
          this={tab.component} on:update={update}/>
        <Paginate {updating} {pages}
          bind:pageSize bind:page records={info.records.length}
          total={info.totalRecords} on:update={update}/>
      {/if}
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
    color: rgb(19, 87, 87)
  }

  #container.dark-mode :global(.link) {
    color: rgb(31, 144, 144)
  }
</style>
