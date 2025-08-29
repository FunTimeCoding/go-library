package relational

import (
	"github.com/funtimecoding/go-library/pkg/postgres"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func NewAdministrator() *Database {
	return New(environment.Get(postgres.AdministratorLocatorEnvironment))
}
