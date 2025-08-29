package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/relational"
)

func main() {
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
