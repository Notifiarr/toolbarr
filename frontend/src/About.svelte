<script>
    import Fa from "svelte-fa"
    import {faGithub, faDiscord} from "@fortawesome/free-brands-svg-icons"
    import {BrowserOpenURL} from "../wailsjs/runtime"
    import {Container, Row, Table, Col} from "sveltestrap"
</script>

<Container>
  <Row>
    <h1>About Toolbarr</h1>
    <p>
      Toolbarr fixes problems with Starr apps. It comes with a five starr rating from <a href="#top" on:click={() => (BrowserOpenURL("https://toys-arr.us"))}>Toys Arr Us</a>!
    </p>
    <Col>
      <Table class="table">
        <tr>
          <td style="width:180px;"><a href="#top" on:click={() => (BrowserOpenURL("https://github.com/Notifiarr/toolbarr"))}><Fa icon={faGithub} /> Toolbarr GitHub</a></td> 
          <td>Take a visit to the sausage factory.</td>
        </tr>
        <tr>
          <td><a href="#top" on:click={() => (BrowserOpenURL("https://notifiarr.com/discord"))}><Fa fw icon={faDiscord} /> Notifiarr Discord</a></td>
          <td>For all your notifications needs.</td>
        </tr>
        <tr>
          <td><a href="#top" on:click={() => (BrowserOpenURL("https://golift.io/discord"))}><Fa fw icon={faDiscord} /> Go Lift Discord</a></td>
          <td>Where the codes get cooked.</td>
        </tr>
    </Table>
  </Col>
    <h3>Attribution</h3>
    <p>
      <li>
        Created by <a href="#top" on:click={() => (BrowserOpenURL("https://golift.io"))}>Go Lift</a> 
        for <a href="#top" on:click={() => (BrowserOpenURL("https://notifiarr.com"))}>Notifiarr</a>.
      </li>
      <li>
        Backgrounds by <a href="#top" on:click={() => (BrowserOpenURL("https://rawpixel.com"))}>rawpixel.com</a> 
        on <a href="#top" on:click={() => (BrowserOpenURL("https://www.freepik.com/author/rawpixel-com"))}>Freepik</a>.
      </li>
      <li>
        Written in <a href="#top" on:click={() => (BrowserOpenURL("https://svelte.dev"))}>Svelte</a> 
        using <a href="#top" on:click={() => (BrowserOpenURL("https://wails.io"))}>Wails</a>.
      </li>
    </p>
  </Row>
</Container>

