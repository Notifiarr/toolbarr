package starrs

import (
	"fmt"

	"golift.io/starr"
	"golift.io/starr/lidarr"
	"golift.io/starr/radarr"
	"golift.io/starr/readarr"
	"golift.io/starr/sonarr"
)

func (s *Starrs) RootFolders(config *AppConfig) (any, error) {
	s.log.Tracef("Call:RootFolders(%s, %s)", config.App, config.Name)

	list, err := s.rootFolders(config)
	if err != nil {
		msg := s.log.Translate("Getting root folders: %v", err.Error())
		s.log.Wails.Error(msg)

		return nil, fmt.Errorf(msg)
	}

	return list, nil
}

func (s *Starrs) rootFolders(config *AppConfig) (any, error) {
	instance, err := s.newAPIinstance(config)
	if err != nil {
		return nil, err
	}

	switch starr.App(config.App) {
	case starr.Lidarr:
		return lidarr.New(instance.Config).GetRootFoldersContext(s.ctx)
	case starr.Radarr:
		return radarr.New(instance.Config).GetRootFoldersContext(s.ctx)
	case starr.Readarr:
		return readarr.New(instance.Config).GetRootFoldersContext(s.ctx)
	case starr.Sonarr:
		return sonarr.New(instance.Config).GetRootFoldersContext(s.ctx)
	case starr.Whisparr:
		return sonarr.New(instance.Config).GetRootFoldersContext(s.ctx)
	default:
		return nil, fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}
