package starrs

import (
	"fmt"

	"golift.io/starr"
	"golift.io/starr/lidarr"
	"golift.io/starr/radarr"
	"golift.io/starr/readarr"
	"golift.io/starr/sonarr"
)

func (s *Starrs) ImportLists(config *AppConfig) (any, error) {
	s.log.Tracef("Call:ImportLists(%s, %s)", config.App, config.Name)

	list, err := s.importList(config)
	if err != nil {
		return "", fmt.Errorf(s.log.Translate("Getting Import List: %v", err.Error()))
	}

	return list, nil
}

func (s *Starrs) importList(config *AppConfig) (any, error) {
	instance, err := s.newAPIinstance(config)
	if err != nil {
		return nil, err
	}

	switch starr.App(config.App) {
	case starr.Lidarr:
		return lidarr.New(instance.Config).GetImportListsContext(s.ctx)
	case starr.Radarr:
		return radarr.New(instance.Config).GetImportListsContext(s.ctx)
	case starr.Readarr:
		return readarr.New(instance.Config).GetImportListsContext(s.ctx)
	case starr.Sonarr:
		return sonarr.New(instance.Config).GetImportListsContext(s.ctx)
	case starr.Whisparr:
		return sonarr.New(instance.Config).GetImportListsContext(s.ctx)
	default:
		return nil, fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}

func (s *Starrs) Exclusions(config *AppConfig) (any, error) {
	s.log.Tracef("Call:Exclusions(%s, %s)", config.App, config.Name)

	list, err := s.exclusions(config)
	if err != nil {
		return "", fmt.Errorf(s.log.Translate("Getting Exclusions: %v", err.Error()))
	}

	return list, nil
}

func (s *Starrs) exclusions(config *AppConfig) (any, error) {
	instance, err := s.newAPIinstance(config)
	if err != nil {
		return nil, err
	}

	switch starr.App(config.App) {
	case starr.Lidarr:
		return lidarr.New(instance.Config).GetExclusionsContext(s.ctx)
	case starr.Radarr:
		return radarr.New(instance.Config).GetExclusionsContext(s.ctx)
	case starr.Readarr:
		return readarr.New(instance.Config).GetExclusionsContext(s.ctx)
	case starr.Sonarr:
		return sonarr.New(instance.Config).GetExclusionsContext(s.ctx)
	case starr.Whisparr:
		return sonarr.New(instance.Config).GetExclusionsContext(s.ctx)
	default:
		return nil, fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}
