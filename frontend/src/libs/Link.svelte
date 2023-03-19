<svelte:options tag="a-link"/>

<!-- This allows us to open an external link in a browser. -->
<script type="ts">
  import { BrowserOpenURL } from "../../wailsjs/runtime/runtime"
  export let url
  const openUrl = (e) => { e.preventDefault(); BrowserOpenURL(url) }
</script>

<a color={color} href="#top" on:click={openUrl}><slot></slot></a>
