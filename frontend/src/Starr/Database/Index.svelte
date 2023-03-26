<script>
  export let starrApp

  import { Input, InputGroup, InputGroupText, Button, Form, Alert, FormGroup, Badge, TabContent, TabPane } from "sveltestrap"
  import { app, conf } from "../../libs/config.js"
  import { port } from "../../libs/info.js"
  import Fa from "svelte-fa"
  import { faFolderOpen, faLock, faUnlock, faEye, faEyeSlash } from "@fortawesome/free-solid-svg-icons"
  import { PickFile, SaveInstance, TestInstance, RemoveInstance } from "../../../wailsjs/go/app/App.js"
  import { toast } from "../../libs/funcs.js"
  import { onDestroy } from "svelte"
  import { _ } from "../../libs/Translate.svelte"
  import Menu from "./Menu.svelte"

  let width
</script>

<svelte:window bind:innerWidth={width}/>

<div class="container">
  <TabContent vertical={width>767} pills>
    {#if $conf.Instances[starrApp] != undefined}
    {#each $conf.Instances[starrApp] as instance}
      {#if instance.DBPath != ""}
        <TabPane tabId="Config" tab={instance.Name} active><Menu {instance} /></TabPane>
      {/if}
    {/each}
    {/if}
  </TabContent>
</div>

<style>
  .container {
    margin-top: 5px;
  }
</style>

