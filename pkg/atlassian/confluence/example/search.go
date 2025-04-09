package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
)

func Search() {
	c := confluence.NewEnvironment()
	f := constant.Dense

	if true {
		for _, r := range c.Search(
			fmt.Sprintf("space=%s", constant.OperationsSpace),
		) {
			fmt.Println(r.Format(f))
		}
	}

	if false {
		// Working syntax examples
		c.Search("favorite=currentUser()")
		c.Search(`label IN ("ExampleLabel")`)
		c.Search("creator IN (currentUser())")
		c.Search("creator=currentUser()")
		c.Search("creator=currentUser()")
		c.Search("space=EXAMPLE")
		c.Search("space=EXAMPLE")

		// No result, watcher does not work, use favorite instead
		c.Search("watcher=currentUser()")
	}

	if false {
		fmt.Println("SearchKaos")

		for _, o := range c.SearchKaos("favorite=currentUser()") {
			fmt.Println(o.Format(f))
		}
	}

	if false {
		fmt.Println("SearchVirtomize")

		for _, o := range c.SearchVirtomize("favorite=currentUser()") {
			fmt.Println(o.Format(f))
		}
	}

	if false {
		fmt.Println("SearchTreminio")

		for _, o := range c.SearchTreminio("favorite=currentUser()") {
			fmt.Println(o.Title)
		}
	}
}
