package main

import (
	"github.com/funtimecoding/go-library/pkg/text/dictionary/check/missing"
	"github.com/funtimecoding/go-library/pkg/text/dictionary/check/order"
)

func main() {
	order.Run()
	missing.Run()
}
