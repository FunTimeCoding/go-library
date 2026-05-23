package usage_result

import "time"

type Result struct {
	SessionPercent      int
	SessionReset        string
	WeeklyAllPercent    int
	WeeklyAllReset      string
	WeeklySonnetPercent int
	WeeklyDesignPercent int
	RoutineRuns         string
	CreditSpent         string
	CreditReset         string
	CreditPercent       int
	LastUpdated         time.Time
}
