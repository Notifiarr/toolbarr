<script lang="ts">
  export let tab: Tab
  export let info
  export let instance

  import type { Tab } from "./fragments/tabs.svelte"
  import { _ } from "../../libs/Translate.svelte"
  import Footer from "./footer.svelte"
  import Dropdown from "./fragments/dropdown.svelte"
  import TDInput from "./fragments/tableInput.svelte"
  import ModalInput from "./fragments/modalInput.svelte"
  import ConfigModal from "./fragments/configModal.svelte"
  import FieldInput from "./fragments/fieldInput.svelte"
  import SelectAll from "./fragments/selectAllHeader.svelte"
  import SelectRow from "./fragments/selectAllRow.svelte"
  import { fixFieldValues } from "./methods"
  import { Table, Tooltip, Icon } from "sveltestrap"

  let isOpen: any = {}           // Modal toggle control.
  let updating: boolean = false  // True while doing updates.
  let all: boolean = false       // Toggle for select-all link.
  let selected: any = {}         // Rows selected by key: ID.
  let str: string = fixFieldValues(info) // Used for equivalence comparison.
  let form: any = JSON.parse(str)        // Form changes go here.
  let starrApp = instance.App
</script>

<Table bordered>
  <tr>
    <SelectAll bind:all={all} bind:selected={selected} bind:updating={updating}/>
    <th class="d-none d-sm-table-cell">
      <Tooltip target="dc{starrApp}Type">{$_("words.Protocol")}</Tooltip>
      <span id="dc{instance.App}Type">{$_("words.Type")}
        <Icon class="small text-info" name="box-arrow-up-right"/>
      </span>
    </th>
    <th>{$_("words.Name")}</th>
    {#if starrApp == "Readarr"}
      <th><Dropdown bind:form={form} {updating} {starrApp} field="enable" name="configvalues.Enabled"/></th>
    {:else}
      <th><Dropdown bind:form={form} {updating} {starrApp} field="removeCompletedDownloads"/></th>
      <th><Dropdown bind:form={form} {updating} {starrApp} field="removeFailedDownloads"/></th>
    {/if}
    <th><span>{$_("words.Priority")}</span></th>
  </tr>

  {#each info as client, idx}
    {#if client} <!-- When deleting a client, this protects an error condition. -->
    <ConfigModal {info} {form} {idx} {str} id={client.id} name={client.implementation} bind:isOpen={isOpen[idx]}
      disabled={starrApp=="Prowlarr"?$_("instances.ProwlarrNotSupported"):""}>
      <ModalInput {info} bind:form={form} {idx} field="name" name="words.Name" type="text"/>
      <ModalInput {info} bind:form={form} {idx} field="priority" name="words.Priority" type="number"/>
      {#each info[idx].fields as item, itemIdx}
        <FieldInput {item} {itemIdx} {info} {idx} bind:form={form}/>
      {/each}
    </ConfigModal>

    <SelectRow {updating} {selected} id={info[idx].id} item={client}>
      <td class={JSON.stringify(form[idx]) != JSON.stringify(info[idx])?"border-warning":""}>
        <a href="/" style="padding-left:0" on:click|preventDefault={() => isOpen[idx]=!updating}>{client.name}</a>
      </td>
      {#if instance.App == "Readarr"}
      <TDInput {idx} {info} {updating} bind:form={form} field="enable" type="switch"/>
      {:else}
      <TDInput {idx} {info} {updating} bind:form={form} field="removeCompletedDownloads" type="switch"/>
      <TDInput {idx} {info} {updating} bind:form={form} field="removeFailedDownloads" type="switch"/>
      {/if}
      <TDInput {idx} {info} {updating} bind:form={form} field="priority" type="range"/>
    </SelectRow>
    {/if}<!-- /if (client) -->
  {/each}<!-- /each info as client, idx -->
</Table>

{#if instance.App == "Prowlarr"}
  {$_("instances.ProwlarrNotSupported")}
{:else}
  <Footer {instance} bind:selected {tab} bind:updating bind:info bind:form bind:str/>
{/if}
