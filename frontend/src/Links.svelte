<script>
  export let hidden = false
  import { Container, Row } from "sveltestrap"
  import BGLogo from "./libs/BackgroundLogo.svelte"
  import { _ } from "./libs/Translate.svelte"
  import { Alert } from "sveltestrap"
</script>

<BGLogo {hidden}  url="notifiarr">
  <Container>
    <Row>
      <h1>{$_("words.Resources")}</h1>
      <Alert color="danger">{$_("incompletePage")}</Alert>
      <p>
        Starr apps, unpackerr, notifiarr, autobrr, others?<br>
        We can make this load dynamically from toys-arr.us.
      </p>
    </Row>
  </Container>
</BGLogo>
