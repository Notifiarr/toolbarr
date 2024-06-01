package starrs

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Notifiarr/toolbarr/pkg/mnd"
	"github.com/mitchellh/go-homedir"
	wr "github.com/wailsapp/wails/v2/pkg/runtime"
	"golift.io/cache"
	"golift.io/starr/lidarr"
	"golift.io/starr/prowlarr"
	"golift.io/starr/radarr"
	"golift.io/starr/readarr"
	"golift.io/starr/sonarr"
)

/* helper functions to abstract import and export code */

// lastPickedDir makes the open/save dialog always start in the last picked folder.
var lastPickedDir = getSavePath() //nolint:gochecknoglobals

// getExportInstance abstracts some logic away from each export method.
func (s *Starrs) getExportInstance(config *AppConfig, selected Selected, item string) (*instance, error) {
	s.log.Tracef("Call:Export%s%s(%v)", config.App, item, selected)

	instance, err := s.newAPIinstance(config)
	if err != nil {
		wr.LogError(s.ctx, err.Error())
		return nil, err
	}

	return instance, nil
}

// filterListItemsByID removes items from a slice that are not selected by the user.
func filterListItemsByID[N any](items []*N, selected Selected) []*N { //nolint:funlen,cyclop
	itemID := func(i any) int64 {
		switch item := i.(type) {
		case *lidarr.IndexerOutput:
			return item.ID
		case *prowlarr.IndexerOutput:
			return item.ID
		case *radarr.IndexerOutput:
			return item.ID
		case *readarr.IndexerOutput:
			return item.ID
		case *sonarr.IndexerOutput:
			return item.ID
		case *lidarr.ImportListOutput:
			return item.ID
		case *radarr.ImportListOutput:
			return item.ID
		case *readarr.ImportListOutput:
			return item.ID
		case *sonarr.ImportListOutput:
			return item.ID
		case *lidarr.DownloadClientOutput:
			return item.ID
		case *prowlarr.DownloadClientOutput:
			return item.ID
		case *radarr.DownloadClientOutput:
			return item.ID
		case *readarr.DownloadClientOutput:
			return item.ID
		case *sonarr.DownloadClientOutput:
			return item.ID
		case *lidarr.Exclusion:
			return item.ID
		case *radarr.Exclusion:
			return item.ID
		case *readarr.Exclusion:
			return item.ID
		case *sonarr.Exclusion:
			return item.ID
		default:
			return 0
		}
	}

	idx := 0

	for _, v := range items {
		if selected[itemID(v)] {
			items[idx] = v
			idx++
		}
	}

	for i := idx; i < len(items); i++ {
		items[i] = nil
	}

	return items[:idx]
}

// exportItems saves things like import lists, download clients and indexers to a json file.
func (s *Starrs) exportItems(item string, config *AppConfig, data any, count int, err error) (string, error) {
	if err != nil {
		wr.LogError(s.ctx, err.Error())
		return "", fmt.Errorf(s.log.Translate("Getting %s from %s: %v", item, config.Name, err))
	}

	filePath, err := wr.SaveFileDialog(s.ctx, wr.SaveDialogOptions{
		DefaultDirectory:           lastPickedDir,
		DefaultFilename:            fmt.Sprintf("%d%s%s.json", count, config.App, item),
		Title:                      s.log.Translate("Save %d %s %s", count, config.App, item),
		ShowHiddenFiles:            false,
		CanCreateDirectories:       true,
		TreatPackagesAsDirectories: false,
	})
	if err != nil {
		wr.LogError(s.ctx, err.Error())
		return "", fmt.Errorf(s.log.Translate("Opening file browser: %v", err))
	} else if filePath == "" {
		return "", nil
	}

	lastPickedDir = filepath.Dir(filePath)

	fileOpen, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, mnd.Mode0640)
	if err != nil {
		wr.LogError(s.ctx, err.Error())
		return "", fmt.Errorf(s.log.Translate("Opening file: %v", err))
	}
	defer fileOpen.Close()

	err = json.NewEncoder(fileOpen).Encode(data)
	if err != nil {
		wr.LogError(s.ctx, err.Error())
		return "", fmt.Errorf(s.log.Translate("Encoding and writing file: %v", err))
	}

	return s.log.Translate("Saved %d %s to %s", count, item, filePath), nil
}

// getSavePath finds an appropriate place to save exported json files.
func getSavePath() string {
	homedir, err := homedir.Dir()
	if err != nil {
		return "/"
	}

	filePath := filepath.Join(homedir, "Documents")
	if _, err := os.Stat(filePath); err != nil {
		return homedir
	}

	return filePath
}

// importItems is the initial call when a user wants to import a json file for indexers, downloaders, etc.
// This prompts the user for a file, decodes the contents, and returns them to the front-end where
// the user can pick and choose the things they want to import.
func importItems[N any](s *Starrs, item string, config *AppConfig, input []N) (map[string]any, error) { //nolint:varnamelen,lll
	s.log.Tracef("Call:Import%s%s()", config.App, item)

	filePath, err := wr.OpenFileDialog(s.ctx, wr.OpenDialogOptions{
		DefaultDirectory: lastPickedDir,
		DefaultFilename:  fmt.Sprintf("%s%s.json", config.App, item),
		Title:            s.log.Translate("Select Export File"),
		Filters:          []wr.FileFilter{{DisplayName: "JSON (*.json)", Pattern: "*.json"}},
	})
	if err != nil {
		wr.LogError(s.ctx, err.Error())
		return nil, fmt.Errorf(s.log.Translate("Opening file browser: %v", err))
	} else if filePath == "" {
		return map[string]any{"msg": ""}, nil
	}

	// Update this so next time we pick from the same location.
	lastPickedDir = filepath.Dir(filePath)

	fileOpen, err := os.Open(filePath)
	if err != nil {
		wr.LogError(s.ctx, err.Error())
		return nil, fmt.Errorf(s.log.Translate("Opening file: %v", err))
	}
	defer fileOpen.Close()

	if err := json.NewDecoder(fileOpen).Decode(&input); err != nil {
		wr.LogError(s.ctx, err.Error())
		return nil, fmt.Errorf(s.log.Translate("Decoding input file failed: %v", err))
	}

	// Save the data we read from the file into a memory cache.
	// So after the user chooses which items to import, we don't have to re-read the file.
	s.cache.Save(cacheKey(filePath, item), input, cache.Options{})

	return map[string]any{
		"item": item,
		"file": filePath,
		"data": input,
		"msg": fmt.Sprintf("Importing is not working yet, but there are %d %s in your backup file: %s",
			len(input), item, filePath),
	}, nil
}

// cacheKey allows us to make a cache key for stored file contents.
func cacheKey(filePath string, item string) string {
	return fmt.Sprint(filePath, item)
}
