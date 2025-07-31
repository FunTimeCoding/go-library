package item

func (i *Item) ageScore() float64 {
	return i.ageHours() * i.collector.Source.AgeMultiplier
}
