package starrs

import (
	"errors"
	"fmt"

	wr "github.com/wailsapp/wails/v2/pkg/runtime"
	"golift.io/starr"
	"golift.io/starr/lidarr"
	"golift.io/starr/radarr"
	"golift.io/starr/readarr"
	"golift.io/starr/sonarr"
)

const ImportLists = "ImportLists"

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

func (s *Starrs) TestLidarrImportList(config *AppConfig, list *lidarr.ImportListInput) (string, error) {
	s.log.Tracef("Call:TestLidarrImportList(%s, %s, %d)", config.App, config.Name, list.ID)
	return s.testImportListReply(config.Name, list.Name, list.ID, s.testImportList(config, list))
}

func (s *Starrs) TestRadarrImportList(config *AppConfig, list *radarr.ImportListInput) (string, error) {
	s.log.Tracef("Call:TestRadarrImportList(%s, %s, %d)", config.App, config.Name, list.ID)
	return s.testImportListReply(config.Name, list.Name, list.ID, s.testImportList(config, list))
}

func (s *Starrs) TestReadarrImportList(config *AppConfig, list *readarr.ImportListInput) (string, error) {
	s.log.Tracef("Call:TestReadarrImportList(%s, %s, %d)", config.App, config.Name, list.ID)
	return s.testImportListReply(config.Name, list.Name, list.ID, s.testImportList(config, list))
}

func (s *Starrs) TestSonarrImportList(config *AppConfig, list *sonarr.ImportListInput) (string, error) {
	s.log.Tracef("Call:TestSonarrImportList(%s, %s, %d)", config.App, config.Name, list.ID)
	return s.testImportListReply(config.Name, list.Name, list.ID, s.testImportList(config, list))
}

func (s *Starrs) TestWhisparrImportList(config *AppConfig, list *sonarr.ImportListInput) (string, error) {
	s.log.Tracef("Call:TestWhisparrImportList(%s, %s, %d)", config.App, config.Name, list.ID)
	return s.testImportListReply(config.Name, list.Name, list.ID, s.testImportList(config, list))
}

func (s *Starrs) testImportList(config *AppConfig, list any) error {
	instance, err := s.newAPIinstance(config)
	if err != nil {
		return err
	}

	switch data := list.(type) {
	case *lidarr.ImportListInput:
		return lidarr.New(instance.Config).TestImportListContextt(s.ctx, data)
	case *radarr.ImportListInput:
		return radarr.New(instance.Config).TestImportListContextt(s.ctx, data)
	case *readarr.ImportListInput:
		return readarr.New(instance.Config).TestImportListContextt(s.ctx, data)
	case *sonarr.ImportListInput:
		return sonarr.New(instance.Config).TestImportListContextt(s.ctx, data)
	default:
		return fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}

func (s *Starrs) testImportListReply(name, listName string, listID int64, err error) (string, error) {
	if err == nil {
		msg := s.log.Translate("Tested %s import list %s (%d).", name, listName, listID)
		s.log.Wails.Info(msg)

		return msg, nil
	}

	reqError := &starr.ReqError{}

	if errors.As(err, &reqError) && reqError.Msg != "" {
		err = fmt.Errorf("%s: %s", reqError.Name, reqError.Msg)
	}

	msg := s.log.Translate("Testing %s import list: %s (%d): %s", name, listName, listID, err.Error())
	s.log.Wails.Error(msg)

	return "", fmt.Errorf(msg)
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

func (s *Starrs) ExportLidarrImportLists(config *AppConfig, selected Selected) (string, error) {
	s.log.Tracef("Call:ExportLidarrImportList(%v)", selected)

	instance, err := s.newAPIinstance(config)
	if err != nil {
		wr.LogError(s.ctx, err.Error())
		return "", err
	}

	items, err := lidarr.New(instance.Config).GetImportListsContext(s.ctx)

	return s.exportItems(ImportLists, config, filterListItemsByID(items, selected), selected.Count(), err)
}

func (s *Starrs) ExportRadarrImportLists(config *AppConfig, selected Selected) (string, error) {
	s.log.Tracef("Call:ExportRadarrImportList(%v)", selected)

	instance, err := s.newAPIinstance(config)
	if err != nil {
		wr.LogError(s.ctx, err.Error())
		return "", err
	}

	items, err := radarr.New(instance.Config).GetImportListsContext(s.ctx)

	return s.exportItems(ImportLists, config, filterListItemsByID(items, selected), selected.Count(), err)
}

func (s *Starrs) ExportReadarrImportLists(config *AppConfig, selected Selected) (string, error) {
	s.log.Tracef("Call:ExportReadarrImportList(%v)", selected)

	instance, err := s.newAPIinstance(config)
	if err != nil {
		wr.LogError(s.ctx, err.Error())
		return "", err
	}

	items, err := readarr.New(instance.Config).GetImportListsContext(s.ctx)

	return s.exportItems(ImportLists, config, filterListItemsByID(items, selected), selected.Count(), err)
}

func (s *Starrs) ExportSonarrImportLists(config *AppConfig, selected Selected) (string, error) {
	s.log.Tracef("Call:ExportSonarrImportList(%v)", selected)

	instance, err := s.newAPIinstance(config)
	if err != nil {
		wr.LogError(s.ctx, err.Error())
		return "", err
	}

	items, err := sonarr.New(instance.Config).GetImportListsContext(s.ctx)

	return s.exportItems(ImportLists, config, filterListItemsByID(items, selected), selected.Count(), err)
}

func (s *Starrs) ExportWhisparrImportLists(config *AppConfig, selected Selected) (string, error) {
	s.log.Tracef("Call:ExportWhisparrImportList(%v)", selected)

	instance, err := s.newAPIinstance(config)
	if err != nil {
		wr.LogError(s.ctx, err.Error())
		return "", err
	}

	items, err := sonarr.New(instance.Config).GetImportListsContext(s.ctx)

	return s.exportItems(ImportLists, config, filterListItemsByID(items, selected), selected.Count(), err)
}

func (s *Starrs) ImportLidarrImportLists(config *AppConfig) (map[string]any, error) {
	var input []lidarr.ImportListInput
	return importItems(s, ImportLists, config, input)
}

func (s *Starrs) ImportRadarrImportLists(config *AppConfig) (map[string]any, error) {
	var input []radarr.ImportListInput
	return importItems(s, ImportLists, config, input)
}

func (s *Starrs) ImportReadarrImportLists(config *AppConfig) (map[string]any, error) {
	var input []readarr.ImportListInput
	return importItems(s, ImportLists, config, input)
}

func (s *Starrs) ImportSonarrImportLists(config *AppConfig) (map[string]any, error) {
	var input []sonarr.ImportListInput
	return importItems(s, ImportLists, config, input)
}

func (s *Starrs) ImportWhisparrImportLists(config *AppConfig) (map[string]any, error) {
	var input []sonarr.ImportListInput
	return importItems(s, ImportLists, config, input)
}
