package goclaude

import (
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/goclaude/constant"
)

func sessionIdentifierFromEnvironment() string {
	if i := environment.Optional(
		constant.SessionIdentifierEnvironment,
	); i != "" {
		return i
	}

	return readHookInput().SessionID
}
