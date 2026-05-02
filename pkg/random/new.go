package random

import (
	"math/rand"
	"time"
)

func New() *rand.Source {
	return new(rand.NewSource(time.Now().UnixNano()))
}
