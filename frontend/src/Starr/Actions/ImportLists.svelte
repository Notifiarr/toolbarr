<script lang="ts">
  export let info: any
  export let instance: Instance
  export let tab: Tab
  export let updating: boolean

  import type { Instance } from "/src/libs/config"
  import type { Tab } from "./fragments/tabs.svelte"
  import { toast } from "/src/libs/funcs"
  import { _ } from "/src/libs/Translate.svelte"
  import Footer from "./fragments/footer.svelte"
  import { fixFieldValues } from "./methods"
  import {
    QualityProfiles,
    MetadataProfiles,
    RootFolders,
    Exclusions as GetExclusions,
  } from "/wailsjs/go/starrs/Starrs"
  import Exclusions from "./Exclusions.svelte"
  import ModalInput from "./fragments/modalInput.svelte"
  import Dropdown from "./fragments/dropdown.svelte"
  import ConfigModal from "./fragments/configModal.svelte"
  import TDInput from "./fragments/tableInput.svelte"
  import FieldInput from "./fragments/fieldInput.svelte"
  import SelectAll from "./fragments/selectAllHeader.svelte"
  import SelectRow from "./fragments/selectAllRow.svelte"
  import { Card, Table, Tooltip, Icon, Collapse } from "@sveltestrap/sveltestrap"

  let isOpen: any = {}       // Modal toggle control.
  let all: boolean = false       // Toggle for select-all link.
  let selected: any = {}     // Rows selected by key: ID.
  let str: string = fixFieldValues(info) // Used for equivalence comparison.
  let form: any = JSON.parse(str)     // Form changes go here.
  let starrApp = instance.App
  let applyAll: string = $_("instances.applyAllimportList", {values:{starrApp: instance.Name}})
  let qualityProfiles: any
  let metadataProfiles: any
  let rootFolders: any
  let exclusions: any
  let exclusionsOpen = false

  // Fetch extra data to populate the form.
  $: if (instance && instance.URL != "") {
    GetExclusions(instance).then(resp => exclusions = resp, err => { toast("error", err) })
    QualityProfiles(instance).then(resp => qualityProfiles = resp, err => { toast("error", err) })
    RootFolders(instance).then(
      resp => rootFolders = Object.keys(resp).map((key) => resp[key].path),
      err  => toast("error", err)
    )
    if (["Lidarr","Readarr"].includes(starrApp)) // Metadata profiles only exist on Lidarr and Readarr.
      MetadataProfiles(instance).then( resp => metadataProfiles = resp, err => { toast("error", err) })
  }
</script>

<Table bordered>
  <tr>
    <SelectAll bind:all bind:selected bind:updating icon="check2-all"/>
    <th class="d-none d-sm-table-cell">
      <Tooltip target="il{instance.App}Type">{$_("words.Implementation")}</Tooltip>
      <span id="il{instance.App}Type">{$_("words.Type")}
        <Icon class="small text-info" name="box-arrow-up-right"/>
      </span>
    </th>
    <th>{$_("words.Name")}</th>
    <th>
      <Dropdown bind:form {updating} {starrApp} field="enableAutomaticAdd" name="instances.enableAutoTitle"/>
    </th>
    {#if starrApp != "Sonarr"}
      <th>
        {#if starrApp == "Radarr"}
          <Dropdown bind:form {updating} {starrApp} field="enabled"/>
        {:else}
          <Dropdown bind:form {updating} {starrApp} field="shouldMonitorExisting"/>
        {/if}
      </th>
    {/if}
    <th>
      {#if starrApp == "Sonarr"}
        <Dropdown bind:form {updating} {starrApp} field="seasonFolder"/>
      {:else}
        <Dropdown bind:form {updating} {starrApp} field="shouldSearch"/>
      {/if}
    </th>
  </tr>

  {#each info as list, idx}
    {#if list} <!-- When deleting a client, this protects an error condition. -->
    <ConfigModal {info} {form} {idx} {str} id={list.id} name={list.implementation} bind:isOpen={isOpen[idx]}>
      {#if list.message && list.message.message}
        <p><Card class="p-1" color="warning">{list.message.message}</Card></p>
      {/if}
      <ModalInput {info} bind:form {idx} field="name" name="words.Name" type="text"/>
      <ModalInput {info} bind:form {idx} field="qualityProfileId" type="select" {applyAll}>
        {#if qualityProfiles}
          {#each qualityProfiles as profile}
            <option value={profile.id}>{profile.name}</option>
          {/each}
        {:else}
          <option value={list.qualityProfileId}>{list.qualityProfileId}</option>
        {/if}
      </ModalInput>
      <ModalInput {info} bind:form {idx} {applyAll} field="metadataProfileId" type="select">
        {#if metadataProfiles}
          {#each metadataProfiles as profile}
            <option value={profile.id}>{profile.name}</option>
          {/each}
        {:else}
          <option value={list.metadataProfileId}>{list.metadataProfileId}</option>
        {/if}
      </ModalInput>
      <ModalInput {info} bind:form {idx} {applyAll} field="rootFolderPath" type="select">
        {#if rootFolders}
          {#each rootFolders as dir}
            <option value={dir}>{dir}</option>
          {/each}
          {#if !rootFolders.includes(list.rootFolderPath)}
            <option value={list.rootFolderPath}>{list.rootFolderPath} (invalid)</option>
          {/if}
        {:else}
          <option value={list.rootFolderPath}>{list.rootFolderPath}</option>
        {/if}
      </ModalInput>
      <ModalInput {info} bind:form {idx} {applyAll} field="monitorNewItems" type="select">
        <option value="all">{$_("words.All")}</option>
        <option value="none">{$_("words.None")}</option>
        <option value="new">{$_("words.New")}</option>
        {#if !["all","none","new"].includes(list.monitorNewItems)}
          <option value={list.monitorNewItems}>{list.monitorNewItems}</option>
        {/if}
      </ModalInput>
      <ModalInput {info} bind:form {idx} {applyAll} {starrApp} field="enableAutomaticAdd"
        desc="instances.enableAutoDesc" name="instances.enableAutoTitle" type="checkbox"/>
      <ModalInput {info} bind:form {idx} {applyAll} {starrApp} field="enableAuto" type="checkbox"/>
      <ModalInput {info} bind:form {idx} {applyAll} {starrApp} field="shouldSearch" type="checkbox"/>
      <ModalInput {info} bind:form {idx} {applyAll} starrApp={instance.Name} field="shouldMonitorExisting" type="checkbox"/>
      <ModalInput {info} bind:form {idx} {applyAll} {starrApp} field="monitor" type="select">
        {#if starrApp == "Radarr"}
        <option value="movieOnly">{$_("instances.MovieOnly")}</option>
        <option value="movieAndCollection">{$_("instances.MovieandCollection")}</option>
        <option value="none">{$_("words.None")}</option>
        {/if}
      </ModalInput>
      <ModalInput {info} bind:form {idx} {applyAll} {starrApp} field="seriesType" type="select">
        {#if starrApp == "Sonarr"}
        <option value="standard">{$_("words.Standard")}</option>
        <option value="daily">{$_("words.Daily")}</option>
        <option value="anime">{$_("words.Anime")}</option>
        {/if} <!-- whisparr? -->
      </ModalInput>
      <ModalInput {info} bind:form {idx} {applyAll} {starrApp} field="shouldMonitor" name="instances.monitorTitle" type="select">
        {#if starrApp == "Lidarr"}
          <option value="entireArtist">{$_("instances.AllArtistAlbums")}</option>
          <option value="specificAlbum">{$_("instances.SpecificAlbum")}</option>
          <option value="none">{$_("words.None")}</option>
        {:else if starrApp == "Readarr"}
          <option value="specificBook">{$_("instances.SpecificBook")}</option>
          <option value="entireAuthor">{$_("instances.EntireAuthor")}</option>
          <option value="none">{$_("words.None")}</option>
        {:else if starrApp == "Sonarr"}
          <option value="future">{$_("instances.FutureEpisodes")}</option>
          <option value="missing">{$_("instances.MissingEpisodes")}</option>
          <option value="existing">{$_("instances.ExistingEpisodes")}</option>
          <option value="pilot">{$_("instances.PilotEpisode")}</option>
          <option value="firstSeason">{$_("instances.OnlyFirstSeason")}</option>
          <option value="latestSeason">{$_("instances.OnlyLatestSeason")}</option>
          <option value="monitorSpecials">{$_("instances.MonitorSpecials")}</option>
          <option value="unmonitorSpecials">{$_("instances.UnmonitorSpecials")}</option>
          <option value="none">{$_("words.None")}</option>
        {:else if starrApp == "Whisparr"}
          <option value="none">{$_("words.None")}</option>
          {#if list.shouldMonitor != "none"}
            <option value={list.shouldMonitor}>{list.shouldMonitor}</option>
          {/if}
        {/if}
      </ModalInput>
      <ModalInput {info} bind:form {idx} {applyAll} {starrApp} field="minimumAvailability" type="select">
        {#if starrApp == "Radarr"}
        <option value="announced">{$_("instances.Announced")}</option>
        <option value="released">{$_("instances.Released")}</option>
        <option value="inCinemas">{$_("instances.InCinemas")}</option>
        {/if}
      </ModalInput>

      {#each info[idx].fields as item, itemIdx}
        <FieldInput {item} {itemIdx} {info} {idx} bind:form/>
      {/each}
    </ConfigModal>

    <SelectRow {updating} bind:selected id={info[idx].id} item={list}>
      <td class={JSON.stringify(form[idx]) != JSON.stringify(list)?"border-warning":""}>
        <a href="/" style="padding-left:0" on:click|preventDefault={() => isOpen[idx]=!updating}>{list.name}</a>
      </td>
      <!-- The order here must match the dropdown menus at the top. -->
      <TDInput {idx} {info} {updating} bind:form field="enableAutomaticAdd" type="switch"/>
      <TDInput {idx} {info} {updating} bind:form field="enableAuto" type="switch"/>
      <TDInput {idx} {info} {updating} bind:form field="shouldMonitorExisting" type="switch"/>
      <TDInput {idx} {info} {updating} bind:form field="enabled" type="switch"/>
      <TDInput {idx} {info} {updating} bind:form field="shouldSearch" type="switch"/>
      <TDInput {idx} {info} {updating} bind:form field="searchOnAdd" type="switch"/>
      <TDInput {idx} {info} {updating} bind:form field="seasonFolder" type="switch"/>
    </SelectRow>
    {/if}<!-- /if (client) -->
  {/each}<!-- /each info as client, idx -->
</Table>

<Footer {instance} bind:selected {tab} bind:updating bind:info bind:form bind:str/>

<Card class="p-1" color="secondary">
  <!-- <Tooltip target="exclusions">Close this card to reset the Exclusions form.</Tooltip> -->
  <h5 style="cursor:pointer" on:keyup={()=>{exclusionsOpen=!exclusionsOpen}} on:click={()=>{exclusionsOpen=!exclusionsOpen}}>
    <span id="exclusions">Import List Exclusions</span>
    <Icon class="text-info" name={exclusionsOpen?"caret-down-fill":"caret-up"}/>
  </h5>
  <Collapse isOpen={exclusions != undefined && exclusionsOpen}>
    <Exclusions bind:updating info={exclusions} {instance}/>
  </Collapse>
</Card>