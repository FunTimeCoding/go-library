package relational

import (
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"strings"
)

func (d *Database) CreateUser(
	user string,
	password string,
) {
	if strings.Contains(user, "'") {
		log.Panicf("user cannot contain single quotes")
	}

	if d.UserExists(user) {
		log.Panicf("user '%s' already exists", user)
	}

	d.Execute(
		fmt.Sprintf(
			"CREATE USER %s WITH PASSWORD '%s'",
			pgx.Identifier{user}.Sanitize(),
			password,
		),
	)
}
