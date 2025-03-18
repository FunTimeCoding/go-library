package metric

func New(metric string) *Metric {
	result := &Metric{}

	if metric != "" {
		name, labels := parse(metric)
		result.name = name
		result.labels = labels
	}

	return result
}
