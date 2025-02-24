package packages

import "fmt"

func UploadLink(
	locator string,
	project string,
	name string,
	tag string,
	file string,
) string {
	return fmt.Sprintf(
		"%s/projects/%s/packages/generic/%s/%s/%s",
		locator,
		project,
		name,
		tag,
		file,
	)
}
