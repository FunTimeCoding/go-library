package relational

import (
	"github.com/funtimecoding/go-library/pkg/relational/postgres"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Database {
	return New(environment.Required(postgres.LocatorEnvironment))
}
