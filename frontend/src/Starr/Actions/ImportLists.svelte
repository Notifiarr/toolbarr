<script>
  export let info
  export let instance

  import { toast } from "../../libs/funcs"
  import T, { _ } from "../../libs/Translate.svelte"
  import Fa from "svelte-fa"
  import Footer from "./footer.svelte"
  import { fixFieldValues } from "./methods"
  import { faCircleInfo, faArrowUpRightFromSquare, faTrashAlt } from "@fortawesome/free-solid-svg-icons"
  import { QualityProfiles, MetadataProfiles, RootFolders } from "../../../wailsjs/go/starrs/Starrs"
  import {
    Badge,
    Button,
    Card,
    Dropdown,
    DropdownItem,
    DropdownMenu,
    DropdownToggle,
    Input,
    InputGroup,
    InputGroupText,
    Modal,
    ModalBody,
    ModalFooter,
    ModalHeader,
    Table,
    Tooltip,
  } from "sveltestrap"

  let isOpen = {}       // Modal toggle control.
  let updating = false  // True while doing updates.
  let all = false       // Toggle for select-all link.
  let selected = {}     // Rows selected by key: ID.
  let str = fixFieldValues(info) // Used for equivalence comparison.
  let form = JSON.parse(str)     // Form changes go here.

  let qualityProfiles
  let metadataProfiles
  let rootFolders

  // Fetch extra data to populate the form.
  $: if (instance && instance.URL != "") {
    QualityProfiles(instance).then(resp => qualityProfiles = resp, err => { toast("error", err) })
    RootFolders(instance).then(resp => rootFolders = resp, err => { toast("error", err) })
    if (["Lidarr","Readarr"].includes(instance.App)) // Metadata profiles only exist on Lidarr and Readarr.
      MetadataProfiles(instance).then( resp => metadataProfiles = resp, err => { toast("error", err) })
  }

  function selectAll() {
    all = !all
    if (!updating) Object.keys(selected).forEach(k => selected[k] = all)
  }

  function toggleAll(key, on) { form.forEach((_, i) => { form[i][key] = on }) }

  function onkeydown(e) { if (e.key == "Escape") e.preventDefault() }
  function onkeyup(e) {
    if (e.key != "Escape") return
    e.preventDefault()
    // Close all modals.
    Object.keys(isOpen).forEach(k => isOpen[k] = false)
  }
</script>

<svelte:window on:keyup={onkeyup} on:keydown={onkeydown}/>

<div id="container">
  <Table bordered>
    <tr>
      <th><span> <Fa size="sm" icon={faTrashAlt}/> 
        <span class={updating?"":"link"} on:keypress={selectAll} on:click={selectAll}>{$_("words.All")}</span>
      </span></th>
      <th class="d-none d-md-table-cell">ID</th>
      <th class="d-none d-sm-table-cell">
        <Tooltip target="il{instance.App}type">{$_("words.Implementation")}</Tooltip>
        <span id="il{instance.App}type">{$_("words.Type")} <Fa size="xs" primaryColor="darkCyan" icon={faArrowUpRightFromSquare}/></span>
      </th>
      <th>{$_("words.Name")}</th>
      <th>
        <Dropdown size="sm">
          <Tooltip placement="top" target="il{instance.App}enableAutomaticAdd"><T id="instances.AutomaticAddDesc" starrApp={instance.App}/></Tooltip>
          <DropdownToggle id="il{instance.App}enableAutomaticAdd" tag="span" class="link">A <Fa primaryColor="darkCyan" icon={faCircleInfo}/></DropdownToggle>
          <DropdownMenu>
            <DropdownItem header>{$_("instances.EnableAutoAdd")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("enableAutomaticAdd", false)}>{$_("instances.DisableAll")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("enableAutomaticAdd", true)}>{$_("instances.EnableAll")}</DropdownItem>
          </DropdownMenu>
        </Dropdown>
      </th>
      {#if instance.App != "Sonarr"}
      <th>
        {#if instance.App == "Radarr"}
        <Dropdown size="sm">
          <Tooltip placement="top" target="il{instance.App}enabled"><T id="instances.EnableListDesc" starrApp={instance.App}/></Tooltip>
          <DropdownToggle id="il{instance.App}enabled" tag="span" class="link">E <Fa primaryColor="darkCyan" icon={faCircleInfo}/></DropdownToggle>
          <DropdownMenu>
            <DropdownItem header>{$_("configvalues.Enabled")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("enabled", false)}>{$_("instances.DisableAll")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("enabled", true)}>{$_("instances.EnableAll")}</DropdownItem>
          </DropdownMenu>
        </Dropdown>
        {:else}
        <Dropdown size="sm">
          <Tooltip placement="top" target="il{instance.App}shouldMonitorExisting"><T id="instances.MonitorExistingDesc" starrApp={instance.App}/></Tooltip>
          <DropdownToggle id="il{instance.App}shouldMonitorExisting" tag="span" class="link">M <Fa primaryColor="darkCyan" icon={faCircleInfo}/></DropdownToggle>
          <DropdownMenu>
            <DropdownItem header>{$_("instances.MonitorExisting")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("shouldMonitorExisting", false)}>{$_("instances.DisableAll")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("shouldMonitorExisting", true)}>{$_("instances.EnableAll")}</DropdownItem>
          </DropdownMenu>
        </Dropdown>
        {/if}
      </th>
      {/if}
      <th>
        <Dropdown size="sm">
          {#if instance.App == "Sonarr"}
          <Tooltip placement="top" target="il{instance.App}seasonFolder"><T id="instances.ListSeasonFolderDesc"/></Tooltip>
          <DropdownToggle id="il{instance.App}seasonFolder" tag="span" class="link">S <Fa primaryColor="darkCyan" icon={faCircleInfo}/></DropdownToggle>
          <DropdownMenu>
            <DropdownItem header>{$_("instances.SeasonFolder")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("seasonFolder", false)}>{$_("instances.DisableAll")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("seasonFolder", true)}>{$_("instances.EnableAll")}</DropdownItem>
          </DropdownMenu>
          {:else}
          <Tooltip placement="top" target="il{instance.App}shouldSearch"><T id="instances.ListSearchDesc"/></Tooltip>
          <DropdownToggle id="il{instance.App}shouldSearch" tag="span" class="link">S <Fa primaryColor="darkCyan" icon={faCircleInfo}/></DropdownToggle>
          <DropdownMenu>
            <DropdownItem header>{$_("instances.SearchNewItems")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("shouldSearch", false)}>{$_("instances.DisableAll")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("shouldSearch", true)}>{$_("instances.EnableAll")}</DropdownItem>
          </DropdownMenu>
          {/if}
        </Dropdown>
      </th>
    </tr>

    {#each info as client, idx}
      {#if client} <!-- When deleting a client, this protects an error condition. -->
      <tr class={selected[info[idx].id]?" bg-secondary":""}>
        <td><div class="switch"><Input disabled={updating} invalid={selected[info[idx].id]} type="switch" bind:checked={selected[info[idx].id]}/></div></td>

        <td class="d-none d-md-table-cell">{client.id}</td>
        <td class="d-none d-sm-table-cell">
          <Tooltip target="il{instance.App}type{idx}">{client.listType} {info[idx].infoLink}</Tooltip>
          <span id="il{instance.App}type{idx}">
            <Badge color="info"><open-browser href={info[idx].infoLink}>{client.implementation}</open-browser></Badge>
          </span>
        </td>
        <td class={JSON.stringify(form[idx]) != JSON.stringify(info[idx])?"border-warning":""}>
          <a href="/" style="padding-left:0" on:click|preventDefault={() => isOpen[idx]=!updating}>{client.name}</a>
          <Modal class="modal-settings" body size="lg" scrollable isOpen={isOpen[idx]}>
            <ModalHeader toggle={() => {isOpen[idx] = false}}>
              <Badge color="info">{client.id}</Badge> {client.implementation}
            </ModalHeader>

            <ModalBody>
              {#if info[idx].message && info[idx].message.message}
                <p><Card style="padding:0 3px 0 3px" color="warning">{info[idx].message.message}</Card></p>
              {/if}
              <InputGroup>
                <InputGroupText class="setting-name">{$_("words.Name")}</InputGroupText>
                <Input invalid={form[idx].name != info[idx].name} type="text" bind:value={form[idx].name} />
              </InputGroup>
              <InputGroup>
                <InputGroupText class="setting-name">{$_("instances.QualityProfile")}</InputGroupText>
                <Tooltip target="il{instance.App}qualityProfileId"><T id="instances.ListQPDesc"/></Tooltip>
                <Input id="il{instance.App}qualityProfileId" invalid={form[idx].qualityProfileId != info[idx].qualityProfileId} type="select" bind:value={form[idx].qualityProfileId}>
                  {#if qualityProfiles}
                    {#each qualityProfiles as profile}
                      <option value={profile.id}>{profile.name}</option>
                    {/each}
                  {:else}
                    <option value={info[idx].qualityProfileId}>{info[idx].qualityProfileId}</option>
                  {/if}
                </Input>
              </InputGroup>
              {#if form[idx].metadataProfileId !== undefined} <InputGroup>
                <InputGroupText class="setting-name">{$_("instances.MetadataProfile")}</InputGroupText>
                <Tooltip target="il{instance.App}metadataProfileId"><T id="instances.ListMDDesc"/></Tooltip>
                <Input id="il{instance.App}metadataProfileId" invalid={form[idx].metadataProfileId != info[idx].metadataProfileId} type="select" bind:value={form[idx].metadataProfileId}>
                  {#if metadataProfiles}
                    {#each metadataProfiles as profile}
                      <option value={profile.id}>{profile.name}</option>
                    {/each}
                  {:else}
                    <option value={info[idx].metadataProfileId}>{info[idx].metadataProfileId}</option>
                  {/if}
                </Input>
              </InputGroup> {/if}
              <InputGroup>
                <InputGroupText class="setting-name"><T id="instances.RootFolder"/></InputGroupText>
                <Tooltip target="il{instance.App}rootFolderPath"><T id="instances.ListRFDesc"/></Tooltip>
                <Input id="il{instance.App}rootFolderPath" invalid={form[idx].rootFolderPath != info[idx].rootFolderPath} type="select" bind:value={form[idx].rootFolderPath}>
                  {#if rootFolders}
                    {#each rootFolders as dir}
                      <option value={dir.path}>{dir.path}</option>
                    {/each}
                  {:else}
                    <option value={info[idx].rootFolderPath}>{info[idx].rootFolderPath}</option>
                  {/if}
                </Input>
              </InputGroup>
              {#if form[idx].monitorNewItems !== undefined} <InputGroup>
                <InputGroupText class="setting-name"><T id="instances.MonitorNew"/></InputGroupText>
                <Tooltip target="il{instance.App}monitorNewItems"><T id="instances.ListMNDesc"/></Tooltip>
                <Input id="il{instance.App}monitorNewItems" invalid={form[idx].monitorNewItems != info[idx].monitorNewItems} type="select" bind:value={form[idx].monitorNewItems}>
                  <option value="all">{$_("words.All")}</option>
                  <option value="none">{$_("words.None")}</option>
                  <option value="new">{$_("words.New")}</option>
                  {#if !["all","none","new"].includes(info[idx].monitorNewItems)}
                    <option value={info[idx].monitorNewItems}>{info[idx].monitorNewItems}</option>
                  {/if}
                </Input>
              </InputGroup> {/if}
              {#if form[idx].enableAutomaticAdd !== undefined} <InputGroup>
                <InputGroupText class="setting-name"><T id="instances.EnableAutoAdd"/></InputGroupText>
                <Tooltip placement="top" target="il{instance.App}enableAutomaticAdd2"><T id="instances.ListAutoDesc" starrApp={instance.App}/></Tooltip>
                <Input id="il{instance.App}enableAutomaticAdd2" invalid={form[idx].enableAutomaticAdd != info[idx].enableAutomaticAdd}
                  type="select" bind:value={form[idx].enableAutomaticAdd}>
                  <option value={true}>{$_("configvalues.Enabled")}</option>
                  <option value={false}>{$_("configvalues.Disabled")}</option>
                </Input>
              </InputGroup> {/if}
              {#if form[idx].enableAuto !== undefined} <InputGroup>
                <InputGroupText class="setting-name"><T id="instances.EnableAutoAdd"/></InputGroupText>
                <Tooltip placement="top" target="il{instance.App}enableAuto"><T id="instances.ListAutoDesc" starrApp={instance.App}/></Tooltip>
                <Input id="il{instance.App}enableAuto" invalid={form[idx].enableAuto != info[idx].enableAuto}
                  type="select" bind:value={form[idx].enableAuto}>
                  <option value={true}>{$_("configvalues.Enabled")}</option>
                  <option value={false}>{$_("configvalues.Disabled")}</option>
                </Input>
              </InputGroup> {/if}
              {#if form[idx].shouldSearch !== undefined} <InputGroup>
                <InputGroupText class="setting-name"><T id="instances.SearchNewItems"/></InputGroupText>
                <Tooltip target="il{instance.App}shouldSearch2"><T id="instances.ListSearchDesc"/></Tooltip>
                <Input id="il{instance.App}shouldSearch2" invalid={form[idx].shouldSearch != info[idx].shouldSearch}
                  type="select" bind:value={form[idx].shouldSearch}>
                  <option value={true}>{$_("configvalues.Enabled")}</option>
                  <option value={false}>{$_("configvalues.Disabled")}</option>
                </Input>
              </InputGroup> {/if}
              {#if form[idx].shouldMonitorExisting !== undefined} <InputGroup>
                <InputGroupText class="setting-name"><T id="instances.MonitorExisting"/></InputGroupText>
                <Tooltip target="il{instance.App}shouldMonitorExisting2"><T id="instances.ListMonitorExistDesc" starrApp={instance.App}/></Tooltip>
                <Input id="il{instance.App}shouldMonitorExisting2" invalid={form[idx].shouldMonitorExisting != info[idx].shouldMonitorExisting}
                  type="select" bind:value={form[idx].shouldMonitorExisting}>
                  <option value={true}>{$_("configvalues.Enabled")}</option>
                  <option value={false}>{$_("configvalues.Disabled")}</option>
                </Input>
              </InputGroup> {/if}
              {#if form[idx].monitor !== undefined} <InputGroup>
                <InputGroupText class="setting-name"><T id="instances.SearchNewItems"/></InputGroupText>
                <Tooltip target="il{instance.App}monitorNewItems"><T id="instances.ListMonitorNewDesc"/></Tooltip>
                <Input id="il{instance.App}monitorNewItems" invalid={form[idx].monitor != info[idx].monitor} type="select" bind:value={form[idx].monitor}>
                  {#if instance.App == "Radarr"}
                    <option value="movieOnly">{$_("instances.MovieOnly")}</option>
                    <option value="movieAndCollection">{$_("instances.MovieandCollection")}</option>
                    <option value="none">{$_("words.None")}</option>
                  {/if}
                </Input>
              </InputGroup> {/if}
              {#if form[idx].seriesType !== undefined} <InputGroup>
                <InputGroupText class="setting-name">{$_("instances.SeriesType")}</InputGroupText>
                <Input invalid={form[idx].seriesType != info[idx].seriesType}
                  type="select" bind:value={form[idx].seriesType}>
                  {#if instance.App == "Sonarr"}
                    <option value="standard">{$_("words.Standard")}</option>
                    <option value="daily">{$_("words.Daily")}</option>
                    <option value="anime">{$_("words.Anime")}</option>
                  {/if} <!-- whisparr? -->
                </Input>
              </InputGroup> {/if}
              {#if form[idx].shouldMonitor !== undefined} <InputGroup>
                <InputGroupText class="setting-name"><T id="instances.Monitor"/></InputGroupText>
                <Tooltip target="il{instance.App}shouldMonitor"><T id="instances.ListMonitorNewDesc"/></Tooltip>
                <Input id="il{instance.App}shouldMonitor" invalid={form[idx].shouldMonitor != info[idx].shouldMonitor}
                  type="select" bind:value={form[idx].shouldMonitor}>
                  {#if instance.App == "Lidarr"}
                    <option value="entireArtist">{$_("instances.AllArtistAlbums")}</option>
                    <option value="specificAlbum">{$_("instances.SpecificAlbum")}</option>
                    <option value="none">{$_("words.None")}</option>
                  {:else if instance.App == "Readarr"}
                    <option value="specificBook">{$_("instances.SpecificBook")}</option>
                    <option value="entireAuthor">{$_("instances.EntireAuthor")}</option>
                    <option value="none">{$_("words.None")}</option>
                  {:else if instance.App == "Sonarr"}
                    <option value="future">{$_("instances.FutureEpisodes")}</option>
                    <option value="missing">{$_("instances.MissingEpisodes")}</option>
                    <option value="existing">{$_("instances.ExistingEpisodes")}</option>
                    <option value="pilot">{$_("instances.PilotEpisode")}</option>
                    <option value="firstSeason">{$_("instances.OnlyFirstSeason")}</option>
                    <option value="latestSeason">{$_("instances.OnlyLatestSeason")}</option>
                    <option value="monitorSpecials">{$_("instances.MonitorSpecials")}</option>
                    <option value="unmonitorSpecials">{$_("instances.UnmonitorSpecials")}</option>
                    <option value="none">{$_("words.None")}</option>
                  {:else if instance.App == "Whisparr"}
                    <option value="none">{$_("words.None")}</option>
                    {#if info[idx].shouldMonitor != "none"} 
                      <option value={info[idx].shouldMonitor}>{info[idx].shouldMonitor}</option>
                    {/if}
                  {/if}
                </Input>
              </InputGroup> {/if}
              {#if form[idx].minimumAvailability !== undefined} <InputGroup>
                <InputGroupText class="setting-name">{$_("instances.MinimumAvailability")}</InputGroupText>
                <Input invalid={form[idx].minimumAvailability != info[idx].minimumAvailability}
                  type="select" bind:value={form[idx].minimumAvailability}>
                  {#if instance.App == "Radarr"}
                    <option value="announced">{$_("instances.Announced")}</option>
                    <option value="released">{$_("instances.Released")}</option>
                    <option value="inCinemas">{$_("instances.InCinemas")}</option>
                  {/if}
                </Input>
              </InputGroup> {/if}
              {#each info[idx].fields as item, itemIdx}
                {#if item.helpText}
                <Tooltip target="il{instance.App}{item.name.replace(".", "")}{idx}{itemIdx}">{item.helpText}</Tooltip>
                {/if}
                <InputGroup id="il{instance.App}{item.name.replace(".", "")}{idx}{itemIdx}">
                  <InputGroupText class="setting-name">{item.label?item.label:item.name}</InputGroupText>
                  {#if item.type == "select"}
                  <Input type="select" bind:value={form[idx].fields[itemIdx].value}
                    invalid={info[idx].fields[itemIdx].value != form[idx].fields[itemIdx].value}>
                    {#each info[idx].fields[itemIdx].selectOptions as val, valIdx}
                      <option value={valIdx}>{val.name}</option>
                    {/each}
                  </Input>
                  {:else if item.type == "checkbox"}
                  <Input
                    invalid={info[idx].fields[itemIdx].value != form[idx].fields[itemIdx].value}
                    type="select" bind:value={form[idx].fields[itemIdx].value}>
                    <option value={true}>{$_("configvalues.Enabled")}</option>
                    <option value={false}>{$_("configvalues.Disabled")}</option>
                  </Input>
                  {:else}
                  <Input
                    invalid={info[idx].fields[itemIdx].value != form[idx].fields[itemIdx].value}
                    type={item.type} bind:value={form[idx].fields[itemIdx].value}/>
                  {/if}<!-- /if (item.type) -->
                </InputGroup>
              {/each}<!-- /each info[idx].fields as item, itemIdx -->
            </ModalBody>

            <ModalFooter>
              <Button size="sm" disabled={str==JSON.stringify(form)} color="danger" 
                on:click={() => {form[idx] = JSON.parse(JSON.stringify(info[idx]))}}>{$_("words.Reset")}</Button>
              <Button size="sm" color="info" on:click={() => {isOpen[idx] = false}}>{$_("words.Close")}</Button>
            </ModalFooter>
          </Modal>

        </td>
        {#if form[idx].enableAutomaticAdd!==undefined} <td class={form[idx].enableAutomaticAdd!=info[idx].enableAutomaticAdd?"bg-warning":""}>
          <div class="switch"><Input disabled={updating} type="switch" bind:checked={form[idx].enableAutomaticAdd} /></div>
        </td> {/if}
        {#if form[idx].enableAuto!==undefined} <td class={form[idx].enableAuto!=info[idx].enableAuto?"bg-warning":""}>
          <div class="switch"><Input disabled={updating} type="switch" bind:checked={form[idx].enableAuto} /></div>
        </td> {/if}
        {#if form[idx].shouldMonitorExisting!==undefined} <td class={form[idx].shouldMonitorExisting!=info[idx].shouldMonitorExisting?"bg-warning":""}>
          <div class="switch"><Input disabled={updating} type="switch" bind:checked={form[idx].shouldMonitorExisting} /></div>
        </td> {/if}
        {#if form[idx].enabled!==undefined} <td class={form[idx].enabled!=info[idx].enabled?"bg-warning":""}>
          <div class="switch"><Input disabled={updating} type="switch" bind:checked={form[idx].enabled} /></div>
        </td> {/if}
        {#if form[idx].shouldSearch!==undefined} <td class={form[idx].shouldSearch!=info[idx].shouldSearch?"bg-warning":""}>
          <div class="switch"><Input disabled={updating} type="switch" bind:checked={form[idx].shouldSearch} /></div>
        </td> {/if}
        {#if form[idx].searchOnAdd!==undefined} <td class={form[idx].searchOnAdd!=info[idx].searchOnAdd?"bg-warning":""}>
          <div class="switch"><Input disabled={updating} type="switch" bind:checked={form[idx].searchOnAdd} /></div>
        </td> {/if}
        {#if form[idx].seasonFolder!==undefined} <td class={form[idx].seasonFolder!=info[idx].seasonFolder?"bg-warning":""}>
          <div class="switch"><Input disabled={updating} type="switch" bind:checked={form[idx].seasonFolder} /></div>
        </td> {/if}
      </tr>
      {/if}<!-- /if (client) -->
    {/each}<!-- /each info as client, idx -->
  </Table>
  <Footer {instance} bind:selected={selected} identifier="ImportLists"
    bind:updating={updating}  bind:info={info} bind:form={form} bind:str={str}/>
</div><!-- id="container" -->

<style>
  .switch {
    height:20px;
    margin-top: -10px !important;
  }

  :global(.modal-settings .setting-name) {
    min-width:max-content !important;
    width:160px !important;
  }
</style>
