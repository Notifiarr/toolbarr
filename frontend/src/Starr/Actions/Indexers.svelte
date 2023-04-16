<script>
  export let info
  export let instance

  import T, { _ } from "../../libs/Translate.svelte"
  import Fa from "svelte-fa"
  import Footer from "./footer.svelte"
  import { fixFieldValues } from "./methods.js"
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

  function toggleAll(key, on) {
    if (!updating) form.forEach((_, i) => {form[i][key] = on})
  }

  function selectAll() {
    all = !all
    if (!updating) Object.keys(selected).forEach(k => selected[k] = all)
  }

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
      <th>
        <span class="link">
          <Fa size="sm" icon={faTrashAlt}/> <span on:keyup={selectAll} on:click={selectAll}>{$_("words.All")}</span>
        </span>
      </th>
      <th class="d-none d-md-table-cell">ID</th>
      <th class="d-none d-sm-table-cell">
        <Tooltip target="indexer{instance.App}Type">{$_("words.Protocol")}</Tooltip>
        <span id="indexer{instance.App}Type">{$_("words.Type")} <Fa size="xs" primaryColor="darkCyan" icon={faArrowUpRightFromSquare}/></span>
      </th>
      <th>{$_("words.Name")}</th>
      <th>
        <Dropdown size="sm">
          <Tooltip placement="top" target="indexer{instance.App}enableRss"><T id="instances.EnableRSSDesc" starrApp={instance.App}/></Tooltip>
          <DropdownToggle id="indexer{instance.App}enableRss" tag="span" class="link">RS <Fa primaryColor="darkCyan" icon={faCircleInfo}/></DropdownToggle>
          <DropdownMenu>
            <DropdownItem header>{$_("instances.EnableRSS")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("enableRss", false)}>{$_("instances.DisableAll")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("enableRss", true)}>{$_("instances.EnableAll")}</DropdownItem>
          </DropdownMenu>
        </Dropdown>
      </th>
      <th>
        <Dropdown size="sm">
          <Tooltip placement="top" target="indexer{instance.App}enableInteractiveSearch">{$_("instances.InteractiveSearchDesc")}</Tooltip>
          <DropdownToggle id="indexer{instance.App}enableInteractiveSearch" tag="span" class="link">IS <Fa primaryColor="darkCyan" icon={faCircleInfo}/></DropdownToggle>
          <DropdownMenu>
            <DropdownItem header>{$_("instances.EnableInteractiveSearch")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("enableInteractiveSearch", false)}>{$_("instances.DisableAll")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("enableInteractiveSearch", true)}>{$_("instances.EnableAll")}</DropdownItem>
          </DropdownMenu>
        </Dropdown>
      </th>
      <th>
        <Dropdown size="sm">
          <Tooltip placement="top" target="indexer{instance.App}enableAutomaticSearch"><T id="instances.AutomaticSearchDesc" starrApp={instance.App}/></Tooltip>
          <DropdownToggle id="indexer{instance.App}enableAutomaticSearch" tag="span" class="link">AS <Fa primaryColor="darkCyan" icon={faCircleInfo}/></DropdownToggle>
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
        <td><div class="switch"><Input disabled={updating} invalid={selected[info[idx].id]} type="switch" bind:checked={selected[info[idx].id]}/></div></td>

        <td class="d-none d-md-table-cell">{indexer.id}</td>
        <td class="d-none d-sm-table-cell">
          <Tooltip target="indexer{instance.App}Type{idx}">{indexer.protocol} {info[idx].infoLink}</Tooltip>
          <span id="indexer{instance.App}Type{idx}">
            <Badge color="info"><open-browser href={info[idx].infoLink}>{indexer.implementation}</open-browser></Badge>
          </span>
        </td>
        <td class={JSON.stringify(form[idx]) != JSON.stringify(info[idx])?"border-warning":""}>
          <a href="/" style="padding-left:0" on:click|preventDefault={() => isOpen[idx]=!updating}>{indexer.name}</a>
          <Modal class="modal-settings" body size="lg" scrollable isOpen={isOpen[idx]}>
            <ModalHeader toggle={() => {isOpen[idx] = false}}>
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
                  <Tooltip target="indexer{instance.App}{item.name.replace(".", "")}{idx}{itemIdx}">{item.helpText}</Tooltip>
                {/if}
                <InputGroup id="indexer{instance.App}{item.name.replace(".", "")}{idx}{itemIdx}">
                  <InputGroupText class="setting-name">{item.label?item.label:item.name}</InputGroupText>
                  {#if item.type == "select" || item.type == "tagSelect"}
                    {#if item.selectOptions && typeof item.value != "object"}
                      <Input invalid={info[idx].fields[itemIdx].value != form[idx].fields[itemIdx].value}
                        type="select" bind:value={form[idx].fields[itemIdx].value}>
                          {#each info[idx].fields[itemIdx].selectOptions as val}
                            <option value={val.name}>{val.name}</option>
                          {/each}
                      </Input>
                    {:else}<!-- /else (item.selectOptions) -->
                      <select multiple disabled style="width:calc(100% - 160px)" bind:value={form[idx].fields[itemIdx].value}>
                      {#if typeof item.value == "object"}
                        {#each info[idx].fields[itemIdx].value as val}
                          <option value={val}>{val}</option>
                        {/each}
                      {:else}
                        <option value={info[idx].fields[itemIdx].value}>{info[idx].fields[itemIdx].value}</option>
                      {/if}
                      </select>
                    {/if}<!-- /if (item.selectOptions) -->
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

  {#if instance.App == "Prowlarr"}
    {$_("instances.ProwlarrNotSupported")}
  {:else}
    <Footer {instance} identifier="Indexers" bind:selected={selected}
      bind:updating={updating} bind:info={info} bind:form={form} bind:str={str}/>
  {/if}<!-- /if (instance.App) -->
</div>

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
