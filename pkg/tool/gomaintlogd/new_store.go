package gomaintlogd

import (
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/option"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"
)

func newStore(o *option.Log) *store.Store {
	if o.PostgresLocator != "" {
		return store.NewPostgres(o.PostgresLocator)
	}

	if o.SQLitePath != "" {
		return store.NewLite(o.SQLitePath)
	}

	panic("set POSTGRES_LOCATOR or SQLITE_PATH")
}
