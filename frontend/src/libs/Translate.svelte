<!-- this component allows a slightly easier syntax for html translations with values. -->

<script context="module">
  import { _ as t } from "./locale"
  export const _ = t // pass it through.
</script>

<script>
  export let id
</script>
{@html $t(id,{values:{...$$props}})}
