package face

import (
	"github.com/funtimecoding/go-library/pkg/habitica/request"
	"github.com/funtimecoding/go-library/pkg/habitica/response"
)

type HabiticaSource interface {
	Allocate(stat string) (response.Stats, error)
	CreateTask(body request.CreateTaskBody) (response.Task, error)
	Cron() (response.CronResult, error)
	Equip(key string) error
	Score(taskID string, direction string) (response.ScoreResult, error)
	Tags() ([]response.Tag, error)
	Tasks(taskType string) ([]response.Task, error)
	UserGear() (response.Gear, error)
	UserStats() (response.Stats, error)
	MustCreateTask(body request.CreateTaskBody) response.Task
	MustCron() response.CronResult
	MustScore(taskID string, direction string) response.ScoreResult
	MustTags() []response.Tag
	MustTasks(taskType string) []response.Task
	MustUserStats() response.Stats
}
