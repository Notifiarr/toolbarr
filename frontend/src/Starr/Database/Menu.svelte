
<script>
  export let instance

  import { Button, Container, FormGroup, Input, InputGroup, InputGroupText, Row,} from "sveltestrap"
  import { app, conf } from "../../libs/config.js"
  import { port } from "../../libs/info.js"
  import Fa from "svelte-fa"
  import { faFolderOpen, faLock, faUnlock, faEye, faEyeSlash } from "@fortawesome/free-solid-svg-icons"
  import { PickFile, SaveInstance, TestInstance, RemoveInstance } from "../../../wailsjs/go/app/App.js"
  import { toast } from "../../libs/funcs.js"
  import { onDestroy } from "svelte"
  import { _ } from "../../libs/Translate.svelte"
  import Inspector from "./Inspector.svelte"
  import Migrate from "./Migrate.svelte"

  let activeTab=Migrate
</script>

<!-- 
<FormGroup floating>
  <div class="text-danger input-label" slot="label" style="display:{reset.Pass != instance.Pass?"block":"none"}">{$_("words.Unsaved")}</div>
  <InputGroup>
    <InputGroupText class="setting-name"></InputGroupText>
    <Input invalid={reset.Pass != instance.Pass} type={passLocked?"password":"text"} id="Pass" placeholder="******" bind:value={instance.Pass}/>
  </InputGroup>
</FormGroup> -->

<Input type="select" bind:value={activeTab}>
  <option value={Migrate}>Migrate Filesystem Paths</option>
  <option value={Inspector}>SQLite3 Database Explorer</option>
</Input>

<svelte:component this={activeTab} {instance}/>
