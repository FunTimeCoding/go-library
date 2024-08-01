package console

import (
	"bufio"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
	"strings"
)

func AskConfirmation(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/N]: ", s)
		response, e := reader.ReadString('\n')
		errors.PanicOnError(e)
		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else {
			return false
		}
	}
}
