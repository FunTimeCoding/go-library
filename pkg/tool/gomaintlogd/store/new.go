package store

import "github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/option"

func New(o *option.Log) *Store {
	if o.PostgresLocator != "" {
		return NewPostgres(o.PostgresLocator)
	}

	if o.SQLitePath != "" {
		return NewLite(o.SQLitePath)
	}

	panic("set POSTGRES_LOCATOR or SQLITE_PATH")
}
