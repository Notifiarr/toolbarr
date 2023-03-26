<script>
    export let instance

    import { Card } from "sveltestrap"
    import { _ } from "../../libs/Translate.svelte"
</script>

<h4>{instance.App} DB Migrations</h4>
<p>
  This tool allows you to update the paths for the media in your {instance.App} library.
  Use this before moving your library from one server to another if the media content paths are changing.
</p>
<hr>

{#if instance.DBPath == ""}
  <Card body color="danger">
    The selected {instance.App} instance ({instance.Name}) has no SQLite3 database file path configured.
  </Card>
{:else}
  Migration form for <b>{instance.Name}</b> goes here.<br>DB Path: <code>{instance.DBPath}</code>
{/if}
