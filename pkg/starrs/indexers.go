package starrs

import (
	"errors"
	"fmt"

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
