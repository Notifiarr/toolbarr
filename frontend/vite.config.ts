import { defineConfig } from "vite"
import { svelte } from "@sveltejs/vite-plugin-svelte"


// https://github.com/sveltejs/vite-plugin-svelte/issues/270#issuecomment-1033190138
export default defineConfig({
  plugins: [ svelte({
    dynamicCompileOptions({code}) {
      return { customElement: isWebComponentSvelte(code) }
    },
  })]
})

// A .svelte file with this tag creates a web component.
// <svelte:options tag="some-thing">
function isWebComponentSvelte(code) {
  const svelteOptionsIdx = code.indexOf('<svelte:options ')
  if(svelteOptionsIdx < 0) return false

  // Check if `tag=` exists as a prop in svelte:options.
  const tagOptionIdx = code.indexOf('customElement=', svelteOptionsIdx)
  const svelteOptionsEndIdx = code.indexOf('>',svelteOptionsIdx);
  return tagOptionIdx > svelteOptionsIdx && tagOptionIdx < svelteOptionsEndIdx
}