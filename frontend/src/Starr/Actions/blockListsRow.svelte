<script lang="ts">
  export let list: any
  export let idx: number
  export let starrApp: string
  export let qualityProfiles: any
  export let metadataProfiles: any

  import { Table, Popover, Badge, Icon, Tooltip } from "sveltestrap"
  import { _, date } from "../../libs/Translate.svelte"
  import { conf } from "../../libs/config"

  function poster(data): string {
    let url = "."
    data.forEach(item => {
      if (item.coverType == "poster") url = item.url?item.url:item.remoteUrl
    })
    return url
  }

  function getName(id, list) {
    let name = id
    list.forEach((item) => {
      if (item.id == id) {
        name = item.name
        return name
      }
    })
    return name
  }

  // max ellipses the middle of a string to a max length.
  function max(max: number, str: string): string {
    if (str.length < max+4) return str
    return str.substring(0, max/2) + `<i class="text-warning">...</i>` +
      str.substring(str.length-max/2)
  }

  let item: any
  let name: string
  if (starrApp == "Lidarr") {
    item = list.artist
    name = list.artist.artistName
  } else if (starrApp == "Radarr") {
    item = list.movie
    name = list.movie.title
  } else if (starrApp == "Readarr") {
    item = list.author
    name = list.author.authorName
  } else if (starrApp == "Sonarr") {
    item = list.series
    name = list.series.title
  }

  let id = `blRecord${starrApp}${idx}`
  let width: number
  $: maxTDlen = width < 1000 ? 30 : width < 1400 ? 50 : 80
</script>

<svelte:window bind:innerWidth={width}/>

<!-- This is 1 table cell. It contains a popover that containers another table. -->
<td class="pop nowrap"><span {id}>{name}
  <div class="popover-content {$conf.Dark?"dark-mode":""}">
    <Popover container="inline" trigger="hover" placement="top" target={id}>
      <div slot="title"><Icon name={item.monitored?"bookmark-fill":"bookmark"}/> {name}</div>
      <div style="float:left;width:103px;"><img width="100px" alt="poster" src={poster(item.images)}/><br><Badge>{item.status}</Badge></div>
      <div style="float:right;width:calc(100% - 103px)">
      <Table striped size="sm" class="m-0">
        <tbody>
          <tr><th class="nowrap">{$_("words.Quality")}</th><td>{list.quality.quality.name}</td></tr>
          <tr><th class="nowrap">{$_("words.Source")}</th><td style="max-width:700px" class="break">{list.sourceTitle}</td></tr>
          <tr><th class="nowrap">{$_("words.Protocol")}</th><td>{list.protocol}</td></tr>
          <tr><th class="nowrap">{$_("words.Indexer")}</th><td>{list.indexer}</td></tr>
          {#if qualityProfiles !== undefined}<!-- Lidarr / Readarr -->
            <tr>
              <th class="nowrap">{$_("instances.qualityProfileIdTitle")}</th>
              <td>{getName(item.qualityProfileId, qualityProfiles)}</td>
            </tr>
          {/if}
          {#if metadataProfiles !== undefined}<!-- Lidarr / Readarr -->
            <tr>
              <th class="nowrap">{$_("instances.metadataProfileIdTitle")}</th>
              <td>{getName(item.metadataProfileId, qualityProfiles)}</td>
            </tr>
          {/if}
          {#if list.languages}<!-- Sonarr / Radarr -->
            <tr>
              <th class="nowrap">{$_("words.Languages")}</th>
              <td>{#each list.languages as lang}<Badge>{lang.name}</Badge> {/each}</td>
            </tr>
          {/if}
          <tr><th class="nowrap">{$_("words.Message")}</th><td style="max-width:500px">{list.message}</td></tr>
        </tbody>
      </Table>
      </div>
    </Popover>
  </div>
</span></td>

<!-- This one disappears on smaller widths. -->
<td class="nowrap d-none d-md-table-cell" id="sourceTitle{id}">
  {@html max(maxTDlen, list.sourceTitle)}
  {#if list.sourceTitle.length > maxTDlen+3}
    <Tooltip target="sourceTitle{id}">{list.sourceTitle}</Tooltip>
  {/if}
</td>
<td class="nowrap">{list.quality.quality.name}</td>
<td class="nowrap">{date(list.date)}</td>

<style>
  .nowrap {
    white-space: nowrap;
    width: max-content;
  }

  .pop {
    cursor: context-menu;
    text-decoration-style: dashed;
    text-decoration-line: underline;
    text-decoration-color: rgb(113, 156, 247);
  }

  .break {
    word-break: break-all;
  }

  .popover-content.dark-mode :global(.popover) {
    background-color: #527d6e;
  }
  .popover-content.dark-mode :global(.popover-header) {
    background-color: #2c6757;
  }

  .popover-content :global(.popover-body) {
    padding: 5px;
  }

  .popover-content :global(.popover) {
    background-color: #cdf6dd;
    max-width: calc(100vw - 15px);
    width: max-content;
    min-width: 465px;
    margin: 5px !important;
  }

  .popover-content :global(.popover-header) {
    background-color: rgb(142, 214, 178);
  }
</style>
