package store

import "github.com/funtimecoding/go-library/pkg/tool/gopgd/config"

func (s *Store) Instance(name string) (*config.Instance, bool) {
	for i := range s.configuration.Instances {
		if s.configuration.Instances[i].Name == name {
			return &s.configuration.Instances[i], true
		}
	}

	return nil, false
}
