package environment

import "sync"

type Environment struct {
	base    []string
	overlay map[string]string
	mutex   sync.RWMutex
}
