<script lang="ts">
  export let info: any
  export let instance: Instance
  export let tab: Tab
  export let updating = false
  export let sortKey: string
  export let sortDir: boolean

  import type { Tab } from "./fragments/tabs.svelte"
  import type { Instance } from "/src/libs/config"
  import { toast } from "/src/libs/funcs"
  import { Table, Icon } from "@sveltestrap/sveltestrap"
  import { _ } from "/src/libs/Translate.svelte"
  import SelectAll from "./fragments/selectAllHeader.svelte"
  import SelectRow from "./fragments/selectAllRow.svelte"
  import Footer from "./fragments/footer.svelte"
  import { QualityProfiles, MetadataProfiles } from "/wailsjs/go/starrs/Starrs"
  import { createEventDispatcher } from "svelte"
  import BlockList from "./blockListsRow.svelte"

  let qualityProfiles: any
  let metadataProfiles: any
  // Fetch extra data to populate the modals.
  $: if (instance && instance.URL != "") {
    QualityProfiles(instance).then(resp => qualityProfiles = resp, err => { toast("error", err) })
    if (["Lidarr","Readarr"].includes(starrApp))
      MetadataProfiles(instance).then(resp => metadataProfiles = resp, err => { toast("error", err) })
  }

  const dispatch = createEventDispatcher()
  const starrApp = instance.App
  const caret = sortDir ? "caret-up" : "caret-down"
  let all = false
  let selected: {[key: string]: boolean} = {}
  let str: string = JSON.stringify(info.records)
  let form: any = JSON.parse(str)

  let id: string = ""
  let cv: string = ""
  // Set app-specific values.
  if (starrApp == "Lidarr") {
    id = "artists.sortName"
    cv = $_("configvalues.ArtistName")
  } else if (starrApp == "Radarr") {
    id = "movies.sortTitle"
    cv = $_("configvalues.MovieTitle")
  } else if (starrApp == "Readarr") {
    id = "authorMetadata.sortName"
    cv = $_("configvalues.AuthorName")
  } else if (starrApp == "Sonarr") {
    id = "series.sortTitle"
    cv = $_("configvalues.SeriesTitle")
  }

  function sort(e: any) {
    if (e.target && e.target.id == sortKey) sortDir = !sortDir
    sortKey = e.target.id
    dispatch("update", true)
  }
</script>

<Table striped responsive size="sm">
  <thead>
    <tr>
    <SelectAll bind:all bind:selected bind:updating />

    <!-- App Specific -->
    <th>
      <span class="link" {id} on:keyup={sort} on:click={sort} role="link" tabindex="-1">{cv}</span>
      <Icon name={sortKey==id?caret:""}/>
    </th>

    <!-- All Apps -->
    <th class="d-none d-md-table-cell">
      <span class="link" id="sourceTitle" on:keyup={sort} on:click={sort} role="link" tabindex="-2">{$_("configvalues.SourceTitle")}</span>
      <Icon name={sortKey=="sourceTitle"?caret:""}/>
    </th>
    <th>{$_("words.Quality")}</th>
      <th>
        <span class="link" id="date" on:keyup={sort} on:click={sort} role="link" tabindex="-3">{$_("words.Date")}</span>
        <Icon name={sortKey=="date"?caret:""}/>
      </th>
    </tr>
  </thead>

  <tbody>
  {#each info.records as list, idx}
    {#if list} <!-- When deleting an exclusion, this protects an error condition. -->
    <SelectRow {updating} bind:selected id={info.records[idx].id} item={list}>
      <!-- This sub page contains the content specific to each app. -->
      <BlockList {idx} {list} {starrApp} {qualityProfiles} {metadataProfiles}/>
    </SelectRow>
    {/if}
  {/each}
  </tbody>
</Table>

<Footer noForce {instance} {tab}
  bind:selected bind:updating bind:form bind:str
  bind:info={info.records} on:delete={()=>info.totalRecords--}/>
