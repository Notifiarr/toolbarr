package starrs

import (
	"errors"
	"fmt"
	"time"

	"golift.io/starr"
	"golift.io/starr/lidarr"
	"golift.io/starr/radarr"
	"golift.io/starr/readarr"
	"golift.io/starr/sonarr"
)

const Exclusions = "Exclusions"

func (s *Starrs) Exclusions(config *AppConfig) (any, error) {
	s.log.Tracef("Call:Exclusions(%s, %s)", config.App, config.Name)

	exclusion, err := s.exclusions(config)
	if err != nil {
		msg := s.log.Translate("Getting import list exclusions: %v", err.Error())
		s.log.Wails.Error(msg)

		return "", fmt.Errorf(msg)
	}

	return exclusion, nil
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

func (s *Starrs) DeleteExclusion(config *AppConfig, exclusionID int64) (any, error) {
	s.log.Tracef("Call:DeleteExclusion(%s, %s, %+v)", config.App, config.Name, exclusionID)

	if err := s.deleteExclusion(config, exclusionID); err != nil {
		msg := s.log.Translate("Deleting %s import list exclusion: %d: %v", config.Name, exclusionID, err.Error())
		s.log.Wails.Error(msg)

		return nil, fmt.Errorf(msg)
	}

	return s.log.Translate("Deleted %s import list exclusion with ID %d.", config.Name, exclusionID), nil
}

func (s *Starrs) deleteExclusion(config *AppConfig, exclusionID int64) error {
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
		return lidarr.New(instance.Config).DeleteExclusionsContext(s.ctx, []int64{exclusionID})
	case starr.Radarr:
		return radarr.New(instance.Config).DeleteExclusionsContext(s.ctx, []int64{exclusionID})
	case starr.Readarr:
		return readarr.New(instance.Config).DeleteExclusionsContext(s.ctx, []int64{exclusionID})
	case starr.Sonarr:
		return sonarr.New(instance.Config).DeleteExclusionsContext(s.ctx, []int64{exclusionID})
	case starr.Whisparr:
		return sonarr.New(instance.Config).DeleteExclusionsContext(s.ctx, []int64{exclusionID})
	default:
		return fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}

func (s *Starrs) UpdateLidarrExclusion(
	config *AppConfig,
	exclusion *lidarr.Exclusion,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateLidarrExclusion(%s, %s, %d)", config.App, config.Name, exclusion.ID)
	data, err := s.updateExclusion(config, exclusion)

	return s.updateExclusionReply(config.Name, exclusion.ArtistName, exclusion.ID, data, err)
}

func (s *Starrs) UpdateRadarrExclusion(
	config *AppConfig,
	exclusion *radarr.Exclusion,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateRadarrExclusion(%s, %s, %d)", config.App, config.Name, exclusion.ID)
	data, err := s.updateExclusion(config, exclusion)

	return s.updateExclusionReply(config.Name, exclusion.Title, exclusion.ID, data, err)
}

func (s *Starrs) UpdateReadarrExclusion(
	config *AppConfig,
	exclusion *readarr.Exclusion,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateReadarrExclusion(%s, %s, %d)", config.App, config.Name, exclusion.ID)
	data, err := s.updateExclusion(config, exclusion)

	return s.updateExclusionReply(config.Name, exclusion.AuthorName, exclusion.ID, data, err)
}

func (s *Starrs) UpdateSonarrExclusion(
	config *AppConfig,
	exclusion *sonarr.Exclusion,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateSonarrExclusion(%s, %s, %d)", config.App, config.Name, exclusion.ID)
	data, err := s.updateExclusion(config, exclusion)

	return s.updateExclusionReply(config.Name, exclusion.Title, exclusion.ID, data, err)
}

func (s *Starrs) UpdateWhisparrExclusion(
	config *AppConfig,
	exclusion *sonarr.Exclusion,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateWhisparrExclusion(%s, %s, %d)", config.App, config.Name, exclusion.ID)
	data, err := s.updateExclusion(config, exclusion)

	return s.updateExclusionReply(config.Name, exclusion.Title, exclusion.ID, data, err)
}

func (s *Starrs) updateExclusion(config *AppConfig, exclusion any) (any, error) {
	instance, err := s.newAPIinstance(config)
	if err != nil {
		return nil, err
	}

	end := time.Now().Add(waitTime)
	// We use `end` and this `defer` to make every request last at least 1 second.
	// Svelte just won't update some reactive variables if you return quickly.
	defer func() { time.Sleep(time.Until(end)) }()

	switch data := exclusion.(type) {
	case *lidarr.Exclusion:
		return lidarr.New(instance.Config).UpdateExclusionContext(s.ctx, data)
	case *radarr.Exclusion:
		return radarr.New(instance.Config).UpdateExclusionContext(s.ctx, data)
	case *readarr.Exclusion:
		return readarr.New(instance.Config).UpdateExclusionContext(s.ctx, data)
	case *sonarr.Exclusion:
		return sonarr.New(instance.Config).UpdateExclusionContext(s.ctx, data)
	default:
		return nil, fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}

func (s *Starrs) updateExclusionReply(
	name, exclusionName string,
	exclusionID int64,
	data any,
	err error,
) (*DataReply, error) {
	if err == nil {
		msg := s.log.Translate("Updated %s import list exclusion %s (%d).", name, exclusionName, exclusionID)
		s.log.Wails.Info(msg)

		return &DataReply{Msg: msg, Data: data}, nil
	}

	reqError := &starr.ReqError{}

	if errors.As(err, &reqError) && reqError.Msg != "" {
		err = fmt.Errorf("%s: %s", reqError.Name, reqError.Msg)
	}

	msg := s.log.Translate("Updating %s import list exclusion: %s (%d): %s", name, exclusionName, exclusionID, err.Error())
	s.log.Wails.Error(msg)

	return nil, fmt.Errorf(msg)
}

func (s *Starrs) ExportExclusions(config *AppConfig, selected Selected) (string, error) {
	instance, err := s.getExportInstance(config, selected, Exclusions)
	if err != nil {
		return "", err
	}

	switch config.App {
	case starr.Lidarr.String():
		items, err := lidarr.New(instance.Config).GetExclusionsContext(s.ctx)
		return s.exportItems(Exclusions, config, filterListItemsByID(items, selected), selected.Count(), err)
	case starr.Radarr.String():
		items, err := radarr.New(instance.Config).GetExclusionsContext(s.ctx)
		return s.exportItems(Exclusions, config, filterListItemsByID(items, selected), selected.Count(), err)
	case starr.Readarr.String():
		items, err := readarr.New(instance.Config).GetExclusionsContext(s.ctx)
		return s.exportItems(Exclusions, config, filterListItemsByID(items, selected), selected.Count(), err)
	case starr.Sonarr.String():
		items, err := sonarr.New(instance.Config).GetExclusionsContext(s.ctx)
		return s.exportItems(Exclusions, config, filterListItemsByID(items, selected), selected.Count(), err)
	case starr.Whisparr.String():
		items, err := sonarr.New(instance.Config).GetExclusionsContext(s.ctx)
		return s.exportItems(Exclusions, config, filterListItemsByID(items, selected), selected.Count(), err)
	}

	return "", ErrInvalidApp
}

func (s *Starrs) ImportExclusions(config *AppConfig) (*DataReply, error) {
	switch config.App {
	case starr.Lidarr.String():
		var input []lidarr.Exclusion
		return importItems(s, Exclusions, config, input)
	case starr.Radarr.String():
		var input []radarr.Exclusion
		return importItems(s, Exclusions, config, input)
	case starr.Readarr.String():
		var input []readarr.Exclusion
		return importItems(s, Exclusions, config, input)
	case starr.Sonarr.String():
		var input []sonarr.Exclusion
		return importItems(s, Exclusions, config, input)
	case starr.Whisparr.String():
		var input []sonarr.Exclusion
		return importItems(s, Exclusions, config, input)
	}

	return nil, ErrInvalidApp
}

func (s *Starrs) AddLidarrExclusion(config *AppConfig, exclusion *lidarr.Exclusion) (*DataReply, error) {
	data, err := s.addExclusion(config, exclusion)

	return &DataReply{
		Data: data,
		Msg:  fmt.Sprintf("Imported Exclusion '%s' into %s", exclusion.ArtistName, config.Name),
	}, err
}

func (s *Starrs) AddRadarrExclusion(config *AppConfig, exclusion *radarr.Exclusion) (*DataReply, error) {
	data, err := s.addExclusion(config, exclusion)
	if err != nil {
		return nil, err
	}

	return &DataReply{Data: data, Msg: fmt.Sprintf("Imported Exclusion '%s' into %s", exclusion.Title, config.Name)}, err
}

func (s *Starrs) AddReadarrExclusion(config *AppConfig, exclusion *readarr.Exclusion) (*DataReply, error) {
	data, err := s.addExclusion(config, exclusion)

	return &DataReply{
		Data: data,
		Msg:  fmt.Sprintf("Imported Exclusion '%s' into %s", exclusion.AuthorName, config.Name),
	}, err
}

func (s *Starrs) AddSonarrExclusion(config *AppConfig, exclusion *sonarr.Exclusion) (*DataReply, error) {
	data, err := s.addExclusion(config, exclusion)
	return &DataReply{Data: data, Msg: fmt.Sprintf("Imported Exclusion '%s' into %s", exclusion.Title, config.Name)}, err
}

func (s *Starrs) AddWhisparrExclusion(config *AppConfig, exclusion *sonarr.Exclusion) (*DataReply, error) {
	data, err := s.addExclusion(config, exclusion)
	return &DataReply{Data: data, Msg: fmt.Sprintf("Imported Exclusion '%s' into %s", exclusion.Title, config.Name)}, err
}

func (s *Starrs) addExclusion(config *AppConfig, exclusion any) (any, error) {
	instance, err := s.newAPIinstance(config)
	if err != nil {
		return nil, err
	}

	end := time.Now().Add(waitTime)
	// We use `end` and this `defer` to make every request last at least 1 second.
	// Svelte just won't update some reactive variables if you return quickly.
	defer func() { time.Sleep(time.Until(end)) }()

	switch data := exclusion.(type) {
	case *lidarr.Exclusion:
		return lidarr.New(instance.Config).AddExclusionContext(s.ctx, data)
	case *radarr.Exclusion:
		return radarr.New(instance.Config).AddExclusionContext(s.ctx, data)
	case *readarr.Exclusion:
		return readarr.New(instance.Config).AddExclusionContext(s.ctx, data)
	case *sonarr.Exclusion:
		return sonarr.New(instance.Config).AddExclusionContext(s.ctx, data)
	default:
		return nil, fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}
