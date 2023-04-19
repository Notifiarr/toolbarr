package starrs

import (
	"errors"
	"fmt"

	"golift.io/starr"
	"golift.io/starr/lidarr"
	"golift.io/starr/radarr"
	"golift.io/starr/readarr"
	"golift.io/starr/sonarr"
)

func (s *Starrs) MetadataProfiles(config *AppConfig) (any, error) {
	s.log.Tracef("Call:MetadataProfiles(%s, %s)", config.App, config.Name)

	profiles, err := s.metadataProfiles(config)
	if err != nil {
		msg := s.log.Translate("Getting metadata profiles: %v", err.Error())
		s.log.Wails.Error(msg)

		return nil, fmt.Errorf(msg)
	}

	return profiles, nil
}

func (s *Starrs) metadataProfiles(config *AppConfig) (any, error) {
	instance, err := s.newAPIinstance(config)
	if err != nil {
		return nil, err
	}

	switch starr.App(config.App) {
	case starr.Lidarr:
		return lidarr.New(instance.Config).GetMetadataProfilesContext(s.ctx)
	case starr.Readarr:
		return readarr.New(instance.Config).GetMetadataProfilesContext(s.ctx)
	default:
		return nil, fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}

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

func (s *Starrs) UpdateLidarrQualityProfile(
	config *AppConfig,
	profile *lidarr.QualityProfile,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateLidarrQualityProfile(%s, %s, %d)", config.App, config.Name, profile.ID)
	data, err := s.updateQualityProfile(config, profile)

	return s.updateQualityProfileReply(config.Name, profile.Name, profile.ID, data, err)
}

func (s *Starrs) UpdateRadarrQualityProfile(
	config *AppConfig,
	profile *radarr.QualityProfile,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateRadarrQualityProfile(%s, %s, %d)", config.App, config.Name, profile.ID)
	data, err := s.updateQualityProfile(config, profile)

	return s.updateQualityProfileReply(config.Name, profile.Name, profile.ID, data, err)
}

func (s *Starrs) UpdateReadarrQualityProfile(
	config *AppConfig,
	profile *readarr.QualityProfile,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateReadarrQualityProfile(%s, %s, %d)", config.App, config.Name, profile.ID)
	data, err := s.updateQualityProfile(config, profile)

	return s.updateQualityProfileReply(config.Name, profile.Name, profile.ID, data, err)
}

func (s *Starrs) UpdateSonarrQualityProfile(
	config *AppConfig,
	profile *sonarr.QualityProfile,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateSonarrQualityProfile(%s, %s, %d)", config.App, config.Name, profile.ID)
	data, err := s.updateQualityProfile(config, profile)

	return s.updateQualityProfileReply(config.Name, profile.Name, profile.ID, data, err)
}

func (s *Starrs) UpdateWhisparrQualityProfile(
	config *AppConfig,
	profile *sonarr.QualityProfile,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateWhisparrQualityProfile(%s, %s, %d)", config.App, config.Name, profile.ID)
	data, err := s.updateQualityProfile(config, profile)

	return s.updateQualityProfileReply(config.Name, profile.Name, profile.ID, data, err)
}

func (s *Starrs) updateQualityProfile(config *AppConfig, profile any) (any, error) {
	instance, err := s.newAPIinstance(config)
	if err != nil {
		return nil, err
	}

	switch data := profile.(type) {
	case *lidarr.QualityProfile:
		return lidarr.New(instance.Config).UpdateQualityProfileContext(s.ctx, data)
	case *radarr.QualityProfile:
		return radarr.New(instance.Config).UpdateQualityProfileContext(s.ctx, data)
	case *readarr.QualityProfile:
		return readarr.New(instance.Config).UpdateQualityProfileContext(s.ctx, data)
	case *sonarr.QualityProfile:
		return sonarr.New(instance.Config).UpdateQualityProfileContext(s.ctx, data)
	default:
		return nil, fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}

func (s *Starrs) updateQualityProfileReply(
	name, profile string,
	profileID int64,
	data any,
	err error,
) (*DataReply, error) {
	if err == nil {
		msg := s.log.Translate("Updated %s quality profile %s (%d).", name, profile, profileID)
		s.log.Wails.Info(msg)

		return &DataReply{Msg: msg, Data: data}, nil
	}

	reqError := &starr.ReqError{}

	if errors.As(err, &reqError) && reqError.Msg != "" {
		err = fmt.Errorf("%s: %s", reqError.Name, reqError.Msg)
	}

	msg := s.log.Translate("Updating %s quality profile: %s (%d): %s", name, profile, profileID, err.Error())
	s.log.Wails.Error(msg)

	return nil, fmt.Errorf(msg)
}
