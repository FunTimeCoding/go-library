package result

import "fmt"

func (r *Result) Print() {
	if r.ExitCode != 0 || r.Error != nil {
		fmt.Printf("Error (%d): %v\n", r.ExitCode, r.Error)
	}

	if r.OutputString != "" {
		fmt.Println("Stdout:")
		fmt.Println(r.OutputString)
	}

	if r.ErrorString != "" {
		fmt.Println("Stderr:")
		fmt.Println(r.ErrorString)
	}
}
