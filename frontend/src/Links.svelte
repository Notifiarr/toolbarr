<script lang="ts">
  export let hidden: boolean = false

  import { Container, Row } from "sveltestrap"
  import BGLogo from "/src/libs/BackgroundLogo.svelte"
  import { _ } from "/src/libs/Translate.svelte"
  import { Card } from "sveltestrap"
</script>

<BGLogo {hidden}  url="notifiarr">
  <Container>
    <Row>
      <h1>{$_("words.Resources")}</h1>
      <Card body color="danger">{$_("incompletePage")}</Card>
      <p>
        Starr apps, unpackerr, notifiarr, autobrr, others?<br>
        We can make this load dynamically from toys-arr.us.
      </p>
    </Row>
  </Container>
</BGLogo>
