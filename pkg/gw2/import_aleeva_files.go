package gw2

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
)

func ImportAleevaFiles() {
	downloads := system.Join(system.Home(), constant.DownloadsPath)

	for _, file := range system.FilesMatching(downloads, MembersPrefix) {
		source := system.Join(downloads, file)
		destination := system.Join(
			constant.Temporary,
			fmt.Sprintf(
				"%s%s%s",
				AleevaPrefix,
				dateFromMembersFile(file),
				NotationSuffix,
			),
		)
		fmt.Printf("Import: %s => %s\n", source, destination)
		system.Move(source, destination)
	}
}
