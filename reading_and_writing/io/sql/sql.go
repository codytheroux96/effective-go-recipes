/*
- Allowing access to data from an API
- Writing a function that accepts a table name and a slice of columns and returns SQL as a string
- AKA using bytes.Buffer to generate SQL
- Genereate the following output:
SELECT
	id,
	time,
	duration
FROM rides;
*/

package sql

import (
	"bytes"
	"fmt"
)

func genSelect(table string, columns []string) (string, error) {
	var buf bytes.Buffer

	if len(columns) == 0 {
		return "", fmt.Errorf("empty select")
	}

	fmt.Fprintln(&buf, "SELECT")
	for i, col := range columns {
		suffix := ","
		if i == len(columns)-1 {
			suffix = "" // this is because no trailing commas in SQL
		}
		fmt.Fprintf(&buf, "     %s%s\n", col, suffix)
	}

	fmt.Fprintf(&buf, "FROM %s;", table)
	return buf.String(), nil
}