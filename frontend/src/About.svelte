<script>
  import Fa from "svelte-fa"
  import { faGithub, faDiscord } from "@fortawesome/free-brands-svg-icons"
  import { faGear } from "@fortawesome/free-solid-svg-icons"
  import { BrowserOpenURL, EventsOn, EventsOff } from "../wailsjs/runtime"
  import { Version, CheckUpdate, DownloadUpdate, LaunchInstaller, OpenFolder } from "../wailsjs/go/app/App"
  import { Container, Row, Table, Col, Card, Tooltip, Button } from   "sveltestrap"
  import { tweened } from 'svelte/motion'
  import BGLogo from "./BackgroundLogo.svelte"
  import { dark } from './Settings/store.js'
  import { toast } from "./funcs";

  let version
  let timer
  Version().then(result => {
    version = result
    timer = tweened(version.Running)
  })

  let update = {}
  let release = {}
  let progress = 0.0

  function checkUpdate(e) {
    e.preventDefault()
    update.Downloading = "Checking for update..."
    CheckUpdate().then(result => {
      release = result
      update.Checked = true
      update.Downloading = ""
    }, (error) => {
      update.Failed = "Error checking for update"
      toast("error", "", error, $dark)
      update.Downloading = ""
    })
  }

  function installUpdate(e) {
    e.preventDefault()
    update.Downloading = "Launching installer.."
    LaunchInstaller(release.FilePath).then(result => {
      update.Downloading = result
    })
  }

  function openFolder(e) {
    e.preventDefault()
    OpenFolder(release.FilePath).then(msg => (toast("warning", "", msg, $dark)))
  }

  function downloadUpdate(e) {
    e.preventDefault()
    update.Downloading = "Downloading the update... "
    DownloadUpdate().then(data => {
      update.Checked = true
      release.FilePath = data.Path
      toast("success", "", data.Path, $dark)
      EventsOn("downloadFinished", (data) => {
        EventsOff("downloadProgress", "downloadFinished")
        update.Downloaded = "Open Update File"
        update.Downloading = ""
        progress = 0.0
      })
      EventsOn("downloadProgress", (data) => {
        progress = data
      })
    }, (error) => {
      update.Failed = "Error checking for update"
      toast("error", "", error, $dark)
      update.Downloading = ""
    })
  }

  setInterval(() => {
    if ($timer > 0) $timer++;
  }, 1000);
  /* All of this is to create a "running" timer. */
  $: days = Math.floor($timer / 86400);
  $: hours = Math.floor(($timer - (days * 86400)) / 3600);
  $: minutes = Math.floor(($timer - (days * 86400) - (hours * 3600)) / 60);
  $: seconds = Math.floor(($timer - (days * 86400) - (hours * 3600) - (minutes * 60)))
  $: uptime = (days > 0 ? days + 'd ' : '') + (hours > 0 ? hours + 'h ' : '') + 
        (minutes > 0 ? minutes + 'm ' : '') + (seconds > 0 ? seconds + 's ' : '')
</script>

<BGLogo>
  <Container>
    <Row>
      <h1>About Toolbarr</h1>
      <p>
        Toolbarr fixes problems with Starr apps. It comes with a five starr rating from <a href="#top" on:click={() => (BrowserOpenURL("https://toys-arr.us"))}>Toys Arr Us</a>!
      </p>
      <Col md="6">
        <h3>Development</h3>
        <Table dark={$dark} responsive>
          <tr>
            <td style="width:180px;"><a href="#top" on:click={() => (BrowserOpenURL("https://github.com/Notifiarr/toolbarr"))}><Fa icon={faGithub} /> Toolbarr GitHub</a></td> 
            <td>Visit the sausage factory.</td>
          </tr>
          <tr>
            <td><a href="#top" on:click={() => (BrowserOpenURL("https://notifiarr.com/discord"))}><Fa fw icon={faDiscord} /> Notifiarr Discord</a></td>
            <td>For your notifications needs.</td>
          </tr>
          <tr>
            <td><a href="#top" on:click={() => (BrowserOpenURL("https://golift.io/discord"))}><Fa fw icon={faDiscord} /> Go Lift Discord</a></td>
            <td>Code cookin' collaborators.</td>
          </tr>
        </Table>
      </Col>
      <Col md="6">
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
    </Col>
    </Row>
    {#if version}
    <Col md="6">
      <h3>App Info</h3>
      <Card  color={$dark ? 'dark' : 'light'} body>
        <Table dark={$dark} responsive>
          <tr><td>Version</td><td>v{version.Version}-{version.Revision} ({version.GoVersion})</td></tr>
          <tr><td>Branch</td><td>{version.Branch}</td></tr>
          <tr><td>Created</td><td>{version.BuildDate} by {version.BuildUser}</td></tr>
          <tr>
            <td>Running</td>
            <Tooltip target="runningTime" placement="top">Running Since:<br>{version.Started}</Tooltip>
            <td id="runningTime">{uptime}</td>
          </tr>
          <tr><td colspan="2">
            {#if update.Failed}
            <Button block disabled size="sm" color="danger">{update.Failed}</Button>
            {:else if update.Downloading}
            <Button block disabled size="sm" color="danger">{update.Downloading} <Fa spin primaryColor="yellow" icon={faGear} /></Button>
            {:else if update.Downloaded}
            <Button style="width:49%" outline on:click={installUpdate} size="sm" color="primary">{update.Downloaded}</Button>
            <Button style="width:49%" outline on:click={openFolder} size="sm" color="info">Open Folder</Button>
            {:else if release.Outdate}
            <Button block outline on:click={downloadUpdate} size="sm" color="warning">Download update: v{release.Current}</Button>
            {:else if update.Checked}
            <Button block color="success" disabled size="sm">Up to date! v{release.Current}</Button>
            {:else}
            <Button block outline size="sm" color="info" on:click={checkUpdate}>Check for update</Button>
            {/if}
          </td></tr>
        </Table>
        {#if progress}
        <progress value={progress}></progress>
        {/if}
      </Card>
    </Col>
    {/if}
  </Container>
</BGLogo>

<style>
  progress {
    width: 100%;
  }
</style>