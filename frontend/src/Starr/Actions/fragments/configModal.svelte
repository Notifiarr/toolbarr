<script lang="ts">
  export let form: any
  export let isOpen = false
  export let id: number | string
  export let name: string
  export let idx: number
  export let info: any
  export let str: string
  export let disabled = ""

  import { _ } from "../../../libs/Translate.svelte"
  import { Badge, Button, Modal, ModalBody, ModalFooter, ModalHeader } from "sveltestrap"

  function reset(e) {
    e.preventDefault()
    form[idx] = JSON.parse(JSON.stringify(info[idx]))
  }

  function toggle(e?: any) {
    e?.preventDefault()
    isOpen = false
  }

  function onkeydown(e) {
    if (e.key == "Escape") if (isOpen) e.preventDefault()
  }

  function onkeyup(e) {
    if (e.key == "Escape") isOpen = false
  }
</script>

<svelte:window on:keyup={onkeyup} on:keydown={onkeydown}/>

<Modal body size="lg" scrollable isOpen={isOpen}>
  <ModalHeader {toggle}>
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
      {$_("words.Close")}
    </Button>
  </ModalFooter>
</Modal>
