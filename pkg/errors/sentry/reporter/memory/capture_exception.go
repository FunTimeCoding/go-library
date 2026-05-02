package memory

func (m *Memory) CaptureException(e error) string {
	m.events = append(m.events, &Event{Error: e})

	return ""
}
