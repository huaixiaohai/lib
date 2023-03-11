package time_plus

import "time"

const (
	Second = 1
	Minute = 60 * Second
	House  = 60 * Minute
	Day    = 24 * House
)

func DailyRange(date time.Time) (dailyStart, dailyEnd time.Time) {
	return DailyStart(date), DailyEnd(date)
}

func DailyStart(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())

}

func DailyEnd(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 24, 0, 0, 0, date.Location())
}

func TodayRange() (dailyStart, dailyEnd time.Time) {
	return TodayStart(), TodayEnd()
}

func TodayStart() time.Time {
	return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Now().Location())

}

func TodayEnd() time.Time {
	return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 24, 0, 0, 0, time.Now().Location())
}

func WeeklyRange(date time.Time) (weeklyStart, weeklyEnd time.Time) {
	return WeeklyStart(date), WeeklyEnd(date)
}

func WeeklyStart(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location()).AddDate(0, 0, 0-int(date.Weekday()))

}

func WeeklyEnd(date time.Time) time.Time {
	return WeeklyStart(date).AddDate(0, 0, 6)
}

func ThisWeeklyRange() (weeklyStart, weeklyEnd time.Time) {
	return ThisWeeklyStart(), ThisWeeklyEnd()
}

func ThisWeeklyStart() time.Time {
	date := time.Now()
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location()).AddDate(0, 0, 0-int(date.Weekday()))

}

func ThisWeeklyEnd() time.Time {
	return ThisWeeklyStart().AddDate(0, 0, 6)
}
