package format

func (t *table) addRow(values []string) {
	for i, v := range values {
		if i < len(t.columns) && len(v) > t.columns[i].width {
			t.columns[i].width = len(v)
		}
	}

	t.rows = append(t.rows, values)
}
