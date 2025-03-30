package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/ollama"
)

func AssessAlert() {
	o := ollama.New()
	alert := "ECCMemoryError"
	// DiskNearFull answer is consistent not broken
	// ECCMemoryError answer is inconsistent
	r := o.GenerateNotation(
		fmt.Sprintf(
			"Answer a JSON object with 2 strings: Reason and Answer. Assess in one short sentence first, then answer already-broken or not-yet-broken. Does this Prometheus alert indicate something is already-broken or not-yet-broken: %s\nAlready-broken examples: DiskFull, Timeout\nNot-yet-broken examples: DiskNearFull, HighLatency",
			alert,
		),
	)
	fmt.Println(r.Text)
	r.Print()
}
