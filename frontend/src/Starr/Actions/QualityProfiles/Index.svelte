<script>
  export let info
  export let instance

  import { _ } from "../../../libs/Translate.svelte"
  import Fa from "svelte-fa"
  import Footer from "../footer.svelte"
  import { faCircleInfo, faTrashAlt } from "@fortawesome/free-solid-svg-icons"
  import { count } from "../methods.js"
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
  let str = JSON.stringify(info) // Used for equivalence comparison.
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
      <th>{$_("words.Name")}</th>
      <th>
        <Dropdown size="sm">
          <Tooltip placement="top" target="qp{instance.App}upgradeAllowed">{$_("instances.EnableQPDesc")}</Tooltip>
          <DropdownToggle id="qp{instance.App}upgradeAllowed" tag="span" class="link">UA <Fa primaryColor="darkCyan" icon={faCircleInfo}/></DropdownToggle>
          <DropdownMenu>
            <DropdownItem header>{$_("instances.UpgradeAllowed")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("upgradeAllowed", false)}>{$_("instances.DisableAll")}</DropdownItem>
            <DropdownItem disabled={updating} on:click={() => toggleAll("upgradeAllowed", true)}>{$_("instances.EnableAll")}</DropdownItem>
          </DropdownMenu>
        </Dropdown>
      </th>
      <th>{$_("instances.UpgradeUntil")}</th>
    </tr>

    {#each info as profile, idx}
      {#if profile} <!-- When deleting a profile, this protects an error condition. -->
      <tr class={selected[info[idx].id]?" bg-secondary":""}>
        <td><div class="switch"><Input disabled={updating} invalid={selected[info[idx].id]} type="switch" bind:checked={selected[info[idx].id]}/></div></td>

        <td class="d-none d-md-table-cell">{profile.id}</td>
        <td class={JSON.stringify(form[idx]) != JSON.stringify(info[idx])?"border-warning":""}>
          <span>
            <Badge color="primary" class="superbadge">{count(profile.items, "allowed")}</Badge>
            <a href="/" style="padding-left:0" on:click|preventDefault={() => isOpen[idx]=!updating}>{profile.name}</a>
          </span>
          <Modal class="modal-settings" body size="lg" scrollable isOpen={isOpen[idx]}>
            <ModalHeader toggle={() => {isOpen[idx] = false}}>
              <Badge color="info">{profile.id}</Badge> {profile.name}
            </ModalHeader>

            <ModalBody>
              <InputGroup>
                <InputGroupText class="setting-name">{$_("words.Name")}</InputGroupText>
                <Input invalid={form[idx].name != info[idx].name} type="text" bind:value={form[idx].name} />
              </InputGroup>
              <h5>{$_("words.Profiles")} <div style="float:right"><Badge>{$_("words.Qualities")}</Badge></div></h5>
              {#each info[idx].items as item, itemIdx}
                <InputGroup>
                  <InputGroupText class={item.allowed!=form[idx].items[itemIdx].allowed?"bg-danger":""} style="padding-left:6px;width:45px;">
                    <Input type="switch" bind:checked={form[idx].items[itemIdx].allowed}/>
                  </InputGroupText>
                  {#if item.name}
                    <Tooltip target="qp{instance.App}name{item.id}">ID: {item.id}</Tooltip>
                    <Input id="qp{instance.App}name{item.id}" invalid={item.name != form[idx].items[itemIdx].name} 
                      bind:value={form[idx].items[itemIdx].name}/>
                  {:else}
                    <Tooltip target="qp{instance.App}name{item.quality.id}">ID: {item.quality.id} ({$_("words.Readonly")})</Tooltip>
                    <Input readonly id="qp{instance.App}name{item.quality.id}" value={item.quality.name}/>
                  {/if}
                  <InputGroupText>{item.items?item.items.length:1}</InputGroupText>
                </InputGroup>
              {/each}<!-- /each info[idx].items as item, itemIdx -->
            </ModalBody>

            <ModalFooter>
              <Button size="sm" disabled={str==JSON.stringify(form)} color="danger" 
                on:click={() => {form[idx] = JSON.parse(JSON.stringify(info[idx]))}}>{$_("words.Reset")}</Button>
              <Button size="sm" color="info" on:click={() => {isOpen[idx] = false}}>{$_("words.Close")}</Button>
            </ModalFooter>
          </Modal>

        </td>
        <td class={form[idx].upgradeAllowed!=info[idx].upgradeAllowed?"bg-warning":""}>
          <div class="switch"><Input disabled={updating} type="switch" bind:checked={form[idx].upgradeAllowed} /></div>
        </td>
        <td class={form[idx].cutoff!=info[idx].cutoff?"bg-warning":""}>
          <div class="select">
            <InputGroup size="sm">
              <Input disabled={updating} invalid={form[idx].cutoff!=info[idx].cutoff} type="select" bind:value={form[idx].cutoff}>
                {#each info[idx].items as item}
                  {#if item.allowed}
                    <option disabled={!form[idx].upgradeAllowed} value={item.id!==undefined?item.id:item.quality.id}>
                      {item.name!==undefined?item.name:item.quality.name}
                    </option>
                  {/if}
                {/each}
              </Input>
            </InputGroup>
          </div>
        </td>
      </tr>
      {/if}<!-- /if (profile) -->
    {/each}<!-- /each info as profile, idx -->
  </Table>

  <Footer {instance} noForce identifier="QualityProfiles" bind:selected={selected}
    bind:updating={updating} bind:info={info} bind:form={form} bind:str={str}/>
</div>

<style>
  .switch {
    height:20px;
    margin-top: -10px !important;
  }
  .select {
    padding:0px;
    margin-bottom:-2px;
    margin-top:-3px;
    margin-right:3px;
  }
  :global(.modal-settings .setting-name) {
    min-width:max-content !important;
    width:160px !important;
  }
</style>