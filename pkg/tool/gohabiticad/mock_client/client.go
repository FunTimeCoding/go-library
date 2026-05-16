package mock_client

import (
	"github.com/funtimecoding/go-library/pkg/habitica/gear"
	"github.com/funtimecoding/go-library/pkg/habitica/statistic"
	"github.com/funtimecoding/go-library/pkg/habitica/tag"
	"github.com/funtimecoding/go-library/pkg/habitica/task"
)

type Client struct {
	tasks []*task.Task
	tags  []*tag.Tag
	stats *statistic.Statistic
	gear  *gear.Gear
}
