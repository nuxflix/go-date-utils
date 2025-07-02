package dateutils

import (
	"time"
)

// StartOfDay returns a new time representing the start of the day (00:00:00)
// for the given time, preserving the timezone.
// The time component is set to 00:00:00.000.
func StartOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// EndOfDay returns a new time representing the end of the day (23:59:59.999999999)
// for the given time, preserving the timezone.
// The time component is set to 23:59:59.999999999.
func EndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 999999999, t.Location())
}

// StartOfWeek returns a new time representing the start of the week (Monday 00:00:00)
// for the given time, preserving the timezone.
// The week is considered to start on Monday (ISO 8601 standard).
func StartOfWeek(t time.Time) time.Time {
	// Get the start of the day first
	startOfDay := StartOfDay(t)

	// Calculate days to subtract to get to Monday
	weekday := startOfDay.Weekday()
	daysToSubtract := int(weekday - time.Monday)
	if daysToSubtract < 0 {
		daysToSubtract += 7 // If Sunday, go back 6 days to Monday
	}

	return startOfDay.AddDate(0, 0, -daysToSubtract)
}

// StartOfWeekSunday returns a new time representing the start of the week (Sunday 00:00:00)
// for the given time, preserving the timezone.
// The week is considered to start on Sunday (US standard).
func StartOfWeekSunday(t time.Time) time.Time {
	// Get the start of the day first
	startOfDay := StartOfDay(t)

	// Calculate days to subtract to get to Sunday
	weekday := startOfDay.Weekday()
	daysToSubtract := int(weekday) // Sunday is 0, so this works directly

	return startOfDay.AddDate(0, 0, -daysToSubtract)
}

// EndOfWeek returns a new time representing the end of the week (Sunday 23:59:59.999999999)
// for the given time, preserving the timezone.
// The week is considered to end on Sunday (with Monday as start of week).
func EndOfWeek(t time.Time) time.Time {
	startOfWeek := StartOfWeek(t)
	return EndOfDay(startOfWeek.AddDate(0, 0, 6)) // Add 6 days to get to Sunday
}

// EndOfWeekSaturday returns a new time representing the end of the week (Saturday 23:59:59.999999999)
// for the given time, preserving the timezone.
// The week is considered to end on Saturday (with Sunday as start of week).
func EndOfWeekSaturday(t time.Time) time.Time {
	startOfWeek := StartOfWeekSunday(t)
	return EndOfDay(startOfWeek.AddDate(0, 0, 6)) // Add 6 days to get to Saturday
}

// StartOfMonth returns a new time representing the start of the month (1st day 00:00:00)
// for the given time, preserving the timezone.
func StartOfMonth(t time.Time) time.Time {
	year, month, _ := t.Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth returns a new time representing the end of the month (last day 23:59:59.999999999)
// for the given time, preserving the timezone.
func EndOfMonth(t time.Time) time.Time {
	// Get the first day of next month, then subtract one nanosecond
	year, month, _ := t.Date()
	firstOfNextMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, t.Location())
	return firstOfNextMonth.Add(-time.Nanosecond)
}

// StartOfYear returns a new time representing the start of the year (January 1st 00:00:00)
// for the given time, preserving the timezone.
func StartOfYear(t time.Time) time.Time {
	year, _, _ := t.Date()
	return time.Date(year, time.January, 1, 0, 0, 0, 0, t.Location())
}

// EndOfYear returns a new time representing the end of the year (December 31st 23:59:59.999999999)
// for the given time, preserving the timezone.
func EndOfYear(t time.Time) time.Time {
	year, _, _ := t.Date()
	return time.Date(year, time.December, 31, 23, 59, 59, 999999999, t.Location())
}

// StartOfHour returns a new time representing the start of the hour (XX:00:00)
// for the given time, preserving the timezone.
// The minute, second, and nanosecond components are set to 0.
func StartOfHour(t time.Time) time.Time {
	year, month, day := t.Date()
	hour, _, _ := t.Clock()
	return time.Date(year, month, day, hour, 0, 0, 0, t.Location())
}

// EndOfHour returns a new time representing the end of the hour (XX:59:59.999999999)
// for the given time, preserving the timezone.
func EndOfHour(t time.Time) time.Time {
	year, month, day := t.Date()
	hour, _, _ := t.Clock()
	return time.Date(year, month, day, hour, 59, 59, 999999999, t.Location())
}

// StartOfMinute returns a new time representing the start of the minute (XX:XX:00)
// for the given time, preserving the timezone.
// The second and nanosecond components are set to 0.
func StartOfMinute(t time.Time) time.Time {
	year, month, day := t.Date()
	hour, minute, _ := t.Clock()
	return time.Date(year, month, day, hour, minute, 0, 0, t.Location())
}

// EndOfMinute returns a new time representing the end of the minute (XX:XX:59.999999999)
// for the given time, preserving the timezone.
func EndOfMinute(t time.Time) time.Time {
	year, month, day := t.Date()
	hour, minute, _ := t.Clock()
	return time.Date(year, month, day, hour, minute, 59, 999999999, t.Location())
}
