<script type="ts">
  import Fa from "svelte-fa"
  import { faGithub, faDiscord } from "@fortawesome/free-brands-svg-icons"
  import { faGear } from "@fortawesome/free-solid-svg-icons"
  import { EventsOn, EventsOff } from "../wailsjs/runtime"
  import { Version, CheckUpdate, DownloadUpdate, LaunchInstaller, OpenFolder } from "../wailsjs/go/app/App"
  import { Container, Row, Table, Col, Card, Tooltip, Button, Progress, Badge } from   "sveltestrap"
  import { tweened } from "svelte/motion"
  import BGLogo from "./libs/BackgroundLogo.svelte"
  import A from "./libs/Link.svelte"
  import { conf } from "./libs/config.js"
  import { toast, onOnce, onInterval } from "./libs/funcs";

  let version = {
    Version: "",
    Revision: "",
    GoVersion: "",
    Running: 0,
    BuildUser: "",
    BuildDate: "",
    Branch: "",
    Started: "",
  }

  let timer = tweened(1)
  $: if (version.Running == 0) {
    Version().then(result => {
      version = result
      timer = tweened(version.Running)
      if ($conf.DevMode) {
        toast("warning", "This message should only show up when you load the page, "+
        "and again if the page gets reloaded or the window wakes from sleep. "+
        "Let captain know if it shows up a lot.", "Debug")
      }
    })
  }

  let update = {
    Downloading: "",
    Checked: false,
    Failed: "",
    Downloaded: "",
  }

  let release = {
    Outdate: false,
    FilePath: "",
    Current: "",
    Size: "",
  }

  let progress = 0.0
  let msg = "just waitin' for somethin' to happen"

  function checkUpdate(e) {
    e.preventDefault()
    update.Downloading = "Checking for update..."

    CheckUpdate().then(result => {
      release = result
      update.Checked = true
      update.Downloading = ""
      msg = "that was pretty cool"
      if (release.Outdate) msg = $conf.IsLinux ?
        "Use your package manager to install the update." :
        "Update available! Click the button to download it."
    }, (error) => {
      update.Failed = "Error checking for update"
      toast("error", error)
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
    OpenFolder(release.FilePath).then(msg => (toast("warning", msg)))
  }

  function downloadUpdate(e) {
    e.preventDefault()
    update.Downloading = "Downloading the update... "

    DownloadUpdate().then(data => {
      update.Checked = true
      release.FilePath = data.Path
      toast("success", "Downloading: "+data.Path)

      EventsOn("downloadFinished", (data) => {
        EventsOff("downloadProgress", "downloadFinished")
        update.Downloaded = "Open "+($conf.IsMac ? "DMG" : "Installer")
        update.Downloading = ""
        msg = ($conf.IsMac ? "Disk image" : "Installer") + " downloaded! Click a button to use it."
        onOnce(() => (progress = 0.0), 0.5)
      })

      EventsOn("downloadProgress", (data) => (progress = data))
    }, (error) => {
      update.Failed = "Error checking for update"
      toast("error", error)
      update.Downloading = ""
    })
  }

  let lastTime = (new Date()).getTime();
   onInterval(() => {
    const current = (new Date()).getTime()
    if ($timer > 0) {
      $timer++
      if (current > (lastTime + 1100)) version.Running = 0
    }
    lastTime = current
  }, 1000)


  /* All of this is to create a "running" timer. */
  $: days = Math.floor($timer / 86400);
  $: hours = Math.floor(($timer - (days * 86400)) / 3600);
  $: minutes = Math.floor(($timer - (days * 86400) - (hours * 3600)) / 60);
  $: seconds = Math.floor(($timer - (days * 86400) - (hours * 3600) - (minutes * 60)))
  $: uptime = (days > 0 ? days + "d " : "") + (hours > 0 ? hours + "h " : "") + 
        (minutes > 0 ? minutes + "m " : "") + (seconds > 0 ? seconds + "s " : "")
</script>

<BGLogo url="golift">
  <Container>
    <Row>
      <h1>About Toolbarr</h1>
      <p>
        Toolbarr fixes problems with Starr apps. It comes with a five starr rating from <A href="https://toys-arr.us">Toys Arr Us</A>!
      </p>
      <Col md="6">
        <h3>Development</h3>
        <Table dark={$conf.Dark} responsive>
          <tr>
            <td style="width:180px;"><A href="https://github.com/Notifiarr/toolbarr"><Fa icon={faGithub} /> Toolbarr GitHub</A></td> 
            <td>Visit the sausage factory.</td>
          </tr>
          <tr>
            <td><A href="https://notifiarr.com/discord"><Fa fw icon={faDiscord} /> Notifiarr Discord</A></td>
            <td>For your notifications needs.</td>
          </tr>
          <tr>
            <td><A href="https://golift.io/discord"><Fa fw icon={faDiscord} /> Go Lift Discord</A></td>
            <td>Code cookin' collaborators.</td>
          </tr>
        </Table>
      </Col>
      <Col md="6">
      <h3>Attribution</h3>
      <p>
        <li>
          Created by <A href="https://golift.io">Go Lift</A> 
          for <A href="https://notifiarr.com">Notifiarr</A>.
        </li>
        <li>
          Backgrounds by <A href="https://rawpixel.com">rawpixel.com</A> 
          on <A href="https://www.freepik.com/author/rawpixel-com">Freepik</A>.
        </li>
        <li>
          Written in <A href="https://svelte.dev">Svelte</A> 
          using <A href="https://wails.io">Wails</A>.
        </li>
      </p>
    </Col>
    </Row>

    <!-- version update card at the bottom of the About page -->
    {#if Object.keys(version).length > 0}
    <Col md="6">
      <h3>App Info</h3><!-- following line shows an error but actually works. -->
      <Card color={$conf.Dark ? "secondary" : "light"} body>
        <Table dark={$conf.Dark} responsive>
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
            <Tooltip target="installButton">{release.FilePath}</Tooltip>
            <Tooltip target="folderButton">{(release.FilePath).split(/[\\/]/).slice(0,-1).join($conf.IsWindows?"\\":"/")}</Tooltip>
            <Button id="installButton" style="width:49%" outline on:click={installUpdate} size="sm" color="primary">{update.Downloaded}</Button>
            <Button id="folderButton" style="width:49%" outline on:click={openFolder} size="sm" color="info">Open Folder</Button>
            {:else if release.Outdate}
              {#if $conf.IsLinux}
              <Button block outline disabled size="sm" color="warning">Update available! v{release.Current}</Button>
              {:else}
              <Button block outline on:click={downloadUpdate} size="sm" color="warning">Download: v{release.Current} ({release.Size})</Button>
              {/if}
            {:else if update.Checked}
            <Button block color="success" disabled size="sm">Up to date! Current: v{release.Current}</Button>
            {:else}
            <Button block outline size="sm" color="info" on:click={checkUpdate}>Check for update</Button>
            {/if}
          </td></tr>
        </Table>
        {#if progress}
        <Progress striped animated color="success" value={progress*100}>{(progress*100).toFixed(0)}%</Progress>
        {:else}
        <Badge class="about-badge" color="dark">{msg}</Badge>
        {/if}
      </Card>
    </Col>
    {/if}
  </Container>
</BGLogo>

<style>
  :global(.about-badge) {
    --bs-bg-opacity: 0.5 !important;
  }
</style>
