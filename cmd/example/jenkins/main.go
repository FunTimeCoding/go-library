package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/jenkins"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
)

func main() {
	j := jenkins.NewEnvironment()

	if false {
		plugins := j.Plugins()
		fmt.Printf("Count: %d\n", len(plugins))
		for _, p := range plugins {
			fmt.Printf("%s %s\n", p.ShortName, p.Version)

			if !p.Enabled {
				fmt.Println("    DISABLED")
			}
		}
	}

	if false {
		plugins := j.Basic().Get(
			"/pluginManager/api/json?tree=plugins[shortName,version,hasUpdate,enabled]",
		)
		fmt.Println(plugins)
	}

	if false {
		root := system.Join(
			system.WorkingDirectory(),
			constant.Temporary,
			"jenkins",
		)
		system.EnsurePathExists(root)

		for _, o := range j.Jobs() {
			name := o.GetName()
			fmt.Println(name)
			b := j.Basic().Get(
				fmt.Sprintf("job/%s/config.xml", name),
			)
			system.SaveFile(
				system.Join(root, fmt.Sprintf("%s.xml", name)),
				b,
			)
		}
	}

	if false {
		j.JobsJSON()
	}
}
