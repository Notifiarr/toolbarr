package app

import (
	"fmt"
	"path"
	"strings"

	"github.com/Notifiarr/toolbarr/pkg/starrs"
	wr "github.com/wailsapp/wails/v2/pkg/runtime"
)

// TableFileMap is a map of table name => item ID => item Path
// Used to display info or update paths on a table.
type TableFileMap map[starrs.TableColumn][]*starrs.Entry

// TableCountMap is a map of table name => root folder => count.
// Used to count items mapped to a root folder from a specific table.
type TableCountMap map[string]map[string]int

// MigratorInfo is full of info about the current root folders.
type MigratorInfo struct {
	Table       AppTable                   // Names of main and sub tables.
	RootFolders []string                   // Configured Root Folders.
	Invalid     map[string][]*starrs.Entry // Derived from items not located in root folders.
	Folders     TableCountMap              // Derived from both tables.
	Recycle     string                     // Recycle Bin
}

// AppTable is a simple name map.
type AppTable map[string]starrs.TableColumn //nolint:revive

// AppTables returns the primary item table for an app.
func AppTables(app string) AppTable { //nolint:revive
	return map[string]AppTable{
		"Lidarr": {
			"Artists":     {Table: "Artists", Column: "Path", Name: `""`},
			"TrackFiles":  {Table: "TrackFiles", Column: "Path", Name: "OriginalFilePath"},
			"ImportLists": {Table: "ImportLists", Column: "RootFolderPath", Name: "Name"},
		},
		"Radarr": {
			"Movies":      {Table: "Movies", Column: "Path", Name: `""`},
			"ImportLists": {Table: "ImportLists", Column: "RootFolderPath", Name: "Name"},
			"Collections": {Table: "Collections", Column: "RootFolderPath", Name: "Title"},
		},
		"Readarr": {
			"Authors":     {Table: "Authors", Column: "Path", Name: `""`},
			"BookFiles":   {Table: "BookFiles", Column: "Path", Name: "OriginalFilePath"},
			"ImportLists": {Table: "ImportLists", Column: "RootFolderPath", Name: "Name"},
		},
		"Sonarr": {
			"Series":      {Table: "Series", Column: "Path", Name: `""`},
			"ImportLists": {Table: "ImportLists", Column: "RootFolderPath", Name: "Name"},
		},
		"Whisparr": {
			"Series":      {Table: "Series", Column: "Path", Name: `""`},
			"ImportLists": {Table: "ImportLists", Column: "RootFolderPath", Name: "Name"},
		},
	}[app]
}

// UpdateRootFolders is the response when changing or deleting root folders.
type UpdateRootFolders struct {
	Msg  string
	Info *MigratorInfo
}

// DeleteDBRootFolder removes a root folder.
func (a *App) DeleteDBRootFolder(instance *starrs.Instance, folder string) (*UpdateRootFolders, error) {
	a.log.Tracef("Call:DeleteDBRootFolder(%s,%s)", instance.Name, folder)

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

	return a.returnMessage(sql, instance, a.log.Translate("Success! Deleted root folder: %s", folder))
}

// UpdateDBRootFolder changes the path for a root folder. It updates all the item with the folder.
func (a *App) UpdateDBRootFolder(instance *starrs.Instance, oldPath, newPath string) (*UpdateRootFolders, error) {
	a.log.Tracef("Call:UpdateDBRootFolder(%s,%s,%s)", instance.Name, oldPath, newPath)

	sql, err := starrs.NewSQL(a.log, instance)
	if err != nil {
		return nil, fmt.Errorf(a.log.Translate("Querying Sqlite3 DB: %v", err.Error()))
	}
	defer sql.Close()

	msg, err := a.updateDBRootFolder(AppTables(instance.App), sql, oldPath, newPath)
	if err != nil {
		return nil, fmt.Errorf(a.log.Translate("Querying Sqlite3 DB: %v", err.Error()))
	}

	return a.returnMessage(sql, instance, msg)
}

func (a *App) UpdateDBInvalidItems(
	instance *starrs.Instance,
	table string,
	newPath string,
	ids map[int64]bool,
) (*UpdateRootFolders, error) {
	a.log.Tracef("Call:UpdateDBInvalidItems(%s,%s,%d)", instance.Name, table, len(ids))

	if newPath == "" {
		return nil, fmt.Errorf(a.log.Translate("No path provided."))
	}

	sql, err := starrs.NewSQL(a.log, instance)
	if err != nil {
		return nil, fmt.Errorf(a.log.Translate("Querying Sqlite3 DB: %v", err.Error()))
	}
	defer sql.Close()

	counts := map[string]int{table: 0}

	for _, v := range ids {
		if v {
			counts[table]++
		}
	}

	wr.EventsEmit(a.ctx, "DBitemTotals", counts)

	column := AppTables(instance.App)[table]

	fn := a.updateDBInvalidBasePath // column.Column == "Path"
	if column.Column == "RootFolderPath" {
		fn = a.updateDBInvalidRootFolder
	}

	msg, err := fn(sql, &column, newPath, ids)
	if err != nil {
		return nil, fmt.Errorf(a.log.Translate("Querying Sqlite3 DB: %v", err.Error()))
	}

	return a.returnMessage(sql, instance, msg)
}

func (a *App) updateDBInvalidBasePath(
	sql *starrs.SQLConn,
	column *starrs.TableColumn,
	newPath string,
	ids map[int64]bool,
) (string, error) {
	current, err := sql.GetEntries(a.ctx, column)
	if err != nil {
		return "", err
	}

	counter := 0
	rows := int64(0)

	for _, entry := range current {
		if !ids[int64(entry.ID)] {
			continue
		}

		counter++
		wr.EventsEmit(a.ctx, "DBfileCount", map[string]int{column.Table: counter})

		res, err := sql.Update(column.Table, column.Column, newPath+path.Base(entry.Path), fmt.Sprint("Id=", entry.ID))
		if err != nil {
			return "", err
		}

		r, _ := res.RowsAffected()
		rows += r
	}

	return a.log.Translate("Updated %d rows in table %s.", rows, column.Table), nil
}

func (a *App) updateDBInvalidRootFolder(
	sql *starrs.SQLConn,
	column *starrs.TableColumn,
	newPath string,
	ids map[int64]bool,
) (string, error) {
	counter := 0
	rows := int64(0)

	for rowID, ok := range ids {
		if !ok {
			continue
		}

		counter++
		wr.EventsEmit(a.ctx, "DBfileCount", map[string]int{column.Table: counter})

		res, err := sql.Update(column.Table, column.Column, newPath, fmt.Sprint("Id=", rowID))
		if err != nil {
			return "", err
		}

		r, _ := res.RowsAffected()
		rows += r
	}

	return a.log.Translate("Updated %d rows in table %s.", rows, column.Table), nil
}

func (a *App) returnMessage(sql *starrs.SQLConn, instance *starrs.Instance, msg string) (*UpdateRootFolders, error) {
	info, err := a.getMigratorInfo(sql, instance)
	if err != nil {
		return nil, fmt.Errorf(a.log.Translate("Querying Sqlite3 DB: %v", err.Error()))
	}

	return &UpdateRootFolders{Msg: msg, Info: info}, nil
}

func (a *App) UpdateDBRecycleBin(instance *starrs.Instance, newPath string) (*UpdateRootFolders, error) {
	a.log.Tracef("Call:UpdateDBRecycleBin(%s,%s)", instance.Name, newPath)

	if newPath == "" {
		question := a.log.Translate("Really unset %s recycle bin?\n", instance.Name)
		if !a.Ask(a.log.Translate("Unset Recycle Bin"), question) {
			return &UpdateRootFolders{}, nil
		}
	}

	sql, err := starrs.NewSQL(a.log, instance)
	if err != nil {
		return nil, fmt.Errorf(a.log.Translate("Querying Sqlite3 DB: %v", err.Error()))
	}
	defer sql.Close()

	count, err := sql.UpdateRecyclebin(a.ctx, newPath)
	if err != nil {
		return nil, fmt.Errorf(a.log.Translate("Querying Sqlite3 DB: %v", err.Error()))
	}

	msg := a.log.Translate("Changed recycle bin path! Rows Updated: %d", count)
	if newPath == "" {
		msg = a.log.Translate("Unset recycle bin path! Rows Updated: %d", count)
	}

	return a.returnMessage(sql, instance, msg)
}

func (a *App) updateDBRootFolder(appTable AppTable, sql *starrs.SQLConn, oldPath, newPath string) (string, error) {
	counts, tables, err := a.getDBCounts(appTable, sql)
	if err != nil {
		return "", err
	}

	wr.EventsEmit(a.ctx, "DBitemTotals", counts)

	_, err = sql.Update("RootFolders", "Path", newPath, fmt.Sprintf("Path=%q", oldPath))
	if err != nil {
		return "", err
	}

	msg := a.log.Translate("Success! Changed Root Folder from '%s' to '%s'.", oldPath, newPath)

	for table, items := range tables {
		rows, err := a.updateDBFilesRootFolder(sql, items, table, oldPath, newPath)
		if err != nil {
			return "", err
		}

		msg += " " + a.log.Translate("Updated %d rows in table %s.", rows, table.Table)
	}

	return msg, nil
}

// getDBCounts returns "totals" to the UI so it can display progress bars.
// It also returns an item list so the calling function may loop the items to update them.
func (a *App) getDBCounts(table AppTable, sql *starrs.SQLConn) (map[string]int, TableFileMap, error) {
	output := make(TableFileMap)
	counts := map[string]int{}

	for name := range table {
		column := table[name]
		counts[name] = 0
		output[column] = []*starrs.Entry{}

		items, err := sql.GetEntries(a.ctx, &column)
		if err != nil {
			return nil, nil, err
		}

		output[column] = items
		counts[name] = len(items)
	}

	return counts, output, nil
}

func (a *App) updateDBFilesRootFolder(
	sql *starrs.SQLConn,
	files []*starrs.Entry,
	table starrs.TableColumn,
	oldPath, newPath string,
) (int64, error) {
	var (
		counter      int64
		rowsAffected int64
	)

	for _, entry := range files {
		counter++
		wr.EventsEmit(a.ctx, "DBfileCount", map[string]int64{table.Table: counter})

		if !strings.HasPrefix(entry.Path, oldPath) {
			continue
		}

		update := strings.Replace(entry.Path, oldPath, newPath, 1)

		rep, err := sql.Update(table.Table, table.Column, update, "Id="+fmt.Sprint(entry.ID))
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

	sql, err := starrs.NewSQL(a.log, instance)
	if err != nil {
		return nil, fmt.Errorf(a.log.Translate("Querying Sqlite3 DB: %v", err.Error()))
	}
	defer sql.Close()

	info, err := a.getMigratorInfo(sql, instance)
	if err != nil {
		return nil, fmt.Errorf(a.log.Translate("Querying Sqlite3 DB: %v", err.Error()))
	}

	return info, nil
}

func (a *App) getMigratorInfo(sql *starrs.SQLConn, instance *starrs.Instance) (*MigratorInfo, error) {
	table := AppTables(instance.App)
	info := &MigratorInfo{
		Table:   table,
		Folders: make(TableCountMap),
		Invalid: make(map[string][]*starrs.Entry),
	}

	var err error

	files := make(TableFileMap)

	for name := range table {
		column := table[name]

		files[column], err = sql.GetEntries(a.ctx, &column)
		if err != nil {
			return nil, err
		}
	}

	if info.RootFolders, err = sql.RootFolders(a.ctx); err != nil {
		return nil, err
	}

	if info.Recycle, err = sql.Recyclebin(a.ctx); err != nil {
		return nil, err
	}

	return info.derive(files), nil
}

func (i *MigratorInfo) derive(files TableFileMap) *MigratorInfo {
	for table, items := range files {
		if i.Folders[table.Table] == nil {
			i.Folders[table.Table] = make(map[string]int)
		}

		for _, entry := range items {
			dirs, found := itemInRoot(entry.Path, i.RootFolders)
			if !found {
				i.Invalid[table.Table] = append(i.Invalid[table.Table], entry)
			}

			for _, dir := range dirs {
				i.Folders[table.Table][dir]++
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
