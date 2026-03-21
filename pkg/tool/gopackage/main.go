package gopackage

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/build"
	sentry "github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/system/join"
	"os"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	if c := environment.Optional(sentry.LocatorEnvironment); c != "" {
		r := reporter.New("gopackage", c, "", version)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
	}

	monitor.ParseBind(version, gitHash, buildDate)
	var runs int

	for _, name := range build.OutputDirectories() {
		fmt.Printf("Name: %s\n", name)
		outputDirectory := join.Relative(constant.Temporary, name)
		fmt.Printf("Output directory: %s\n", outputDirectory)

		for _, systemArchitecture := range build.SystemArchitectures() {
			if build.GuessBinaryPath(name, systemArchitecture) != "" {
				build.Archive(name, systemArchitecture)
				runs++
			}
		}
	}

	if runs == 0 {
		fmt.Println("No archive created")
		os.Exit(1)
	}
}
