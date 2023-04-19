<svelte:options tag="open-browser"/>

<!-- This allows us to open an external link in a browser. -->
<script lang="ts">
  import { BrowserOpenURL } from "../../wailsjs/runtime/runtime"
  export let href
  const openUrl = (e) => { e.preventDefault(); BrowserOpenURL(href) }
</script>

<a href="#top" {...$$props} on:click={openUrl}><slot></slot></a>
