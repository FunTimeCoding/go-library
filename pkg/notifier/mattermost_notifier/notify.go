package mattermost_notifier

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

	n.mattermost.PostSimple(n.channel, key_value.Space(n.prefix, f))
}
