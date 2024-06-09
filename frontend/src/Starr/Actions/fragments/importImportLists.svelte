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

{Object.keys(info[0])}
