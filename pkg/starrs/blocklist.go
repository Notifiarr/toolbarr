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

func (s *Starrs) DeleteBlockList(config *AppConfig, listID int64) (any, error) {
	s.log.Tracef("Call:DeleteBlockList(%s, %s, %v)", config.App, config.Name, listID)

	if err := s.deleteBlockList(config, listID); err != nil {
		msg := s.log.Translate("Deleting %s block list: %d: %v", config.Name, listID, err.Error())
		s.log.Wails.Error(msg)

		return nil, fmt.Errorf(msg)
	}

	return s.log.Translate("Deleted %s block list with ID %d.", config.Name, listID), nil
}

func (s *Starrs) deleteBlockList(config *AppConfig, listID int64) error {
	instance, err := s.newAPIinstance(config)
	if err != nil {
		return err
	}

	switch starr.App(config.App) {
	case starr.Lidarr:
		return lidarr.New(instance.Config).DeleteBlockListsContext(s.ctx, []int64{listID})
	case starr.Radarr:
		return radarr.New(instance.Config).DeleteBlockListsContext(s.ctx, []int64{listID})
	case starr.Readarr:
		return readarr.New(instance.Config).DeleteBlockListsContext(s.ctx, []int64{listID})
	case starr.Sonarr:
		return sonarr.New(instance.Config).DeleteBlockListsContext(s.ctx, []int64{listID})
	case starr.Whisparr:
		return sonarr.New(instance.Config).DeleteBlockListsContext(s.ctx, []int64{listID})
	default:
		return fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}
