package starrs

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Notifiarr/toolbarr/pkg/mnd"
	wr "github.com/wailsapp/wails/v2/pkg/runtime"
)

/* Root Folders filesystem paths migrator for sqlite3 db. */

// TableColumn is used to pull paths from specific columns in specific tables.
type TableColumn struct {
	Table  string
	Column string
	Name   string
}

// TableFileMap is a map of table name => item ID => item Path
// Used to display info or update paths on a table.
// XXX: This is too complicated.
type TableFileMap map[TableColumn][]*Entry

// TableCountMap is a map of table name => root folder => count.
// Used to count items mapped to a root folder from a specific table.
type TableCountMap map[string]map[string]int

// MigratorInfo is full of info about the current root folders.
// This data gets sent to the front end any time there is a change.
type MigratorInfo struct {
	Table       AppTable            // Names of main and sub tables.
	RootFolders []string            // Configured Root Folders.
	Invalid     map[string][]*Entry // Derived from items not located in root folders.
	Folders     TableCountMap       // Derived from both tables.
	Recycle     string              // Recycle Bin
}

// AppTable is a simple name map.
type AppTable map[string]TableColumn

// AppTables returns the primary item table for an app.
func AppTables(app string) AppTable {
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

// RootFolders is the response to the front end when changing or deleting root folders.
type RootFolders struct {
	Msg  string
	Info *MigratorInfo
}

// MigratorInfo returns the current info about root folders.
func (s *Starrs) MigratorInfo(config *AppConfig) (*MigratorInfo, error) {
	s.log.Tracef("Call:GetMigratorInfo(%s)", config.DBPath)

	sql, err := s.newSQL(config)
	if err != nil {
		return nil, err
	}
	defer sql.Close()

	info, err := s.migratorInfo(sql, config)
	if err != nil {
		return nil, fmt.Errorf(s.log.Translate("Querying Sqlite3 DB: %v", err.Error()))
	}

	return info, nil
}

// DeleteRootFolder removes a root folder.
func (s *Starrs) DeleteRootFolder(config *AppConfig, folder string) (*RootFolders, error) {
	s.log.Tracef("Call:DeleteDBRootFolder(%s,%s)", config.Name, folder)

	question := s.log.Translate(
		"Really delete %s root folder?\nPath: %s\nThis Action cannot be undone, and this application does not make backups.",
		config.Name, folder)
	if !s.app.Ask(s.log.Translate("Delete Root Folder"), question) {
		return &RootFolders{}, nil
	}

	sql, err := s.newSQL(config)
	if err != nil {
		return nil, err
	}
	defer sql.Close()

	_, err = sql.Delete("RootFolders", fmt.Sprintf("Path='%s'", folder))
	if err != nil {
		return nil, fmt.Errorf(s.log.Translate("Querying Sqlite3 DB: %v", err.Error()))
	}

	return s.returnMessage(sql, config, s.log.Translate("Success! Deleted root folder: %s", folder))
}

// UpdateRootFolder changes the path for a root folder. It updates all the items with the folder.
func (s *Starrs) UpdateRootFolder(config *AppConfig, oldPath, newPath string) (*RootFolders, error) {
	s.log.Tracef("Call:UpdateDBRootFolder(%s,%s,%s)", config.Name, oldPath, newPath)

	sql, err := s.newSQL(config)
	if err != nil {
		return nil, err
	}
	defer sql.Close()

	msg, err := s.updateRootFolder(AppTables(config.App), sql, oldPath, newPath)
	if err != nil {
		return nil, fmt.Errorf(s.log.Translate("Querying Sqlite3 DB: %v", err.Error()))
	}

	return s.returnMessage(sql, config, msg)
}

func (s *Starrs) UpdateRecycleBin(config *AppConfig, newPath string) (*RootFolders, error) {
	s.log.Tracef("Call:UpdateDBRecycleBin(%s,%s)", config.Name, newPath)

	if newPath == "" {
		question := s.log.Translate("Really unset %s recycle bin?\n", config.Name)
		if !s.app.Ask(s.log.Translate("Unset Recycle Bin"), question) {
			return &RootFolders{}, nil
		}
	}

	sql, err := s.newSQL(config)
	if err != nil {
		return nil, err
	}
	defer sql.Close()

	count, err := sql.UpdateRecyclebin(s.ctx, newPath)
	if err != nil {
		return nil, fmt.Errorf(s.log.Translate("Querying Sqlite3 DB: %v", err.Error()))
	}

	msg := s.log.Translate("Changed recycle bin path! Rows Updated: %d", count)
	if newPath == "" {
		msg = s.log.Translate("Unset recycle bin path! Rows Updated: %d", count)
	}

	return s.returnMessage(sql, config, msg)
}

func (s *Starrs) UpdateInvalidItems(
	config *AppConfig,
	table string,
	newPath string,
	ids map[int64]bool,
) (*RootFolders, error) {
	s.log.Tracef("Call:UpdateDBInvalidItems(%s,%s,%d)", config.Name, table, len(ids))

	if newPath == "" {
		return nil, fmt.Errorf(s.log.Translate("No path provided."))
	}

	sql, err := s.newSQL(config)
	if err != nil {
		return nil, err
	}
	defer sql.Close()

	counts := map[string]int{table: 0}

	for _, v := range ids {
		if v {
			counts[table]++
		}
	}

	wr.EventsEmit(s.ctx, "DBitemTotals", counts)

	column := AppTables(config.App)[table]

	fn := s.updateInvalidBasePath // column.Column == "Path"
	if column.Column == "RootFolderPath" {
		fn = s.updateInvalidRootFolders
	}

	msg, err := fn(sql, &column, newPath, ids)
	if err != nil {
		return nil, fmt.Errorf(s.log.Translate("Querying Sqlite3 DB: %v", err.Error()))
	}

	return s.returnMessage(sql, config, msg)
}

func (s *Starrs) returnMessage(sql *sqlConn, config *AppConfig, msg string) (*RootFolders, error) {
	info, err := s.migratorInfo(sql, config)
	if err != nil {
		return nil, fmt.Errorf(s.log.Translate("Querying Sqlite3 DB: %v", err.Error()))
	}

	return &RootFolders{Msg: msg, Info: info}, nil
}

// updateInvalidBasePath handles one column type (Path).
func (s *Starrs) updateInvalidBasePath(
	sql *sqlConn,
	column *TableColumn,
	newPath string,
	ids map[int64]bool,
) (string, error) {
	current, err := sql.GetEntries(s.ctx, column)
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
		wr.EventsEmit(s.ctx, "DBfileCount", map[string]int{column.Table: counter})

		updatePath := newPath + filepath.Base(FromSlash(entry.Path))
		idStr := "Id=" + strconv.FormatUint(entry.ID, mnd.Base10)

		res, err := sql.Update(column.Table, column.Column, updatePath, idStr)
		if err != nil {
			return "", err
		}

		r, _ := res.RowsAffected()
		rows += r
	}

	return s.log.Translate("Updated %d rows in table %s.", rows, column.Table), nil
}

// updateInvalidRootFolders handles the other column type (RootFolderPath).
func (s *Starrs) updateInvalidRootFolders(
	sql *sqlConn,
	column *TableColumn,
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
		wr.EventsEmit(s.ctx, "DBfileCount", map[string]int{column.Table: counter})

		res, err := sql.Update(column.Table, column.Column, newPath, fmt.Sprint("Id=", rowID))
		if err != nil {
			return "", err
		}

		r, _ := res.RowsAffected()
		rows += r
	}

	return s.log.Translate("Updated %d rows in table %s.", rows, column.Table), nil
}

func (s *Starrs) updateRootFolder(appTable AppTable, sql *sqlConn, oldPath, newPath string) (string, error) {
	counts, tables, err := s.getDBCounts(appTable, sql)
	if err != nil {
		return "", err
	}

	wr.EventsEmit(s.ctx, "DBitemTotals", counts)

	_, err = sql.Update("RootFolders", "Path", newPath, fmt.Sprintf("Path='%s'", oldPath))
	if err != nil {
		return "", err
	}

	msg := s.log.Translate("Success! Changed Root Folder from '%s' to '%s'.", oldPath, newPath)

	for table, items := range tables {
		rows, err := s.updateFilesRootFolder(sql, items, table, trailingSlash(oldPath), newPath)
		if err != nil {
			return "", err
		}

		msg += " " + s.log.Translate("Updated %d rows in table %s.", rows, table.Table)
	}

	return msg, nil
}

// getDBCounts returns "totals" to the UI so it can display progress bars.
// It also returns an item list so the calling function may loop the items to update them.
func (s *Starrs) getDBCounts(table AppTable, sql *sqlConn) (map[string]int, TableFileMap, error) {
	output := make(TableFileMap)
	counts := map[string]int{}

	for name := range table {
		column := table[name]
		counts[name] = 0
		output[column] = []*Entry{}

		items, err := sql.GetEntries(s.ctx, &column)
		if err != nil {
			return nil, nil, err
		}

		output[column] = items
		counts[name] = len(items)
	}

	return counts, output, nil
}

func (s *Starrs) updateFilesRootFolder(
	sql *sqlConn,
	files []*Entry,
	table TableColumn,
	oldPath, newPath string,
) (int64, error) {
	var (
		counter      int64
		rowsAffected int64
	)

	for _, entry := range files {
		counter++
		wr.EventsEmit(s.ctx, "DBfileCount", map[string]int64{table.Table: counter})

		if !strings.HasPrefix(entry.Path, oldPath) {
			s.log.Debugf("Skipping path (wrong prefix): %s", entry.Path)
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

func (s *Starrs) migratorInfo(sql *sqlConn, config *AppConfig) (*MigratorInfo, error) {
	table := AppTables(config.App)
	info := &MigratorInfo{
		Table:   table,
		Folders: make(TableCountMap),
		Invalid: make(map[string][]*Entry),
	}

	var err error

	files := make(TableFileMap)

	for name := range table {
		column := table[name]

		files[column], err = sql.GetEntries(s.ctx, &column)
		if err != nil {
			return nil, err
		}
	}

	if info.RootFolders, err = sql.RootFolders(s.ctx); err != nil {
		return nil, err
	}

	if info.Recycle, err = sql.Recyclebin(s.ctx); err != nil {
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
		if strings.HasPrefix(item, trailingSlash(folder)) {
			dirs = append(dirs, folder)
			found = true
		}
	}

	return dirs, found
}

func trailingSlash(str string) string {
	if strings.HasSuffix(str, "/") || strings.HasSuffix(str, `\`) {
		return str
	}

	return str + pickSlash(str)
}

func pickSlash(str string) string {
	if len(str) == 0 || str[0] == '/' {
		return "/"
	}

	forward := strings.Count(str, `/`)
	reverse := strings.Count(str, `\`)

	if reverse > forward {
		return `\`
	}

	return "/"
}

// FromSlash returns the result of replacing each slash character in path with a separator character.
func FromSlash(path string) string {
	const sep = filepath.Separator
	return strings.ReplaceAll(strings.ReplaceAll(path, `\`, string(sep)), "/", string(sep))
}
