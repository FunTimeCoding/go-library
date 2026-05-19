package model_context_tester

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"

func (s *Session) Announce(
	name string,
	topic string,
) string {
	s.T.Helper()

	return s.MustCallTool(
		constant.Announce,
		map[string]any{
			constant.SessionName: name,
			constant.Topic:       topic,
		},
	)
}
