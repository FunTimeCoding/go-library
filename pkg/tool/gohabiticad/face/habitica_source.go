package face

import (
	"github.com/funtimecoding/go-library/pkg/habitica/cron"
	"github.com/funtimecoding/go-library/pkg/habitica/gear"
	"github.com/funtimecoding/go-library/pkg/habitica/request"
	"github.com/funtimecoding/go-library/pkg/habitica/score"
	"github.com/funtimecoding/go-library/pkg/habitica/statistic"
	"github.com/funtimecoding/go-library/pkg/habitica/tag"
	"github.com/funtimecoding/go-library/pkg/habitica/task"
)

type HabiticaSource interface {
	Allocate(stat string) (*statistic.Statistic, error)
	CreateTask(*request.CreateTaskBody) (*task.Task, error)
	Cron() (*cron.Cron, error)
	Equip(key string) error
	Score(
		taskID string,
		direction string,
	) (*score.Score, error)
	Tags() ([]*tag.Tag, error)
	Tasks(taskType string) ([]*task.Task, error)
	UserGear() (*gear.Gear, error)
	UserStats() (*statistic.Statistic, error)
	MustCreateTask(*request.CreateTaskBody) *task.Task
	MustCron() *cron.Cron
	MustScore(
		taskID string,
		direction string,
	) *score.Score
	MustTags() []*tag.Tag
	MustTasks(taskType string) []*task.Task
	MustUserStats() *statistic.Statistic
}
