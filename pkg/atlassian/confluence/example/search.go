package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
)

func Search() {
	c := confluence.NewEnvironment()
	f := constant.DenseFormat

	if false {
		// Working syntax examples
		fmt.Println(c.SearchBasic("favourite=currentUser()"))
		fmt.Println(c.SearchBasic(`label IN ("ExampleLabel")`))
		fmt.Println(c.SearchBasic("creator IN (currentUser())"))
		fmt.Println(c.SearchBasic("creator=currentUser()"))
		c.SearchVirtomize("creator=currentUser()")
		c.SearchVirtomize("space=EXAMPLE")
		c.SearchKaos("space=EXAMPLE")

		// No result, watcher does not work, use favorite instead
		c.SearchVirtomize("watcher=currentUser()")
	}

	if false {
		fmt.Println("SearchKaos")

		for _, o := range c.SearchKaos("favorite=currentUser()") {
			fmt.Println(o.Format(f))
		}
	}

	if true {
		fmt.Println("SearchTreminio")

		for _, o := range c.SearchTreminio("favourite=currentUser()") {
			fmt.Printf("  Search: %s\n", o.Title)
		}
	}
}
