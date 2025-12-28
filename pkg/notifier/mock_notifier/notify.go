package mock_notifier

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/strings/join/key_value"
)

func (n *Notifier) Notify(
	f string,
	a ...any,
) {
	if len(a) > 0 {
		f = fmt.Sprintf(f, a...)
	}

	n.Notified = append(n.Notified, key_value.Space(n.prefix, f))
}
