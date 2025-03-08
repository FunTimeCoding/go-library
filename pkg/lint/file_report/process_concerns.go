package file_report

import "fmt"

func (r *Report) ProcessConcerns(fix bool) {
	if r.HasConcerns() {
		for _, c := range r.Concerns {
			fmt.Printf("Concern: %+v\n", c)
		}

		if fix && r.Fix != nil {
			r.Fix()
		}
	}
}
