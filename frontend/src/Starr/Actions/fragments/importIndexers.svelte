<script lang="ts">
  // The instances being imported into.
  export let instance: Instance
  // The import data read from a file. Loop this.
  export let info: any[]
  // This goes true when a button is clicked.
  export let updating: boolean
  export let selected: {[key: string]: boolean} // Rows selected by key: ID.

  import { _ } from "/src/libs/Translate.svelte"
  import ModalInput from "./modalInput.svelte"
  import ConfigModal from "./configModal.svelte"
  import FieldInput from "./fieldInput.svelte"
  import SelectAll from "./selectAllHeader.svelte"
  import SelectRow from "./selectAllRow.svelte"
  import { Input, Table, Tooltip } from "@sveltestrap/sveltestrap"
  import type { Instance } from "/src/libs/config"

  let all = false
  let isOpen: {[key: number]: boolean} = {} // Modal toggle control.
</script>

<!-- {Object.keys(info[0])} -->

<Table bordered>
  <tr>
    <SelectAll bind:all bind:selected updating={false} icon="check2-all" />
    <th class="d-none d-sm-table-cell">
      <Tooltip target="indexer{instance.App}Type">{$_("words.Protocol")}</Tooltip>
      <span id="indexer{instance.App}Type">{$_("words.Type")}</span>
    </th>
    <th>{$_("words.Name")}</th>
    <th>
      <span id="enableRss">{$_(`instances.enableRssTitle`).match(/\b(\w)/g)?.join('').substring(0,2)}</span>
      <Tooltip target="enableRss">{$_("instances.enableRssDesc")}</Tooltip>
    </th>
    <th>
      <span id="enableInteractiveSearch">{$_(`instances.enableInteractiveSearchTitle`).match(/\b(\w)/g)?.join('').substring(0,2)}</span>
      <Tooltip target="enableInteractiveSearch">{$_("instances.enableInteractiveSearchDesc")}</Tooltip>        
    </th>
    <th>
      <span id="enableAutomaticSearch">{$_(`instances.enableAutomaticSearchTitle`).match(/\b(\w)/g)?.join('').substring(0,2)}</span>
      <Tooltip target="enableAutomaticSearch">{$_("instances.enableAutomaticSearchDesc")}</Tooltip>
    </th>
  </tr>

  {#each info as indexer, idx}
    <SelectRow {updating} bind:selected id={idx} item={indexer}>
      <td>
        <a href="/" style="padding-left:0" on:click|preventDefault={() => isOpen[idx]=!updating}>{indexer.name}</a>
        <ConfigModal {info} disabled="Details for importable indexer." {idx} id={idx} name={indexer.implementation} bind:isOpen={isOpen[idx]}>
          <ModalInput {info} form={info} disabled {idx} field="name" name="words.Name" type="text"/>
          <ModalInput {info} form={info} disabled {idx} field="priority" name="words.Priority" type="number"/>
          {#each info[idx].fields as item, itemIdx}
            <FieldInput {item} {itemIdx} {info} {idx} form={info} disabled />
          {/each}
        </ConfigModal>
      </td>
      <td><div class="switch"><Input type="checkbox" disabled bind:checked={indexer.enableRss}/></div></td>
      <td><div class="switch"><Input type="checkbox" disabled bind:checked={indexer.enableInteractiveSearch}/></div></td>
      <td><div class="switch"><Input type="checkbox" disabled bind:checked={indexer.enableAutomaticSearch}/></div></td>
    </SelectRow>
  {/each}<!-- /each info as indexer, idx -->
</Table>

<style>
  .switch {
    height:20px;
    margin-top: -5px !important;
  }
</style>
