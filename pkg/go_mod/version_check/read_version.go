package version_check

import (
	"bufio"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"regexp"
	"strings"
)

func readVersion(mod string) string {
	f := system.Open(mod)
	defer errors.PanicClose(f)
	r := regexp.MustCompile(`^go\s+(\d+\.\d+(?:\.\d+)?)`)
	s := bufio.NewScanner(f)

	for s.Scan() {
		if m := r.FindStringSubmatch(strings.TrimSpace(s.Text())); m != nil {
			return m[1]
		}
	}

	errors.PanicOnError(s.Err())

	return ""
}
