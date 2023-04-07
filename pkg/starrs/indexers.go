//nolint:dupl
package starrs

import (
	"fmt"
	"strings"

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

func (s *Starrs) DeleteIndexers(config *AppConfig, ids []int64) (any, error) {
	s.log.Tracef("Call:DeleteIndexers(%s, %s)", config.App, config.Name)

	var errs []string

	for _, id := range ids {
		if err := s.deleteIndexers(config, id); err != nil {
			errs = append(errs, s.log.Translate("Deleting indexer: %d: %v", id, err.Error()))
		}
	}

	if count := len(errs); count != 0 {
		errors := strings.Join(errs, ", ")
		msg := s.log.Translate("%d errors: %s", count, errors)
		s.log.Wails.Error(msg)

		return nil, fmt.Errorf(msg)
	}

	count := len(ids) // so it says "{Count}" in the translation string.

	return s.log.Translate("Deleted %d indexers.", count), nil
}

func (s *Starrs) deleteIndexers(config *AppConfig, indexerID int64) error {
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
