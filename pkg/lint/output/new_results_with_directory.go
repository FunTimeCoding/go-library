package output

import "fmt"

func NewResultsWithDirectory(directory string) *Results {
	if directory == "" {
		return &Results{}
	}

	if directory[len(directory)-1] != '/' {
		directory = fmt.Sprintf("%s/", directory)
	}

	return &Results{workingDirectory: directory}
}
