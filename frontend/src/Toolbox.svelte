<script>
  import { Container, Row } from "sveltestrap"
  import BGLogo from "./libs/BackgroundLogo.svelte"
  import { _ } from "./libs/Translate.svelte"
</script>

<BGLogo url="notifiarr">
  <Container>
    <Row>
      <h1>{$_("OtherTools")}</h1>
      <p>No idea what goes here. Gimme ideas!</p>
    </Row>
  </Container>
  </BGLogo>
