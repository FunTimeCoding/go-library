package goprocessd

import (
	"crypto/sha256"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/tool/goprocessd/constant"
	"path/filepath"
)

func socketPath(procfilePath string) string {
	absolute := system.AbsolutePath(filepath.Dir(procfilePath))
	hash := sha256.Sum256([]byte(absolute))
	name := fmt.Sprintf("%x", hash[:8])

	return filepath.Join(
		constant.SocketDirectory(),
		fmt.Sprintf("%s.sock", name),
	)
}
