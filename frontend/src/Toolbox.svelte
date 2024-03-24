<script lang="ts">
  export let hidden: boolean = false

  import { Container, Row } from "@sveltestrap/sveltestrap"
  import BGLogo from "/src/libs/BackgroundLogo.svelte"
  import { _ } from "/src/libs/Translate.svelte"
  import { Card } from "@sveltestrap/sveltestrap"
</script>

<BGLogo {hidden} url="notifiarr">
  <Container>
    <Row>
      <h1>{$_("OtherTools")}</h1>
      <Card body color="danger">{$_("incompletePage")}</Card>
      <p>No idea what goes here. <open-browser href="https://github.com/Notifiarr/toolbarr/issues/new">Gimme ideas!</open-browser></p>
    </Row>
  </Container>
</BGLogo>
