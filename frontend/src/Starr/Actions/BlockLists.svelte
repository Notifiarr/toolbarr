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
  import { QualityProfiles, MetadataProfiles } from "../../../wailsjs/go/starrs/Starrs"
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
  let metadataProfiles: any

  // Fetch extra data to populate the form.
  $: if (instance && instance.URL != "") {
    QualityProfiles(instance).then(resp => qualityProfiles = resp, err => { toast("error", err) })
    if (["Lidarr","Readarr"].includes(starrApp))
      MetadataProfiles(instance).then(resp => metadataProfiles = resp, err => { toast("error", err) })
  }
</script>

<Table striped responsive size="sm">
  <thead>
    <tr>
    <SelectAll bind:all bind:selected bind:updating />

    <!-- App Specific -->
    {#if starrApp == "Lidarr"}
    <th>
      <span class="link" id="artists.sortName" on:keyup={sort} on:click={sort}>{$_("configvalues.ArtistName")}</span>
      {#if sortKey == "artists.sortName"} <Icon name="caret-{sortDir?"up":"down"}"/> {/if}
    </th>
    {:else if starrApp == "Radarr"}
    <th>
      <span class="link" id="movies.sortTitle" on:keyup={sort} on:click={sort}>{$_("configvalues.MovieTitle")}</span>
      {#if sortKey == "movies.sortTitle"} <Icon name="caret-{sortDir?"up":"down"}"/> {/if}
    </th>
    {:else if starrApp == "Readarr"}
      <th>
        <span class="link" id="authorMetadata.sortName" on:keyup={sort} on:click={sort}>{$_("configvalues.AuthorName")}</span>
        {#if sortKey == "authorMetadata.sortName"} <Icon name="caret-{sortDir?"up":"down"}"/> {/if}
      </th>
    {:else if starrApp == "Sonarr"}
    <th>
      <span class="link" id="series.sortTitle" on:keyup={sort} on:click={sort}>{$_("configvalues.SeriesTitle")}</span>
      {#if sortKey == "series.sortTitle"} <Icon name="caret-{sortDir?"up":"down"}"/> {/if}
    </th>
    {/if}

    <!-- All Apps -->
    <th class="d-none d-md-table-cell">
      <span class="link" id="sourceTitle" on:keyup={sort} on:click={sort}>{$_("configvalues.SourceTitle")}</span>
      {#if sortKey == "sourceTitle"} <Icon name="caret-{sortDir?"up":"down"}"/> {/if}
    </th>
    <th>{$_("words.Quality")}</th>
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
      <BlockList {idx} {list} {starrApp} {qualityProfiles} {metadataProfiles}/>
    </SelectRow>
    {/if}
  {/each}
  </tbody>
</Table>

<Footer noForce {instance} {tab}
  bind:selected bind:updating bind:form bind:str
  bind:info={info.records} on:delete={()=>info.totalRecords--}/>
