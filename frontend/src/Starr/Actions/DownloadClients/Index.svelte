<script>
  export let info
  export let instance

  import T, { _ } from "../../../libs/Translate.svelte"
  import Fa from "svelte-fa"
  import { toast } from "../../../libs/funcs"
  import Loading from "../loading.svelte"
  import { faCircleInfo, faArrowUpRightFromSquare, faTrashAlt } from "@fortawesome/free-solid-svg-icons"
  import {
    Alert,
    Badge,
    Button,
    Collapse,
    Dropdown,
    DropdownItem,
    DropdownMenu,
    DropdownToggle,
    Fade,
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
  import {
    UpdateLidarrDownloadClient,
    UpdateProwlarrDownloadClient,
    UpdateRadarrDownloadClient,
    UpdateReadarrDownloadClient,
    UpdateSonarrDownloadClient,
    UpdateWhisparrDownloadClient,
    DeleteDownloader,
  } from "../../../../wailsjs/go/starrs/Starrs"

  let isOpen = {}       // Modal toggle control.
  let modalOpen = false // True if any modal is open.
  let updating = false  // True while doing updates.
  let all = false       // Toggle for select-all link.
  let goodMsg = ""      // Card color="success"
  let badMsg = ""       // Card color="danger"
  let selected = {}     // Rows selected by key: ID.

  let str = JSON.stringify(info) // Used for equivalence comparison.
  info = JSON.parse(str)         // Use this to compare for changes.
  let form = JSON.parse(str)     // Form changes go here.
  $: unSaved = JSON.stringify(form) !== str // True when something changed.
  $: selectedCount = count(selected)        // How many items are selected.

  let updateMethod // Each app has a unique update method.
  switch(instance.App) {
    case "Lidarr":
      updateMethod = UpdateLidarrDownloadClient; break
    case "Prowlarr":
      updateMethod = UpdateProwlarrDownloadClient; break
    case "Radarr":
      updateMethod = UpdateRadarrDownloadClient; break
    case "Readarr":
      updateMethod = UpdateReadarrDownloadClient; break
    case "Sonarr":
      updateMethod = UpdateSonarrDownloadClient; break
    case "Whisparr":
      updateMethod = UpdateWhisparrDownloadClient; break
  }

  // save id (id=id), save all (id=true), save selected (id=false) (not used)
  async function update(id, force) {
    toast("info", $_("instances.UpdatingDownloadClients"))
    goodMsg = badMsg = ""
    updating = true

    for (var idx = 0; idx < form.length; idx++) {
      modalOpen = isOpen[idx] = false // close all nodals
      if (id === false && !selected[form[idx].id]) continue // not selected
      if (![true,false,form[idx].id].includes(id)) continue // not a save all
      if (JSON.stringify(form[idx]) == JSON.stringify(info[idx])) continue // not changed

      await updateMethod(instance, force, form[idx]).then(
        (resp) => showMsg(idx, resp.Msg, resp.Data), (err) => showError(idx, err)
      )
    }

    updating = false
    str = JSON.stringify(info)
    Object.keys(selected).forEach(k => selected[k] = false)
  }

  async function deleteDownloaders() {
    toast("info", $_("instances.DeletingDownloadClients", {values:{"count": count(selected)}}))
    goodMsg = badMsg = ""
    updating = true

    for (var idx = form.length-1; idx >= 0; idx--) {
      if (!selected[form[idx].id]) continue // Not selected.
      await DeleteDownloader(instance, form[idx].id).then(
        (msg) => showMsg(idx, msg), (err) => showError(idx, err)
      )
    }

    updating = false
    Object.keys(selected).forEach(k => selected[k] = false)
  }

  function showMsg(idx, msg, data) {
    goodMsg += `<li>${$_("instances.SuccessMsg", {values:{"msg": msg}})}</li>`
    if (data) { // update client (repalce in place)
      const istr = JSON.stringify(data)
      info[idx] = JSON.parse(istr)
      form[idx] = JSON.parse(istr)
    } else { // delete client (delete from list)
      delete form[idx]
      str = JSON.stringify(form)
      info = JSON.parse(str)
      form = JSON.parse(str)
    }
  }

  function showError(idx, err) {
    reset(idx)
    badMsg += `<li>${$_("instances.ErrorMsg", {values:{"msg": err}})}</li>`
  }

  function reset(idx) {
    modalOpen = isOpen[idx] = false
    form[idx] = JSON.parse(JSON.stringify(info[idx]))
  }

  function count(selected) {
    let counter = 0
    for (var k in selected) if (selected[k]) counter++
    return counter
  }

  function selectAll() {
    if (updating) return
    all = !all
    Object.keys(selected).forEach(k => selected[k] = all)
  }

  function toggleAll(key, on) {
    form.forEach((_, i) => {form[i][key] = on})
  }
</script>

<div id="container">
  <Table bordered>
    <tr>
      <th><span> <Fa size="sm" icon={faTrashAlt}/> 
        <span class={updating?"":"link"} on:keypress={selectAll} on:click={selectAll}>{$_("words.All")}</span>
      </span></th>
      <th class="d-none d-md-table-cell">ID</th>
      <th class="d-none d-sm-table-cell">
        <Tooltip target="type">{$_("words.Protocol")}</Tooltip>
        <span id="type">{$_("words.Type")} <Fa size="xs" primaryColor="darkCyan" icon={faArrowUpRightFromSquare}/></span>
      </th>
      <th>{$_("words.Name")}</th>
      <th>
        <Dropdown size="sm">
          <DropdownToggle tag="span" class="link">CD <Fa primaryColor="darkCyan" icon={faCircleInfo}/></DropdownToggle>
          <DropdownMenu>
            <DropdownItem header>{$_("instances.RemoveCompletedDownloads")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("removeCompletedDownloads", false)}>{$_("instances.DisableAll")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("removeCompletedDownloads", true)}>{$_("instances.EnableAll")}</DropdownItem>
          </DropdownMenu>
        </Dropdown>
      </th>
      <th>
        <Dropdown size="sm">
          <DropdownToggle tag="span" class="link">FD <Fa primaryColor="darkCyan" icon={faCircleInfo}/></DropdownToggle>
          <DropdownMenu>
            <DropdownItem header>{$_("instances.RemoveFailedDownloads")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("removeFailedDownloads", false)}>{$_("instances.DisableAll")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("removeFailedDownloads", true)}>{$_("instances.EnableAll")}</DropdownItem>
          </DropdownMenu>
        </Dropdown>
      </th>
      <th>
        <span>{$_("words.Priority")}</span>
      </th>
    </tr>

    {#each info as client, idx}
      {#if client} <!-- When deleting a client, this protects an error condition. -->
      <tr class={selected[info[idx].id]?" bg-secondary":""}>
        <td><div class="switch"><Input disabled={updating} type="switch" bind:checked={selected[info[idx].id]}/></div></td>

        <td class="d-none d-md-table-cell">{client.id}</td>
        <td class="d-none d-sm-table-cell">
          <Tooltip target="type{idx}">{client.protocol} {info[idx].infoLink}</Tooltip>
          <span id="type{idx}">
            <Badge color="info" class="superbadge">
              <open-browser href={info[idx].infoLink}>{client.implementation}</open-browser>
            </Badge>
          </span>
        </td>
        <td>
          <a href="/" style="padding-left:0" on:click|preventDefault={() => modalOpen=isOpen[idx]=!updating}>{client.name}</a>
          <Modal class="modal-settings" body size="lg" scrollable isOpen={isOpen[idx]}>
            <ModalHeader toggle={() => reset(idx)}>
              <Badge color="info">{client.id}</Badge> {client.implementation}
            </ModalHeader>

            <ModalBody>
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
                  <InputGroupText class="setting-name">{item.label?item.label:item.name}</InputGroupText>
                  {#if item.type == "select"}
                  <Input type="select" bind:value={form[idx].fields[itemIdx].value}>
                    {#each info[idx].fields[itemIdx].selectOptions as val, valIdx}
                      <option value={valIdx}>{val.name}</option>
                    {/each}
                  </Input>
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
                  {/if}<!-- /if (item.type) -->
                </InputGroup>
              {/each}<!-- /each info[idx].fields as item, itemIdx -->
            </ModalBody>

            <ModalFooter>
              {#if instance.App == "Prowlarr"}
                <p>{$_("instances.ProwlarrNotSupported")}</p>
              {:else}
                <Button size="sm" color="success" on:click={() => update(form[idx].id, false)}>{$_("instances.TestandSave")}</Button>
                <Tooltip target="forceSave"><T id="instances.ForceSaveDesc" starrApp={instance.App}/></Tooltip>
                <Button size="sm" id="forceSave" color="info" on:click={() => update(form[idx].id, true)}>{$_("instances.ForceSave")}</Button>
              {/if}
              <Button size="sm" color="danger" on:click={() => reset(idx)}>{$_("words.Cancel")}</Button>
            </ModalFooter>
          </Modal>

        </td>
        <td class={form[idx].removeCompletedDownloads!=info[idx].removeCompletedDownloads?"bg-warning":""}>
          <div class="switch"><Input disabled={updating} type="switch" bind:checked={form[idx].removeCompletedDownloads} /></div>
        </td>
        <td class={form[idx].removeFailedDownloads!=info[idx].removeFailedDownloads?"bg-warning":""}>
          <div class="switch"><Input disabled={updating} type="switch" bind:checked={form[idx].removeFailedDownloads} /></div>
        </td>
        <Tooltip target="priority{idx}">{form[idx].priority}</Tooltip>
        <td id="priority{idx}" class={form[idx].priority!=info[idx].priority?"bg-warning":""}>
          <div class="range"> <Input disabled={updating} bsSize="sm" type="range" min={1} max={100} bind:value={form[idx].priority} /></div>
        </td>
      </tr>
      {/if}<!-- /if (client) -->
    {/each}<!-- /each info as client, idx -->
  </Table>
</div><!-- id="container" -->

<div id="footer">
  <Alert isOpen={goodMsg != ""} dismissible color="success">{@html goodMsg}</Alert>
  <Alert isOpen={badMsg != ""} dismissible color="danger">{@html badMsg}</Alert>
  <Loading isOpen={updating}/>

  {#if instance.App == "Prowlarr"}
    {$_("instances.ProwlarrNotSupported")}
  {:else}
    <Collapse isOpen={!updating && !modalOpen && (unSaved||selectedCount > 0)}>
      <Fade style="display:inline-block" isOpen={unSaved}>
        <Button class="actions" color="success" on:click={() => update(true, false)}>{$_("instances.TestandSave")}</Button>
        <Tooltip target="forceSave"><T id="instances.ForceSaveDesc" starrApp={instance.App}/></Tooltip>
        <Button id="forceSave" class="actions" color="info" on:click={() => update(true, true)}>{$_("instances.ForceSave")}</Button>
      </Fade>
      <Fade style="display:inline-block" isOpen={selectedCount > 0}>
        <Button class="actions" color="danger" on:click={deleteDownloaders}><T id="instances.DeleteSelected" count={selectedCount}/></Button>
      </Fade>
    </Collapse>
  {/if}<!-- /if (instance.App) -->
</div> <!-- id="footer" -->

<style>
  .switch {
    height:20px;
    margin-top: -10px !important;
  }

  .range {
    height:20px;
    margin-top: -25px !important;
  }

   :global(.modal-settings .setting-name) {
    min-width:max-content !important;
    width:160px !important;
  }

  #footer :global(button.actions) {
    padding:0px 6px 0px 6px;
    margin:4px 1px;
  }

  #container :global(.link) {
    cursor: pointer;
    text-decoration: underline;
  }
</style>