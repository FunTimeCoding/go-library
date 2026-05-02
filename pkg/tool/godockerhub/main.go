package godockerhub

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/docker/hub"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	if c := environment.Optional(constant.LocatorEnvironment); c != "" {
		r := reporter.New("godockerhub", c, "", version)
		r.Start()
		defer func() { r.RecoverFlush(recover()) }()
	}

	pflag.String(
		argument.Image,
		"",
		"Image to list tags for (e.g. library/golang)",
	)
	monitor.ParseBind(version, gitHash, buildDate)
	image := viper.GetString(argument.Image)

	if image == "" {
		fmt.Println("--image is required")

		return
	}

	c := hub.New()
	tags := c.Tags(image)
	limit := len(tags)

	if limit > maxDisplay {
		limit = maxDisplay
	}

	for _, t := range tags[:limit] {
		updated := ""

		if t.LastUpdated != nil {
			updated = t.LastUpdated.Format("2006-01-02 15:04")
		}

		fmt.Printf("%-40s %s\n", t.Name, updated)
	}
}
