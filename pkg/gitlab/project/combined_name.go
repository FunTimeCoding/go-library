package project

import "github.com/funtimecoding/go-library/pkg/separator"

func (p *Project) CombinedName() string {
	return p.Namespace + separator.Slash + p.Name
}
