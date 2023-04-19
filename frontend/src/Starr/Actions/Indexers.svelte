<script>
  export let tab
  export let info
  export let instance

  import { _ } from "../../libs/Translate.svelte"
  import Footer from "./footer.svelte"
  import ModalInput from "./fragments/modalInput.svelte"
  import Dropdown from "./fragments/dropdown.svelte"
  import ConfigModal from "./fragments/configModal.svelte"
  import TDInput from "./fragments/tableInput.svelte"
  import FieldInput from "./fragments/fieldInput.svelte"
  import SelectAll from "./fragments/selectAllHeader.svelte"
  import SelectRow from "./fragments/selectAllRow.svelte"
  import { fixFieldValues } from "./methods.js"
  import { Table, Tooltip, Icon } from "sveltestrap"

  let isOpen = {}       // Modal toggle control.
  let updating = false  // True while doing updates.
  let all = false       // Toggle for select-all link.
  let selected = {}     // Rows selected by key: ID.
  let str = fixFieldValues(info) // Used for equivalence comparison.
  let form = JSON.parse(str)     // Form changes go here.
  let starrApp = instance.App
</script>

<Table bordered>
  <tr>
    <SelectAll bind:all={all} bind:selected={selected} bind:updating={updating}/>
    <th class="d-none d-sm-table-cell">
      <Tooltip target="indexer{starrApp}Type">{$_("words.Protocol")}</Tooltip>
      <span id="indexer{starrApp}Type">
        {$_("words.Type")} <Icon class="small text-info" name="box-arrow-up-right"/>
      </span>
    </th>
    <th>{$_("words.Name")}</th>
    <th><Dropdown bind:form={form} {updating} {starrApp} field="enableRss"/></th>
    <th><Dropdown bind:form={form} {updating} {starrApp} field="enableInteractiveSearch"/></th>
    <th><Dropdown bind:form={form} {updating} {starrApp} field="enableAutomaticSearch"/></th>
  </tr>

  {#each info as indexer, idx}
    {#if indexer} <!-- When deleting an indexer, this protects an error condition. -->
    <SelectRow {updating} {selected} id={info[idx].id} item={indexer}>
      <td class={JSON.stringify(form[idx]) != JSON.stringify(info[idx])?"border-warning":""}>
        <a href="/" style="padding-left:0" on:click|preventDefault={() => isOpen[idx]=!updating}>{indexer.name}</a>
        <ConfigModal {info} {form} {idx} {str} id={indexer.id} name={indexer.implementation} bind:isOpen={isOpen[idx]}
          disabled={starrApp=="Prowlarr"?$_("instances.ProwlarrNotSupported"):""}>
          <ModalInput {info} bind:form={form} {idx} field="name" name="words.Name" type="text"/>
          <ModalInput {info} bind:form={form} {idx} field="priority" name="words.Priority" type="number"/>
          {#each info[idx].fields as item, itemIdx}
            <FieldInput {item} {itemIdx} {info} {idx} bind:form={form}/>
          {/each}
        </ConfigModal>
      </td>
      <TDInput {idx} {info} {updating} bind:form={form} field="enableRss" type="switch"/>
      <TDInput {idx} {info} {updating} bind:form={form} field="enableInteractiveSearch" type="switch"/>
      <TDInput {idx} {info} {updating} bind:form={form} field="enableAutomaticSearch" type="switch"/>
    </SelectRow>
    {/if}<!-- /if (indexer) -->
  {/each}<!-- /each info as indexer, idx -->
</Table>

{#if instance.App == "Prowlarr"}
  {$_("instances.ProwlarrNotSupported")}
{:else}
  <Footer {instance} {tab} bind:selected={selected} bind:updating={updating} bind:info={info} bind:form={form} bind:str={str}/>
{/if}<!-- /if (instance.App) -->
