package main

import "github.com/funtimecoding/go-library/pkg/provision/ansible"

func main() {
	c := ansible.NewEnvironment()
	c.Execute(c.Playbook("site.yml"))
}
