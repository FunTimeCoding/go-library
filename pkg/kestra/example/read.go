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

			for _, e := range k.Executions(n) {
				fmt.Printf("  Execution: %+v\n", e)

				for _, l := range k.Logs(e.Id) {
					fmt.Printf("    Log: %+v\n", l)
				}
			}
		}
	}

	if false {
		// 404
		fmt.Printf("Login: %+v\n", k.Login())
	}

	if false {
		// 404
		fmt.Println("Tenants")

		for _, t := range k.Tenants() {
			fmt.Printf("  %+v\n", t)
		}
	}

	if false {
		// panic: no value given for required property deprecated
		fmt.Println("Plugins")

		for _, f := range k.Plugins() {
			fmt.Printf("  %+v\n", f)
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
