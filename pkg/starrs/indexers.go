package starrs

import (
	"errors"
	"fmt"
	"time"

	"golift.io/starr"
	"golift.io/starr/lidarr"
	"golift.io/starr/prowlarr"
	"golift.io/starr/radarr"
	"golift.io/starr/readarr"
	"golift.io/starr/sonarr"
)

const Indexers = "Indexers"

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

	end := time.Now().Add(waitTime)
	// We use `end` and this `defer` to make every request last at least 1 second.
	// Svelte just won't update some reactive variables if you return quickly.
	defer func() { time.Sleep(time.Until(end)) }()

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

	end := time.Now().Add(waitTime)
	// We use `end` and this `defer` to make every request last at least 1 second.
	// Svelte just won't update some reactive variables if you return quickly.
	defer func() { time.Sleep(time.Until(end)) }()

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

	end := time.Now().Add(waitTime)
	// We use `end` and this `defer` to make every request last at least 1 second.
	// Svelte just won't update some reactive variables if you return quickly.
	defer func() { time.Sleep(time.Until(end)) }()

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

func (s *Starrs) ExportIndexer(config *AppConfig, selected Selected) (string, error) {
	instance, err := s.getExportInstance(config, selected, Indexers)
	if err != nil {
		return "", err
	}

	switch config.App {
	case starr.Lidarr.String():
		items, err := lidarr.New(instance.Config).GetIndexersContext(s.ctx)
		return s.exportItems(Indexers, config, filterListItemsByID(items, selected), selected.Count(), err)
	case starr.Prowlarr.String():
		items, err := prowlarr.New(instance.Config).GetIndexersContext(s.ctx)
		return s.exportItems(Indexers, config, filterListItemsByID(items, selected), selected.Count(), err)
	case starr.Radarr.String():
		items, err := radarr.New(instance.Config).GetIndexersContext(s.ctx)
		return s.exportItems(Indexers, config, filterListItemsByID(items, selected), selected.Count(), err)
	case starr.Readarr.String():
		items, err := readarr.New(instance.Config).GetIndexersContext(s.ctx)
		return s.exportItems(Indexers, config, filterListItemsByID(items, selected), selected.Count(), err)
	case starr.Sonarr.String():
		items, err := sonarr.New(instance.Config).GetIndexersContext(s.ctx)
		return s.exportItems(Indexers, config, filterListItemsByID(items, selected), selected.Count(), err)
	case starr.Whisparr.String():
		items, err := sonarr.New(instance.Config).GetIndexersContext(s.ctx)
		return s.exportItems(Indexers, config, filterListItemsByID(items, selected), selected.Count(), err)
	}

	return "", ErrInvalidApp
}

func (s *Starrs) ImportIndexer(config *AppConfig) (*DataReply, error) {
	switch config.App {
	case starr.Lidarr.String():
		var input []lidarr.IndexerOutput
		return importItems(s, Indexers, config, input)
	case starr.Prowlarr.String():
		var input []prowlarr.IndexerOutput
		return importItems(s, Indexers, config, input)
	case starr.Radarr.String():
		var input []radarr.IndexerOutput
		return importItems(s, Indexers, config, input)
	case starr.Readarr.String():
		var input []readarr.IndexerOutput
		return importItems(s, Indexers, config, input)
	case starr.Sonarr.String():
		var input []sonarr.IndexerOutput
		return importItems(s, Indexers, config, input)
	case starr.Whisparr.String():
		var input []sonarr.IndexerOutput
		return importItems(s, Indexers, config, input)
	}

	return nil, ErrInvalidApp
}

func (s *Starrs) AddLidarrIndexer(config *AppConfig, indexer *lidarr.IndexerInput) (*DataReply, error) {
	data, err := s.addIndexer(config, indexer)
	return &DataReply{Data: data, Msg: fmt.Sprintf("Imported Indexer '%s' into %s", indexer.Name, config.Name)}, err
}

func (s *Starrs) AddProwlarrIndexer(config *AppConfig, indexer *prowlarr.IndexerInput) (*DataReply, error) {
	data, err := s.addIndexer(config, indexer)
	return &DataReply{Data: data, Msg: fmt.Sprintf("Imported Indexer '%s' into %s", indexer.Name, config.Name)}, err
}

func (s *Starrs) AddRadarrIndexer(config *AppConfig, indexer *radarr.IndexerInput) (*DataReply, error) {
	data, err := s.addIndexer(config, indexer)
	return &DataReply{Data: data, Msg: fmt.Sprintf("Imported Indexer '%s' into %s", indexer.Name, config.Name)}, err
}

func (s *Starrs) AddReadarrIndexer(config *AppConfig, indexer *readarr.IndexerInput) (*DataReply, error) {
	data, err := s.addIndexer(config, indexer)
	return &DataReply{Data: data, Msg: fmt.Sprintf("Imported indexer '%s' into %s", indexer.Name, config.Name)}, err
}

func (s *Starrs) AddSonarrIndexer(config *AppConfig, indexer *sonarr.IndexerInput) (*DataReply, error) {
	data, err := s.addIndexer(config, indexer)
	return &DataReply{Data: data, Msg: fmt.Sprintf("Imported Indexer '%s' into %s", indexer.Name, config.Name)}, err
}

func (s *Starrs) AddWhisparrIndexer(config *AppConfig, indexer *sonarr.IndexerInput) (*DataReply, error) {
	data, err := s.addIndexer(config, indexer)
	return &DataReply{Data: data, Msg: fmt.Sprintf("Imported Indexer '%s' into %s", indexer.Name, config.Name)}, err
}

func (s *Starrs) addIndexer(config *AppConfig, indexer any) (any, error) {
	instance, err := s.newAPIinstance(config)
	if err != nil {
		return nil, err
	}

	end := time.Now().Add(waitTime)
	// We use `end` and this `defer` to make every request last at least 1 second.
	// Svelte just won't update some reactive variables if you return quickly.
	defer func() { time.Sleep(time.Until(end)) }()

	switch data := indexer.(type) {
	case *lidarr.IndexerInput:
		return lidarr.New(instance.Config).AddIndexerContext(s.ctx, data)
	case *prowlarr.IndexerInput:
		return prowlarr.New(instance.Config).AddIndexerContext(s.ctx, data)
	case *radarr.IndexerInput:
		return radarr.New(instance.Config).AddIndexerContext(s.ctx, data)
	case *readarr.IndexerInput:
		return readarr.New(instance.Config).AddIndexerContext(s.ctx, data)
	case *sonarr.IndexerInput:
		return sonarr.New(instance.Config).AddIndexerContext(s.ctx, data)
	default:
		return nil, fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}
