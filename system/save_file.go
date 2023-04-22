package system

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
	errors.PanicOnError(directoryFail)

	file, fileFail := os.Create(path.Join(directory, name))
	errors.PanicOnError(fileFail)

	writer := bufio.NewWriter(file)
	_, writeFail := writer.WriteString(text)
	errors.PanicOnError(writeFail)

	errors.PanicOnError(writer.Flush())

	errors.PanicOnError(file.Close())
}
