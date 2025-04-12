package toast

import "fmt"

func (t *Toast) Format() string {
	return fmt.Sprintf("%d %s", t.Identifier, t.Text)
}
