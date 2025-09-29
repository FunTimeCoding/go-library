package pod

func (p *Pod) Annotation(key string) string {
	for k, v := range p.Raw.Annotations {
		if k == key {
			return v
		}
	}

	return ""
}
