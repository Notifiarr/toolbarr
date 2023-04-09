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

	profiles, err := s.qualityProfiles(config)
	if err != nil {
		msg := s.log.Translate("Getting quality profiles: %v", err.Error())
		s.log.Wails.Error(msg)

		return nil, fmt.Errorf(msg)
	}

	return profiles, nil
}

func (s *Starrs) qualityProfiles(config *AppConfig) (any, error) {
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
	case starr.Whisparr:
		return sonarr.New(instance.Config).GetQualityProfilesContext(s.ctx)
	default:
		return nil, fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}

func (s *Starrs) DeleteQualityProfile(config *AppConfig, profileID int64) (any, error) {
	s.log.Tracef("Call:DeleteQualityProfile(%s, %s, %v)", config.App, config.Name, profileID)

	if err := s.deleteQualityProfile(config, profileID); err != nil {
		msg := s.log.Translate("Deleting %s quality profile: %d: %v", config.Name, profileID, err.Error())
		s.log.Wails.Error(msg)

		return nil, fmt.Errorf(msg)
	}

	return s.log.Translate("Deleted %s quality profile with ID %d.", config.Name, profileID), nil
}

func (s *Starrs) deleteQualityProfile(config *AppConfig, profileID int64) error {
	instance, err := s.newAPIinstance(config)
	if err != nil {
		return err
	}

	switch starr.App(config.App) {
	case starr.Lidarr:
		return lidarr.New(instance.Config).DeleteQualityProfileContext(s.ctx, profileID)
	case starr.Radarr:
		return radarr.New(instance.Config).DeleteQualityProfileContext(s.ctx, profileID)
	case starr.Readarr:
		return readarr.New(instance.Config).DeleteQualityProfileContext(s.ctx, profileID)
	case starr.Sonarr:
		return sonarr.New(instance.Config).DeleteQualityProfileContext(s.ctx, profileID)
	case starr.Whisparr:
		return sonarr.New(instance.Config).DeleteQualityProfileContext(s.ctx, profileID)
	default:
		return fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}
