package metric

func (m *Metric) Label(key string) string {
	return m.labels[key]
}
