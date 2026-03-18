package gopg

import (
	"fmt"
	sentry "github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/relational"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	if c := environment.Optional(sentry.LocatorEnvironment); c != "" {
		r := reporter.New("gopg", c, "", version)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
	}

	monitor.ParseBind(version, gitHash, buildDate)
	a := relational.NewAdministrator()
	defer a.Close()

	for _, d := range a.Databases() {
		fmt.Printf("Database: %s\n", d)
	}

	user := "jdoe"
	password := "jdoe"
	dbname := "jdoe"

	if false {
		if !a.UserExists(user) {
			a.CreateUser(user, password)
		}

		if !a.DatabaseExists(dbname) {
			a.CreateDatabase(dbname)
		}

		if !a.HasConnectPrivilege(user, dbname) {
			a.GrantPrivileges(user, dbname)
		}
	}

	if false {
		a.CreateDatabase(dbname)
		a.CreateUser(user, password)
	}

	r := relational.NewEnvironment()
	defer r.Close()

	s := r.Client().Stat()
	fmt.Printf("Statistics: %+v\n", s)
}
