package starrs

import (
	"fmt"

	"golift.io/starr"
	"golift.io/starr/lidarr"
	"golift.io/starr/radarr"
	"golift.io/starr/readarr"
	"golift.io/starr/sonarr"
)

func (s *Starrs) QualityProfiles(config *AppConfig) (any, error) {
	s.log.Tracef("Call:QualityProfiles(%s, %s)", config.App, config.Name)

	instance, err := s.newAPIinstance(config)
	if err != nil {
		return nil, err
	}

	switch starr.App(config.App) {
	case starr.Lidarr:
		return lidarr.New(instance.Config).GetQualityProfilesContext(s.ctx)
	case starr.Radarr:
		return radarr.New(instance.Config).GetQualityProfilesContext(s.ctx)
	case starr.Readarr:
		return readarr.New(instance.Config).GetQualityProfilesContext(s.ctx)
	case starr.Sonarr:
		return sonarr.New(instance.Config).GetQualityProfilesContext(s.ctx)
	case "Whisparr":
		return sonarr.New(instance.Config).GetQualityProfilesContext(s.ctx)
	default:
		return nil, fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}
