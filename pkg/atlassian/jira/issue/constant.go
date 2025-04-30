package issue

import "regexp"

const (
	NoStatus      = "no status"
	NoDescription = "no description"
	NoLink        = "no link"
	NoAge         = "no age"
	NoScore       = "no score"

	BlockedBy = "is blocked by"

	NilValue = "nil value"

	UnknownField = "unknown field"
	UnknownValue = "unknown value"

	DefaultPriority = "Default"
)

// Issue type
const (
	BugType     = "Bug"
	EpicType    = "Epic"
	StoryType   = "Story"
	TaskType    = "Task"
	SubTaskType = "Sub-task"
)

var KeyMatch = regexp.MustCompile(`[A-Z]+-\d+`)
