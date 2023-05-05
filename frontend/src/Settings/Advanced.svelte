<script lang="ts">
  import { Input, InputGroup, InputGroupText } from "sveltestrap"
  import ConfigInput from "/src/libs/Input.svelte"
  import { app } from "/src/libs/config"
  import { _ } from "/src/libs/Translate.svelte"
</script>

<p>{$_("advanced_application_settings")}</p>
<InputGroup>
  <ConfigInput type="select" id="DevMode">
    <option value={true}>{$_("configvalues.Enabled")}</option>
    <option value={false}>{$_("configvalues.Disabled")}</option>
  </ConfigInput>
</InputGroup>
<InputGroup>
  <ConfigInput type="select" id="Updates">
    <option value="production">{$_("configvalues.Production")}</option>
    <option value="unstable">{$_("configvalues.Unstable")}</option>
  </ConfigInput>
</InputGroup>
<InputGroup>
  <InputGroupText class="setting-name">{$_("configinput.AppPath")}</InputGroupText>
  <Input disabled type="text" value={$app.Exe} />
</InputGroup>
