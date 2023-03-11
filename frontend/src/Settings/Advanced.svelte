<script>
  import {Input,InputGroup,InputGroupText,Tooltip} from "sveltestrap"
  import {GetConfig} from "../../wailsjs/go/app/App.js"
	import {devMode,dark} from './store.js'
  import { saveValue } from "../funcs";

  let validProps = {}
  let invalidProps = {}
  $: dm = $devMode + ""
  let conf = {}
  GetConfig().then(result => conf = result)

  function saveInput(e) {
    if (e.target.id == "DevMode") {
      $devMode = e.target.value == "false" ? false : true
    }
    validProps[e.target.id] = saveValue(e.target.name, e.target.value, $dark, false)
    invalidProps[e.target.id] = !validProps[e.target.id]
    setInterval(() => {validProps[e.target.id]=false}, 5000)
  }
</script>

<InputGroup>
  <InputGroupText class="setting-name">Dev Mode</InputGroupText>
  <Input valid={validProps.DevMode} bind:invalid={invalidProps.DevMode} on:change={saveInput} value={dm} type="select" name="Advanced.DevMode" id="DevMode">
    <option value="true">Enabled</option>
    <option value="false">Disabled</option>
  </Input>
  <Tooltip target="DevMode" placement="top">Enable this when a developer instructs you to do so.</Tooltip>
</InputGroup>
<InputGroup>
  <InputGroupText class="setting-name">Updates</InputGroupText>
  <Input valid={validProps.Updates} bind:invalid={invalidProps.Updates} value={conf.Updates} on:change={saveInput} type="select" name="Advanced.Updates" id="Updates">
    <option value="production">Production</option>
    <option value="unstable">Unstable</option>
  </Input>
</InputGroup>