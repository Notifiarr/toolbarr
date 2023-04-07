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

func (s *Starrs) Downloaders(config *AppConfig) (any, error) {
	s.log.Tracef("Call:Downloaders(%s, %s)", config.App, config.Name)

	downloaders, err := s.downloaders(config)
	if err != nil {
		msg := s.log.Translate("Getting download clients: %v", err.Error())
		s.log.Wails.Error(msg)

		return nil, fmt.Errorf(msg)
	}

	return downloaders, nil
}

func (s *Starrs) downloaders(config *AppConfig) (any, error) {
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
	case starr.Whisparr:
		return sonarr.New(instance.Config).GetDownloadClientsContext(s.ctx)
	default:
		return nil, fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}

func (s *Starrs) DeleteDownloaders(config *AppConfig, ids []int64) (any, error) {
	s.log.Tracef("Call:DeleteDownloaders(%s, %s, %+v)", config.App, config.Name, ids)

	var errs []string

	for _, id := range ids {
		if err := s.deleteDownloader(config, id); err != nil {
			errs = append(errs, s.log.Translate("Deleting download client: %d: %v", id, err.Error()))
		}
	}

	if count := len(errs); count != 0 {
		errors := strings.Join(errs, ", ")
		msg := s.log.Translate("%d errors: %s", count, errors)
		s.log.Wails.Error(msg)

		return nil, fmt.Errorf(msg)
	}

	count := len(ids) // so it says "{Count}" in the translation string.

	return s.log.Translate("Deleted %d download clients.", count), nil
}

func (s *Starrs) deleteDownloader(config *AppConfig, clientID int64) error {
	instance, err := s.newAPIinstance(config)
	if err != nil {
		return err
	}

	switch starr.App(config.App) {
	case starr.Lidarr:
		return lidarr.New(instance.Config).DeleteDownloadClientContext(s.ctx, clientID)
	case starr.Prowlarr:
		return prowlarr.New(instance.Config).DeleteDownloadClientContext(s.ctx, clientID)
	case starr.Radarr:
		return radarr.New(instance.Config).DeleteDownloadClientContext(s.ctx, clientID)
	case starr.Readarr:
		return readarr.New(instance.Config).DeleteDownloadClientContext(s.ctx, clientID)
	case starr.Sonarr:
		return sonarr.New(instance.Config).DeleteDownloadClientContext(s.ctx, clientID)
	case starr.Whisparr:
		return sonarr.New(instance.Config).DeleteDownloadClientContext(s.ctx, clientID)
	default:
		return fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}
