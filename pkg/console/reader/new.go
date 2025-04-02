package reader

import (
	"bufio"
	"os"
)

func New() *Reader {
	return &Reader{reader: bufio.NewReader(os.Stdin)}
}
