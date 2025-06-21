package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

func main() {
	// TODO: Design an object to hold a markdown runbook
	base := environment.Get("WIKI_PATH")

	for _, n := range system.Files(base) {
		fmt.Printf("File: %s\n", n)
		source := system.ReadBytes(system.Join(base, n))
		fmt.Println(
			console.Cyan(
				"%s",
				strings.PrefixMultiline(string(source), "> "),
			),
		)

		p := goldmark.DefaultParser()
		o := p.Parse(text.NewReader(source))

		if false {
			walk(&source, o)
		}

		if true {
			walkTree(&source, o)
		}

		if true {
			break
		}
	}
}

func walkTree(
	s *[]byte,
	n ast.Node,
) {
	errors.PanicOnError(
		ast.Walk(
			n,
			func(
				n ast.Node,
				entering bool,
			) (ast.WalkStatus, error) {
				if !entering {
					return ast.WalkContinue, nil
				}

				printNode(s, n)

				if n.Kind() == ast.KindText {
					textNode := n.(*ast.Text)
					segment := textNode.Segment
					fmt.Printf(
						"Text kind: %s\n",
						string(segment.Value(*s)),
					)
				}

				return ast.WalkContinue, nil
			},
		),
	)
}

func walk(
	s *[]byte,
	n ast.Node,
) {
	if n == nil {
		return
	}

	printNode(s, n)

	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		walk(s, c)
	}
}

func printNode(
	s *[]byte,
	n ast.Node,
) {
	switch o := n.(type) {
	case *ast.Document:
		v := value(s, o)

		if v == "" {
			return
		}

		fmt.Printf("Document: %s\n", v)
	case *ast.Heading:
		fmt.Printf("Heading: %s\n", value(s, o))
	case *ast.Paragraph:
		fmt.Printf("Paragraph: %s\n", value(s, o))
	case *ast.Text:
		fmt.Printf("Text: %s\n", string(o.Value(*s)))
	case *ast.FencedCodeBlock:
		fmt.Printf("FencedCodeBlock: %s\n", value(s, o))
	default:
		fmt.Printf("Unknown node type: %T\n", o)
	}
}

func value(
	s *[]byte,
	n ast.Node,
) string {
	return string(n.Lines().Value(*s))
}
