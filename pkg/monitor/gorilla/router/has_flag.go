package router

func (r *Router) HasFlag(identifier string) bool {
	for _, l := range r.Flag {
		if l.Identifier == identifier {
			return true
		}
	}

	return false
}
