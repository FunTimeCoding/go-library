package multi_line

func (m *MultiLine) Add(line string) {
	if line == "" {
		return
	}

	m.lines = append(m.lines, line)
}
