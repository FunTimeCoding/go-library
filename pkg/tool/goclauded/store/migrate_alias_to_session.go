package store

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"gorm.io/gorm"
)

func migrateAliasToSession(d *gorm.DB) {
	if !d.Migrator().HasTable(constant.Alias) {
		return
	}

	d.Exec(
		`
		UPDATE session SET
			alias = (SELECT name FROM alias WHERE alias.session_identifier = session.identifier),
			description = (SELECT description FROM alias WHERE alias.session_identifier = session.identifier)
		WHERE identifier IN (SELECT session_identifier FROM alias)
	`,
	)
	d.Exec(
		`
		INSERT INTO session (identifier, alias, description)
		SELECT session_identifier, name, description FROM alias
		WHERE session_identifier NOT IN (SELECT identifier FROM session)
	`,
	)
	d.Exec("DROP TABLE alias")
}
