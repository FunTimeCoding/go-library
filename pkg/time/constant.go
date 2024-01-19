package time

const (
	DateMinute = "2006-01-02 15:04"
	DateSecond = "2006-01-02 15:04:05"
)

const (
	HourInMinutes  = int(60)
	DayInMinutes   = HourInMinutes * 24
	WeekInMinutes  = DayInMinutes * 7
	MonthInMinutes = WeekInMinutes * 4
	YearInMinutes  = MonthInMinutes * 12
)

const (
	MinuteInSeconds = int(60)
	HourInSeconds   = HourInMinutes * MinuteInSeconds
	DayInSeconds    = DayInMinutes * MinuteInSeconds
	WeekInSeconds   = WeekInMinutes * MinuteInSeconds
	MonthInSeconds  = MonthInMinutes * MinuteInSeconds
	YearInSeconds   = YearInMinutes * MinuteInSeconds
)

const MonthInWeeks = int(4)

const (
	WeekInDays  = int(7)
	MonthInDays = WeekInDays * MonthInWeeks
)
