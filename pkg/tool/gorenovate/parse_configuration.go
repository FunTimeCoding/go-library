package gorenovate

import (
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/system"
)

func parseConfiguration() *Configuration {
	if !system.FileExists(RenovateFile) {
		return nil
	}

	var result Configuration
	notation.DecodeBytesStrict(
		system.ReadBytesUnsafe(RenovateFile),
		&result,
		false,
	)

	return &result
}
