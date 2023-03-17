<script>
  import { InputGroup, InputGroupText } from "sveltestrap"
  import ConfigInput from "../libs/Input.svelte"
  import { conf } from "../libs/config.js"
</script>

<p>General application settings.</p>
<InputGroup>
  <InputGroupText class="setting-name">Language</InputGroupText>
  <ConfigInput type="select" id="Language" tooltip="Not working yet.">
    <option value="{$conf.Language}">{$conf.Language}</option>
  </ConfigInput>
</InputGroup>
<InputGroup>
  <InputGroupText class="setting-name">Time Zone</InputGroupText>
  <ConfigInput type="select" id="TimeZone" tooltip="Not woring yet.">
    <option value="{$conf.TimeZone}">{$conf.TimeZone}</option>
  </ConfigInput>
</InputGroup>
