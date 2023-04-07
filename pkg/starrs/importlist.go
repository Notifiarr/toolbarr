package starrs

import (
	"fmt"
	"strings"

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
		msg := s.log.Translate("Getting import list: %v", err.Error())
		s.log.Wails.Error(msg)

		return nil, fmt.Errorf(msg)
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

func (s *Starrs) DeleteImportLists(config *AppConfig, ids []int64) (any, error) {
	s.log.Tracef("Call:DeleteImportLists(%s, %s)", config.App, config.Name)

	var errs []string

	for _, id := range ids {
		if err := s.deleteImportList(config, id); err != nil {
			errs = append(errs, s.log.Translate("Deleting import list: %d: %v", id, err.Error()))
		}
	}

	if count := len(errs); count != 0 {
		errors := strings.Join(errs, ", ")
		msg := s.log.Translate("%d errors: %s", count, errors)
		s.log.Wails.Error(msg)

		return nil, fmt.Errorf(msg)
	}

	count := len(ids) // so it says "{Count}" in the translation string.

	return s.log.Translate("Deleted %d import lists.", count), nil
}

func (s *Starrs) deleteImportList(config *AppConfig, listID int64) error {
	instance, err := s.newAPIinstance(config)
	if err != nil {
		return err
	}

	switch starr.App(config.App) {
	case starr.Lidarr:
		return lidarr.New(instance.Config).DeleteImportListContext(s.ctx, listID)
	case starr.Radarr:
		return radarr.New(instance.Config).DeleteImportListContext(s.ctx, []int64{listID})
	case starr.Readarr:
		return readarr.New(instance.Config).DeleteImportListContext(s.ctx, listID)
	case starr.Sonarr:
		return sonarr.New(instance.Config).DeleteImportListContext(s.ctx, listID)
	case starr.Whisparr:
		return sonarr.New(instance.Config).DeleteImportListContext(s.ctx, listID)
	default:
		return fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}

func (s *Starrs) Exclusions(config *AppConfig) (any, error) {
	s.log.Tracef("Call:Exclusions(%s, %s)", config.App, config.Name)

	list, err := s.exclusions(config)
	if err != nil {
		msg := s.log.Translate("Getting exclusions: %v", err.Error())
		s.log.Wails.Error(msg)

		return "", fmt.Errorf(msg)
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

func (s *Starrs) DeleteExclusions(config *AppConfig, ids []int64) (any, error) {
	s.log.Tracef("Call:DeleteExclusions(%s, %s)", config.App, config.Name)

	var errs []string

	for _, id := range ids {
		if err := s.deleteExclusion(config, id); err != nil {
			errs = append(errs, s.log.Translate("Deleting import list exclusion: %d: %v", id, err.Error()))
		}
	}

	if count := len(errs); count != 0 {
		errors := strings.Join(errs, ", ")
		msg := s.log.Translate("%d errors: %s", count, errors)
		s.log.Wails.Error(msg)

		return nil, fmt.Errorf(msg)
	}

	count := len(ids) // so it says "{Count}" in the translation string.

	return s.log.Translate("Deleted %d import list exclusions.", count), nil
}

func (s *Starrs) deleteExclusion(config *AppConfig, exclusionID int64) error {
	instance, err := s.newAPIinstance(config)
	if err != nil {
		return err
	}

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
