package format

func newTable(headers []string) *table {
	result := make([]column, len(headers))

	for i, h := range headers {
		result[i] = column{header: h, width: len(h)}
	}

	return &table{columns: result}
}
