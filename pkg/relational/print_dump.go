package relational

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func (d *Database) PrintDump() {
	fmt.Println(notation.MarshallIndent(d.Dump()))
}
