package example

func (e *Entry) FindByName(name string) bool {
	return e.Name == name
}
