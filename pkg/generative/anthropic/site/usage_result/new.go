package usage_result

import "time"

func New(
	sessionPercent int,
	sessionReset string,
	weeklyAllPercent int,
	weeklyAllReset string,
	weeklySonnetPercent int,
	weeklyDesignPercent int,
	routineRuns string,
	creditSpent string,
	creditReset string,
	creditPercent int,
) *Result {
	return &Result{
		SessionPercent:      sessionPercent,
		SessionReset:        sessionReset,
		WeeklyAllPercent:    weeklyAllPercent,
		WeeklyAllReset:      weeklyAllReset,
		WeeklySonnetPercent: weeklySonnetPercent,
		WeeklyDesignPercent: weeklyDesignPercent,
		RoutineRuns:         routineRuns,
		CreditSpent:         creditSpent,
		CreditReset:         creditReset,
		CreditPercent:       creditPercent,
		LastUpdated:         time.Now(),
	}
}
