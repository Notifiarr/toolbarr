<script>
  import mainLogo from "./assets/images/notifiarr.png"
  import { Container, Row } from "sveltestrap"
  import BGLogo from "./libs/BackgroundLogo.svelte"
  import { _ } from "./libs/Translate.svelte"
</script>

<BGLogo>
  <Container>
    <Row>
        <img alt="{$_("words.Main")} {$_("words.Logo")}" id="logo" src="{mainLogo}" />
        <p class="main">{$_("theStarroftheshow")}</p>
    </Row>
  </Container>
</BGLogo>

<style>
  .main {
    text-align: center;
  }

  #logo {
    display: block;
    margin: auto;
    padding: 50px 0 0;
    width:75%;
    max-width:512px;
  }
</style>
