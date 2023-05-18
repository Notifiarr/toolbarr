<script lang="ts">
  export let app: string
  export let size: number | string | undefined
  export let style = ""
  
  import logo from "/src/assets/images/logo.png"
  import toolboxLogo from "/src/assets/images/Toolbox.png"
  import settingsLogo from "/src/assets/images/settings.png"
  import lidarrLogo from "/src/assets/images/Lidarr.png"
  import prowlarrLogo from "/src/assets/images/Prowlarr.png"
  import radarrLogo from "/src/assets/images/Radarr.png"
  import readarrLogo from "/src/assets/images/Readarr.png"
  import sonarrLogo from "/src/assets/images/Sonarr.png"
  import whisparrLogo from "/src/assets/images/Whisparr.png"
  import Fa from "svelte-fa"
  import { faBookBible, faLink } from "@fortawesome/free-solid-svg-icons"
  import { _ } from "/src/libs/Translate.svelte"
</script>

{#if app == "Lidarr"}
  <img alt="{app} {$_("words.Logo")}" height={size} src={lidarrLogo} {style}/>
{:else if app == "Prowlarr"}
  <img alt="{app} {$_("words.Logo")}" height={size} src={prowlarrLogo} {style}/>
{:else if app == "Radarr"}
  <img alt="{app} {$_("words.Logo")}" height={size} src={radarrLogo} {style}/>
{:else if app == "Readarr"}
  <img alt="{app} {$_("words.Logo")}" height={size} src={readarrLogo} {style}/>
{:else if app == "Sonarr"}
  <img alt="{app} {$_("Lwords.ogo")}" height={size} src={sonarrLogo} {style}/>
{:else if app == "Whisparr"}
  <img alt="{app} {$_("words.Logo")}" height={size} src={whisparrLogo} {style}/>
{:else if app == "Toolbox"}
  <img alt="{app} {$_("words.Logo")}" height={size} src={toolboxLogo} {style}/>
{:else if app == "Settings"}
  <img alt="{app} {$_("words.Logo")}" height={size} src={settingsLogo} {style}/>
{:else if app == "Links"}
  <Fa primaryColor="dodgerblue" icon="{faLink}" {style}/>
{:else if app == "About"}
  <Fa primaryColor="mediumpurple" icon="{faBookBible}" {style}/>
{:else}
  <img alt="{app} {$_("words.Logo")}" height={size} src={logo} {style}/>
{/if}
