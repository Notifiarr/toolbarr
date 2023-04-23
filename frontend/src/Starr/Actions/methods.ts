import {
  DeleteBlockList,
  DeleteDownloader,
  DeleteExclusion,
  DeleteIndexer,
  DeleteQualityProfile,
  DeleteImportList,
  UpdateLidarrDownloadClient,
  UpdateLidarrImportList,
  UpdateLidarrIndexer,
  UpdateLidarrQualityProfile,
  UpdateProwlarrDownloadClient,
  UpdateProwlarrIndexer,
  UpdateRadarrDownloadClient,
  UpdateRadarrImportList,
  UpdateRadarrIndexer,
  UpdateRadarrQualityProfile,
  UpdateReadarrDownloadClient,
  UpdateReadarrImportList,
  UpdateReadarrIndexer,
  UpdateReadarrQualityProfile,
  UpdateSonarrDownloadClient,
  UpdateSonarrImportList,
  UpdateSonarrIndexer,
  UpdateSonarrQualityProfile,
  UpdateWhisparrDownloadClient,
  UpdateWhisparrImportList,
  UpdateWhisparrIndexer,
  UpdateWhisparrQualityProfile,
  UpdateLidarrExclusion,
  UpdateRadarrExclusion,
  UpdateReadarrExclusion,
  UpdateSonarrExclusion,
  UpdateWhisparrExclusion,
} from "../../../wailsjs/go/starrs/Starrs"

export function fixFieldValues(info) {
  info.forEach((item, idx) => {
    if (!item.fields) return
    item.fields.forEach((field, itemIdx) => {
      if (field.value === undefined || field.value === null) info[idx].fields[itemIdx].value = ""
    })
  })

  return JSON.stringify(info)
}

export const remove = {
  "BlockLists": {
    "Lidarr":   DeleteBlockList,
    "Prowlarr": DeleteBlockList,
    "Radarr":   DeleteBlockList,
    "Readarr":  DeleteBlockList,
    "Sonarr":   DeleteBlockList,
    "Whisparr": DeleteBlockList,
  },
  "DownloadClients": {
    "Lidarr":   DeleteDownloader,
    "Prowlarr": DeleteDownloader,
    "Radarr":   DeleteDownloader,
    "Readarr":  DeleteDownloader,
    "Sonarr":   DeleteDownloader,
    "Whisparr": DeleteDownloader,
  },
  "Indexers": {
    "Lidarr":   DeleteIndexer,
    "Prowlarr": DeleteIndexer,
    "Radarr":   DeleteIndexer,
    "Readarr":  DeleteIndexer,
    "Sonarr":   DeleteIndexer,
    "Whisparr": DeleteIndexer,
  },
  "QualityProfiles": {
    "Lidarr":   DeleteQualityProfile,
    "Prowlarr": DeleteQualityProfile,
    "Radarr":   DeleteQualityProfile,
    "Readarr":  DeleteQualityProfile,
    "Sonarr":   DeleteQualityProfile,
    "Whisparr": DeleteQualityProfile,
  },
  "ImportLists": {
    "Lidarr":   DeleteImportList,
    "Prowlarr": DeleteImportList,
    "Radarr":   DeleteImportList,
    "Readarr":  DeleteImportList,
    "Sonarr":   DeleteImportList,
    "Whisparr": DeleteImportList,
  },
  "Exclusions": {
    "Lidarr": DeleteExclusion,
    "Radarr": DeleteExclusion,
    "Readarr": DeleteExclusion,
    "Sonarr": DeleteExclusion,
    "Whisparr": DeleteExclusion,
  }
}

export const update = {
  "DownloadClients": {
    "Lidarr":   UpdateLidarrDownloadClient,
    "Prowlarr": UpdateProwlarrDownloadClient,
    "Radarr":   UpdateRadarrDownloadClient,
    "Readarr":  UpdateReadarrDownloadClient,
    "Sonarr":   UpdateSonarrDownloadClient,
    "Whisparr": UpdateWhisparrDownloadClient,
  },
  "Indexers": {
    "Lidarr":   UpdateLidarrIndexer,
    "Prowlarr": UpdateProwlarrIndexer,
    "Radarr":   UpdateRadarrIndexer,
    "Readarr":  UpdateReadarrIndexer,
    "Sonarr":   UpdateSonarrIndexer,
    "Whisparr": UpdateWhisparrIndexer,
  },
  "QualityProfiles": {
    "Lidarr": UpdateLidarrQualityProfile,
    "Radarr": UpdateRadarrQualityProfile,
    "Readarr": UpdateReadarrQualityProfile,
    "Sonarr": UpdateSonarrQualityProfile,
    "Whisparr": UpdateWhisparrQualityProfile,
  },
  "ImportLists": {
    "Lidarr": UpdateLidarrImportList,
    "Radarr": UpdateRadarrImportList,
    "Readarr": UpdateReadarrImportList,
    "Sonarr": UpdateSonarrImportList,
    "Whisparr": UpdateWhisparrImportList,
  },
  "Exclusions": {
    "Lidarr": UpdateLidarrExclusion,
    "Radarr": UpdateRadarrExclusion,
    "Readarr": UpdateReadarrExclusion,
    "Sonarr": UpdateSonarrExclusion,
    "Whisparr": UpdateWhisparrExclusion,
  },
}