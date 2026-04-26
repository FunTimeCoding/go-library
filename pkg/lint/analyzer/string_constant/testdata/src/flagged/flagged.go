package flagged

import "fmt"

func Register() {
	fmt.Println("name")  // want `string literal "name" has constant constant.Name`
	fmt.Println("query") // want `string literal "query" has constant constant.Query`
}
