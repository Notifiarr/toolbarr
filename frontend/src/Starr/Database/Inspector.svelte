<script>
    export let instance
    // import { Alert, Card, Input, InputGroup, InputGroupText,} from "sveltestrap"
    // import { app, conf } from "../../libs/config.js"
    // import { port } from "../../libs/info.js"
    // import Fa from "svelte-fa"
    // import { faFolderOpen, faLock, faUnlock, faEye, faEyeSlash } from "@fortawesome/free-solid-svg-icons"
    // import { PickFile, SaveInstance, TestInstance, RemoveInstance } from "../../../wailsjs/go/app/App.js"
    // import { toast } from "../../libs/funcs.js"
    // import { onDestroy } from "svelte"
    import { Card } from "sveltestrap"
    import { _ } from "../../libs/Translate.svelte"
</script>

<h4>{instance.App} DB Inspector</h4>
{#if instance.DBPath == ""}
  <Card body color="danger">
    The selected {instance.App} instance ({instance.Name}) has no SQLite3 database file path configured.
  </Card>
{:else}
  DB explorer for {instance.Name} with DB Path {instance.DBPath} goes here....
{/if}