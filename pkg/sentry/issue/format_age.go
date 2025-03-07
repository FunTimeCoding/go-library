package issue

import "github.com/docker/go-units"

func (i *Issue) formatAge() string {
	return units.HumanDuration(i.Age())
}
