package starrs

import (
	"fmt"

	"golift.io/starr"
	"golift.io/starr/lidarr"
	"golift.io/starr/radarr"
	"golift.io/starr/readarr"
	"golift.io/starr/sonarr"
)

const getBlockLists = 10000

// DataReply is a generic reply with a message and data.
type DataReply struct {
	Msg  string
	Data any
}

func (s *Starrs) BlockList(config *AppConfig) (any, error) {
	s.log.Tracef("Call:BlockList(%s, %s)", config.App, config.Name)

	list, err := s.blockList(config)
	if err != nil {
		msg := s.log.Translate("Getting block lists: %v", err.Error())
		s.log.Wails.Error(msg)

		return nil, fmt.Errorf(msg)
	}

	return list, nil
}

func (s *Starrs) blockList(config *AppConfig) (any, error) {
	instance, err := s.newAPIinstance(config)
	if err != nil {
		return nil, err
	}

	switch starr.App(config.App) {
	case starr.Lidarr:
		return lidarr.New(instance.Config).GetBlockListContext(s.ctx, getBlockLists)
	case starr.Radarr:
		return radarr.New(instance.Config).GetBlockListContext(s.ctx, getBlockLists)
	case starr.Readarr:
		return readarr.New(instance.Config).GetBlockListContext(s.ctx, getBlockLists)
	case starr.Sonarr:
		return sonarr.New(instance.Config).GetBlockListContext(s.ctx, getBlockLists)
	case starr.Whisparr:
		return sonarr.New(instance.Config).GetBlockListContext(s.ctx, getBlockLists)
	default:
		return nil, fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}

func (s *Starrs) DeleteBlockLists(config *AppConfig, ids []int64) (any, error) {
	s.log.Tracef("Call:DeleteBlockLists(%s, %s, %+v)", config.App, config.Name, ids)

	output, err := s.deleteBlockLists(config, ids)
	if err != nil {
		msg := s.log.Translate("Deleting block lists: %v", err.Error())
		s.log.Wails.Error(msg)

		return nil, fmt.Errorf(msg)
	}

	return output, nil
}

func (s *Starrs) deleteBlockLists(config *AppConfig, ids []int64) (any, error) {
	instance, err := s.newAPIinstance(config)
	if err != nil {
		return nil, err
	}

	count := len(ids)
	msg := s.log.Translate("Deleted %d block list items.", count)

	switch starr.App(config.App) {
	case starr.Lidarr:
		return msg, lidarr.New(instance.Config).DeleteBlockListsContext(s.ctx, ids)
	case starr.Radarr:
		return msg, radarr.New(instance.Config).DeleteBlockListsContext(s.ctx, ids)
	case starr.Readarr:
		return msg, readarr.New(instance.Config).DeleteBlockListsContext(s.ctx, ids)
	case starr.Sonarr:
		return msg, sonarr.New(instance.Config).DeleteBlockListsContext(s.ctx, ids)
	case starr.Whisparr:
		return msg, sonarr.New(instance.Config).DeleteBlockListsContext(s.ctx, ids)
	default:
		return nil, fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}
