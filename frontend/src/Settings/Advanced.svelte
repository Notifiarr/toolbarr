<script>
  import { InputGroup, InputGroupText } from "sveltestrap"
  import ConfigInput from "../libs/Input.svelte"
</script>

<p>These settings control advanced aspects of this application, and should probably not be changed.</p>
<InputGroup>
  <InputGroupText class="setting-name">Dev Mode</InputGroupText>
  <ConfigInput type="select" id="Advanced.DevMode" name="DevMode"
    tooltip="Enable this when a developer instructs you to do so.">
    <option value={true}>Enabled</option>
    <option value={false}>Disabled</option>
  </ConfigInput>
</InputGroup>
<InputGroup>
  <InputGroupText class="setting-name">Updates</InputGroupText>
  <ConfigInput type="select" id="Advanced.Updates" name="Updates"
    tooltip="Production gets updates from github, unstable gets them from unstable.golift.io">
    <option value="production">Production</option>
    <option value="unstable">Unstable</option>
  </ConfigInput>
</InputGroup>
<InputGroup>
  <InputGroupText class="setting-name">App Path</InputGroupText>
  <ConfigInput locked type="text"id="App.Exe" name="Exe" />
</InputGroup>
