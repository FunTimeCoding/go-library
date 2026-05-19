package chunk

import (
	"go/parser"
	"go/token"
)

func chunkGoSource(content string) []Chunk {
	set := token.NewFileSet()
	file, e := parser.ParseFile(set, "", content, parser.ParseComments)

	if e != nil {
		return nil
	}

	boundaries := collectBoundaries(set, file)

	if len(boundaries) == 0 {
		return nil
	}

	return splitAtBoundaries(content, boundaries)
}
