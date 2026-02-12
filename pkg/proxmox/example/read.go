package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/proxmox"
)

func Read() {
	p := proxmox.NewEnvironment()
	fmt.Printf("Version: %+v\n", p.Version())
	fmt.Printf("Cluster: %+v\n", p.Cluster())

	for _, n := range p.Nodes() {
		fmt.Printf("Node list: %+v\n", n)
		o := p.Node(n.Node)
		fmt.Printf("Node: %+v\n", o)

		for _, m := range p.Machines(o) {
			fmt.Printf("  Machine list: %+v\n", m)
			fmt.Printf(
				"  Machine: %+v\n",
				p.Machine(o, int(m.VMID)),
			)
		}

		for _, c := range p.Containers(o) {
			fmt.Printf("  Container list: %+v\n", c)
			fmt.Printf(
				"  Container: %+v\n",
				p.Container(o, int(c.VMID)),
			)
		}
	}

	for _, d := range p.Domains() {
		fmt.Printf("Domain list: %+v\n", d)
		fmt.Printf("Domain: %+v\n", p.Domain(d.Realm))
	}

	for _, g := range p.Groups() {
		fmt.Printf("Group list: %+v\n", g)
		fmt.Printf("Group: %+v\n", p.Group(g.GroupID))
	}

	for _, o := range p.Pools() {
		fmt.Printf("Pool list: %+v\n", o)
		fmt.Printf("Pool: %+v\n", p.Pool(o.PoolID))
	}

	for _, r := range p.Roles() {
		fmt.Printf("Role list: %+v\n", r)
		fmt.Printf("Role: %+v\n", p.Role(r.RoleID))
	}

	for _, u := range p.Users() {
		fmt.Printf("User list: %+v\n", u)
		fmt.Printf("User: %+v\n", p.User(u.UserID))
	}
}
