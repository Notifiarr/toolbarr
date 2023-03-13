<script lang="ts">
  export let app

  import { Container, Nav, NavLink, Row } from "sveltestrap"
  import Instances from "./Instances.svelte"
  import Database from "./Database.svelte"
  import Source from "./Source.svelte"
  import BGLogo from "../libs/BackgroundLogo.svelte"

  let tab = Instances
</script>

<BGLogo url="starr">
  <Nav tabs fill>
    <NavLink href="#" on:click={() => (tab = Instances)} active={tab == Instances}>{app} Tools</NavLink>
    <NavLink href="#" on:click={() => (tab = Database)} active={tab == Database}>Database</NavLink>
    <NavLink href="#" on:click={() => (tab = Source)} active={tab == Source}>Source</NavLink>
  </Nav>
  <br />
  <Container>
    <Row>
      <svelte:component this={tab} {app}/>
    </Row>
  </Container>
</BGLogo>
