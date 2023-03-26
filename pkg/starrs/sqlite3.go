package starrs

import (
	"context"
	"database/sql"
	"fmt"
)

func getSQLLiteRowStringSlice(ctx context.Context, conn *sql.DB, sql string) ([]string, error) {
	slice := []string{}

	rows, err := conn.QueryContext(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("%s: running DB query: %w", sql, err)
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: reading DB rows: %w", sql, err)
	}

	for rows.Next() {
		var text string
		if err := rows.Scan(&text); err != nil {
			return nil, fmt.Errorf("%s: reading DB query: %w", sql, err)
		}

		slice = append(slice, text)
	}

	return slice, nil
}

func getSQLLiteRowString(ctx context.Context, conn *sql.DB, sql string) (string, int) {
	text := "<no data returned>"
	count := 0

	rows, err := conn.QueryContext(ctx, sql)
	if err != nil {
		return fmt.Sprintf("%s: running DB query: %v", text, err), 0
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return fmt.Sprintf("%s: reading DB rows: %v", text, err), 0
	}

	for rows.Next() {
		if err := rows.Scan(&text); err != nil {
			return fmt.Sprintf("%s: reading DB query: %v", text, err), 0
		}

		count++
	}

	return text, count
}

// func getSQLLiteRowInt64(
// 	ctx context.Context,
// 	conn *sql.DB,
// 	sql string,
// ) int64 {
// 	rows, err := conn.QueryContext(ctx, sql)
// 	if err != nil {
// 		return 0
// 	}
// 	defer rows.Close()

// 	if err := rows.Err(); err != nil {
// 		return 0
// 	}

// 	if rows.Next() {
// 		var i int64
// 		if err := rows.Scan(&i); err != nil {
// 			return 0
// 		}

// 		return i
// 	}

// 	return 0
// }
