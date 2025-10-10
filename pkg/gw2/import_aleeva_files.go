package gw2

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gw2/constant"
	"github.com/funtimecoding/go-library/pkg/system"
	systemConstant "github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/funtimecoding/go-library/pkg/system/join"
)

func ImportAleevaFiles() {
	downloads := join.Absolute(system.Home(), systemConstant.DownloadsPath)

	for _, file := range system.FilesMatching(
		downloads,
		constant.MembersPrefix,
	) {
		source := join.Absolute(downloads, file)
		destination := join.Absolute(
			systemConstant.Temporary,
			fmt.Sprintf(
				"%s%s%s",
				constant.AleevaPrefix,
				dateFromMembersFile(file),
				constant.NotationSuffix,
			),
		)
		fmt.Printf("Import: %s => %s\n", source, destination)
		system.Move(source, destination)
	}
}
