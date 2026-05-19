package importer

import (
	"crypto/md5"
	"fmt"
)

func contentHash(content string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(content)))
}
