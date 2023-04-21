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

func (s *Starrs) DeleteImportList(config *AppConfig, listID int64) (any, error) {
	s.log.Tracef("Call:DeleteImportList(%s, %s, %v)", config.App, config.Name, listID)

	if err := s.deleteImportList(config, listID); err != nil {
		msg := s.log.Translate("Deleting %s import list: %d: %v", config.Name, listID, err.Error())
		s.log.Wails.Error(msg)

		return nil, fmt.Errorf(msg)
	}

	return s.log.Translate("Deleted %s import list with ID %d.", config.Name, listID), nil
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

func (s *Starrs) UpdateLidarrImportList(
	config *AppConfig,
	force bool,
	list *lidarr.ImportListInput,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateLidarrImportList(%s, %s, %d)", config.App, config.Name, list.ID)
	data, err := s.updateImportList(config, force, list)

	return s.updateImportListReply(config.Name, list.Name, list.ID, data, err)
}

func (s *Starrs) UpdateRadarrImportList(
	config *AppConfig,
	force bool,
	list *radarr.ImportListInput,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateRadarrImportList(%s, %s, %d)", config.App, config.Name, list.ID)
	data, err := s.updateImportList(config, force, list)

	return s.updateImportListReply(config.Name, list.Name, list.ID, data, err)
}

func (s *Starrs) UpdateReadarrImportList(
	config *AppConfig,
	force bool,
	list *readarr.ImportListInput,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateReadarrImportList(%s, %s, %d)", config.App, config.Name, list.ID)
	data, err := s.updateImportList(config, force, list)

	return s.updateImportListReply(config.Name, list.Name, list.ID, data, err)
}

func (s *Starrs) UpdateSonarrImportList(
	config *AppConfig,
	force bool,
	list *sonarr.ImportListInput,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateSonarrImportList(%s, %s, %d)", config.App, config.Name, list.ID)
	data, err := s.updateImportList(config, force, list)

	return s.updateImportListReply(config.Name, list.Name, list.ID, data, err)
}

func (s *Starrs) UpdateWhisparrImportList(
	config *AppConfig,
	force bool,
	list *sonarr.ImportListInput,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateWhisparrImportList(%s, %s, %d)", config.App, config.Name, list.ID)
	data, err := s.updateImportList(config, force, list)

	return s.updateImportListReply(config.Name, list.Name, list.ID, data, err)
}

func (s *Starrs) updateImportList(config *AppConfig, force bool, list any) (any, error) {
	instance, err := s.newAPIinstance(config)
	if err != nil {
		return nil, err
	}

	switch data := list.(type) {
	case *lidarr.ImportListInput:
		return lidarr.New(instance.Config).UpdateImportListContext(s.ctx, data, force)
	case *radarr.ImportListInput:
		return radarr.New(instance.Config).UpdateImportListContext(s.ctx, data, force)
	case *readarr.ImportListInput:
		return readarr.New(instance.Config).UpdateImportListContext(s.ctx, data, force)
	case *sonarr.ImportListInput:
		return sonarr.New(instance.Config).UpdateImportListContext(s.ctx, data, force)
	default:
		return nil, fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}

func (s *Starrs) updateImportListReply(
	name, listName string,
	listID int64,
	data any,
	err error,
) (*DataReply, error) {
	if err == nil {
		msg := s.log.Translate("Updated %s import list %s (%d).", name, listName, listID)
		s.log.Wails.Info(msg)

		return &DataReply{Msg: msg, Data: data}, nil
	}

	reqError := &starr.ReqError{}

	if errors.As(err, &reqError) && reqError.Msg != "" {
		err = fmt.Errorf("%s: %s", reqError.Name, reqError.Msg)
	}

	msg := s.log.Translate("Updating %s import list: %s (%d): %s", name, listName, listID, err.Error())
	s.log.Wails.Error(msg)

	return nil, fmt.Errorf(msg)
}
