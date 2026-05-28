package face

import (
	"io"
	"os"
)

type Process interface {
	StdoutReader() io.ReadCloser
	StderrReader() io.ReadCloser
	Wait() error
	Signal(s os.Signal) error
}
