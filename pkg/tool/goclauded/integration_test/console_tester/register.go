package console_tester

import "github.com/funtimecoding/go-library/pkg/tool/goclaude"

func (o *Tester) Register(session string) string {
	o.t.Helper()

	return goclaude.RunRegister(o.client, session)
}
