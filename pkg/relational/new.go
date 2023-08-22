package relational

import "context"

func New(locator string) *Database {
	d := Database{context: context.Background()}
	d.connect(locator)

	return &d
}
