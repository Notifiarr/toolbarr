<script>
  export let app
  export let size
  import logo from "../assets/images/logo.png"
  import toolboxLogo from "../assets/images/Toolbox.png"
  import settingsLogo from "../assets/images/settings.png"
  import lidarrLogo from "../assets/images/Lidarr.png"
  import prowlarrLogo from "../assets/images/Prowlarr.png"
  import radarrLogo from "../assets/images/Radarr.png"
  import readarrLogo from "../assets/images/Readarr.png"
  import sonarrLogo from "../assets/images/Sonarr.png"
  import whisparrLogo from "../assets/images/Whisparr.png"
  import Fa from "svelte-fa"
  import { faBookBible, faLink } from "@fortawesome/free-solid-svg-icons"
  import { _ } from "./locale"

</script>

{#if app == "Lidarr"}
  <img alt="{app} {$_("words.Logo")}" height={size} src={lidarrLogo} />
{:else if app == "Prowlarr"}
  <img alt="{app} {$_("words.Logo")}" height={size} src={prowlarrLogo} />
{:else if app == "Radarr"}
  <img alt="{app} {$_("words.Logo")}" height={size} src={radarrLogo} />
{:else if app == "Readarr"}
  <img alt="{app} {$_("words.Logo")}" height={size} src={readarrLogo} />
{:else if app == "Sonarr"}
  <img alt="{app} {$_("Lwords.ogo")}" height={size} src={sonarrLogo} />
{:else if app == "Whisparr"}
  <img alt="{app} {$_("words.Logo")}" height={size} src={whisparrLogo} />
{:else if app == "Toolbox"}
  <img alt="{app} {$_("words.Logo")}" height={size} src={toolboxLogo} />
{:else if app == "Settings"}
  <img alt="{app} {$_("words.Logo")}" height={size} src={settingsLogo} />
{:else if app == "Links"}
  <Fa primaryColor="dodgerblue" icon="{faLink}" />
{:else if app == "About"}
  <Fa primaryColor="mediumpurple" icon="{faBookBible}" />
{:else}
  <img alt="{app} {$_("words.Logo")}" height={size} src={logo} />
{/if}
