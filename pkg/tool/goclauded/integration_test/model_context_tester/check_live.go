package model_context_tester

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/client"
)

func (s *Session) CheckLive() *client.CheckResponse {
	s.T.Helper()
	response, e := s.RestClient.GetCheckWithResponse(
		s.Context,
		&client.GetCheckParams{Session: s.UUID},
	)
	assert.FatalOnError(s.T, e)

	return response.JSON200
}
