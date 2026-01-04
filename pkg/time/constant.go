package time

import "time"

const (
	HourMinute       = "15:04"
	HourMinuteSecond = "15:04:05"
	DateMinute       = "2006-01-02 15:04"
	DateSecond       = "2006-01-02 15:04:05"
	Micro            = "15:04:05.000000"

	HourInMinutes  int = 60
	DayInMinutes       = HourInMinutes * 24
	WeekInMinutes      = DayInMinutes * 7
	MonthInMinutes     = WeekInMinutes * 4
	YearInMinutes      = MonthInMinutes * 12

	MinuteInSeconds int = 60
	HourInSeconds       = HourInMinutes * MinuteInSeconds
	DayInSeconds        = DayInMinutes * MinuteInSeconds
	WeekInSeconds       = WeekInMinutes * MinuteInSeconds
	MonthInSeconds      = MonthInMinutes * MinuteInSeconds
	YearInSeconds       = YearInMinutes * MinuteInSeconds

	MonthInWeeks int = 4

	WeekInDays  int = 7
	MonthInDays     = WeekInDays * MonthInWeeks
)

var PublicHolidays = []time.Time{
	NewDay(2024, time.November, 1),  // All Saints' Day
	NewDay(2024, time.December, 25), // Christmas Day
	NewDay(2024, time.December, 26), // St. Stephen's Day
	NewDay(2025, time.January, 1),   // New Year's Day
	NewDay(2025, time.January, 6),   // Epiphany
	NewDay(2025, time.April, 18),    // Good Friday
	NewDay(2025, time.April, 21),    // Easter Monday
	NewDay(2025, time.May, 1),       // Labour Day
	NewDay(2025, time.May, 29),      // Ascension Day
	NewDay(2025, time.June, 9),      // Whit Monday
	NewDay(2025, time.June, 19),     // Corpus Christi
	NewDay(2025, time.October, 3),   // German Unity Day
	NewDay(2025, time.November, 1),  // All Saints' Day
	NewDay(2025, time.December, 25), // Christmas Day
	NewDay(2025, time.December, 26), // St. Stephen's Day
}
