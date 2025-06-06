package main

import (
	"bufio"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/go_mod"
	"github.com/funtimecoding/go-library/pkg/runtime"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	var skip []string
	pflag.String(
		argument.Skip,
		"",
		"Directory skip matches, comma-separated",
	)
	argument.ParseBind()

	if s := viper.GetString(argument.Skip); s != "" {
		skip = split.Comma(s)
	}

	scan(
		argument.PositionalFallback(0, "."),
		skip,
		runtime.ExecutableVersion().String(),
	)
}

func exclude(
	path string,
	skip []string,
) bool {
	for _, s := range skip {
		if strings.Contains(filepath.Base(path), s) {
			return true
		}

		if strings.Contains(path, s) {
			return true
		}
	}

	return false
}

func scan(
	root string,
	skip []string,
	installedVersion string,
) {
	for _, e := range system.ReadDirectory(root) {
		if !e.IsDir() || exclude(e.Name(), skip) {
			continue
		}

		depth1 := system.Join(root, e.Name())
		checkMod(
			depth1,
			system.Join(".", e.Name()),
			installedVersion,
			root,
		)

		for _, e2 := range system.ReadDirectory(depth1) {
			if !e2.IsDir() {
				continue
			}

			relative := system.Join(e.Name(), e2.Name())

			if exclude(e2.Name(), skip) || exclude(relative, skip) {
				continue
			}

			checkMod(
				system.Join(depth1, e2.Name()),
				system.Join(".", relative),
				installedVersion,
				root,
			)
		}
	}
}

func checkMod(
	path string,
	display string,
	installedVersion string,
	root string,
) {
	p := system.Join(path, go_mod.ModFile)

	if system.FileExists(p) {
		v := readVersion(p)

		if v != installedVersion {
			fmt.Printf("%s %s\n", system.Join(root, display), v)
		}
	}
}

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
