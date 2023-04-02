package starrs

import (
	"fmt"

	"golift.io/starr"
	"golift.io/starr/lidarr"
	"golift.io/starr/prowlarr"
	"golift.io/starr/radarr"
	"golift.io/starr/readarr"
	"golift.io/starr/sonarr"
)

func (s *Starrs) Downloaders(config *AppConfig) (any, error) {
	s.log.Tracef("Call:Downloaders(%s, %s)", config.App, config.Name)

	instance, err := s.newAPIinstance(config)
	if err != nil {
		return nil, err
	}

	switch starr.App(config.App) {
	case starr.Lidarr:
		return lidarr.New(instance.Config).GetDownloadClientsContext(s.ctx)
	case starr.Prowlarr:
		return prowlarr.New(instance.Config).GetDownloadClientsContext(s.ctx)
	case starr.Radarr:
		return radarr.New(instance.Config).GetDownloadClientsContext(s.ctx)
	case starr.Readarr:
		return readarr.New(instance.Config).GetDownloadClientsContext(s.ctx)
	case starr.Sonarr:
		return sonarr.New(instance.Config).GetDownloadClientsContext(s.ctx)
	case "Whisparr":
		return sonarr.New(instance.Config).GetDownloadClientsContext(s.ctx)
	default:
		return nil, fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}
