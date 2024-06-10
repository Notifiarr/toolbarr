<script context="module" lang="ts">
  import { Exclusions } from "/wailsjs/go/starrs/Starrs"
  import type { Tab } from "./fragments/tabs.svelte"

  export const tab: Tab = { data: Exclusions, id: "Exclusions" }
  const max = new Date().getFullYear()+25
</script>

<script lang="ts">
  export let info: any
  export let instance: Instance
  export let updating: boolean
  export let footer = true
  export let selected: {[key: string]: boolean} = {} // Rows selected by key: ID.
  export let str: string = ""                        // Used for equivalence comparison.
  export let form: any = undefined                   // Form changes go here.

  import type { Instance } from "/src/libs/config"
  import { _ } from "/src/libs/Translate.svelte"
  import Footer from "./fragments/footer.svelte"
  import TDInput from "./fragments/tableInput.svelte"
  import SelectAll from "./fragments/selectAllHeader.svelte"
  import SelectRow from "./fragments/selectAllRow.svelte"
  import { Table } from "@sveltestrap/sveltestrap"
  import self from "./Exclusions.svelte"
  import { onMount } from "svelte";

  let all: boolean = false      // Toggle for select-all link.
  str = JSON.stringify(info)
  form = JSON.parse(str)
  let starrApp = instance.App

  // The module can only reference itself after mounting itself.
  onMount(() => { tab.component = self });
</script>

<Table bordered>
  <tr>
    <SelectAll bind:all bind:selected bind:updating icon="check2-all" />
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
      {@const loopid = footer?exclusion.id:idx}
      <SelectRow {updating} bind:selected id={loopid} item={exclusion}>
        {#if starrApp == "Sonarr" || starrApp == "Whisparr"}
        <TDInput {idx} {info} {updating} bind:form field="tvdbId" type="number" />
        <TDInput {idx} {info} {updating} bind:form field="title" type="text" />
        {:else if starrApp == "Radarr"}
        <TDInput {idx} {info} {updating} bind:form field="tmdbId" type="number" />
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

{#if footer}
<Footer noForce {tab} {instance} bind:selected bind:updating bind:info bind:form bind:str importable />
{/if}