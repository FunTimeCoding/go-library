package gofix

import (
	"fmt"
	"go/token"
	"os"
	"sort"
	"strings"
)

func applyEdits(
	fileSet *token.FileSet,
	edits []edit,
	diff bool,
) {
	grouped := groupByFile(fileSet, edits)

	for path, fileEdits := range grouped {
		sort.Slice(fileEdits, func(i, j int) bool {
			return fileEdits[i].offset > fileEdits[j].offset
		})

		original, e := os.ReadFile(path)

		if e != nil {
			fmt.Fprintf(os.Stderr, "read %s: %s\n", path, e)

			continue
		}

		modified := make([]byte, len(original))
		copy(modified, original)

		for _, fe := range fileEdits {
			modified = splice(modified, fe.offset, fe.length, []byte(fe.newText))
		}

		if diff {
			printDiff(path, original, modified)
		} else {
			e = os.WriteFile(path, modified, 0644)

			if e != nil {
				fmt.Fprintf(os.Stderr, "write %s: %s\n", path, e)
			}
		}
	}
}

type fileEdit struct {
	offset  int
	length  int
	newText string
}

func groupByFile(
	fileSet *token.FileSet,
	edits []edit,
) map[string][]fileEdit {
	result := make(map[string][]fileEdit)
	workingDirectory, _ := os.Getwd()

	for _, e := range edits {
		position := fileSet.Position(e.position)
		endPosition := fileSet.Position(e.end)
		path := position.Filename

		if !strings.HasPrefix(path, workingDirectory) {
			continue
		}

		result[path] = append(result[path], fileEdit{
			offset:  position.Offset,
			length:  endPosition.Offset - position.Offset,
			newText: e.newText,
		})
	}

	return result
}

func splice(
	b []byte,
	offset int,
	length int,
	replacement []byte,
) []byte {
	var result []byte
	result = append(result, b[:offset]...)
	result = append(result, replacement...)
	result = append(result, b[offset+length:]...)

	return result
}

func printDiff(path string, original []byte, modified []byte) {
	if string(original) == string(modified) {
		return
	}

	fmt.Printf("--- %s\n+++ %s\n", path, path)

	originalLines := splitLines(original)
	modifiedLines := splitLines(modified)

	for i := 0; i < len(originalLines) || i < len(modifiedLines); i++ {
		var o, m string

		if i < len(originalLines) {
			o = originalLines[i]
		}

		if i < len(modifiedLines) {
			m = modifiedLines[i]
		}

		if o != m {
			if o != "" {
				fmt.Printf("-%s\n", o)
			}

			if m != "" {
				fmt.Printf("+%s\n", m)
			}
		}
	}
}

func splitLines(b []byte) []string {
	var result []string
	start := 0

	for i, c := range b {
		if c == '\n' {
			result = append(result, string(b[start:i]))
			start = i + 1
		}
	}

	if start < len(b) {
		result = append(result, string(b[start:]))
	}

	return result
}
