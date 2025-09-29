package pod

func (p *Pod) Label(key string) string {
	for k, v := range p.Raw.Labels {
		if k == key {
			return v
		}
	}

	return ""
}
