package memory

func (m *Memory) CaptureWithContext(
	e error,
	_ string,
	context map[string]any,
) string {
	m.events = append(m.events, &Event{Error: e, Context: context})

	return ""
}
