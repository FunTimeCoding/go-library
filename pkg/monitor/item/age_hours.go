package item

func (i *Item) ageHours() float64 {
	return i.Age().Hours()
}
