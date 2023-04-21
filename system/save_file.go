package file

import (
	"bufio"
	"github.com/funtimecoding/go-library/errors"
	"os"
	"path"
)

func SaveFile(
	name string,
	text string,
) {
	directory, directoryFail := os.Getwd()
	errors.FatalOnError(directoryFail)

	file, fileFail := os.Create(path.Join(directory, name))
	errors.FatalOnError(fileFail)

	writer := bufio.NewWriter(file)
	_, writeFail := writer.WriteString(text)
	errors.FatalOnError(writeFail)

	errors.FatalOnError(writer.Flush())

	errors.FatalOnError(file.Close())
}
