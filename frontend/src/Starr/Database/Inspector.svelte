
<script>
    export let instance

    import { Input, InputGroup, InputGroupText,} from "sveltestrap"
    import { app, conf } from "../../libs/config.js"
    import { port } from "../../libs/info.js"
    import Fa from "svelte-fa"
    import { faFolderOpen, faLock, faUnlock, faEye, faEyeSlash } from "@fortawesome/free-solid-svg-icons"
    import { PickFile, SaveInstance, TestInstance, RemoveInstance } from "../../../wailsjs/go/app/App.js"
    import { toast } from "../../libs/funcs.js"
    import { onDestroy } from "svelte"
    import { _ } from "../../libs/Translate.svelte"
</script>

inspect {instance.DBPath} on this page
