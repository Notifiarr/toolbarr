package starrs

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Notifiarr/toolbarr/pkg/logs"
	"github.com/jmoiron/sqlx"
)

// SQLConn is used to query a sqllite3 db.
type SQLConn struct {
	logger   *logs.Logger
	instance *Instance
	conn     *sqlx.DB
}

// Close must called when you're done with the sql.
func (s *SQLConn) Close() {
	_ = s.conn.Close()
}

type TableColumn struct {
	Table  string
	Column string
	Name   string
}

type Entry struct {
	ID   uint64
	Name *string
	Path string
}

// NewSQL provides a sql connection to make queries.
// Must call Close() when finished.
func NewSQL(logger *logs.Logger, instance *Instance) (*SQLConn, error) {
	conn, err := sqlx.Open("sqlite", instance.DBPath)
	if err != nil {
		return nil, err
	}

	return &SQLConn{
		logger:   logger,
		instance: instance,
		conn:     conn,
	}, nil
}

// Close must called when you're done with the sql.
func (s *SQLConn) Update(table, name, val, where string) (sql.Result, error) {
	query := fmt.Sprintf("UPDATE %s SET %s='%s' WHERE %s", table, name, Escape(val), where)
	s.logger.Debugf("Running Query: %s", query)

	return s.conn.Exec(query)
}

// Delete rows from a database table.
func (s *SQLConn) Delete(table, where string) (sql.Result, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s", table, where)
	s.logger.Debugf("Running Query: %s", query)

	return s.conn.Exec(query)
}

// RootFolders returns the root folders.
func (s *SQLConn) RootFolders(ctx context.Context) ([]string, error) {
	return s.RowsStringSlice(ctx, "SELECT Path FROM RootFolders")
}

// Recyclebin returns the recycle bin.
func (s *SQLConn) Recyclebin(ctx context.Context) (string, error) {
	return s.RowString(ctx, "SELECT Value FROM Config WHERE Key = 'recyclebin'")
}

func (s *SQLConn) UpdateRecyclebin(ctx context.Context, path string) (int64, error) {
	query := fmt.Sprintf("INSERT INTO Config (Key, Value) VALUES ('recyclebin', %[1]q)"+
		" ON CONFLICT(Key) DO UPDATE SET Value = %[1]q", Escape(path))

	if path == "" {
		query = "DELETE FROM Config WHERE Key='recyclebin'"
	}

	s.logger.Debugf("Running Query: %s", query)

	rows, err := s.conn.ExecContext(ctx, query)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", query, err)
	}

	count, _ := rows.RowsAffected()

	return count, nil
}

// ItemPaths returns the ID=>Path mapping from any table.
func (s *SQLConn) ItemPaths(ctx context.Context, table, column string) (map[int64]string, error) {
	return s.RowsIDString(ctx, "SELECT Id, "+column+" FROM "+table)
}

// ItemPaths returns the ID=>Path mapping from any table.
func (s *SQLConn) GetEntries(ctx context.Context, tcd *TableColumn) ([]*Entry, error) {
	sql := fmt.Sprintf("SELECT Id AS id, %s AS name, %s As path FROM %s", tcd.Name, tcd.Column, tcd.Table)
	s.logger.Debugf("Running Query: %s", sql)

	rows, err := s.conn.QueryxContext(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", sql, err)
	}
	defer rows.Close()

	output := []*Entry{}

	for rows.Next() {
		var row Entry
		if err = rows.StructScan(&row); err != nil {
			return nil, fmt.Errorf("%s: %w", sql, err)
		}

		output = append(output, &row)
	}

	return output, nil
}

// TableCount returns the row count for a table.
func (s *SQLConn) TableCount(ctx context.Context, table string) (int64, error) {
	return s.RowInt64(ctx, "SELECT count(1) FROM "+table)
}

// RowsIDString returns 2 columns from N rows as a map of ID (int64) => item (string).
func (s *SQLConn) RowsIDString(ctx context.Context, sql string) (map[int64]string, error) {
	output := make(map[int64]string)

	rows, err := s.conn.QueryContext(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", sql, err)
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", sql, err)
	}

	for rows.Next() {
		row := struct {
			ID  int64
			Col string
		}{}
		if err := rows.Scan(&row.ID, &row.Col); err != nil {
			return nil, fmt.Errorf("%s: %w", sql, err)
		}

		output[row.ID] = row.Col
	}

	return output, nil
}

// RowsStringSlice returns 1 column from N rows as a string slice.
func (s *SQLConn) RowsStringSlice(ctx context.Context, sql string) ([]string, error) {
	slice := []string{}

	rows, err := s.conn.QueryContext(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", sql, err)
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", sql, err)
	}

	for rows.Next() {
		var text string
		if err := rows.Scan(&text); err != nil {
			return nil, fmt.Errorf("%s: %w", sql, err)
		}

		slice = append(slice, text)
	}

	return slice, nil
}

// RowString returns 1 column from 1 row as a string.
func (s *SQLConn) RowString(ctx context.Context, sql string) (string, error) {
	rows, err := s.conn.QueryContext(ctx, sql)
	if err != nil {
		return "", fmt.Errorf("%s: %w", sql, err)
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return "", fmt.Errorf("%s: %w", sql, err)
	}

	if rows.Next() {
		var text string
		if err := rows.Scan(&text); err != nil {
			return "", fmt.Errorf("%s: %w", sql, err)
		}

		return text, nil
	}

	return "", nil
}

// RowInt64 returns 1 column from 1 row as an Int64.
func (s *SQLConn) RowInt64(ctx context.Context, sql string) (int64, error) {
	rows, err := s.conn.QueryContext(ctx, sql)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", sql, err)
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return 0, fmt.Errorf("%s: %w", sql, err)
	}

	if rows.Next() {
		var i int64
		if err := rows.Scan(&i); err != nil {
			return 0, fmt.Errorf("%s: %w", sql, err)
		}

		return i, nil
	}

	return 0, nil
}

// Escape tries to make queries values safer.
func Escape(sql string) string {
	dest := make([]byte, 0, 2*len(sql)) //nolint:gomnd

	for pos := 0; pos < len(sql); pos++ {
		switch sql[pos] {
		case 0: /* Must be escaped for 'mysql' */
			dest = append(dest, '\\', '0')
		case '\n': /* Must be escaped for logs */
			dest = append(dest, '\\', '\n')
		case '\r':
			dest = append(dest, '\\', '\r')
		case '\\':
			dest = append(dest, '\\', '\\')
		case '\'':
			dest = append(dest, '\'', '\'')
		case '"': /* Better safe than sorry */
			dest = append(dest, '\\', '"')
		case '\032': // 十进制26,八进制32,十六进制1a, /* This gives problems on Win32 */
			dest = append(dest, '\\', 'Z')
		default:
			dest = append(dest, sql[pos])
		}
	}

	return string(dest)
}
