<script lang="ts">
  export let instance: Instance
  export let info: any
  export let form: any
  export let str: string
  export let updating: boolean
  export let selected: {[key: string]: boolean}
  export let tab: Tab
  export let noForce = false

  import type { Tab } from "./tabs.svelte"
  import { Alert, Button, Collapse, Fade, Tooltip, Icon, Card } from "sveltestrap"
  import Loading from "../../loading.svelte"
  import T, { _ } from "../../../libs/Translate.svelte"
  import { toast, count } from "../../../libs/funcs"
  import { update, remove, fixFieldValues } from "../methods"
  import type { Instance } from "../../../libs/config"
  import { createEventDispatcher } from "svelte"

  const dispatch = createEventDispatcher()
  let badMsg = ""
  let goodMsg = ""
  $: selectedCount = count(selected)        // How many items are selected.
  $: unSaved = JSON.stringify(form) !== str // True when something changed.
  let button

  function showMsg(idx, msg, data) {
    goodMsg += `<li>${$_("instances.SuccessMsg", {values:{"msg": msg}})}</li>`
    let kind = "update"
    
    if (data) { // update client (repalce in place)
      form[idx] = JSON.parse(JSON.stringify(data))
    } else {   // delete list item (remove in place)
      form.splice(idx, 1)
      kind = "delete"
    }

    str = JSON.stringify(form)
    info = JSON.parse(str)
    dispatch(kind)
  }

  function showError(idx, err) {
    form[idx] = JSON.parse(JSON.stringify(info[idx]))
    badMsg += `<li>${$_("instances.ErrorMsg", {values:{"msg": err}})}</li>`
  }

  async function updateItems(force) {
    toast("info", $_("instances.Updating"+tab.id))
    goodMsg = badMsg = ""
    updating = true

    for (var idx = 0; idx < form.length; idx++) {
      if (JSON.stringify(form[idx]) == JSON.stringify(info[idx])) continue // not changed
      if (noForce) {
        await update[tab.id][instance.App](instance, form[idx]).then(
          (resp) => showMsg(idx, resp.Msg, resp.Data),
          (err) => showError(idx, err)
        )
      } else {
        await update[tab.id][instance.App](instance, force, form[idx]).then(
          (resp) => showMsg(idx, resp.Msg, resp.Data),
          (err) => showError(idx, err)
        )
      }
    }

    updating = false
    str = fixFieldValues(info)
    form = JSON.parse(str)
    Object.keys(selected).forEach(k => selected[k] = false)
  }

  async function deleteItem() {
    toast("info", $_("instances.Deleting"+tab.id, { values:{"count": count(selected)} }))
    goodMsg = badMsg = ""
    updating = true

    for (var idx = info.length-1; idx >= 0; idx--) {
      if (!selected[info[idx].id]) continue // Not selected.
      await remove[tab.id][instance.App](instance, info[idx].id).then(
        (msg) => showMsg(idx, msg, false),
        (err) => showError(idx, err)
      )
    }

    updating = false
    Object.keys(selected).forEach(k => selected[k] = false)
  }
</script>

<div id="footer">
  <Collapse isOpen={goodMsg != ""}>
    <Alert dismissible color="success" fade={false}>{@html goodMsg}</Alert>
  </Collapse>
  <Collapse isOpen={badMsg != ""}>
    <Alert dismissible color="danger" fade={false}>{@html badMsg}</Alert>
  </Collapse>
  <Loading isOpen={updating}/>

  <Collapse isOpen={!updating && (unSaved||selectedCount > 0)}>
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
    <Fade style="display:inline-block" isOpen={selectedCount > 0}>
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