package relational

import (
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
)

func (d *Database) GrantPrivileges(
	user string,
	dbname string,
) {
	if !d.UserExists(user) {
		log.Panicf("user '%s' does not exist", user)
	}

	if !d.DatabaseExists(dbname) {
		log.Panicf("database '%s' does not exist", dbname)
	}

	d.Execute(
		fmt.Sprintf(
			"GRANT ALL PRIVILEGES ON DATABASE %s TO %s",
			pgx.Identifier{dbname}.Sanitize(),
			pgx.Identifier{user}.Sanitize(),
		),
	)
}
