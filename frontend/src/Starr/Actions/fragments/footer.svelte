<script lang="ts">
  export let instance: Instance
  export let info: any
  export let form: any
  export let str: string
  export let updating: boolean
  export let selected: {[key: string]: boolean}
  export let tab: Tab
  export let noForce = false
  export let importable = false

  import type { Tab } from "./tabs.svelte"
  import { Alert, Button, Collapse, Fade, Tooltip, Icon, Card } from "@sveltestrap/sveltestrap"
  import Loading from "/src/Starr/loading.svelte"
  import T, { _ } from "/src/libs/Translate.svelte"
  import { toast, count } from "/src/libs/funcs"
  import { update, remove, test, exportFile, importFile, add, fixFieldValues } from "../methods"
  import type { Instance } from "/src/libs/config"
  import { createEventDispatcher } from "svelte"
  import ConfigModal from "./configModal.svelte"

  const dispatch = createEventDispatcher()
  let badMsg = ""
  let goodMsg = ""
  $: selectedCount = count(selected)        // How many items are selected.
  $: unSaved = JSON.stringify(form) !== str // True when something changed.
  let button: HTMLElement
  let importing: {[key: string]: boolean} = {}

  // Used by the importer.
  let importData: any = undefined
  let importForm: any = undefined

  async function showError(idx: number, err: string) {
    form[idx] = JSON.parse(JSON.stringify(info[idx]))
    badMsg += `<li>${$_("instances.ErrorMsg", {values:{"msg": err}})}</li>`
  }

  async function updateItems(force: boolean) {
    toast("info", $_("instances.Updating"+tab.id))
    goodMsg = badMsg = ""
    updating = true

    for (var idx = 0; idx < form.length; idx++) {
      if (JSON.stringify(form[idx]) == JSON.stringify(info[idx])) continue // not changed
      if (noForce) {
        await update[tab.id][instance.App](instance, form[idx]).then(
          async (resp: any) => {
            goodMsg += `<li>${$_("instances.SuccessMsg", {values:{"msg": resp.Msg}})}</li>`
            form[idx] = JSON.parse(JSON.stringify(resp.Data))
            dispatch("update")
            str = fixFieldValues(form)
          },
          async (err: string) => await showError(idx, err)
        )
      } else {
        await update[tab.id][instance.App](instance, force, form[idx]).then(
          async (resp: any) => {
            goodMsg += `<li>${$_("instances.SuccessMsg", {values:{"msg": resp.Msg}})}</li>`
            form[idx] = JSON.parse(JSON.stringify(resp.Data))
            dispatch("update")
            str = fixFieldValues(form)
          },
          async (err: string) => await showError(idx, err)
        )
      }
    }

    updating = false
    info = JSON.parse(str)
    Object.keys(selected).forEach(k => selected[k] = false)
  }

  async function deleteItem() {
    toast("info", $_("instances.Deleting"+tab.id, { values:{"count": count(selected)} }), "", 4)
    goodMsg = badMsg = ""
    updating = true

    for (var idx = info.length-1; idx >= 0; idx--) {
      if (!selected[info[idx].id]) continue // Not selected.
      selected[info[idx].id] = false
      await remove[tab.id][instance.App](instance, info[idx].id).then(
        (msg: string) => {
          goodMsg += `<li>${$_("instances.SuccessMsg", {values:{"msg": msg}})}</li>`
          selected[info[idx].id] = false
          form.splice(idx, 1)
          info.splice(idx, 1)
          dispatch("delete")
        },
        (err: string) => showError(idx, err)
      )
    }

    Object.keys(selected).forEach(k => selected[k] = false)
    str = fixFieldValues(form)
    updating = false
  }

  async function testItem() {
    toast("info", $_("instances.Testing"+tab.id, { values:{"count": count(selected)} }), "", 3)
    goodMsg = badMsg = ""
    updating = true

    for (var idx = info.length-1; idx >= 0; idx--) {
      if (!selected[info[idx].id]) continue // Not selected.
      await test[tab.id][instance.App](instance, info[idx]).then(
        (msg: string) => {goodMsg += `<li>${$_("instances.SuccessMsg", {values:{"msg": msg}})}</li>`},
        (err: string) => {badMsg += `<li>${$_("instances.ErrorMsg", {values:{"msg": err}})}</li>`},
      )
    }

    updating = false
  }

  async function exportItems() {
    toast("info", $_("instances.Exporting"+tab.id, { values:{"count": count(selected)} }), "", 3)
    goodMsg = badMsg = ""
    updating = true

    await exportFile[tab.id](instance, selected).then(
      (msg: string) => {if (msg!="") goodMsg += `<li>${$_("instances.SuccessMsg", {values:{"msg": msg}})}</li>`},
      (err: string) => {badMsg += `<li>${$_("instances.ErrorMsg", {values:{"msg": err}})}</li>`},
    )

    updating = false
  }

  async function importItems() {
    goodMsg = badMsg = ""
    updating = true
    importData = undefined
    importForm = undefined

    await importFile[tab.id](instance).then(
      (resp: any) => {
        if (resp && resp.Msg) {
          goodMsg += `<li>${$_("instances.SuccessMsg", {values:{"msg": resp.Msg}})}</li>`
          //  This opens a modal with list of stuff and checkboxes to select things to import.
          importData = resp.Data
          importForm = resp.Data
        }
      },
      (err: string) => {badMsg += `<li>${$_("instances.ErrorMsg", {values:{"msg": err}})}</li>`},
    )

    updating = false
  }

  // This gets called when the import items modal gets closed.
  async function addSelectedItems() {
    updating = true

    for (var idx = importForm.length-1; idx >= 0; idx--) {
      if (!importing[idx]) continue // not selected
      importing[idx] = false

      await add[tab.id][instance.App](instance, importForm[idx]).then(
        (resp: any) => {
          goodMsg += `<li>${$_("instances.SuccessMsg", {values:{"msg": resp.Msg}})}</li>`
          if (!hasID(info, resp.Data.id)) {
            info.push(resp.Data)
            form.push(resp.Data)
            str = fixFieldValues(form)
          }
        },
        (err: string) => {badMsg += `<li>${$_("instances.ErrorMsg", {values:{"msg": err}})}</li>`},
      )
    }

    importData = undefined
    updating = false
  }

  // Sometimes when adding new items, we get 'success' and an old item returned.
  // This function allows checking if the old item is already in the form list.
  function hasID(info: any, id: number): boolean {
    for (var idx = info.length-1; idx >= 0; idx--) {
      if (info[idx].id == id) return true
    }
    return false
  }
</script>

<!-- This modal is used to import data. Only some actions support importing. -->
<ConfigModal info={undefined} form={undefined} str="" idx=""
  name={$_("instances.ImportSelection")} closeButton={$_("words.Import")}
  id={$_("instances."+tab.id)} disabled={$_("instances.CloseImportModal")}
  isOpen={importData != undefined} callback={addSelectedItems}>
    <svelte:component this={tab.component} footer={false} {instance}
    bind:info={importData} bind:form={importForm} bind:updating bind:selected={importing} applyAll="" />
</ConfigModal>

<div id="footer">
  <Collapse isOpen={goodMsg != ""}>
    <Alert dismissible color="success" fade={false}>{@html goodMsg}</Alert>
  </Collapse>
  <Collapse isOpen={badMsg != ""}>
    <Alert dismissible color="danger" fade={false}>{@html badMsg}</Alert>
  </Collapse>
  <Loading isOpen={updating}/>

  <Collapse isOpen={!updating}>
    {#if instance.App == "Whisparr"}
    <Card color="warning" class="p-1">
      Toolbarr has not been tested with {instance.App}.
      Make a backup and save with caution!
    </Card>
    {/if}

    <Fade style="display:inline-block" isOpen={unSaved}>
      <span class="text-warning">
        <Icon class="text-danger" name="exclamation-circle"/>
        {$_("configvalues.UnsavedChanges")}
      </span><br>
      <!-- Save and Force Save Buttons -->
      <Button class="actions" color="success" on:click={() => updateItems(false)}>
        {$_(noForce?"words.Save":"instances.TestandSave")}
      </Button>
      <Tooltip target={button}><T id="instances.ForceSaveDesc" starrApp={instance.App}/></Tooltip>
      <span bind:this={button}>
        <Button disabled={noForce} class="actions" color="info" on:click={() => updateItems(true)}>
          {$_("instances.ForceSave")}
        </Button>
      </span>
    </Fade>

    {#if importable}
    <Button class="actions" color="warning" on:click={importItems}>
      <T id="instances.ImportFile" />
    </Button>
    {/if}

    <Fade style="display:inline-block" isOpen={selectedCount > 0}>
      {#if importable}
      <Button class="actions" color="secondary" on:click={exportItems}>
        <T id="instances.ExportSelected" count={selectedCount}/>
      </Button>
      {/if}

      <!-- Test and Delete Buttons-->
      {#if !noForce}
      <Button class="actions" color="primary" on:click={testItem}>
        <T id="instances.TestSelected" count={selectedCount}/>
      </Button>
      {/if}

      <Button class="actions" color="danger" on:click={deleteItem}>
        <T id="instances.DeleteSelected" count={selectedCount}/>
      </Button>
    </Fade>
  </Collapse>
</div>

<style>
  #footer :global(button.actions) {
    padding:0px 6px 0px 6px;
    margin:4px 1px;
  }
</style>
