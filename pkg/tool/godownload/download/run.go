package download

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/gitlab/packages"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/tool/common"
	"github.com/funtimecoding/go-library/pkg/tool/godownload/download/option"
	"os"
)

func Run(o *option.Download) {
	validate(o)
	client := gitlab.New(o.Host, o.Token)
	project := common.FindProjectOrExit(client, o.Owner, o.Repository)
	p := packages.FindVersionOrLatest(
		packages.FilterByName(
			client.Packages(project.Identifier, false),
			o.Package,
		),
		o.PackageVersion,
	)

	if p == nil {
		fmt.Printf("package not found: %s\n", o.Package)

		os.Exit(1)
	}

	if o.PackageVersion != LatestVersion && p.Version != o.PackageVersion {
		fmt.Printf(
			"warning: version %s not found, using latest\n",
			o.PackageVersion,
		)
	}

	f := packages.SystemMatchingFile(
		client.PackageFiles(project.Identifier, p.ID),
	)

	if f == nil {
		fmt.Printf("architecture and system match not found")

		os.Exit(1)
	}

	if !system.FileExists(o.Package) {
		if !system.FileExists(f.FileName) {
			packages.Download(
				packages.Link(o.Host, project, p, f, o.Verbose),
				o.Token,
				f.FileName,
			)
		}

		system.ExtractZip(f.FileName, system.WorkingDirectory())
		system.DeleteFile(f.FileName)
	} else {
		fmt.Printf("file already exists: %s\n", o.Package)
	}

	system.Executable(o.Package)

	if o.Output != DefaultOutput {
		system.Move(o.Package, system.Join(o.Output, o.Package))
	}
}
