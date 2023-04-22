<script lang="ts">
  export let list: any
  export let idx: number
  export let starrApp: string
  export let qualityProfiles: any

  import { Table, Popover, Badge, Icon } from "sveltestrap"
  import { _, date } from "../../libs/Translate.svelte"
  import { conf } from "../../libs/config"


  function qp(id) {
    let name = id
    qualityProfiles.forEach((qp) => {
      if (qp.id == id) { 
        name = qp.name
        return name
      }
    })
    return name
  }
</script>

{#if starrApp == "Lidarr"}<!-- LIDARR PAGE-->
<td id="record{idx}" class="pop nowrap">{list.artist.artistName}
  <div class="popover-content {$conf.Dark?"dark-mode":""}">
    <Popover container="inline" trigger="hover" placement="top" target="record{idx}">
      <span slot="title"><Icon name={list.artist.monitored?"bookmark-fill":"bookmark"}/> {list.artist.artistName}</span>
      <Table striped size="sm">
        <tbody>
          <tr><th class="nowrap">{$_("words.Quality")}</th><td>{list.quality.quality.name}</td></tr>
          <tr><th class="nowrap">{$_("words.Source")}</th><td>{list.sourceTitle}</td></tr>
          <tr><th class="nowrap">{$_("words.Protocol")}</th><td>{list.protocol}</td></tr>
          <tr><th class="nowrap">{$_("words.Indexer")}</th><td>{list.indexer}</td></tr>
          {#if qualityProfiles !== undefined}
            <tr><th class="nowrap">{$_("instances.qualityProfileIdTitle")}</th><td>{qp(list.artist.qualityProfileId)}</td></tr>
          {/if}
          <tr><th class="nowrap">{$_("words.Message")}</th><td>{list.message}</td></tr>
        </tbody>
      </Table>
    </Popover>
  </div>
</td>
<td>{list.sourceTitle}</td>
<td class="nowrap">{list.quality.quality.name}</td>

{:else if starrApp == "Radarr"}<!-- RADARR PAGE-->
<td id="record{idx}" class="pop nowrap">{list.movie.title}</td>
<td>{list.sourceTitle}</td>
<td class="nowrap">{list.quality.quality.name}
  <div class="popover-content {$conf.Dark?"dark-mode":""}">
    <Popover container="inline" trigger="hover" placement="top" target="record{idx}">
      <span slot="title"><Icon name={list.movie.monitored?"bookmark-fill":"bookmark"}/>
        {list.movie.title} ({list.movie.year})
      </span>
      <Table striped size="sm">
        <tbody>
          <tr><th class="nowrap">{$_("words.Source")}</th><td>{list.sourceTitle}</td></tr>
          <tr><th class="nowrap">{$_("words.Protocol")}</th><td>{list.protocol}</td></tr>
          <tr><th class="nowrap">{$_("words.Indexer")}</th><td>{list.indexer}</td></tr>
          {#if qualityProfiles !== undefined}
            <tr><th class="nowrap">{$_("instances.qualityProfileIdTitle")}</th><td>{qp(list.movie.qualityProfileId)}</td></tr>
          {/if}
          <tr>
            <th class="nowrap">{$_("words.Languages")}</th>
            <td>{#each list.languages as lang}<Badge>{lang.name}</Badge> {/each}</td>
          </tr>
          <tr><th class="nowrap">{$_("words.Message")}</th><td>{list.message}</td></tr>
        </tbody>
      </Table>
    </Popover>
  </div>
</td>

{:else if starrApp == "Readarr"}<!-- READARR PAGE-->
<td id="record{idx}" class="pop nowrap">{list.author.authorName}
  <div class="popover-content {$conf.Dark?"dark-mode":""}">
    <Popover container="inline" trigger="hover" placement="top" target="record{idx}">
      <span slot="title"><Icon name={list.author.monitored?"bookmark-fill":"bookmark"}/> {list.author.authorName}</span>
      <Table striped size="sm">
        <tbody>
          <tr><th class="nowrap">{$_("words.Quality")}</th><td>{list.quality.quality.name}</td></tr>
          <tr><th class="nowrap">{$_("words.Source")}</th><td>{list.sourceTitle}</td></tr>
          <tr><th class="nowrap">{$_("words.Protocol")}</th><td>{list.protocol}</td></tr>
          <tr><th class="nowrap">{$_("words.Indexer")}</th><td>{list.indexer}</td></tr>
          {#if qualityProfiles !== undefined}
            <tr><th class="nowrap">{$_("instances.qualityProfileIdTitle")}</th><td>{qp(list.author.qualityProfileId)}</td></tr>
          {/if}
          <tr><th class="nowrap">{$_("words.Message")}</th><td>{list.message}</td></tr>
        </tbody>
      </Table>
    </Popover>
  </div>
</td>
<td>{list.sourceTitle}</td>
<td class="nowrap">{list.quality.quality.name}</td>

{:else if starrApp == "Sonarr"}<!-- SONARR PAGE-->
<td id="record{idx}" class="pop nowrap"> {list.series.title}
  <div class="popover-content {$conf.Dark?"dark-mode":""}">
    <Popover container="inline" trigger="hover" placement="top" target="record{idx}">
      <span slot="title"><Icon name={list.series.monitored?"bookmark-fill":"bookmark"}/>
        {list.series.title} ({list.series.year})
      </span>
      <Table striped size="sm">
        <tbody>
          <tr><th class="nowrap">{$_("words.Source")}</th><td>{list.sourceTitle}</td></tr>
          <tr><th class="nowrap">{$_("words.Protocol")}</th><td>{list.protocol}</td></tr>
          <tr><th class="nowrap">{$_("words.Indexer")}</th><td>{list.indexer}</td></tr>
          {#if qualityProfiles !== undefined}
            <tr><th class="nowrap">{$_("instances.qualityProfileIdTitle")}</th><td>{qp(list.series.qualityProfileId)}</td></tr>
          {/if}
          <tr>
            <th class="nowrap">{$_("words.Languages")}</th>
            <td>{#each list.languages as lang}<Badge>{lang.name}</Badge> {/each}</td>
          </tr>
          <tr><th class="nowrap">{$_("words.Message")}</th><td>{list.message}</td></tr>
        </tbody>
      </Table>
    </Popover>
  </div>
</td>
<td>{list.sourceTitle}</td>
<td class="nowrap">{list.quality.quality.name}</td>
{/if}

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
  
  .popover-content.dark-mode :global(.popover) {
    background-color: #527d6e;
  }
  .popover-content.dark-mode :global(.popover-header) {
    background-color: #2c6757;
  }

  .popover-content :global(.popover) {
    background-color: #cdf6dd;
    max-width: 700px
  }
  .popover-content :global(.popover-header) {
    background-color: rgb(142, 214, 178);
  }
</style>