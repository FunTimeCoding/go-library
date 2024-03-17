package random

import (
	"math/rand"
	"time"
)

func New() *rand.Source {
	s := rand.NewSource(time.Now().UnixNano())

	return &s
}
