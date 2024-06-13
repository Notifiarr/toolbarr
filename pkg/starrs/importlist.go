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

	end := time.Now().Add(waitTime)
	// We use `end` and this `defer` to make every request last at least 1 second.
	// Svelte just won't update some reactive variables if you return quickly.
	defer func() { time.Sleep(time.Until(end)) }()

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

	end := time.Now().Add(waitTime)
	// We use `end` and this `defer` to make every request last at least 1 second.
	// Svelte just won't update some reactive variables if you return quickly.
	defer func() { time.Sleep(time.Until(end)) }()

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
	data, err := s.updateImportList(config, list, force)

	return s.updateImportListReply(config.Name, list.Name, list.ID, data, err)
}

func (s *Starrs) UpdateRadarrImportList(
	config *AppConfig,
	list *radarr.ImportListInput,
	force bool,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateRadarrImportList(%s, %s, %d)", config.App, config.Name, list.ID)
	data, err := s.updateImportList(config, list, force)

	return s.updateImportListReply(config.Name, list.Name, list.ID, data, err)
}

func (s *Starrs) UpdateReadarrImportList(
	config *AppConfig,
	list *readarr.ImportListInput,
	force bool,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateReadarrImportList(%s, %s, %d)", config.App, config.Name, list.ID)
	data, err := s.updateImportList(config, list, force)

	return s.updateImportListReply(config.Name, list.Name, list.ID, data, err)
}

func (s *Starrs) UpdateSonarrImportList(
	config *AppConfig,
	list *sonarr.ImportListInput,
	force bool,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateSonarrImportList(%s, %s, %d)", config.App, config.Name, list.ID)
	data, err := s.updateImportList(config, list, force)

	return s.updateImportListReply(config.Name, list.Name, list.ID, data, err)
}

func (s *Starrs) UpdateWhisparrImportList(
	config *AppConfig,
	list *sonarr.ImportListInput,
	force bool,
) (*DataReply, error) {
	s.log.Tracef("Call:UpdateWhisparrImportList(%s, %s, %d)", config.App, config.Name, list.ID)
	data, err := s.updateImportList(config, list, force)

	return s.updateImportListReply(config.Name, list.Name, list.ID, data, err)
}

func (s *Starrs) updateImportList(config *AppConfig, list any, force bool) (any, error) {
	instance, err := s.newAPIinstance(config)
	if err != nil {
		return nil, err
	}

	end := time.Now().Add(waitTime)
	// We use `end` and this `defer` to make every request last at least 1 second.
	// Svelte just won't update some reactive variables if you return quickly.
	defer func() { time.Sleep(time.Until(end)) }()

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

func (s *Starrs) ExportImportLists(config *AppConfig, selected Selected) (string, error) {
	instance, err := s.getExportInstance(config, selected, ImportLists)
	if err != nil {
		return "", err
	}

	switch config.App {
	case starr.Lidarr.String():
		items, err := lidarr.New(instance.Config).GetImportListsContext(s.ctx)
		return s.exportItems(ImportLists, config, filterListItemsByID(items, selected), selected.Count(), err)
	case starr.Radarr.String():
		items, err := radarr.New(instance.Config).GetImportListsContext(s.ctx)
		return s.exportItems(ImportLists, config, filterListItemsByID(items, selected), selected.Count(), err)
	case starr.Readarr.String():
		items, err := readarr.New(instance.Config).GetImportListsContext(s.ctx)
		return s.exportItems(ImportLists, config, filterListItemsByID(items, selected), selected.Count(), err)
	case starr.Sonarr.String():
		items, err := sonarr.New(instance.Config).GetImportListsContext(s.ctx)
		return s.exportItems(ImportLists, config, filterListItemsByID(items, selected), selected.Count(), err)
	case starr.Whisparr.String():
		items, err := sonarr.New(instance.Config).GetImportListsContext(s.ctx)
		return s.exportItems(ImportLists, config, filterListItemsByID(items, selected), selected.Count(), err)
	}

	return "", ErrInvalidApp
}

func (s *Starrs) ImportImportLists(config *AppConfig) (*DataReply, error) {
	switch config.App {
	case starr.Lidarr.String():
		var input []lidarr.ImportListInput
		return importItems(s, ImportLists, config, input)
	case starr.Radarr.String():
		var input []radarr.ImportListInput
		return importItems(s, ImportLists, config, input)
	case starr.Readarr.String():
		var input []readarr.ImportListInput
		return importItems(s, ImportLists, config, input)
	case starr.Sonarr.String():
		var input []sonarr.ImportListInput
		return importItems(s, ImportLists, config, input)
	case starr.Whisparr.String():
		var input []sonarr.ImportListInput
		return importItems(s, ImportLists, config, input)
	}

	return nil, ErrInvalidApp
}

func (s *Starrs) AddLidarrImportList(config *AppConfig, list *lidarr.ImportListInput) (*DataReply, error) {
	data, err := s.addImportList(config, list, list.Name)
	return &DataReply{Data: data, Msg: fmt.Sprintf("Imported Import List '%s' into %s", list.Name, config.Name)}, err
}

func (s *Starrs) AddRadarrImportList(config *AppConfig, list *radarr.ImportListInput) (*DataReply, error) {
	data, err := s.addImportList(config, list, list.Name)
	return &DataReply{Data: data, Msg: fmt.Sprintf("Imported Import List '%s' into %s", list.Name, config.Name)}, err
}

func (s *Starrs) AddReadarrImportList(config *AppConfig, list *readarr.ImportListInput) (*DataReply, error) {
	data, err := s.addImportList(config, list, list.Name)
	return &DataReply{Data: data, Msg: fmt.Sprintf("Imported Import List '%s' into %s", list.Name, config.Name)}, err
}

func (s *Starrs) AddSonarrImportList(config *AppConfig, list *sonarr.ImportListInput) (*DataReply, error) {
	data, err := s.addImportList(config, list, list.Name)
	return &DataReply{Data: data, Msg: fmt.Sprintf("Imported Import List '%s' into %s", list.Name, config.Name)}, err
}

func (s *Starrs) AddWhisparrImportList(config *AppConfig, list *sonarr.ImportListInput) (*DataReply, error) {
	data, err := s.addImportList(config, list, list.Name)
	return &DataReply{Data: data, Msg: fmt.Sprintf("Imported Import List '%s' into %s", list.Name, config.Name)}, err
}

func (s *Starrs) addImportList(config *AppConfig, list any, listName string) (any, error) {
	s.log.Tracef("Call:Add%sImportList(%s, %s)", config.App, config.Name, listName)

	instance, err := s.newAPIinstance(config)
	if err != nil {
		return nil, err
	}

	end := time.Now().Add(waitTime)
	// We use `end` and this `defer` to make every request last at least 1 second.
	// Svelte just won't update some reactive variables if you return quickly.
	defer func() { time.Sleep(time.Until(end)) }()

	switch data := list.(type) {
	case *lidarr.ImportListInput:
		return lidarr.New(instance.Config).AddImportListContext(s.ctx, data)
	case *radarr.ImportListInput:
		return radarr.New(instance.Config).AddImportListContext(s.ctx, data)
	case *readarr.ImportListInput:
		return readarr.New(instance.Config).AddImportListContext(s.ctx, data)
	case *sonarr.ImportListInput:
		return sonarr.New(instance.Config).AddImportListContext(s.ctx, data)
	default:
		return nil, fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}
