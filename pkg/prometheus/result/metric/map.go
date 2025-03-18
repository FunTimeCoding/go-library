package metric

func (m *Metric) Map() map[string]string {
	return m.labels
}
