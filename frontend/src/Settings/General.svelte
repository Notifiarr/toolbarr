<script>
  import { InputGroup, InputGroupText } from "sveltestrap"
  import ConfigInput from "../libs/Input.svelte"
  import { conf } from "../libs/config.js"
  import { Languages } from "../../wailsjs/go/app/App.js"
  import { _ } from "../libs/locale"

  let langs = undefined
  const update = () => Languages().then(v => langs = v)
  update()
</script>

<p>{$_("general_application_settings")}</p>

<InputGroup>
  <ConfigInput on:change={update} type="select" id="Lang">
    {#if langs != undefined && $conf.Lang != undefined}
      <option value="{$conf.Lang}">{langs[$conf.Lang]}</option>
      {#each Object.keys(langs) as $id}
        {#if $id != $conf.Lang}
          <option value="{$id}">{langs[$id]}</option>
        {/if}
      {/each}
    {/if}
  </ConfigInput>
</InputGroup>
