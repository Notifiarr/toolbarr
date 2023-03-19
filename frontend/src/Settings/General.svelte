<script lang="ts">
  import { Button, InputGroup, Offcanvas } from "sveltestrap"
  import ConfigInput from "../libs/Input.svelte"
  import { conf } from "../libs/config.js"
  import { Languages } from "../../wailsjs/go/app/App.js"
  import { _ } from "../libs/locale"
  import { faQuestion } from "@fortawesome/free-solid-svg-icons"
  import Fa from "svelte-fa"

  let langHelp = false
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
  <Button on:click={(e) => {e.preventDefault();langHelp = !langHelp}}>
    <Fa primaryColor="cyan" icon="{faQuestion}" />
  </Button>
</InputGroup>

<Offcanvas
  style="width:50%;min-width:390px;max-width:550px"
  class="{$conf.Dark ? "bg-secondary" : "bg-light"}"
  isOpen={langHelp}
  toggle={() => {langHelp = !langHelp}}
  header={$_("TranslationInformation")} placement="end">
  <p>{@html $_("LanguageHelpText", {values: {url: "https://translate.notifiarr.com/projects/toolbarr/"}})}</p>
</Offcanvas>