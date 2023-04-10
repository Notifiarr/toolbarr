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

func (s *Starrs) Tags(config *AppConfig) (map[int]string, error) {
	s.log.Tracef("Call:Tags(%s, %s)", config.App, config.Name)

	tags, err := s.tags(config)
	if err != nil {
		msg := s.log.Translate("Getting tags: %v", err.Error())
		s.log.Wails.Error(msg)

		return nil, fmt.Errorf(msg)
	}

	tagmap := make(map[int]string, len(tags))

	for _, tag := range tags {
		tagmap[tag.ID] = tag.Label
	}

	return tagmap, nil
}

func (s *Starrs) tags(config *AppConfig) ([]*starr.Tag, error) {
	instance, err := s.newAPIinstance(config)
	if err != nil {
		return nil, err
	}

	switch starr.App(config.App) {
	case starr.Lidarr:
		return lidarr.New(instance.Config).GetTagsContext(s.ctx)
	case starr.Prowlarr:
		return prowlarr.New(instance.Config).GetTagsContext(s.ctx)
	case starr.Radarr:
		return radarr.New(instance.Config).GetTagsContext(s.ctx)
	case starr.Readarr:
		return readarr.New(instance.Config).GetTagsContext(s.ctx)
	case starr.Sonarr:
		return sonarr.New(instance.Config).GetTagsContext(s.ctx)
	case starr.Whisparr:
		return sonarr.New(instance.Config).GetTagsContext(s.ctx)
	default:
		return nil, fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}
