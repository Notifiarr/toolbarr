<script>
  import { InputGroup, InputGroupText } from "sveltestrap"
  import ConfigInput from "../libs/Input.svelte"
  import { app, conf } from "../libs/config.js"
</script>

<p>General application settings.</p>
<InputGroup>
  <InputGroupText class="setting-name">Language</InputGroupText>
  <ConfigInput type="select" id="Lang" tooltip="Only English works.">
    <option value="{$conf.Lang}">{$conf.Lang}</option>
    {#each $app.Langs as $lang}
      {#if $lang != $conf.Lang}
        <option value="{$lang}">{$lang}</option>
      {/if}
    {/each}
  </ConfigInput>
</InputGroup>
