package helper

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/web/constant"
)

func Base(host string) string {
	return fmt.Sprintf("%s://%s", constant.SecureScheme, host)
}
