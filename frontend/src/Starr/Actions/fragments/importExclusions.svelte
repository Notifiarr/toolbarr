<script lang="ts">
  // The instances being imported into.
  export let instance: Instance
  // The import data read from a file. Loop this.
  export let info: any[]
  // This goes true when a button is clicked.
  export let updating: boolean
  export let selected: {[key: string]: boolean} // Rows selected by key: ID.

  import { _ } from "/src/libs/Translate.svelte"
  import TDInput from "./tableInput.svelte"
  import SelectAll from "./selectAllHeader.svelte"
  import SelectRow from "./selectAllRow.svelte"
  import { Input, Table, Tooltip } from "@sveltestrap/sveltestrap"
  import type { Instance } from "/src/libs/config"

  let all = false
  let isOpen: {[key: number]: boolean} = {} // Modal toggle control.
</script>

<Table bordered>
  <tr>
    <SelectAll bind:all bind:selected bind:updating />
    {#if instance.App == "Sonarr" || instance.App == "Whisparr"}
      <th>TvDb ID</th>
      <th>{$_("words.Title")}</th>
    {:else if instance.App == "Radarr"}
      <th>TMDb ID</th>
      <th>{$_("configvalues.MovieTitle")}</th>
      <th>{$_("words.Year")}</th>
    {:else if instance.App == "Lidarr"}
      <th>{$_("configvalues.ForeignID")}</th>
      <th>{$_("configvalues.ArtistName")}</th>
    {:else if instance.App == "Readarr"}
      <th>{$_("configvalues.ForeignID")}</th>
      <th>{$_("configvalues.AuthorName")}</th>
    {/if}
  </tr>

  {#each info as exclusion, idx}
    {#if exclusion} <!-- When deleting an exclusion, this protects an error condition. -->
      <SelectRow {updating} bind:selected id={info[idx].id} item={exclusion}>
        {#if instance.App == "Sonarr" || instance.App == "Whisparr"}
        <TDInput {idx} {info} {updating} form={info} disabled field="tvdbId" type="text" />
        <TDInput {idx} {info} {updating} form={info} disabled field="title" type="text" />
        {:else if instance.App == "Radarr"}
        <TDInput {idx} {info} {updating} form={info} disabled field="tmdbId" type="text" />
        <TDInput {idx} {info} {updating} form={info} disabled field="movieTitle" type="text" />
        <TDInput {idx} {info} {updating} form={info} disabled field="movieYear" type="text" />
        {:else if instance.App == "Lidarr"}
        <TDInput {idx} {info} {updating} form={info} disabled field="foreignId" type="text" />
        <TDInput {idx} {info} {updating} form={info} disabled field="artistName" type="text" />
        {:else if instance.App == "Readarr"}
        <TDInput {idx} {info} {updating} form={info} disabled field="foreignId" type="text" />
        <TDInput {idx} {info} {updating} form={info} disabled field="authorName" type="text" />
        {/if}
      </SelectRow>
      {/if}
    {/each}
</Table>