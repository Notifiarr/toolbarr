<!-- this component allows a slightly easier syntax for html translations with values. -->

<script context="module">
  import { derived } from "svelte/store"
  import { _, locale } from "svelte-i18n"
  export { _ } // pass it through
  export const isReady = derived(locale, $locale => typeof $locale === 'string');
</script>

<script>
  export let id
</script>

{#if $isReady == true}
  {@html $_(id,{values:{...$$props}})}
{/if}