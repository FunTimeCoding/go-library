package log

import (
	"log"
	"os"
)

func NewGenericLogger() *log.Logger {
	return log.New(os.Stdout, "", 0)
}
