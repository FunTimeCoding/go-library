package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/argument"
	"github.com/funtimecoding/go-library/pkg/ssh"
	"github.com/funtimecoding/go-library/pkg/system"
)

func main() {
	argument.ParseBind()
	n := argument.RequiredPositional(0, "NODE")
	fmt.Printf("Node: %s\n", n)

	if false {
		s := ssh.New(system.User().Username, n, false)
		defer s.Close()
		r := s.Run("ls")
		fmt.Printf("Run: %s\n", r.OutputString)
	}

	if true {
		s := ssh.NewWithFile(
			system.User().Username,
			n,
			system.Join(system.Home(), ".ssh", "ansible", "id_rsa_insecure"),
			false,
		)
		defer s.Close()
		r := s.Run("ls")
		fmt.Printf("Run: %s\n", r.OutputString)
	}
}
