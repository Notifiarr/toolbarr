<script context="module">
  import { Input } from "sveltestrap"
  import { toast } from "./funcs.js"
  import { GetConfig, SaveConfigItem } from "../../wailsjs/go/app/App.js"
  import { devMode } from '../Settings/settings.js'
</script>

<script>
  // Like `text` or `select`.
  export let type
  // Used to update the valie in Go. Needs full struct path for gorilla/schema.
  export let id
  // Used as the config item name, only struct member name goes here.
  export let name
  // Used if you don't want this value changed.
  export let readonly = false
  // Similar to readonly, but the input dims/greys out.
  export let disabled = false
  // Similar to readonly, but the update method won't work.
  export let locked = false
  // Avoid relaoding running config when saving? Most things need to.
  export let noreload = false
  // Avoid sending a toast when not in dev mode?
  export let notoast = false
  // Optional value. Should only be used for binding.
  export let value = undefined

  GetConfig().then(result => value = result[name])

  let valid = false
  let invalid = false
  
  const checkbox = (type=="switch" || type=="checkbox")

  function saveValue(e) {
    if (locked) return
    /* I really do not know why this works, and nothing else will. */
    const confValue = e.target.value!=""?e.target.value+"":e.target.checked+""
    SaveConfigItem(id, confValue, !noreload).then((msg) => {
      if (!notoast) {toast("success", msg, "CONFIG")}
      if (notoast&&$devMode) {toast("warning", msg, "DEBUG")}
      invalid = false
      valid = true
      setInterval(() => {valid=false}, 5000)
    }, (error) => {
      toast("error", error)
      invalid = true
      valid = false
    })
  }

  // Allows importers to call the exported functions below.
  let input 

  export function update(newVal) {
    value = newVal
    saveValue({target:{value:newVal}})
  }
</script>

{#if value != undefined}
  {#if checkbox}
    <Input {type} {id} {name} 
      {disabled} {valid} {invalid}
      readonly={(readonly||locked)}
      bind:this={input}
      on:change={saveValue}
      bind:checked={value}
    ><slot/></Input>
  {:else}
    <Input {type} {id} {name} 
      {disabled} {valid} {invalid}
      readonly={(readonly||locked)}
      bind:this={input}
      on:change={saveValue}
      bind:value={value}
    ><slot/></Input>
  {/if}
{/if}