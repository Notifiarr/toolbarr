<script>
  export let info
  export let instance

  import T, { _ } from "../../../libs/Translate.svelte"
  import Fa from "svelte-fa"
  import { faCircleInfo, faArrowUpRightFromSquare, faTrashAlt, faClose, faCaretDown, faCaretUp } from "@fortawesome/free-solid-svg-icons"
  import { toast } from "../../../libs/funcs"
  import { conf } from "../../../libs/config"
  import {
    Alert,
    Badge,
    Button,
    Card,
    Collapse,
    Dropdown,
    DropdownItem,
    DropdownMenu,
    DropdownToggle,
    Input,
    InputGroup,
    InputGroupText,
    Modal,
    Spinner,
    Table,
    Tooltip,
  } from "sveltestrap"
  import {
    UpdateLidarrIndexer,
    UpdateProwlarrIndexer,
    UpdateRadarrIndexer,
    UpdateReadarrIndexer,
    UpdateSonarrIndexer,
    UpdateWhisparrIndexer,
    DeleteIndexer,
  } from "../../../../wailsjs/go/starrs/Starrs"

  let isOpen = {} // Modal control
  let modalOpen = false
  let goodMsg = ""
  let badMsg = ""
  let updating = 0
  let rawOpen = false

  let str = JSON.stringify(info)
  info = JSON.parse(str) // use this to compare for changes
  let form = JSON.parse(str) // changes go here.

  $: unSaved = JSON.stringify(form) !== str

  let updateMethod
  switch(instance.App) {
    case "Lidarr":
      updateMethod = UpdateLidarrIndexer; break
    case "Prowlarr":
      updateMethod = UpdateProwlarrIndexer; break
    case "Radarr":
      updateMethod = UpdateRadarrIndexer; break
    case "Readarr":
      updateMethod = UpdateReadarrIndexer; break
    case "Sonarr":
      updateMethod = UpdateSonarrIndexer; break
    case "Whisparr":
      updateMethod = UpdateWhisparrIndexer; break
  }

  // save id (id=id), save all (id=true), save selected (id=false)
  async function update(id, force) {
    toast("info", $_("instances.UpdatingIndexers"))
    goodMsg = badMsg = ""

    for (let i = 0; i < form.length; i++) {
      modalOpen = isOpen[i] = false // close all nodals
      if (id === false && !selected[form[i].id]) continue
      if (![true,false,form[i].id].includes(id)) continue
      if (JSON.stringify(form[i]) == JSON.stringify(info[i])) continue
      updating++

      await updateMethod(instance, force, form[i]).then(
        (resp) => {
          const str = JSON.stringify(resp.Data)
          info[i] = JSON.parse(str)
          form[i] = JSON.parse(str)
          goodMsg +=  "<li>" + $_("instances.SuccessMsg", {values:{"msg": resp.Msg}}) + "</li>"
          updating--
        },
        (err) => {
          reset(i)
          badMsg +=  "<li>" + $_("instances.ErrorMsg", {values:{"msg": err}}) + "</li>"
          updating--
        }
      )
    }

    str = JSON.stringify(info)
    selected = {}
  }

  async function deleteIndexers() {
    toast("info", $_("instances.DeletingIndexers", {values:{"count": count(selected)}}))
    goodMsg = badMsg = ""

    for (let i = form.length-1; i >= 0; i--) {
      if (!selected[form[i].id]) continue
      updating++

      await DeleteIndexer(instance, form[i].id).then(
        (msg) => {
          delete info[i]
          delete form[i]
          str = JSON.stringify(info)
          info = JSON.parse(str)
          form = JSON.parse(str)
          goodMsg +=  "<li>" + $_("instances.SuccessMsg", {values:{"msg": msg}}) + "</li>"
          updating--
        },
        (err) => {
          reset(i)
          badMsg +=  "<li>" + $_("instances.ErrorMsg", {values:{"msg": err}}) + "</li>"
          updating--
        }
      )
    }

    selected = {}
  }

  function reset(idx) {
    modalOpen = isOpen[idx] = false
    form[idx] = JSON.parse(JSON.stringify(info[idx]))
  }

  let selected = {}
  $: selectedCount = count(selected)

  function count(selected) {
    let counter = 0
    for (var k in selected) if (selected[k]) counter++;
    return counter
  }

  function toggleAll(key, on) {
    form.forEach((_, i) => {form[i][key] = on})
  }

  let all = false
  function selectAll() {
    all = !all
    Object.keys(selected).forEach(k => selected[k] = all)
  }
</script>

<div id="container">
{#if !info} <Spinner/> {$_("words.Loading")} {instance.Name} ... {:else}
  <Table bordered>
    <tr>
      <th><span><Fa size="sm" icon={faTrashAlt}/> <span class="link" on:keypress={selectAll} on:click={selectAll}>{$_("words.All")}</span></span></th>
      <th class="d-none d-md-table-cell">ID</th>
      <th class="d-none d-sm-table-cell"><span>{$_("words.Type")} <Fa size="xs" primaryColor="darkCyan" icon={faArrowUpRightFromSquare}/></span></th>
      <th>{$_("words.Name")}</th>
      <th>
        <Dropdown size="sm">
          <DropdownToggle tag="span" class="link">RS <Fa primaryColor="darkCyan" icon={faCircleInfo}/></DropdownToggle>
          <DropdownMenu>
            <DropdownItem header>{$_("instances.EnableRSS")}</DropdownItem>
            <DropdownItem on:click={() => toggleAll("enableRss", false)}>{$_("instances.DisableAll")}</DropdownItem>
            <DropdownItem on:click={() => toggleAll("enableRss", true)}>{$_("instances.EnableAll")}</DropdownItem>
          </DropdownMenu>
        </Dropdown>
      </th>
      <th>
        <Dropdown size="sm">
          <DropdownToggle tag="span" class="link">IS <Fa primaryColor="darkCyan" icon={faCircleInfo}/></DropdownToggle>
          <DropdownMenu>
            <DropdownItem header>{$_("instances.EnableInteractiveSearch")}</DropdownItem>
            <DropdownItem on:click={() => toggleAll("enableInteractiveSearch", false)}>{$_("instances.DisableAll")}</DropdownItem>
            <DropdownItem on:click={() => toggleAll("enableInteractiveSearch", true)}>{$_("instances.EnableAll")}</DropdownItem>
          </DropdownMenu>
        </Dropdown>
      </th>
      <th>
        <Dropdown size="sm">
          <DropdownToggle tag="span" class="link">AS <Fa primaryColor="darkCyan" icon={faCircleInfo}/></DropdownToggle>
          <DropdownMenu>
            <DropdownItem header>{$_("instances.EnableAutomaticSearch")}</DropdownItem>
            <DropdownItem on:click={() => toggleAll("enableAutomaticSearch", false)}>{$_("instances.DisableAll")}</DropdownItem>
            <DropdownItem on:click={() => toggleAll("enableAutomaticSearch", true)}>{$_("instances.EnableAll")}</DropdownItem>
          </DropdownMenu>
        </Dropdown>
      </th>
    </tr>

    {#each info as indexer, idx}
      {#if indexer} <!-- When deleting an indexer, this protects an error condition. -->
      <tr class={selected[indexer.id]?" bg-secondary":""}>
        <td><div class="switch"><Input type="switch" bind:checked={selected[indexer.id]}/></div></td>

        <td class="d-none d-md-table-cell">{indexer.id}</td>
        <td class="d-none d-sm-table-cell"><span><Badge color="info" class="superbadge">
          <open-browser href={info[idx].infoLink}>{indexer.implementation}</open-browser>
        </Badge></span></td>
        <td>
          {#if updating > 0} {indexer.name} {:else}
            <a href="/" style="padding-left:0" on:click|preventDefault={() => modalOpen=isOpen[idx]=true}>{indexer.name}</a>
          {/if}
          <Modal class="modal-settings" body size="lg" scrollable isOpen={isOpen[idx]}>
            <h4>
              <Badge color="info">{indexer.id}</Badge> {indexer.implementation}
              <a href="/" on:click|preventDefault={()=>reset(idx)}><Fa pull="right" icon={faClose} color="red"/></a>
            </h4>
            <InputGroup>
              <InputGroupText class="setting-name">{$_("words.Name")}</InputGroupText>
              <Input invalid={form[idx].name != info[idx].name} type="text" bind:value={form[idx].name} />
            </InputGroup>
            <InputGroup>
              <InputGroupText class="setting-name">{$_("words.Priority")}</InputGroupText>
              <Input invalid={form[idx].priority != info[idx].priority} type="number" bind:value={form[idx].priority} />
            </InputGroup>
            {#each info[idx].fields as item, itemIdx}
              {#if item.helpText}
              <Tooltip target="{item.name.replace(".", "")}{idx}{itemIdx}">{item.helpText}</Tooltip>
              {/if}
              <InputGroup id="{item.name.replace(".", "")}{idx}{itemIdx}">
                <InputGroupText class="setting-name">{item.label}</InputGroupText>
                {#if item.type == "select"}
                <select disabled style="width:calc(99% - 161px)" multiple bind:value={form[idx].fields[itemIdx].value}>
                  {#if typeof info[idx].fields[itemIdx].value == "object"}
                    {#each info[idx].fields[itemIdx].value as val}
                    <option value={val}>{val}</option>
                    {/each}
                  {/if}
                </select>
                {:else if item.type == "checkbox"}
                <Input
                  invalid={info[idx].fields[itemIdx].value != form[idx].fields[itemIdx].value &&
                          (info[idx].fields[itemIdx].value != undefined ||
                          form[idx].fields[itemIdx].value != "")}
                  type="select" bind:value={form[idx].fields[itemIdx].value}>
                  <option value={true}>{$_("configvalues.Enabled")}</option>
                  <option value={false}>{$_("configvalues.Disabled")}</option>
                </Input>
                {:else}
                <Input
                invalid={info[idx].fields[itemIdx].value != form[idx].fields[itemIdx].value &&
                        (info[idx].fields[itemIdx].value != undefined ||
                        form[idx].fields[itemIdx].value != "")}
                type={item.type} bind:value={form[idx].fields[itemIdx].value}/>
                {/if}
              </InputGroup>
            {/each}
            {#if instance.App == "Prowlarr"}
              This tool does not yet work with Prowlarr.<br>
            {:else}
              <Button color="success" on:click={() => update(form[idx].id, false)}>{$_("instances.TestandSave")}</Button>
              <Tooltip target="forceSave"><T id="instances.ForceSaveDesc" starrApp={instance.App}/></Tooltip>
              <Button id="forceSave" color="info" on:click={() => update(form[idx].id, true)}>{$_("instances.ForceSave")}</Button>
            {/if}
            <Button color="danger" on:click={() => reset(idx)}>{$_("words.Cancel")}</Button>
          </Modal>

        </td>
        <td class={form[idx].enableRss!=info[idx].enableRss?"bg-warning":""}>
          <div class="switch"><Input type="switch" bind:checked={form[idx].enableRss} /></div>
        </td>
        <td class={form[idx].enableInteractiveSearch!=info[idx].enableInteractiveSearch?"bg-warning":""}>
          <div class="switch"><Input type="switch" bind:checked={form[idx].enableInteractiveSearch} /></div>
        </td>
        <td class={form[idx].enableAutomaticSearch!=info[idx].enableAutomaticSearch?"bg-warning":""}>
          <div class="switch"><Input type="switch" bind:checked={form[idx].enableAutomaticSearch} /></div>
        </td>
      </tr>
      {/if}
    {/each}
  </Table>

  {#if goodMsg != ""}<Alert dismissible color="success">{@html goodMsg}</Alert>{/if}
  {#if badMsg != ""}<Alert dismissible color="danger">{@html badMsg}</Alert>{/if}
  {#if instance.App != "Prowlarr"}
    {#if updating > 0}
      <Card body color="secondary">
        <span>
          <Spinner size="sm" color="info" />
          <h5 style="display:inline-block">{$_("words.Loading")} ...</h5>
        </span>
      </Card>
    {:else if !modalOpen}
      {#if unSaved}
        <Button class="actions" color="success" on:click={() => update(true, false)}>{$_("instances.TestandSave")}</Button>
        <Tooltip target="forceSave"><T id="instances.ForceSaveDesc" starrApp={instance.App}/></Tooltip>
        <Button id="forceSave" class="actions" color="info" on:click={() => update(true, true)}>{$_("instances.ForceSave")}</Button>
      {/if}
      {#if selectedCount > 0}
        <Button class="actions delete" color="danger" on:click={deleteIndexers}><T id="instances.DeleteSelected" count={selectedCount}/></Button>
      {/if}
    {/if}
  {/if}
{/if}
{#if instance.App == "Prowlarr"}
  {$_("instances.ProwlarrNotSupported")}
{/if}
</div>

{#if $conf.DevMode}
  <hr>
  <Button size="sm" on:click={() => (rawOpen = !rawOpen)} class="mb-1">Raw Data <Fa icon={rawOpen?faCaretDown:faCaretUp}/></Button>
  <Card color="secondary">
    <Collapse isOpen={rawOpen}><code><pre class="code">{JSON.stringify(info, null, 3)}</pre></code></Collapse>
  </Card>
{/if}

<style>
  .switch {
    height:20px;
    margin-top: -10px !important;
  }

   :global(.modal-settings .setting-name) {
    min-width:max-content !important;
    width:160px !important;
  }

  #container :global(button.actions) {
    padding:0px 6px 0px 6px;
    margin:4px 1px;
  }

  #container :global(.link) {
    cursor:pointer;
    text-decoration:underline;
  }
</style>