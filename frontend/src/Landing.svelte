<script lang="ts">
  export let hidden = false

  import mainLogo from "/src/assets/images/logo.png"
  import BGLogo from "/src/libs/BackgroundLogo.svelte"
  import { _ } from "/src/libs/Translate.svelte"
</script>

<BGLogo {hidden}>
  <center>
    <img alt="{$_("words.Main")} {$_("words.Logo")}" id="logo" src={mainLogo} />
    <p>{$_("theStarroftheshow")}</p>
  </center>
</BGLogo>

<style>
  #logo {
    max-width: 100%;
    max-height: calc(100vh - 100px);
  }
</style>
