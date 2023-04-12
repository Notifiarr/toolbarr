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
    UpdateLidarrIndexer,
    UpdateProwlarrIndexer,
    UpdateRadarrIndexer,
    UpdateReadarrIndexer,
    UpdateSonarrIndexer,
    UpdateWhisparrIndexer,
    DeleteIndexer,
  } from "../../../../wailsjs/go/starrs/Starrs"

  let isOpen = {}       // Modal toggle control.
  let modalOpen = false // True if any modal is open.
  let updating = false  // True while doing updates.
  let rawOpen = false   // Dev mode toggle button.
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

  // save id (id=id), save all (id=true), save selected (id=false) (not used)
  async function update(id, force) {
    toast("info", $_("instances.UpdatingIndexers"))
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

  async function deleteIndexers() {
    toast("info", $_("instances.DeletingIndexers", {values:{"count": count(selected)}}))
    goodMsg = badMsg = ""
    updating = true

    for (var idx = form.length-1; idx >= 0; idx--) {
      if (!selected[form[idx].id]) continue // Not selected.
      await DeleteIndexer(instance, form[idx].id).then(
        (msg) => showMsg(idx, msg), (err) => showError(idx, err)
      )
    }

    updating = false
    Object.keys(selected).forEach(k => selected[k] = false)
  }

  function showMsg(idx, msg, data) {
    goodMsg += `<li>${$_("instances.SuccessMsg", {values:{"msg": msg}})}</li>`
    if (data) { // update indexer (repalce in place)
      const istr = JSON.stringify(data)
      info[idx] = JSON.parse(istr)
      form[idx] = JSON.parse(istr)
    } else { // delete indexer (delete from list)
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
      <th class="d-none d-sm-table-cell"><span>{$_("words.Type")} <Fa size="xs" primaryColor="darkCyan" icon={faArrowUpRightFromSquare}/></span></th>
      <th>{$_("words.Name")}</th>
      <th>
        <Dropdown size="sm">
          <DropdownToggle tag="span" class="link">RS <Fa primaryColor="darkCyan" icon={faCircleInfo}/></DropdownToggle>
          <DropdownMenu>
            <DropdownItem header>{$_("instances.EnableRSS")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("enableRss", false)}>{$_("instances.DisableAll")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("enableRss", true)}>{$_("instances.EnableAll")}</DropdownItem>
          </DropdownMenu>
        </Dropdown>
      </th>
      <th>
        <Dropdown size="sm">
          <DropdownToggle tag="span" class="link">IS <Fa primaryColor="darkCyan" icon={faCircleInfo}/></DropdownToggle>
          <DropdownMenu>
            <DropdownItem header>{$_("instances.EnableInteractiveSearch")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("enableInteractiveSearch", false)}>{$_("instances.DisableAll")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("enableInteractiveSearch", true)}>{$_("instances.EnableAll")}</DropdownItem>
          </DropdownMenu>
        </Dropdown>
      </th>
      <th>
        <Dropdown size="sm">
          <DropdownToggle tag="span" class="link">AS <Fa primaryColor="darkCyan" icon={faCircleInfo}/></DropdownToggle>
          <DropdownMenu>
            <DropdownItem header>{$_("instances.EnableAutomaticSearch")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("enableAutomaticSearch", false)}>{$_("instances.DisableAll")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("enableAutomaticSearch", true)}>{$_("instances.EnableAll")}</DropdownItem>
          </DropdownMenu>
        </Dropdown>
      </th>
    </tr>

    {#each info as indexer, idx}
      {#if indexer} <!-- When deleting an indexer, this protects an error condition. -->
      <tr class={selected[info[idx].id]?" bg-secondary":""}>
        <td><div class="switch"><Input disabled={updating} type="switch" bind:checked={selected[info[idx].id]}/></div></td>

        <td class="d-none d-md-table-cell">{indexer.id}</td>
        <td class="d-none d-sm-table-cell"><span><Badge color="info" class="superbadge">
          <open-browser href={info[idx].infoLink}>{indexer.implementation}</open-browser>
        </Badge></span></td>
        <td>
          <a href="/" style="padding-left:0" on:click|preventDefault={() => modalOpen=isOpen[idx]=!updating}>{indexer.name}</a>
          <Modal class="modal-settings" body size="lg" scrollable isOpen={isOpen[idx]}>
            <ModalHeader toggle={() => reset(idx)}>
              <Badge color="info">{indexer.id}</Badge> {indexer.implementation}
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
                  <select disabled style="width:calc(99% - 160px)" multiple bind:value={form[idx].fields[itemIdx].value}>
                    {#if typeof item.value == "object"}
                      {#each info[idx].fields[itemIdx].value as val}
                      <option value={val}>{val}</option>
                      {/each}
                    {:else}
                      <option value={info[idx].fields[itemIdx].value}>{info[idx].fields[itemIdx].value}</option>
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
        <td class={form[idx].enableRss!=info[idx].enableRss?"bg-warning":""}>
          <div class="switch"><Input disabled={updating} type="switch" bind:checked={form[idx].enableRss} /></div>
        </td>
        <td class={form[idx].enableInteractiveSearch!=info[idx].enableInteractiveSearch?"bg-warning":""}>
          <div class="switch"><Input disabled={updating} type="switch" bind:checked={form[idx].enableInteractiveSearch} /></div>
        </td>
        <td class={form[idx].enableAutomaticSearch!=info[idx].enableAutomaticSearch?"bg-warning":""}>
          <div class="switch"><Input disabled={updating} type="switch" bind:checked={form[idx].enableAutomaticSearch} /></div>
        </td>
      </tr>
      {/if}<!-- /if (indexer) -->
    {/each}<!-- /each info as indexer, idx -->
  </Table>

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
        <Button class="actions" color="danger" on:click={deleteIndexers}><T id="instances.DeleteSelected" count={selectedCount}/></Button>
      </Fade>
    </Collapse>
  {/if}<!-- /if (instance.App) -->
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

  #container :global(button.actions) {
    padding:0px 6px 0px 6px;
    margin:4px 1px;
  }

  #container :global(.link) {
    cursor: pointer;
    text-decoration: underline;
  }
</style>