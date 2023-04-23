<script context="module" lang="ts">
  import { Exclusions } from "../../../wailsjs/go/starrs/Starrs"
  import type { Tab } from "./fragments/tabs.svelte"

  export const tab: Tab = { data: Exclusions, id: "Exclusions" }
  const max = new Date().getFullYear()+25
</script>

<script lang="ts">
  export let info: any
  export let instance: Instance
  export let updating: boolean

  import type { Instance } from "../../libs/config"
  import { _ } from "../../libs/Translate.svelte"
  import Footer from "./fragments/footer.svelte"
  import TDInput from "./fragments/tableInput.svelte"
  import SelectAll from "./fragments/selectAllHeader.svelte"
  import SelectRow from "./fragments/selectAllRow.svelte"
  import { Table } from "sveltestrap"

  let all: boolean = false      // Toggle for select-all link.
  let selected: any = {}        // Rows selected by key: ID.
  let str: string = JSON.stringify(info) // Used for equivalence comparison.
  let form: any = JSON.parse(str)        // Form changes go here.
  let starrApp = instance.App
</script>

<Table bordered>
  <tr>
    <SelectAll bind:all bind:selected bind:updating />
    {#if starrApp == "Sonarr" || starrApp == "Whisparr"}
      <th>TvDb ID</th>
      <th>{$_("words.Title")}</th>
    {:else if starrApp == "Radarr"}
      <th>TMDb ID</th>
      <th>{$_("configvalues.MovieTitle")}</th>
      <th>{$_("words.Year")}</th>
    {:else if starrApp == "Lidarr"}
      <th>{$_("configvalues.ForeignID")}</th>
      <th>{$_("configvalues.ArtistName")}</th>
    {:else if starrApp == "Readarr"}
      <th>{$_("configvalues.ForeignID")}</th>
      <th>{$_("configvalues.AuthorName")}</th>
    {/if}
  </tr>

  {#each info as exclusion, idx}
    {#if exclusion} <!-- When deleting an exclusion, this protects an error condition. -->
      <SelectRow {updating} bind:selected id={info[idx].id} item={exclusion}>
        {#if starrApp == "Sonarr" || starrApp == "Whisparr"}
        <TDInput {idx} {info} {updating} bind:form field="tvdbId" type="text" />
        <TDInput {idx} {info} {updating} bind:form field="title" type="text" />
        {:else if starrApp == "Radarr"}
        <TDInput {idx} {info} {updating} bind:form field="tmdbId" type="text" />
        <TDInput {idx} {info} {updating} bind:form field="movieTitle" type="text" />
        <TDInput {idx} {info} {updating} bind:form field="movieYear" min={1895} {max} type="range" />
        {:else if starrApp == "Lidarr"}
        <TDInput {idx} {info} {updating} bind:form field="foreignId" type="text" />
        <TDInput {idx} {info} {updating} bind:form field="artistName" type="text" />
        {:else if starrApp == "Readarr"}
        <TDInput {idx} {info} {updating} bind:form field="foreignId" type="text" />
        <TDInput {idx} {info} {updating} bind:form field="authorName" type="text" />
        {/if}
      </SelectRow>
      {/if}
    {/each}
</Table>

<Footer noForce {tab} {instance} bind:selected bind:updating bind:info bind:form bind:str />
