package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/ollama"
	"github.com/funtimecoding/go-library/pkg/ollama/constant/prompts"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
)

func ClassifyAlert() {
	o := ollama.NewEnvironment()

	if false {
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

	if true {
		p := prompts.ClassifyAlert()
		r := o.GenerateNotation(p.Render())
		fmt.Printf("Response: %+v\n", r)
		response := p.ParseResponse(r.Text)
		fmt.Printf("To save: %s", response)
		base := system.Join(constant.Temporary, "ollama")
		system.EnsurePathExists(base)
		system.SaveFile(
			system.Join(
				base,
				fmt.Sprintf(
					"response-%d.json",
					len(system.Files(base))+1,
				),
			),
			response,
		)
	}
}
