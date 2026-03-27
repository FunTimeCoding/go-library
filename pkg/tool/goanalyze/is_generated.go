package goanalyze

import (
	"go/ast"
	"strings"
)

func isGenerated(file *ast.File) bool {
	for _, group := range file.Comments {
		for _, comment := range group.List {
			if strings.Contains(comment.Text, "Code generated") &&
				strings.Contains(comment.Text, "DO NOT EDIT") {
				return true
			}
		}
	}

	return false
}
