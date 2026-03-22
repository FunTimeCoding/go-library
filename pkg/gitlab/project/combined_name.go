package project

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/strings/separator"
)

func (p *Project) CombinedName() string {
	return fmt.Sprintf("%s%s%s", p.Namespace, separator.Slash, p.Name)
}
