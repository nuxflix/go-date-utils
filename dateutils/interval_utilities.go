package dateutils

import (
	"time"
)

// Interval represents a time interval with start and end times.
// This structure is used by the "Each" functions to define the boundaries
// of time periods to generate.
type Interval struct {
	Start time.Time
	End   time.Time
}

// EachDayOfInterval returns a slice of time.Time representing each day
// within the specified interval. Each day is set to the start of day (00:00:00).
// Returns an empty slice if the interval is invalid (start after end).
// This function mimics date-fns eachDayOfInterval behavior.
func EachDayOfInterval(interval Interval) []time.Time {
	if interval.Start.After(interval.End) {
		return []time.Time{}
	}

	var days []time.Time
	current := StartOfDay(interval.Start)
	end := StartOfDay(interval.End)

	for !current.After(end) {
		days = append(days, current)
		current = AddDays(current, 1)
	}

	return days
}

// EachWeekOfInterval returns a slice of time.Time representing the start
// of each week within the specified interval. Weeks start on Monday (ISO convention).
// Returns an empty slice if the interval is invalid (start after end).
// This function mimics date-fns eachWeekOfInterval behavior.
func EachWeekOfInterval(interval Interval) []time.Time {
	if interval.Start.After(interval.End) {
		return []time.Time{}
	}

	var weeks []time.Time
	current := StartOfWeek(interval.Start)
	end := StartOfWeek(interval.End)

	for !current.After(end) {
		weeks = append(weeks, current)
		current = AddDays(current, 7)
	}

	return weeks
}

// EachWeekOfIntervalSunday returns a slice of time.Time representing the start
// of each week within the specified interval. Weeks start on Sunday (US convention).
// Returns an empty slice if the interval is invalid (start after end).
// This function provides Sunday-based week intervals similar to date-fns.
func EachWeekOfIntervalSunday(interval Interval) []time.Time {
	if interval.Start.After(interval.End) {
		return []time.Time{}
	}

	var weeks []time.Time
	current := StartOfWeekSunday(interval.Start)
	end := StartOfWeekSunday(interval.End)

	for !current.After(end) {
		weeks = append(weeks, current)
		current = AddDays(current, 7)
	}

	return weeks
}

// EachMonthOfInterval returns a slice of time.Time representing the start
// of each month within the specified interval. Each month is set to the first day
// of the month at 00:00:00.
// Returns an empty slice if the interval is invalid (start after end).
// This function mimics date-fns eachMonthOfInterval behavior.
func EachMonthOfInterval(interval Interval) []time.Time {
	if interval.Start.After(interval.End) {
		return []time.Time{}
	}

	var months []time.Time
	current := StartOfMonth(interval.Start)
	end := StartOfMonth(interval.End)

	for !current.After(end) {
		months = append(months, current)
		current = AddMonths(current, 1)
	}

	return months
}

// EachYearOfInterval returns a slice of time.Time representing the start
// of each year within the specified interval. Each year is set to January 1st
// at 00:00:00.
// Returns an empty slice if the interval is invalid (start after end).
// This function mimics date-fns eachYearOfInterval behavior.
func EachYearOfInterval(interval Interval) []time.Time {
	if interval.Start.After(interval.End) {
		return []time.Time{}
	}

	var years []time.Time
	current := StartOfYear(interval.Start)
	end := StartOfYear(interval.End)

	for !current.After(end) {
		years = append(years, current)
		current = AddYears(current, 1)
	}

	return years
}

// EachQuarterOfInterval returns a slice of time.Time representing the start
// of each quarter within the specified interval. Each quarter is set to the first day
// of the quarter at 00:00:00.
// Returns an empty slice if the interval is invalid (start after end).
// This function mimics date-fns eachQuarterOfInterval behavior.
func EachQuarterOfInterval(interval Interval) []time.Time {
	if interval.Start.After(interval.End) {
		return []time.Time{}
	}

	var quarters []time.Time
	current := StartOfQuarter(interval.Start)
	end := StartOfQuarter(interval.End)

	for !current.After(end) {
		quarters = append(quarters, current)
		current = AddMonths(current, 3) // Add 3 months for next quarter
	}

	return quarters
}

// EachHourOfInterval returns a slice of time.Time representing each hour
// within the specified interval. Each hour is set to the start of hour (:00:00).
// Returns an empty slice if the interval is invalid (start after end).
// This function mimics date-fns eachHourOfInterval behavior.
func EachHourOfInterval(interval Interval) []time.Time {
	if interval.Start.After(interval.End) {
		return []time.Time{}
	}

	var hours []time.Time
	current := StartOfHour(interval.Start)
	end := StartOfHour(interval.End)

	for !current.After(end) {
		hours = append(hours, current)
		current = AddHours(current, 1)
	}

	return hours
}

// EachMinuteOfInterval returns a slice of time.Time representing each minute
// within the specified interval. Each minute is set to the start of minute (:00).
// Returns an empty slice if the interval is invalid (start after end).
// This function mimics date-fns eachMinuteOfInterval behavior.
func EachMinuteOfInterval(interval Interval) []time.Time {
	if interval.Start.After(interval.End) {
		return []time.Time{}
	}

	var minutes []time.Time
	current := StartOfMinute(interval.Start)
	end := StartOfMinute(interval.End)

	for !current.After(end) {
		minutes = append(minutes, current)
		current = AddMinutes(current, 1)
	}

	return minutes
}

// EachWeekendOfInterval returns a slice of time.Time representing each weekend day
// (Saturday and Sunday) within the specified interval. Each day is set to the start
// of day (00:00:00).
// Returns an empty slice if the interval is invalid (start after end).
// This function mimics date-fns eachWeekendOfInterval behavior.
func EachWeekendOfInterval(interval Interval) []time.Time {
	if interval.Start.After(interval.End) {
		return []time.Time{}
	}

	var weekends []time.Time
	days := EachDayOfInterval(interval)

	for _, day := range days {
		if IsWeekend(day) {
			weekends = append(weekends, day)
		}
	}

	return weekends
}

// EachBusinessDayOfInterval returns a slice of time.Time representing each business day
// (Monday through Friday) within the specified interval. Each day is set to the start
// of day (00:00:00).
// Returns an empty slice if the interval is invalid (start after end).
// This function provides business day intervals similar to date-fns weekday functions.
func EachBusinessDayOfInterval(interval Interval) []time.Time {
	if interval.Start.After(interval.End) {
		return []time.Time{}
	}

	var businessDays []time.Time
	days := EachDayOfInterval(interval)

	for _, day := range days {
		if IsWeekday(day) {
			businessDays = append(businessDays, day)
		}
	}

	return businessDays
}

// StartOfQuarter returns a new time representing the start of the quarter
// for the given time. The quarter starts at the first day of the quarter month
// at 00:00:00.
// Q1: January, Q2: April, Q3: July, Q4: October
func StartOfQuarter(t time.Time) time.Time {
	quarter := GetQuarter(t)
	var month time.Month

	switch quarter {
	case 1:
		month = time.January
	case 2:
		month = time.April
	case 3:
		month = time.July
	case 4:
		month = time.October
	}

	return time.Date(t.Year(), month, 1, 0, 0, 0, 0, t.Location())
}

// EndOfQuarter returns a new time representing the end of the quarter
// for the given time. The quarter ends at the last day of the quarter month
// at 23:59:59.999999999.
// Q1: March, Q2: June, Q3: September, Q4: December
func EndOfQuarter(t time.Time) time.Time {
	quarter := GetQuarter(t)
	var month time.Month

	switch quarter {
	case 1:
		month = time.March
	case 2:
		month = time.June
	case 3:
		month = time.September
	case 4:
		month = time.December
	}

	// Get the last day of the quarter month
	nextMonth := month + 1
	year := t.Year()
	if nextMonth > time.December {
		nextMonth = time.January
		year++
	}

	firstOfNextMonth := time.Date(year, nextMonth, 1, 0, 0, 0, 0, t.Location())
	lastOfQuarter := firstOfNextMonth.AddDate(0, 0, -1)

	return time.Date(lastOfQuarter.Year(), lastOfQuarter.Month(), lastOfQuarter.Day(),
		23, 59, 59, 999999999, t.Location())
}
