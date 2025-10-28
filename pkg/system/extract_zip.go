package system

func ExtractZip(
	sourceFile string,
	destinationDirectory string,
) {
	for _, f := range ZipReader(sourceFile).File {
		ExtractZipFile(f, destinationDirectory)
	}
}
