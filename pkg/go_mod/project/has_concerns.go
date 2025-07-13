package project

func (p *Project) HasConcerns() bool {
	return len(p.concern) > 0
}
