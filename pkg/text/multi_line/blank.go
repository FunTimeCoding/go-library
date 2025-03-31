package multi_line

func (m *MultiLine) Blank() {
	m.lines = append(m.lines, "")
}
