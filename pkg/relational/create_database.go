package relational

import (
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"strings"
)

func (d *Database) CreateDatabase(
	dbname string,
	owner ...string,
) {
	if strings.Contains(dbname, "'") {
		log.Panicf("database name cannot contain single quotes")
	}

	if d.DatabaseExists(dbname) {
		log.Panicf("database '%s' already exists", dbname)
	}

	var sql string

	if len(owner) > 0 && owner[0] != "" {
		if !d.UserExists(owner[0]) {
			log.Panicf("owner '%s' does not exist", owner[0])
		}

		sql = fmt.Sprintf(
			"CREATE DATABASE %s OWNER %s",
			pgx.Identifier{dbname}.Sanitize(),
			pgx.Identifier{owner[0]}.Sanitize(),
		)
	} else {
		sql = fmt.Sprintf(
			"CREATE DATABASE %s",
			pgx.Identifier{dbname}.Sanitize(),
		)
	}

	d.Execute(sql)
}
