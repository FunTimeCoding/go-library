package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/kestra"
)

func Read() {
	k := kestra.NewEnvironment()

	if true {
		fmt.Println("Namespaces")

		for _, n := range k.Namespaces() {
			fmt.Printf("Namespace: %+v\n", n)

			// TODO: Fails on tutorial namespace
			//  panic: json: cannot unmarshal number into Go struct field _Flow.variables of type map[string]interface {}
			for _, f := range k.Flows(n) {
				fmt.Printf("  Flow: %+v\n", f)
			}

			for _, e := range k.ExecutionFlows(n) {
				fmt.Printf("  Execution flow: %+v\n", e)
				fmt.Printf("  Execution: %+v\n", k.Execution(e.Id))
			}
		}
	}

	if false {
		// 404
		fmt.Println("Users")

		for _, f := range k.Users() {
			fmt.Printf("  %+v\n", f)
		}
	}
}
