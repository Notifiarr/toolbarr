<script lang="ts">
  export let info: any
  export let instance: Instance
  export let tab: Tab
  export let updating = false
  export let sortKey: string
  export let sortDir: boolean

  import type { Tab } from "./fragments/tabs.svelte"
  import type { Instance } from "../../libs/config"
  import { toast } from "../../libs/funcs"
  import { Table, Icon } from "sveltestrap"
  import { _ } from "../../libs/Translate.svelte"
  import SelectAll from "./fragments/selectAllHeader.svelte"
  import SelectRow from "./fragments/selectAllRow.svelte"
  import Footer from "./fragments/footer.svelte"
  import { QualityProfiles } from "../../../wailsjs/go/starrs/Starrs"
  import { createEventDispatcher } from "svelte"
  import BlockList from "./blockListsRow.svelte"

  const dispatch = createEventDispatcher()
  let all = false
  let selected: any = {}
  let str: string = JSON.stringify(info.records)
  let form: any = JSON.parse(str)
  const starrApp = instance.App

  function sort(e) {
    if (e.target.id == sortKey) sortDir = !sortDir 
    sortKey = e.target.id
    dispatch("update", true)
  }

  let qualityProfiles: any

  // Fetch extra data to populate the form.
  $: if (instance && instance.URL != "") {
    QualityProfiles(instance).then(resp => qualityProfiles = resp, err => { toast("error", err) })
  }

  function qp(id) {
    let name = id
    qualityProfiles.forEach((qp) => {
      if (qp.id == id) name = qp.name
    })
    return name
  }
</script>

<Table striped responsive size="sm">
  <thead>
    <tr>
    <SelectAll bind:all bind:selected bind:updating />
    {#if starrApp == "Lidarr"}
      <th>{$_("configvalues.ArtistName")}</th>
      <th>{$_("configvalues.SourceTitle")}</th>
      <th>{$_("words.Quality")}</th>
    {:else if starrApp == "Radarr"}
      <th>{$_("configvalues.MovieTitle")}</th>
      <th>{$_("configvalues.SourceTitle")}</th>
      <th>{$_("words.Quality")}</th>
    {:else if starrApp == "Readarr"}
      <th>{$_("configvalues.AuthorName")}</th>
      <th>{$_("configvalues.SourceTitle")}</th>
      <th>{$_("words.Quality")}</th>
    {:else if starrApp == "Sonarr"}
      <th>Title</th>
      <th>{$_("configvalues.SourceTitle")}</th>
      <th>{$_("words.Quality")}</th>
    {/if}
      <th>
        <span class="link" id="date" on:keyup={sort} on:click={sort}>{$_("words.Date")}</span>
        {#if sortKey == "date"} <Icon name="caret-{sortDir?"up":"down"}"/> {/if}
      </th>
    </tr>
  </thead>

  <tbody>
  {#each info.records as list, idx}
    {#if list} <!-- When deleting an exclusion, this protects an error condition. -->
    <SelectRow {updating} bind:selected id={info.records[idx].id} item={list}>
      <!-- This sub page contains the content specific to each app. -->
      <BlockList {idx} {list} {starrApp} {qualityProfiles}/>
    </SelectRow>
    {/if}
  {/each}
  </tbody>
</Table>

<Footer noForce {instance} {tab}
  bind:selected bind:updating bind:form bind:str 
  bind:info={info.records} on:delete={()=>info.totalRecords--}/>