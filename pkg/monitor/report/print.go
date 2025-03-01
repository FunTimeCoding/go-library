package report

import "fmt"

func (r *Report) Print() {
	fmt.Println(r.Encode())
}
