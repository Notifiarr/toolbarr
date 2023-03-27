package app

import (
	"fmt"
	"strings"

	"github.com/Notifiarr/toolbarr/pkg/starrs"
	wr "github.com/wailsapp/wails/v2/pkg/runtime"
)

// MigratorInfo is full of info about the current root folders.
type MigratorInfo struct {
	Table       AppTable         // Names of main and sub tables.
	RootFolders []string         // Configured Root Folders.
	Items       map[int64]string // Main table, ID=>Path
	Files       map[int64]string // Sub table, ID=>Path
	InvalidI    map[int64]string // Derived from items not located in root folders.
	InvalidF    map[int64]string // Derived from files not located in root folders.
	FoldersI    map[string]int   // Derived from both tables.
	FoldersF    map[string]int   // Derived from both tables.
	Recycle     string           // Recycle Bin
}

// AppTable is a simple name map.
type AppTable struct { //nolint:revive
	Items string
	Files string
}

// AppTables returns the primary item table for an app.
func AppTables(app string) AppTable { //nolint:revive
	return map[string]AppTable{
		"Lidarr":   {Items: "Artists", Files: "TrackFiles"},
		"Radarr":   {Items: "Movies", Files: "MovieFiles"},
		"Readarr":  {Items: "Authors", Files: "BookFiles"},
		"Sonarr":   {Items: "Series", Files: "EpisodeFiles"},
		"Whisparr": {Items: "Series", Files: "EpisodeFiles"},
	}[app]
}

// UpdateRootFolders is the response when changing or deleting root folders.
type UpdateRootFolders struct {
	Msg  string
	Info *MigratorInfo
}

// DeleteDBRootFolder removes a root folder.
func (a *App) DeleteDBRootFolder(instance *starrs.Instance, folder string) (*UpdateRootFolders, error) {
	a.log.Tracef("Call:DeleteDBRootFolder(%s)", folder)

	question := a.log.Translate(
		"Really delete %s root folder?\nPath: %s\nThis Action cannot be undone, and this application does not make backups.",
		instance.Name, folder)
	if !a.Ask(a.log.Translate("Delete Root Folder"), question) {
		return &UpdateRootFolders{}, nil
	}

	sql, err := starrs.NewSQL(a.log, instance)
	if err != nil {
		return nil, fmt.Errorf(a.log.Translate("Querying Sqlite3 DB: %v", err.Error()))
	}
	defer sql.Close()

	_, err = sql.Delete("RootFolders", fmt.Sprintf("Path=%q", folder))
	if err != nil {
		return nil, fmt.Errorf(a.log.Translate("Querying Sqlite3 DB: %v", err.Error()))
	}

	info, err := a.getMigratorInfo(instance)
	if err != nil {
		return nil, fmt.Errorf(a.log.Translate("Querying Sqlite3 DB: %v", err.Error()))
	}

	return &UpdateRootFolders{Msg: a.log.Translate("Success! Deleted root folder: %s", folder), Info: info}, nil
}

// UpdateDBRootFolder changes the path for a root folder. It updates all the item with the folder.
func (a *App) UpdateDBRootFolder(instance *starrs.Instance, oldPath, newPath string) (*UpdateRootFolders, error) {
	a.log.Tracef("Call:UpdateDBRootFolder(%s,%s)", oldPath, newPath)

	sql, err := starrs.NewSQL(a.log, instance)
	if err != nil {
		return nil, fmt.Errorf(a.log.Translate("Querying Sqlite3 DB: %v", err.Error()))
	}
	defer sql.Close()

	msg, err := a.updateDBRootFolder(AppTables(instance.App), sql, oldPath, newPath)
	if err != nil {
		return nil, fmt.Errorf(a.log.Translate("Querying Sqlite3 DB: %v", err.Error()))
	}

	info, err := a.getMigratorInfo(instance)
	if err != nil {
		return nil, fmt.Errorf(a.log.Translate("Querying Sqlite3 DB: %v", err.Error()))
	}

	return &UpdateRootFolders{Msg: msg, Info: info}, nil
}

func (a *App) updateDBRootFolder(table AppTable, sql *starrs.SQLConn, oldPath, newPath string) (string, error) {
	items, err := sql.ItemPaths(a.ctx, table.Items)
	if err != nil {
		return "", err
	}

	files, err := sql.ItemPaths(a.ctx, table.Files)
	if err != nil {
		return "", err
	}

	wr.EventsEmit(a.ctx, "DBitemTotals", map[string]int{
		table.Items: len(items),
		table.Files: len(files),
	})

	_, err = sql.Update("RootFolders", "Path", newPath, fmt.Sprintf("Path=%q", oldPath))
	if err != nil {
		return "", err
	}

	count, err := a.updateDBFilesRootFolder(sql, items, table.Items, oldPath, newPath)
	if err != nil {
		return "", err
	}

	tableName := table.Items
	msg := a.log.Translate("Success! Changed Root Folder from '%s' to '%s'.", oldPath, newPath) +
		" " + a.log.Translate("Updated %d rows in table %s.", count, tableName)

	count, err = a.updateDBFilesRootFolder(sql, files, table.Files, oldPath, newPath)
	if err != nil {
		return "", err
	}

	tableName = table.Files
	msg += " " + a.log.Translate("Updated %d rows in table %s.", count, tableName)

	return msg, nil
}

func (a *App) updateDBFilesRootFolder(
	sql *starrs.SQLConn,
	files map[int64]string,
	table, oldPath, newPath string,
) (int64, error) {
	var (
		counter      int64
		rowsAffected int64
	)

	for pathID, path := range files {
		counter++
		wr.EventsEmit(a.ctx, "DBfileCount", map[string]int64{table: counter})

		if !strings.HasPrefix(path, oldPath) {
			continue
		}

		update := strings.Replace(path, oldPath, newPath, 1)

		rep, err := sql.Update(table, "Path", update, "Id="+fmt.Sprint(pathID))
		if err != nil {
			return 0, err
		}

		c, _ := rep.RowsAffected()
		rowsAffected += c
	}

	return rowsAffected, nil
}

// GetMigratorInfo returns the current info about root folders.
func (a *App) GetMigratorInfo(instance *starrs.Instance) (*MigratorInfo, error) {
	a.log.Tracef("Call:GetMigratorInfo(%s)", instance.DBPath)

	info, err := a.getMigratorInfo(instance)
	if err != nil {
		return nil, fmt.Errorf(a.log.Translate("Querying Sqlite3 DB: %v", err.Error()))
	}

	return info, nil
}

func (a *App) getMigratorInfo(instance *starrs.Instance) (*MigratorInfo, error) {
	sql, err := starrs.NewSQL(a.log, instance)
	if err != nil {
		return nil, err
	}
	defer sql.Close()

	info := &MigratorInfo{
		Table:    AppTables(instance.App),
		FoldersI: make(map[string]int),
		FoldersF: make(map[string]int),
		InvalidI: make(map[int64]string),
		InvalidF: make(map[int64]string),
	}

	if info.Items, err = sql.ItemPaths(a.ctx, info.Table.Items); err != nil {
		return nil, err
	}

	if info.Files, err = sql.ItemPaths(a.ctx, info.Table.Files); err != nil {
		return nil, err
	}

	if info.RootFolders, err = sql.RootFolders(a.ctx); err != nil {
		return nil, err
	}

	if info.Recycle, err = sql.Recyclebin(a.ctx); err != nil {
		return nil, err
	}

	return info.derive(), nil
}

func (i *MigratorInfo) derive() *MigratorInfo {
	for k, v := range i.Items {
		if dirs, found := itemInRoot(v, i.RootFolders); !found {
			i.InvalidI[k] = v
		} else {
			for _, dir := range dirs {
				i.FoldersI[dir]++
			}
		}
	}

	for k, v := range i.Files {
		if dirs, found := itemInRoot(v, i.RootFolders); !found {
			i.InvalidF[k] = v
		} else {
			for _, dir := range dirs {
				i.FoldersF[dir]++
			}
		}
	}

	return i
}

func itemInRoot(item string, folders []string) ([]string, bool) {
	dirs := []string{}
	found := false

	for _, folder := range folders {
		if strings.HasPrefix(item, folder) {
			dirs = append(dirs, folder)
			found = true
		}
	}

	return dirs, found
}
