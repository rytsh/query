package database

import (
	"context"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cast"
)

func Query(ctx context.Context, q string, db *sqlx.DB) (*sqlx.Rows, error) {
	return db.QueryxContext(ctx, q)
}

func Print(rows *sqlx.Rows) error {
	table := tablewriter.NewWriter(os.Stdout)

	headers, err := rows.Columns()
	if err != nil {
		return err
	}

	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(false)
	table.SetHeader(headers)

	for rows.Next() {
		values, err := rows.SliceScan()
		if err != nil {
			return err
		}

		valuesStr := make([]string, 0, len(values))
		for _, v := range values {
			valuesStr = append(valuesStr, cast.ToString(v))
		}

		table.Append(valuesStr)
	}

	table.Render()

	return nil
}
