package relational

import (
	"github.com/funtimecoding/go-library/pkg/postgres"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewEnvironment() *Database {
	return New(environment.Exit(postgres.LocatorEnvironment))
}
