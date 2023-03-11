<script>
  import mainLogo from "./assets/images/notifiarr.png"
  import { Container, Row } from "sveltestrap"
  import BGLogo from "./BackgroundLogo.svelte"
</script>

<BGLogo>
  <Container>
    <Row>
        <img alt="Main logo" id="logo" src="{mainLogo}" />
        <p class="main">The Starr of the show.</p>
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
