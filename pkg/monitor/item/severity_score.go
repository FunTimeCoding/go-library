package item

func (i *Item) severityScore() float64 {
	if weight, okay := i.collector.Source.SeverityWeights[i.Severity]; okay {
		return weight
	}

	return 0.0
}
