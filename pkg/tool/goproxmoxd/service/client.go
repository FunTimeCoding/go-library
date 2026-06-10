package service

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/proxmox"
	"time"
)

func (s *Service) Client(instance string) (*proxmox.Client, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if c, okay := s.clients[instance]; okay {
		return c, nil
	}

	i, okay := s.Instance(instance)

	if !okay {
		return nil, fmt.Errorf("unknown instance: %s", instance)
	}

	var option []proxmox.Option

	if i.User != "" && i.Password != "" {
		option = append(option, proxmox.WithUser(i.User, i.Password))
	} else if i.Token != "" && i.Secret != "" {
		option = append(option, proxmox.WithToken(i.Token, i.Secret))
	}

	if i.Insecure {
		option = append(option, proxmox.WithInsecure())
	}

	if i.Timeout != "" {
		if d, e := time.ParseDuration(i.Timeout); e == nil {
			option = append(option, proxmox.WithTimeout(d))
		}
	}

	c := proxmox.New(i.Host, option...)
	s.clients[instance] = c

	return c, nil
}
