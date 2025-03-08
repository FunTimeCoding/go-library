package file_report

import "fmt"

func (r *Report) ProcessConcerns(fix bool) {
	if r.HasConcerns() {
		if fix && r.Fix != nil {
			r.Fix()
		} else {
			for _, c := range r.Concerns {
				fmt.Printf("%s: %s\n", c.Text, c.Path)
			}
		}
	}
}
