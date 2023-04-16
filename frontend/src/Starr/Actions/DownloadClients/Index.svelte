<script>
  export let info
  export let instance

  import { _ } from "../../../libs/Translate.svelte"
  import Fa from "svelte-fa"
  import Footer from "../footer.svelte"
  import { fixFieldValues } from "../methods"
  import { faCircleInfo, faArrowUpRightFromSquare, faTrashAlt } from "@fortawesome/free-solid-svg-icons"
  import {
    Badge,
    Button,
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
        <Tooltip target="dc{instance.App}Type">{$_("words.Protocol")}</Tooltip>
        <span id="dc{instance.App}Type">{$_("words.Type")} <Fa size="xs" primaryColor="darkCyan" icon={faArrowUpRightFromSquare}/></span>
      </th>
      <th>{$_("words.Name")}</th>
      {#if instance.App == "Readarr"}
      <th>
        <Dropdown size="sm">
          <Tooltip placement="top" target="dc{instance.App}enable">{$_("instances.EnableDLCDesc")}</Tooltip>
          <DropdownToggle id="dc{instance.App}enable" tag="span" class="link">{$_("configvalues.Enabled")} <Fa primaryColor="darkCyan" icon={faCircleInfo}/></DropdownToggle>
          <DropdownMenu>
            <DropdownItem header>{$_("configvalues.Enabled")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("enable", false)}>{$_("instances.DisableAll")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("enable", true)}>{$_("instances.EnableAll")}</DropdownItem>
          </DropdownMenu>
        </Dropdown>
      </th>
      {:else}
      <th>
        <Dropdown size="sm">
          <Tooltip placement="top" target="dc{instance.App}removeCompletedDownloads">{$_("instances.DLCCompleteDesc")}</Tooltip>
          <DropdownToggle id="dc{instance.App}removeCompletedDownloads" tag="span" class="link">CD <Fa primaryColor="darkCyan" icon={faCircleInfo}/></DropdownToggle>
          <DropdownMenu> 
            <DropdownItem header>{$_("instances.RemoveCompletedDownloads")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("removeCompletedDownloads", false)}>{$_("instances.DisableAll")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("removeCompletedDownloads", true)}>{$_("instances.EnableAll")}</DropdownItem>
          </DropdownMenu>
        </Dropdown>
      </th>
      <th>
        <Dropdown size="sm">
          <Tooltip placement="top" target="dc{instance.App}removeFailedDownloads">{$_("instances.DLCFailedDesc")}</Tooltip>
          <DropdownToggle id="dc{instance.App}removeFailedDownloads" tag="span" class="link">FD <Fa primaryColor="darkCyan" icon={faCircleInfo}/></DropdownToggle>
          <DropdownMenu>
            <DropdownItem header>{$_("instances.RemoveFailedDownloads")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("removeFailedDownloads", false)}>{$_("instances.DisableAll")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("removeFailedDownloads", true)}>{$_("instances.EnableAll")}</DropdownItem>
          </DropdownMenu>
        </Dropdown>
      </th>
      {/if}
      <th>
        <span>{$_("words.Priority")}</span>
      </th>
    </tr>

    {#each info as client, idx}
      {#if client} <!-- When deleting a client, this protects an error condition. -->
      <tr class={selected[info[idx].id]?" bg-secondary":""}>
        <td><div class="switch"><Input disabled={updating} invalid={selected[info[idx].id]} type="switch" bind:checked={selected[info[idx].id]}/></div></td>

        <td class="d-none d-md-table-cell">{client.id}</td>
        <td class="d-none d-sm-table-cell">
          <span id="dc{instance.App}Type{idx}">
            <Badge color="info"><open-browser href={info[idx].infoLink}>{client.implementation}</open-browser></Badge>
          </span>
          <Tooltip target="dc{instance.App}Type{idx}">{client.protocol} {info[idx].infoLink}</Tooltip>
        </td>
        <td class={JSON.stringify(form[idx]) != JSON.stringify(info[idx])?"border-warning":""}>
          <a href="/" style="padding-left:0" on:click|preventDefault={() => isOpen[idx]=!updating}>{client.name}</a>
          <Modal class="modal-settings" body size="lg" scrollable isOpen={isOpen[idx]}>
            <ModalHeader toggle={() => {isOpen[idx] = false}}>
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
                <Tooltip target="dc{instance.App}{item.name.replace(".", "")}{idx}{itemIdx}">{item.helpText}</Tooltip>
                {/if}
                <InputGroup id="dc{instance.App}{item.name.replace(".", "")}{idx}{itemIdx}">
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
              {#if instance.App == "Prowlarr"}
                <p>{$_("instances.ProwlarrNotSupported")}</p>
              {:else}
                <Button size="sm" disabled={str==JSON.stringify(form)} color="danger" 
                  on:click={() => {form[idx] = JSON.parse(JSON.stringify(info[idx]))}}>{$_("words.Reset")}</Button>
              {/if}
              <Button size="sm" color="info" on:click={() => {isOpen[idx] = false}}>{$_("words.Close")}</Button>
            </ModalFooter>
          </Modal>
        </td>
        {#if instance.App == "Readarr"}
        <td class={form[idx].enable!=info[idx].enable?"bg-warning":""}>
          <div class="switch"><Input disabled={updating} type="switch" bind:checked={form[idx].enable} /></div>
        </td>
        {:else}
        <td class={form[idx].removeCompletedDownloads!=info[idx].removeCompletedDownloads?"bg-warning":""}>
          <div class="switch"><Input disabled={updating} type="switch" bind:checked={form[idx].removeCompletedDownloads} /></div>
        </td>
        <td class={form[idx].removeFailedDownloads!=info[idx].removeFailedDownloads?"bg-warning":""}>
          <div class="switch"><Input disabled={updating} type="switch" bind:checked={form[idx].removeFailedDownloads} /></div>
        </td>
        {/if}
        <Tooltip target="dc{instance.App}priority{idx}">{form[idx].priority}</Tooltip>
        <td id="dc{instance.App}priority{idx}" class={form[idx].priority!=info[idx].priority?"bg-warning":""}>
          <div class="range"> <Input disabled={updating} bsSize="sm" type="range" min={1} max={100} bind:value={form[idx].priority} /></div>
        </td>
      </tr>
      {/if}<!-- /if (client) -->
    {/each}<!-- /each info as client, idx -->
  </Table>

  {#if instance.App == "Prowlarr"}
    {$_("instances.ProwlarrNotSupported")}
  {:else}
    <Footer {instance} bind:selected={selected} identifier="DownloadClients"
    bind:updating={updating}  bind:info={info} bind:form={form} bind:str={str}/>
  {/if}<!-- /if (instance.App) -->
</div><!-- id="container" -->

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
</style>