package main

import (
	"fmt"
	"github.com/amikos-tech/chroma-go/pkg/api/v2"
	"github.com/funtimecoding/go-library/pkg/generative/chroma"
)

func main() {
	c := chroma.New()
	defer c.Close()

	for _, d := range c.Databases(v2.NewDefaultTenant()) {
		fmt.Printf("Database: %s\n", d.Name())
	}

	for _, o := range c.Collections() {
		fmt.Printf("Collection: %s\n", o.Name())
	}
}
