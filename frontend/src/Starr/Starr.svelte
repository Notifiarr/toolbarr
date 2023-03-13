<script lang="ts">
  export let app

  import { Container, Form, Nav, NavLink, Row } from "sveltestrap"
  import Instances from "./Instances.svelte"
  import Database from "./Database.svelte"
  import Source from "./Source.svelte"
  import BGLogo from "../libs/BackgroundLogo.svelte"

  let tab = Instances
</script>

<BGLogo url="starr">
  <Container>
    <Form class="Starr">
      <Row>
        <Nav tabs fill>
          <NavLink href="#" on:click={() => (tab = Instances)} active={tab == Instances}>{app} Tools</NavLink>
          <NavLink href="#" on:click={() => (tab = Database)} active={tab == Database}>Database</NavLink>
          <NavLink href="#" on:click={() => (tab = Source)} active={tab == Source}>Source</NavLink>
        </Nav>
        <br />
        <svelte:component this={tab} {app} />
      </Row>
    </Form>
  </Container>
</BGLogo>
