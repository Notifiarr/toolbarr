<script lang="ts">
  export let app

  import { Container, Nav, NavLink, Row } from "sveltestrap"
  import Instances from "./Instances.svelte"
  import Database from "./Database.svelte"
  import Source from "./Source.svelte"
  import BGLogo from "../BackgroundLogo.svelte"

  let activeTab = "Instances"
</script>

<BGLogo url="starr">
  <Nav tabs fill>
    <NavLink href="#" on:click={() => (activeTab = "Instances")} active={activeTab == "Instances"}>{app} Tools</NavLink>
    <NavLink href="#" on:click={() => (activeTab = "Database")} active={activeTab == "Database"}>Database Tools</NavLink>
    <NavLink href="#" on:click={() => (activeTab = "Source")} active={activeTab == "Source"}>Source Info</NavLink>
  </Nav>
  <br />
  <Container>
    <Row>
      {#if activeTab == "Instances"}<Instances {app} />{/if}
      {#if activeTab == "Database"}<Database {app} />{/if}
      {#if activeTab == "Source"}<Source {app} />{/if}
    </Row>
  </Container>
</BGLogo>
