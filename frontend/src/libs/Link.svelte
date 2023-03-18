<!-- This allows us to open an external link in a browser. -->
<script type="ts">
  import { BrowserOpenURL } from "../../wailsjs/runtime/runtime";
  export let href
</script>

<a href="#top" on:click={() => (BrowserOpenURL(href))}><slot></slot></a>
