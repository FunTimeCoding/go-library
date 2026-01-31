package flat

func (f *Flat) Add(s string) {
	if s == "" {
		return
	}

	f.content = append(f.content, s)
}
