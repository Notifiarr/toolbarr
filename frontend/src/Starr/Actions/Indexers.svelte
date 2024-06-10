<script lang="ts">
  export let tab: Tab|undefined
  export let info: any
  export let instance: Instance
  export let updating: boolean
  export let selected: {[key: string]: boolean} = {} // Rows selected by key: ID.
  export let str = ""                                // Used for equivalence comparison.
  export let form: any = undefined                   // Form changes go here.

  import type { Tab } from "./fragments/tabs.svelte"
  import { _ } from "/src/libs/Translate.svelte"
  import Footer from "./fragments/footer.svelte"
  import ModalInput from "./fragments/modalInput.svelte"
  import Dropdown from "./fragments/dropdown.svelte"
  import ConfigModal from "./fragments/configModal.svelte"
  import TDInput from "./fragments/tableInput.svelte"
  import FieldInput from "./fragments/fieldInput.svelte"
  import SelectAll from "./fragments/selectAllHeader.svelte"
  import SelectRow from "./fragments/selectAllRow.svelte"
  import { fixFieldValues } from "./methods"
  import { Table, Tooltip, Icon } from "@sveltestrap/sveltestrap"
  import type { Instance } from "/src/libs/config"

  let isOpen: {[key: number]: boolean} = {} // Modal toggle control.
  let all = false                           // Toggle for select-all link.
  str = fixFieldValues(info)
  form = JSON.parse(str)
  let starrApp = instance.App
</script>

<Table bordered>
  <tr>
    <SelectAll bind:all bind:selected bind:updating icon="check2-all" />
    <th class="d-none d-sm-table-cell">
      <Tooltip target="indexer{starrApp}Type">{$_("words.Protocol")}</Tooltip>
      <span id="indexer{starrApp}Type">
        {$_("words.Type")} <Icon class="small text-info" name="box-arrow-up-right"/>
      </span>
    </th>
    <th>{$_("words.Name")}</th>
    <th><Dropdown bind:form {updating} {starrApp} field="enableRss"/></th>
    <th><Dropdown bind:form {updating} {starrApp} field="enableInteractiveSearch"/></th>
    <th><Dropdown bind:form {updating} {starrApp} field="enableAutomaticSearch"/></th>
  </tr>

  {#each info as indexer, idx}
    {#if indexer} <!-- When deleting an indexer, this protects an error condition. -->
    {@const loopid = tab?indexer.id:idx}
    <SelectRow {updating} bind:selected id={loopid} item={indexer}>
      <td class={JSON.stringify(form[idx]) != JSON.stringify(info[idx])?"border-warning":""}>
        <a href="/" style="padding-left:0" on:click|preventDefault={() => isOpen[idx]=!updating}>{indexer.name}</a>
        <ConfigModal {info} {form} {idx} {str} id={loopid} name={indexer.implementation} bind:isOpen={isOpen[idx]}
          disabled={starrApp=="Prowlarr"?$_("instances.ProwlarrNotSupported"):""}>
          <ModalInput {info} bind:form {idx} field="name" name="words.Name" type="text"/>
          <ModalInput {info} bind:form {idx} field="priority" name="words.Priority" type="number"/>
          {#each info[idx].fields as item, itemIdx}
            <FieldInput {item} {itemIdx} {info} {idx} bind:form />
          {/each}
        </ConfigModal>
      </td>
      <TDInput {idx} {info} {updating} bind:form field="enableRss" type="switch"/>
      <TDInput {idx} {info} {updating} bind:form field="enableInteractiveSearch" type="switch"/>
      <TDInput {idx} {info} {updating} bind:form field="enableAutomaticSearch" type="switch"/>
    </SelectRow>
    {/if}<!-- /if (indexer) -->
  {/each}<!-- /each info as indexer, idx -->
</Table>

{#if instance.App == "Prowlarr"}
  {$_("instances.ProwlarrNotSupported")}
{:else if tab}
  <Footer {instance} {tab} bind:selected bind:updating bind:info bind:form bind:str importable />
{/if}<!-- /if (instance.App) -->
