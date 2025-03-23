package macos

import (
	"fmt"
	"github.com/andybrewer/mack"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func Alert(
	subject string,
	body string,
) {
	// Icon: informational, warning, critical
	icon := "informational"
	timeout := "5"
	response, e := mack.Alert(subject, body, icon, timeout)
	errors.PanicOnError(e)
	fmt.Printf("%#v", response)
}
