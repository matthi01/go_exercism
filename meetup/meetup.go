package meetup

import (
	"time"
)

type WeekSchedule int

const (
	First  WeekSchedule = 1
	Second              = 8
	Third               = 15
	Fourth              = 22
	Teenth              = 13
	Last                = -6
)

func Day(week WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	if week == Last {
		month++
	}
	t := time.Date(year, month, int(week), 12, 0, 0, 0, time.UTC)
	return t.Day() + int(weekday-t.Weekday()+7)%7
}
