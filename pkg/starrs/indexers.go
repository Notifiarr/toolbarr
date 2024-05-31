package starrs

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	wr "github.com/wailsapp/wails/v2/pkg/runtime"
	"golift.io/starr"
	"golift.io/starr/lidarr"
	"golift.io/starr/prowlarr"
	"golift.io/starr/radarr"
	"golift.io/starr/readarr"
	"golift.io/starr/sonarr"
)

func (s *Starrs) Indexers(config *AppConfig) (any, error) {
	s.log.Tracef("Call:Indexers(%s, %s)", config.App, config.Name)

	indexers, err := s.indexers(config)
	if err != nil {
		msg := s.log.Translate("Getting indexers: %v", err.Error())
		s.log.Wails.Error(msg)

		return "", fmt.Errorf(msg)
	}

	return indexers, nil
}

func (s *Starrs) indexers(config *AppConfig) (any, error) {
	instance, err := s.newAPIinstance(config)
	if err != nil {
		return nil, err
	}

	switch starr.App(config.App) {
	case starr.Lidarr:
		return lidarr.New(instance.Config).GetIndexersContext(s.ctx)
	case starr.Prowlarr:
		return prowlarr.New(instance.Config).GetIndexersContext(s.ctx)
	case starr.Radarr:
		return radarr.New(instance.Config).GetIndexersContext(s.ctx)
	case starr.Readarr:
		return readarr.New(instance.Config).GetIndexersContext(s.ctx)
	case starr.Sonarr:
		return sonarr.New(instance.Config).GetIndexersContext(s.ctx)
	case starr.Whisparr:
		return sonarr.New(instance.Config).GetIndexersContext(s.ctx)
	default:
		return nil, fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}

func (s *Starrs) DeleteIndexer(config *AppConfig, indexerID int64) (any, error) {
	s.log.Tracef("Call:DeleteIndexer(%s, %s, %v)", config.App, config.Name, indexerID)

	if err := s.deleteIndexer(config, indexerID); err != nil {
		msg := s.log.Translate("Deleting %s indexer: %d: %v", config.Name, indexerID, err.Error())
		s.log.Wails.Error(msg)

		return nil, fmt.Errorf(msg)
	}

	msg := s.log.Translate("Deleted %s indexer with ID %d.", config.Name, indexerID)
	s.log.Wails.Info(msg)

	return msg, nil
}

func (s *Starrs) deleteIndexer(config *AppConfig, indexerID int64) error {
	instance, err := s.newAPIinstance(config)
	if err != nil {
		return err
	}

	switch starr.App(config.App) {
	case starr.Lidarr:
		return lidarr.New(instance.Config).DeleteIndexerContext(s.ctx, indexerID)
	case starr.Prowlarr:
		return prowlarr.New(instance.Config).DeleteIndexerContext(s.ctx, indexerID)
	case starr.Radarr:
		return radarr.New(instance.Config).DeleteIndexerContext(s.ctx, indexerID)
	case starr.Readarr:
		return readarr.New(instance.Config).DeleteIndexerContext(s.ctx, indexerID)
	case starr.Sonarr:
		return sonarr.New(instance.Config).DeleteIndexerContext(s.ctx, indexerID)
	case starr.Whisparr:
		return sonarr.New(instance.Config).DeleteIndexerContext(s.ctx, indexerID)
	default:
		return fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}

func (s *Starrs) TestLidarrIndexer(config *AppConfig, indexer *lidarr.IndexerInput) (string, error) {
	s.log.Tracef("Call:TestLidarrIndexer(%s, %s, %d)", config.App, config.Name, indexer.ID)
	return s.testIndexerReply(config.Name, indexer.Name, indexer.ID, s.testIndexer(config, indexer))
}

func (s *Starrs) TestProwlarrIndexer(config *AppConfig, indexer *prowlarr.IndexerInput) (string, error) {
	s.log.Tracef("Call:TestProwlarrIndexer(%s, %s, %d)", config.App, config.Name, indexer.ID)
	return s.testIndexerReply(config.Name, indexer.Name, indexer.ID, s.testIndexer(config, indexer))
}

func (s *Starrs) TestRadarrIndexer(config *AppConfig, indexer *radarr.IndexerInput) (string, error) {
	s.log.Tracef("Call:TestRadarrIndexer(%s, %s, %d)", config.App, config.Name, indexer.ID)
	return s.testIndexerReply(config.Name, indexer.Name, indexer.ID, s.testIndexer(config, indexer))
}

func (s *Starrs) TestReadarrIndexer(config *AppConfig, indexer *readarr.IndexerInput) (string, error) {
	s.log.Tracef("Call:TestReadarrIndexer(%s, %s, %d)", config.App, config.Name, indexer.ID)
	return s.testIndexerReply(config.Name, indexer.Name, indexer.ID, s.testIndexer(config, indexer))
}

func (s *Starrs) TestSonarrIndexer(config *AppConfig, indexer *sonarr.IndexerInput) (string, error) {
	s.log.Tracef("Call:TestSonarrIndexer(%s, %s, %d)", config.App, config.Name, indexer.ID)
	return s.testIndexerReply(config.Name, indexer.Name, indexer.ID, s.testIndexer(config, indexer))
}

func (s *Starrs) TestWhisparrIndexer(config *AppConfig, indexer *sonarr.IndexerInput) (string, error) {
	s.log.Tracef("Call:TestWhisparrIndexer(%s, %s, %d)", config.App, config.Name, indexer.ID)
	return s.testIndexerReply(config.Name, indexer.Name, indexer.ID, s.testIndexer(config, indexer))
}

func (s *Starrs) testIndexer(config *AppConfig, indexer any) error {
	instance, err := s.newAPIinstance(config)
	if err != nil {
		return err
	}

	switch data := indexer.(type) {
	case *lidarr.IndexerInput:
		return lidarr.New(instance.Config).TestIndexerContext(s.ctx, data)
	case *prowlarr.IndexerInput:
		return prowlarr.New(instance.Config).TestIndexerContext(s.ctx, data)
	case *radarr.IndexerInput:
		return radarr.New(instance.Config).TestIndexerContext(s.ctx, data)
	case *readarr.IndexerInput:
		return readarr.New(instance.Config).TestIndexerContext(s.ctx, data)
	case *sonarr.IndexerInput:
		return sonarr.New(instance.Config).TestIndexerContext(s.ctx, data)
	default:
		return fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}

func (s *Starrs) testIndexerReply(
	name, indexerName string,
	indexerID int64,
	err error,
) (string, error) {
	if err == nil {
		msg := s.log.Translate("Tested %s indexer %s (%d).", name, indexerName, indexerID)
		s.log.Wails.Info(msg)

		return msg, nil
	}

	reqError := &starr.ReqError{}

	if errors.As(err, &reqError) && reqError.Msg != "" {
		err = fmt.Errorf("%s: %s", reqError.Name, reqError.Msg)
	}

	msg := s.log.Translate("Testing %s indexer: %s (%d): %s", name, indexerName, indexerID, err.Error())
	s.log.Wails.Error(msg)

	return "", fmt.Errorf(msg)
}

func (s *Starrs) UpdateLidarrIndexer(
	config *AppConfig,
	force bool,
	indexer *lidarr.IndexerInput,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateLidarrIndexer(%s, %s, %d)", config.App, config.Name, indexer.ID)
	data, err := s.updateIndexer(config, force, indexer)

	return s.updateIndexerReply(config.Name, indexer.Name, indexer.ID, data, err)
}

func (s *Starrs) UpdateProwlarrIndexer(
	config *AppConfig,
	force bool,
	indexer *prowlarr.IndexerInput,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateProwlarrIndexer(%s, %s, %d)", config.App, config.Name, indexer.ID)
	data, err := s.updateIndexer(config, force, indexer)

	return s.updateIndexerReply(config.Name, indexer.Name, indexer.ID, data, err)
}

func (s *Starrs) UpdateRadarrIndexer(
	config *AppConfig,
	force bool,
	indexer *radarr.IndexerInput,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateRadarrIndexer(%s, %s, %d)", config.App, config.Name, indexer.ID)
	data, err := s.updateIndexer(config, force, indexer)

	return s.updateIndexerReply(config.Name, indexer.Name, indexer.ID, data, err)
}

func (s *Starrs) UpdateReadarrIndexer(
	config *AppConfig,
	force bool,
	indexer *readarr.IndexerInput,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateReadarrIndexer(%s, %s, %d)", config.App, config.Name, indexer.ID)
	data, err := s.updateIndexer(config, force, indexer)

	return s.updateIndexerReply(config.Name, indexer.Name, indexer.ID, data, err)
}

func (s *Starrs) UpdateSonarrIndexer(
	config *AppConfig,
	force bool,
	indexer *sonarr.IndexerInput,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateSonarrIndexer(%s, %s, %d)", config.App, config.Name, indexer.ID)
	data, err := s.updateIndexer(config, force, indexer)

	return s.updateIndexerReply(config.Name, indexer.Name, indexer.ID, data, err)
}

func (s *Starrs) UpdateWhisparrIndexer(
	config *AppConfig,
	force bool,
	indexer *sonarr.IndexerInput,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateWhisparrIndexer(%s, %s, %d)", config.App, config.Name, indexer.ID)
	data, err := s.updateIndexer(config, force, indexer)

	return s.updateIndexerReply(config.Name, indexer.Name, indexer.ID, data, err)
}

func (s *Starrs) updateIndexer(config *AppConfig, force bool, indexer any) (any, error) {
	instance, err := s.newAPIinstance(config)
	if err != nil {
		return nil, err
	}

	switch data := indexer.(type) {
	case *lidarr.IndexerInput:
		return lidarr.New(instance.Config).UpdateIndexerContext(s.ctx, data, force)
	case *prowlarr.IndexerInput:
		return prowlarr.New(instance.Config).UpdateIndexerContext(s.ctx, data, force)
	case *radarr.IndexerInput:
		return radarr.New(instance.Config).UpdateIndexerContext(s.ctx, data, force)
	case *readarr.IndexerInput:
		return readarr.New(instance.Config).UpdateIndexerContext(s.ctx, data, force)
	case *sonarr.IndexerInput:
		return sonarr.New(instance.Config).UpdateIndexerContext(s.ctx, data, force)
	default:
		return nil, fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}

func (s *Starrs) updateIndexerReply(
	name, indexerName string,
	indexerID int64,
	data any,
	err error,
) (*DataReply, error) {
	if err == nil {
		msg := s.log.Translate("Updated %s indexer %s (%d).", name, indexerName, indexerID)
		s.log.Wails.Info(msg)

		return &DataReply{Msg: msg, Data: data}, nil
	}

	reqError := &starr.ReqError{}

	if errors.As(err, &reqError) && reqError.Msg != "" {
		err = fmt.Errorf("%s: %s", reqError.Name, reqError.Msg)
	}

	msg := s.log.Translate("Updating %s indexer: %s (%d): %s", name, indexerName, indexerID, err.Error())
	s.log.Wails.Error(msg)

	return nil, fmt.Errorf(msg)
}

func (s *Starrs) ExportLidarrIndexer(config *AppConfig, selected Selected) (string, error) {
	s.log.Tracef("Call:ExportLidarrIndexer(%v)", selected)

	instance, err := s.newAPIinstance(config)
	if err != nil {
		wr.LogError(s.ctx, err.Error())
		return "", err
	}

	indexers, err := lidarr.New(instance.Config).GetIndexersContext(s.ctx)

	return s.exportIndexers(config, filterIndexers(indexers, selected), selected.Count(), err)
}

func (s *Starrs) ExportProwlarrIndexer(config *AppConfig, selected Selected) (string, error) {
	s.log.Tracef("Call:ExportProwlarrIndexer(%v)", selected)

	instance, err := s.newAPIinstance(config)
	if err != nil {
		wr.LogError(s.ctx, err.Error())
		return "", err
	}

	indexers, err := prowlarr.New(instance.Config).GetIndexersContext(s.ctx)

	return s.exportIndexers(config, filterIndexers(indexers, selected), selected.Count(), err)
}

func (s *Starrs) ExportRadarrIndexer(config *AppConfig, selected Selected) (string, error) {
	s.log.Tracef("Call:ExportRadarrIndexer(%v)", selected)

	instance, err := s.newAPIinstance(config)
	if err != nil {
		wr.LogError(s.ctx, err.Error())
		return "", err
	}

	indexers, err := radarr.New(instance.Config).GetIndexersContext(s.ctx)

	return s.exportIndexers(config, filterIndexers(indexers, selected), selected.Count(), err)
}

func (s *Starrs) ExportReadarrIndexer(config *AppConfig, selected Selected) (string, error) {
	s.log.Tracef("Call:ExportReadarrIndexer(%v)", selected)

	instance, err := s.newAPIinstance(config)
	if err != nil {
		wr.LogError(s.ctx, err.Error())
		return "", err
	}

	indexers, err := readarr.New(instance.Config).GetIndexersContext(s.ctx)

	return s.exportIndexers(config, filterIndexers(indexers, selected), selected.Count(), err)
}

func (s *Starrs) ExportSonarrIndexer(config *AppConfig, selected Selected) (string, error) {
	s.log.Tracef("Call:ExportSonarrIndexer(%v)", selected)

	instance, err := s.newAPIinstance(config)
	if err != nil {
		wr.LogError(s.ctx, err.Error())
		return "", err
	}

	indexers, err := sonarr.New(instance.Config).GetIndexersContext(s.ctx)

	return s.exportIndexers(config, filterIndexers(indexers, selected), selected.Count(), err)
}

func (s *Starrs) ExportWhisparrIndexer(config *AppConfig, selected Selected) (string, error) {
	s.log.Tracef("Call:ExportWhisparrIndexer(%v)", selected)

	instance, err := s.newAPIinstance(config)
	if err != nil {
		wr.LogError(s.ctx, err.Error())
		return "", err
	}

	indexers, err := sonarr.New(instance.Config).GetIndexersContext(s.ctx)

	return s.exportIndexers(config, filterIndexers(indexers, selected), selected.Count(), err)
}

func filterIndexers[N any](indexers []*N, selected Selected) []*N {
	id := func(i any) int64 {
		switch indexer := i.(type) {
		case *lidarr.IndexerOutput:
			return indexer.ID
		case *prowlarr.IndexerOutput:
			return indexer.ID
		case *radarr.IndexerOutput:
			return indexer.ID
		case *readarr.IndexerOutput:
			return indexer.ID
		case *sonarr.IndexerOutput:
			return indexer.ID
		default:
			return 0
		}
	}

	i := 0
	for _, v := range indexers {
		if selected[id(v)] {
			indexers[i] = v
			i++
		}
	}

	for j := i; j < len(indexers); j++ {
		indexers[j] = nil
	}

	return indexers[:i]
}

func (s *Starrs) exportIndexers(config *AppConfig, indexers any, count int, err error) (string, error) {
	if err != nil {
		wr.LogError(s.ctx, err.Error())
		return "", fmt.Errorf(s.log.Translate("Getting Indexers from %s: %v", config.Name, err))
	}

	homedir, _ := homedir.Dir()

	filePath, err := wr.SaveFileDialog(s.ctx, wr.SaveDialogOptions{
		DefaultDirectory:           filepath.Join(homedir, "Documents"),
		DefaultFilename:            fmt.Sprintf("%d%sIndexers.json", count, config.App),
		Title:                      s.log.Translate("Save %d %s Indexers", count, config.App),
		ShowHiddenFiles:            false,
		CanCreateDirectories:       true,
		TreatPackagesAsDirectories: false,
	})
	if err != nil {
		wr.LogError(s.ctx, err.Error())
		return "", fmt.Errorf(s.log.Translate("Opening file browser: %v", err))
	}

	if filePath == "" {
		return "", nil
	}

	fileOpen, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0o640)
	if err != nil {
		wr.LogError(s.ctx, err.Error())
		return "", fmt.Errorf(s.log.Translate("Opening file: %v", err))
	}

	err = json.NewEncoder(fileOpen).Encode(indexers)
	if err != nil {
		wr.LogError(s.ctx, err.Error())
		return "", fmt.Errorf(s.log.Translate("Encoding and writing file: %v", err))
	}

	return s.log.Translate("Saved %d indexers to %s", count, filePath), nil
}

func (s *Starrs) ImportLidarrIndexer(config *AppConfig) (string, error) {
	s.log.Tracef("Call:ImportLidarrIndexer()")
	return "not working yet", nil
}
func (s *Starrs) ImportProwlarrIndexer(config *AppConfig) (string, error) {
	s.log.Tracef("Call:ImportProwlarrIndexer()")
	return "not working yet", nil
}
func (s *Starrs) ImportRadarrIndexer(config *AppConfig) (string, error) {
	s.log.Tracef("Call:ImportRadarrIndexer()")
	return "not working yet", nil
}
func (s *Starrs) ImportReadarrIndexer(config *AppConfig) (string, error) {
	s.log.Tracef("Call:ImportReadarrIndexer()")
	return "not working yet", nil
}
func (s *Starrs) ImportSonarrIndexer(config *AppConfig) (string, error) {
	s.log.Tracef("Call:ImportSonarrIndexer()")
	return "not working yet", nil
}
func (s *Starrs) ImportWhisparrIndexer(config *AppConfig) (string, error) {
	s.log.Tracef("Call:ImportWhisparrIndexer()")
	return "not working yet", nil
}
