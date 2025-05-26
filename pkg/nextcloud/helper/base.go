package helper

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/web"
)

func Base(host string) string {
	return fmt.Sprintf("%s://%s", web.SecureScheme, host)
}
