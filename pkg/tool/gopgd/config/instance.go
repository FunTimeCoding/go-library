package config

import "fmt"

type Instance struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func (i *Instance) Locator() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		i.User,
		i.Password,
		i.Host,
		i.Port,
		i.Database,
	)
}
