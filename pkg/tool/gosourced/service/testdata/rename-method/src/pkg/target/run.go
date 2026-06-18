package target

func Run() string {
	v := &Store{}

	return v.FindByName("test")
}
