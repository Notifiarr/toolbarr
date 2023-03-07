<script>
  export let dark
  import { 
    Badge,
    Container,
    Form,
    Input,
    InputGroup,
    InputGroupText,
    Nav,
    NavLink,
    Row,
    Tooltip,
  } from "sveltestrap"
  import Logs from "./Logs.svelte"
  import Advanced from "./Advanced.svelte"
  import {GetConfig} from "../../wailsjs/go/app/App.js"
  
  let activeTab = "Logs"

  let conf = {}
  GetConfig().then(result => conf = result)

</script>

<Container>
  <Row>
    <p>
      This is where the application settings are found.
    </p>
    <Form class="Settings">
      <Tooltip target="ConfigFilePath" placement="bottom">May not be changed</Tooltip>
      <InputGroup id="ConfigFilePath">
        <InputGroupText>Config File</InputGroupText>
        <Input disabled value={conf.File} />
      </InputGroup>
      <br />
      <Nav tabs fill>
        <NavLink href="#" on:click={() => (activeTab = "Logs")} active={activeTab == "Logs"}>Logging</NavLink>
        <NavLink href="#" on:click={() => (activeTab = "Advanced")} active={activeTab == "Advanced"}>Advanced</NavLink>
      </Nav>
      <br />
      {#if activeTab == "Advanced"}<Advanced />{/if}
      {#if activeTab == "Logs"}<Logs {dark} />{/if}
    </Form>
  </Row>
</Container>

<style>
  :global(.Settings .input-group-text) {
    min-width:120px;
    max-width:120px;
  }
</style>