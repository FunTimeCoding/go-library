package project

func (p *Project) CombinedName() string {
	return p.Namespace + "/" + p.Name
}
