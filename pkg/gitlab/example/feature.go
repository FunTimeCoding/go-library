package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gitlab"
)

func Feature() {
	g := gitlab.NewEnvironment()

	for _, f := range g.Features() {
		fmt.Printf("Feature: %+v\n", f)
	}

	for _, d := range g.FeatureDefinitions() {
		fmt.Printf("Definition: %+v\n", d)
	}
}
