package wdutil

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/system/macos/constant"
	"github.com/funtimecoding/go-library/pkg/system/run"
)

func collect() *Result {
	r := run.New()
	r.Start("sudo", constant.Wdutil, constant.Info)

	if false {
		fmt.Printf("Status: %s\n", r.OutputString)
	}

	sequence := parseKey(r.OutputString, "Channel Sequence")

	if false {
		fmt.Printf("Sequence: %s\n", sequence)
	}

	return &Result{Sequence: sequence}
}
