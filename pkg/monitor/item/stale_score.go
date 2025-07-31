package item

func (i *Item) staleScore() float64 {
	if i.collector.Source.StaleMultiplier == 0 {
		return 0
	}

	staleHours := i.staleHours()
	threshold := i.collector.Source.StaleThreshold.Hours()

	if staleHours > threshold {
		return (staleHours - threshold) * i.collector.Source.StaleMultiplier
	}

	return 0
}
