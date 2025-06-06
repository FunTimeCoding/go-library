package goupload

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/build"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/gitlab/packages"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"path/filepath"
)

// TODO: What to do about Debian packages? Detect and upload to Nexus?

func Main() {
	locatorDefault := ""

	if s := os.Getenv(constant.InterfaceLocator); s != "" {
		locatorDefault = s
	}

	projectDefault := ""

	if s := os.Getenv(constant.ProjectIdentifier); s != "" {
		projectDefault = s
	}

	tagDefault := ""

	if s := os.Getenv(constant.CommitTag); s != "" {
		tagDefault = s
	}

	headerDefault := ""

	if s := os.Getenv(constant.JobToken); s != "" {
		headerDefault = fmt.Sprintf(
			"%s=%s",
			constant.JobTokenHeader,
			s,
		)
	}

	pflag.String(argument.Locator, locatorDefault, "GitLab API base URL")
	pflag.String(
		argument.Project,
		projectDefault,
		"Project ID to update to",
	)
	pflag.String(argument.Tag, tagDefault, "Git tag")
	pflag.String(
		argument.Header,
		headerDefault,
		"Header for authentication in key=value format",
	)
	argument.ParseBind()

	locator := argument.RequiredStringFlag(argument.Locator, 1)
	fmt.Printf("Locator: %s\n", locator)

	project := argument.RequiredStringFlag(argument.Project, 1)
	fmt.Printf("Project: %s\n", project)

	tag := argument.RequiredStringFlag(argument.Tag, 1)
	fmt.Printf("Tag: %s\n", tag)

	headers := build.Headers(viper.GetString(argument.Header))

	var runs int

	for _, name := range build.OutputDirectories() {
		for _, systemArchitecture := range build.SystemArchitectures() {
			if p := build.GuessArchivePath(
				name,
				systemArchitecture,
			); p != "" {
				runs++
				file := filepath.Base(p)
				fmt.Printf("Archive: %s\n", file)
				l := packages.UploadLink(locator, project, name, tag, file)
				fmt.Printf("Link: %s\n", l)
				status, body := web.PutFile(l, headers, system.ReadBytes(p))

				if status != http.StatusCreated {
					fmt.Printf("Upload failed: %d %s\n", status, body)

					os.Exit(1)
				}
			}
		}
	}

	if runs == 0 {
		fmt.Println("No archive uploaded")

		os.Exit(1)
	}
}
