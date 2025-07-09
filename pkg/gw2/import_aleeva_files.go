package gw2

import (
	"fmt"
	constant2 "github.com/funtimecoding/go-library/pkg/gw2/constant"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
)

func ImportAleevaFiles() {
	downloads := system.Join(system.Home(), constant.DownloadsPath)

	for _, file := range system.FilesMatching(
		downloads,
		constant2.MembersPrefix,
	) {
		source := system.Join(downloads, file)
		destination := system.Join(
			constant.Temporary,
			fmt.Sprintf(
				"%s%s%s",
				constant2.AleevaPrefix,
				dateFromMembersFile(file),
				constant2.NotationSuffix,
			),
		)
		fmt.Printf("Import: %s => %s\n", source, destination)
		system.Move(source, destination)
	}
}
