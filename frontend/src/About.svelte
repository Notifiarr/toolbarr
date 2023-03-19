<script type="ts">
  import Fa from "svelte-fa"
  import { faGithub, faDiscord } from "@fortawesome/free-brands-svg-icons"
  import { faGear } from "@fortawesome/free-solid-svg-icons"
  import { EventsOn, EventsOff } from "../wailsjs/runtime"
  import { CheckUpdate, DownloadUpdate, LaunchInstaller, OpenFolder } from "../wailsjs/go/app/App"
  import { Container, Row, Table, Col, Card, Tooltip, Button, Progress, Badge } from   "sveltestrap"
  import BGLogo from "./libs/BackgroundLogo.svelte"
  import A from "./libs/Link.svelte"
  import { app, conf } from "./libs/config.js"
  import { toast, onOnce, onInterval } from "./libs/funcs"
  import { _ } from "./libs/locale"

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
  let msg = $_("aboutPage.justWaitin")

  function checkUpdate(e) {
    e.preventDefault()
    update.Downloading = $_("aboutPage.Checkingforupdate")

    CheckUpdate().then(result => {
      release = result
      update.Checked = true
      update.Downloading = ""
      msg = $_("aboutPage.Thatwascool")
      if (release.Outdate) msg = $app.IsLinux ?
        $_("aboutPage.Useyourpackagemanager") :
        ($_("aboutPage.UpdateAvailable")+ " " + $_("aboutPage.Clickthebuttontouseit"))
    }, (error) => {
      update.Failed = $_("aboutPage.Errorchecking")
      toast("error", error)
      update.Downloading = ""
    })
  }

  function installUpdate(e) {
    e.preventDefault()
    update.Downloading = $_("aboutPage.Launchinginstaller")
    LaunchInstaller(release.FilePath).then(result => {
      update.Downloading = result
    })
  }

  function openFolder(e) {
    e.preventDefault()
    OpenFolder(release.FilePath).then(msg => (toast("info", msg)))
  }

  function downloadUpdate(e) {
    e.preventDefault()
    update.Downloading = $_("aboutPage.Downloadingtheupdate")

    DownloadUpdate().then(data => {
      update.Checked = true
      release.FilePath = data.Path
      toast("info", $_("words.Downloading")+": "+data.Path)
      EventsOn("downloadProgress", (data) => (progress = data))
      EventsOn("downloadFinished", () => {
        EventsOff("downloadProgress", "downloadFinished")
        update.Downloaded = "Open "+($app.IsMac ? "DMG" : $_("words.Installer"))
        update.Downloading = ""
        msg = ($app.IsMac ? $_("aboutPage.Diskimagedownloaded") : $_("aboutPage.Installerdownloaded")) +
          " " + $_("aboutPage.Clickabuttontouseit")
        onOnce(() => (progress = 0.0), 0.5)
      })
    }, (error) => {
      update.Failed = $_("aboutPage.Errordownloading")
      toast("error", error)
      update.Downloading = ""
    })
  }

  $: timer = (new Date()).getTime()/1000-$app.StartTime
  onInterval(() => { timer = (new Date()).getTime()/1000-$app.StartTime }, 1)
  /* All of this is to create a "running" timer. */
  $: days = Math.floor(timer / 86400);
  $: hours = Math.floor((timer - (days * 86400)) / 3600);
  $: minutes = Math.floor((timer - (days * 86400) - (hours * 3600)) / 60);
  $: seconds = Math.floor((timer - (days * 86400) - (hours * 3600) - (minutes * 60)))
  $: uptime = (days > 0 ? days + "d " : "") + (hours > 0 ? hours + "h " : "") + 
        (minutes > 0 ? minutes + "m " : "") + (seconds > 0 ? seconds + "s " : "")
</script>

<BGLogo url="golift">
  <Container>
    <Row>
      <h1>{$_("words.About")} {$app.Title}</h1>
      <p>
        {@html $_("aboutPage.toolbarDescription")} <A href="https://toys-arr.us">Toys Arr Us</A>
      </p>
      <Col md="6">
        <h3>{$_("words.Development")} </h3>
        <Table dark={$conf.Dark} responsive>
          <tr>
            <td style="width:180px;"><A href="https://github.com/Notifiarr/toolbarr"><Fa icon={faGithub} /> {$app.Title} GitHub</A></td> 
            <td>{$_("aboutPage.Visitthesausagefactory")}</td>
          </tr>
          <tr>
            <td><A href="https://notifiarr.com/discord"><Fa fw icon={faDiscord} /> Notifiarr Discord</A></td>
            <td>{$_("aboutPage.Foryournotificationsneeds")}</td>
          </tr>
          <tr>
            <td><A href="https://golift.io/discord"><Fa fw icon={faDiscord} /> Go Lift Discord</A></td>
            <td>{$_("aboutPage.Codecookincollaborators")}</td>
          </tr>
        </Table>
      </Col>
      <Col md="6">
      <h3>{$_("words.Attribution")}</h3>
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
    {#if Object.keys(app).length > 0}
    <Col md="6">
      <h3>{$_("aboutPage.AppInfo")}</h3><!-- following line shows an error but actually works. -->
      <Card color={$conf.Dark ? "secondary" : "light"} body>
        <Table dark={$conf.Dark} responsive>
          <tr><td>{$_("words.Version")}</td><td>v{$app.Version}-{$app.Revision} ({$app.GoVersion})</td></tr>
          <tr><td>{$_("words.Branch")}</td><td>{$app.Branch}</td></tr>
          <tr><td>{$_("words.Created")}</td><td>{$app.BuildDate} {$_("words.by")} {$app.BuildUser}</td></tr>
          <tr>
            <td>{$_("words.Running")}</td>
            <Tooltip target="runningTime" placement="top">{$_("aboutPage.RunningSince")}:<br>{$app.Started}</Tooltip>
            <td id="runningTime">{uptime}</td>
          </tr>
          <tr><td colspan="2">
            {#if update.Failed}
            <Button block disabled size="sm" color="danger">{update.Failed}</Button>
            {:else if update.Downloading}
            <Button block disabled size="sm" color="danger">{update.Downloading} <Fa spin primaryColor="yellow" icon={faGear} /></Button>
            {:else if update.Downloaded}
            <Tooltip target="installButton">{release.FilePath}</Tooltip>
            <Tooltip target="folderButton">{(release.FilePath).split(/[\\/]/).slice(0,-1).join($app.IsWindows?"\\":"/")}</Tooltip>
            <Button id="installButton" style="width:49%" outline on:click={installUpdate} size="sm" color="primary">{update.Downloaded}</Button>
            <Button id="folderButton" style="width:49%" outline on:click={openFolder} size="sm" color="info">{$_("aboutPage.OpenFolder")}</Button>
            {:else if release.Outdate}
              {#if $app.IsLinux}
              <Button block outline disabled size="sm" color="warning">{$_("aboutPage.UpdateAvailable")} v{release.Current}</Button>
              {:else}
              <Button block outline on:click={downloadUpdate} size="sm" color="warning">{$_("words.Download")}: v{release.Current} ({release.Size})</Button>
              {/if}
            {:else if update.Checked}
            <Button block color="success" disabled size="sm">{$_("aboutPage.Up-to-date")} {$_("words.Current")}: v{release.Current}</Button>
            {:else}
            <Button block outline size="sm" color="info" on:click={checkUpdate}>{$_("aboutPage.Checkforupdate")}</Button>
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
