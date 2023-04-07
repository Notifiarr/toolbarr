<script>
  export let info
  export let instance

  import { Badge, Button, Dropdown, DropdownItem, DropdownMenu, DropdownToggle, Input, InputGroup, InputGroupText, Modal, Spinner, Table, Tooltip } from "sveltestrap"
  import T, { _ } from "../../../libs/Translate.svelte"
  import { createEventDispatcher } from "svelte"
  import Fa from "svelte-fa";
  import { faCircleInfo } from "@fortawesome/free-solid-svg-icons";

  let isOpen = {} // Modal control
  const dispatch = createEventDispatcher()

  let newInfo = JSON.parse(JSON.stringify(info))
  $: if (!newInfo && info) info = newInfo = JSON.parse(JSON.stringify(info))
  $: unSaved = JSON.stringify(newInfo) !== JSON.stringify(info)

  function save() {
    newInfo = info = undefined
    dispatch('update')
  }

  function reset(idx) {
    newInfo[idx] = JSON.parse(JSON.stringify(info[idx]))
    isOpen[idx] = false
  }

  let selected = {}
  $: selectedCount = count(selected)

  function count(selected) {
    let counter = 0
    for (var k in selected) if (selected[k]) counter++;
    return counter
  }

  function toggleAll(key, on) {
    newInfo.forEach((_, i) => {newInfo[i][key] = on})
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
    <th><span class="link" on:keypress={selectAll} on:click={selectAll}>All</span></th>
    <th>ID <span><Badge color="info" class="superbadge">Type</Badge></span></th>
    <th>Name</th>
    <th>
      <Dropdown size="sm">
        <DropdownToggle tag="span" class="link">RS <Fa primaryColor="darkCyan" icon={faCircleInfo}/></DropdownToggle>
        <DropdownMenu>
          <DropdownItem header>Enable RSS</DropdownItem>
          <DropdownItem on:click={() => toggleAll("enableRss", false)}>Disable All</DropdownItem>
          <DropdownItem on:click={() => toggleAll("enableRss", true)}>Enable All</DropdownItem>
        </DropdownMenu>
      </Dropdown>
    </th>
    <th>
      <Dropdown size="sm">
        <DropdownToggle tag="span" class="link">IS <Fa primaryColor="darkCyan" icon={faCircleInfo}/></DropdownToggle>
        <DropdownMenu>
          <DropdownItem header>Enable Interactive Search</DropdownItem>
          <DropdownItem on:click={() => toggleAll("enableInteractiveSearch", false)}>Disable All</DropdownItem>
          <DropdownItem on:click={() => toggleAll("enableInteractiveSearch", true)}>Enable All</DropdownItem>
        </DropdownMenu>
      </Dropdown>
    </th>
    <th>
      <Dropdown size="sm">
        <DropdownToggle tag="span" class="link">AS <Fa primaryColor="darkCyan" icon={faCircleInfo}/></DropdownToggle>
        <DropdownMenu>
          <DropdownItem header>Enable Automatic Search</DropdownItem>
          <DropdownItem on:click={() => toggleAll("enableAutomaticSearch", false)}>Disable All</DropdownItem>
          <DropdownItem on:click={() => toggleAll("enableAutomaticSearch", true)}>Enable All</DropdownItem>
        </DropdownMenu>
      </Dropdown>
    </th>
  </tr>
  {#each info as indexer, idx}
  <tr class={selected[idx]?" bg-secondary":""}>
    <td><div class="switch"><Input type="switch" bind:checked={selected[idx]}/></div></td>

    <td>{indexer.id} <span><Badge color="info" class="superbadge">
      <open-browser href={info[idx].infoLink}>{indexer.implementation}</open-browser>
    </Badge></span></td>
    <td>
      <a href="/" on:click|preventDefault={() => isOpen[idx]=true}>{indexer.name}</a>
      <Modal class="modal-settings" body size="lg" scrollable isOpen={isOpen[idx]}>
        <h4><Badge color="info">{indexer.id}</Badge> {indexer.implementation}</h4>
        <InputGroup>
          <InputGroupText class="setting-name">Name</InputGroupText>
          <Input type="text" bind:value={newInfo[idx].name} />
        </InputGroup>
        <InputGroup>
          <InputGroupText class="setting-name">Priority</InputGroupText>
          <Input type="number" bind:value={newInfo[idx].priority} />
        </InputGroup>
        {#each info[idx].fields as item, itemIdx}
          {#if item.helpText}
          <Tooltip target="{item.name.replace(".", "")}{idx}{itemIdx}">{item.helpText}</Tooltip>
          {/if}
          <InputGroup id="{item.name.replace(".", "")}{idx}{itemIdx}">
            <InputGroupText class="setting-name">{item.label}</InputGroupText>
            {#if item.type == "select"}
            <select disabled style="width:calc(99% - 161px)" multiple bind:value={newInfo[idx].fields[itemIdx].value}>
              {#each info[idx].fields[itemIdx].value as val}
              <option value={val}>{val}</option>
              {/each}
            </select>
            {:else}
            <Input type={item.type} bind:value={newInfo[idx].fields[itemIdx].value}/>
            {/if}
          </InputGroup>
        {/each}
        <Button color="success" on:click={() => isOpen[idx]=false}>Save</Button>
        <Button color="danger" on:click={() => reset(idx)}>Cancel</Button>
      </Modal>
    </td>
    <td class={newInfo[idx].enableRss!=info[idx].enableRss?"bg-warning":""}>
      <div class="switch"><Input type="switch" bind:checked={newInfo[idx].enableRss} /></div>
    </td>
    <td class={newInfo[idx].enableInteractiveSearch!=info[idx].enableInteractiveSearch?"bg-warning":""}>
      <div class="switch"><Input type="switch" bind:checked={newInfo[idx].enableInteractiveSearch} /></div>
    </td>
    <td class={newInfo[idx].enableAutomaticSearch!=info[idx].enableAutomaticSearch?"bg-warning":""}>
      <div class="switch"><Input type="switch" bind:checked={newInfo[idx].enableAutomaticSearch} /></div>
    </td>
  </tr>
  {/each}
</Table>
<Button color="success">Save All</Button>
{#if selectedCount > 0}
  <Button color="info">Save {selectedCount} Selected</Button>
  <Button color="danger">Delete {selectedCount} Selected</Button>
{/if}
{#if unSaved}&nbsp; <span class="text-danger">Unsaved Changes!</span>{/if}
{/if}
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

  #container :global(.link) {
    cursor:pointer;
    text-decoration:underline;
  }
</style>