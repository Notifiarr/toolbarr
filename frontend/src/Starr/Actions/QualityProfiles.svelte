<script lang="ts">
  export let tab: Tab
  export let info
  export let instance
  export let updating: boolean

  import type { Tab } from "./fragments/tabs.svelte"
  import { _ } from "/src/libs/Translate.svelte"
  import Footer from "./fragments/footer.svelte"
  import { count } from "/src/libs/funcs"
  import Dropdown from "./fragments/dropdown.svelte"
  import TDInput from "./fragments/tableInput.svelte"
  import ModalInput from "./fragments/modalInput.svelte"
  import ConfigModal from "./fragments/configModal.svelte"
  import { Badge, Input, InputGroup, InputGroupText, Table, Tooltip } from "sveltestrap"
  import SelectAll from "./fragments/selectAllHeader.svelte"
  import SelectRow from "./fragments/selectAllRow.svelte"

  let isOpen: any = {}           // Modal toggle control.
  let all: boolean = false       // Toggle for select-all link.
  let selected: any = {}         // Rows selected by key: ID.
  let str: string = JSON.stringify(info) // Used for equivalence comparison.
  let form: any = JSON.parse(str)        // Form changes go here.
</script>

<Table bordered>
  <tr>
    <SelectAll bind:all bind:selected bind:updating/>
    <th>{$_("words.Name")}</th>
    <th><Dropdown bind:form {updating} starrApp={instance.App} field="upgradeAllowed"/></th>
    <th>{$_("instances.UpgradeUntil")}</th>
  </tr>

  {#each info as profile, idx}
    {#if profile} <!-- When deleting a profile, this protects an error condition. -->
    <SelectRow {updating} bind:selected id={info[idx].id}>
      <td class={JSON.stringify(form[idx]) != JSON.stringify(profile)?"border-warning":""}>
        <span>
          <Badge color="primary" class="superbadge">{count(profile.items, "allowed")}</Badge>
          <a href="/" style="padding-left:0" on:click|preventDefault={() => isOpen[idx]=!updating}>{profile.name}</a>
        </span>
        <ConfigModal {info} {form} {idx} {str} id={profile.id} name={profile.name} bind:isOpen={isOpen[idx]}>
          <ModalInput {info} bind:form {idx} field="name" name="words.Name" type="text"/>
          <h5>{$_("words.Profiles")} <div style="float:right"><Badge>{$_("words.Qualities")}</Badge></div></h5>

          {#each profile.items as item, itemIdx}
            <InputGroup>
              <InputGroupText class={item.allowed!=form[idx].items[itemIdx].allowed?"bg-warning":""}>
                <span><Input type="switch" bind:checked={form[idx].items[itemIdx].allowed}/></span>
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
          {/each}<!-- /each profile.items as item, itemIdx -->
        </ConfigModal>
      </td>

      <TDInput {idx} {info} {updating} bind:form field="upgradeAllowed" type="switch"/>
      <TDInput {idx} {info} {updating} bind:form field="cutoff" type="select">
        {#each profile.items as item}
          {#if item.allowed}
            <option disabled={!form[idx].upgradeAllowed} value={item.id!==undefined?item.id:item.quality.id}>
              {item.name!==undefined?item.name:item.quality.name}
            </option>
          {/if}
        {/each}
      </TDInput>
    </SelectRow>
    {/if}<!-- /if (profile) -->
  {/each}<!-- /each info as profile, idx -->
</Table>

<Footer {instance} noForce {tab} bind:selected bind:updating bind:info bind:form bind:str/>
