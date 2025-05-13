package scanner

import (
	"bufio"
	"os"
)

func New() *Scanner {
	return &Scanner{scanner: bufio.NewScanner(os.Stdin)}
}
