import {defineConfig} from "vite"
import {svelte} from "@sveltejs/vite-plugin-svelte"

// https://github.com/sveltejs/vite-plugin-svelte/issues/270#issuecomment-1033190138
export default defineConfig({
  plugins: [
    svelte({
      experimental: {
        dynamicCompileOptions({code}) {
          if(isWebComponentSvelte(code)) {
            return {
              customElement: true
            }
          }
        }
      }
    })
  ]
})

function isWebComponentSvelte(code) {
    const svelteOptionsIdx = code.indexOf('<svelte:options ')
    if(svelteOptionsIdx < 0) {
        return false
    }
    const tagOptionIdx = code.indexOf('tag=', svelteOptionsIdx)
    const svelteOptionsEndIdx = code.indexOf('>',svelteOptionsIdx);
    return tagOptionIdx > svelteOptionsIdx && tagOptionIdx < svelteOptionsEndIdx
}