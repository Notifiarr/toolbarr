<script lang="ts">
  export let form: any = undefined
  export let isOpen = false
  export let id: number | string
  export let name: string
  export let idx: number = 0
  export let info: any = undefined
  export let str: string = ""
  export let disabled = ""
  export let closeButton = $_("words.Close")
  export let callback: VoidFunction|undefined = undefined

  import { _ } from "/src/libs/Translate.svelte"
  import { Badge, Button, Modal, ModalBody, ModalFooter, ModalHeader } from "@sveltestrap/sveltestrap"

  function reset(e: MouseEvent) {
    e.preventDefault()
    form[idx] = JSON.parse(JSON.stringify(info[idx]))
  }

  function toggle() {
    if (callback) callback()
    if (form == undefined && disabled != "") info = undefined
    isOpen = false
  }

  function onkeydown(e: KeyboardEvent) {
    if (e.key == "Escape") if (isOpen) e.preventDefault()
  }

  function onkeyup(e: KeyboardEvent) {
    if (e.key == "Escape") isOpen = false
  }
</script>

<svelte:window on:keyup={onkeyup} on:keydown={onkeydown}/>

<Modal body size="lg" scrollable isOpen={isOpen}>
  <ModalHeader toggle={toggle}>
    <Badge color="info">{id}</Badge> {name}
  </ModalHeader>

  <ModalBody>
    <slot></slot>
  </ModalBody>

  <ModalFooter>
    {#if disabled}
      <p>{disabled}</p>
    {:else}
      <Button size="sm" disabled={str==JSON.stringify(form)} color="warning" on:click={reset}>
        {$_("words.Reset")}
      </Button>
    {/if}
    <Button size="sm" color="info" on:click={toggle}>
      {closeButton}
    </Button>
  </ModalFooter>
</Modal>
