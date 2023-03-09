<script>
    export let app
    export let size
    import logo from "./assets/images/logo.png"
    import toolboxLogo from "./assets/images/Toolbox.png"
    import settingsLogo from "./assets/images/settings.png"
    import lidarrLogo from "./assets/images/Lidarr.png"
    import prowlarrLogo from "./assets/images/Prowlarr.png"
    import radarrLogo from "./assets/images/Radarr.png"
    import readarrLogo from "./assets/images/Readarr.png"
    import sonarrLogo from "./assets/images/Sonarr.png"
    import whisparrLogo from "./assets/images/Whisparr.png"

    let showLogo = logo
    $: if (app == "Lidarr") {
        showLogo = lidarrLogo
    } else if (app == "Prowlarr") {
        showLogo = prowlarrLogo
    } else if (app == "Radarr") {
        showLogo = radarrLogo
    } else if (app == "Readarr") {
        showLogo = readarrLogo
    } else if (app == "Sonarr") {
        showLogo = sonarrLogo
    } else if (app == "Whisparr") {
        showLogo = whisparrLogo
    } else if (app == "Toolbox") {
        showLogo = toolboxLogo
    } else if (app == "Settings") {
        showLogo = settingsLogo
    }
</script>
<img alt="{app} Logo" height={size} src={showLogo} />