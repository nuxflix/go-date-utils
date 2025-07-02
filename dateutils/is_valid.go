package dateutils

import (
	"time"
)

// IsValid checks if a time.Time value represents a valid date and time.
// Returns false for zero time values or invalid dates.
func IsValid(t time.Time) bool {
	return !t.IsZero()
}

// IsLeapYear checks if the given year is a leap year.
// A leap year is divisible by 4, but not by 100, unless it's also divisible by 400.
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// IsWeekend checks if the given time falls on a weekend (Saturday or Sunday).
// Returns true for Saturday and Sunday, false for weekdays.
func IsWeekend(t time.Time) bool {
	weekday := t.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

// IsWeekday checks if the given time falls on a weekday (Monday through Friday).
// Returns true for Monday through Friday, false for weekends.
func IsWeekday(t time.Time) bool {
	return !IsWeekend(t)
}

// IsToday checks if the given time is today in the specified timezone.
// If timezone is nil, uses the time's current timezone.
// Only compares the date part, ignoring the time component.
func IsToday(t time.Time, timezone *time.Location) bool {
	if timezone != nil {
		t = t.In(timezone)
	}

	now := time.Now()
	if timezone != nil {
		now = now.In(timezone)
	}

	return isSameDate(t, now)
}

// IsTomorrow checks if the given time is tomorrow in the specified timezone.
// If timezone is nil, uses the time's current timezone.
// Only compares the date part, ignoring the time component.
func IsTomorrow(t time.Time, timezone *time.Location) bool {
	if timezone != nil {
		t = t.In(timezone)
	}

	tomorrow := time.Now().AddDate(0, 0, 1)
	if timezone != nil {
		tomorrow = tomorrow.In(timezone)
	}

	return isSameDate(t, tomorrow)
}

// IsYesterday checks if the given time is yesterday in the specified timezone.
// If timezone is nil, uses the time's current timezone.
// Only compares the date part, ignoring the time component.
func IsYesterday(t time.Time, timezone *time.Location) bool {
	if timezone != nil {
		t = t.In(timezone)
	}

	yesterday := time.Now().AddDate(0, 0, -1)
	if timezone != nil {
		yesterday = yesterday.In(timezone)
	}

	return isSameDate(t, yesterday)
}

// IsWithinInterval checks if the given time falls within the specified interval.
// The interval is inclusive of both start and end times.
// Returns false if start is after end.
func IsWithinInterval(t, start, end time.Time) bool {
	if start.After(end) {
		return false
	}

	return (t.Equal(start) || t.After(start)) && (t.Equal(end) || t.Before(end))
}

// IsFirstDayOfMonth checks if the given time is the first day of its month.
func IsFirstDayOfMonth(t time.Time) bool {
	return t.Day() == 1
}

// IsLastDayOfMonth checks if the given time is the last day of its month.
func IsLastDayOfMonth(t time.Time) bool {
	// Get the first day of next month, then subtract one day
	nextMonth := time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, t.Location())
	lastDay := nextMonth.AddDate(0, 0, -1)
	return t.Day() == lastDay.Day()
}

// IsFirstDayOfYear checks if the given time is January 1st.
func IsFirstDayOfYear(t time.Time) bool {
	return t.Month() == time.January && t.Day() == 1
}

// IsLastDayOfYear checks if the given time is December 31st.
func IsLastDayOfYear(t time.Time) bool {
	return t.Month() == time.December && t.Day() == 31
}

// IsSameDate checks if two times represent the same calendar date,
// ignoring the time component and timezone.
func IsSameDate(t1, t2 time.Time) bool {
	return isSameDate(t1, t2)
}

// IsSameMonth checks if two times are in the same month and year,
// ignoring the day and time components.
func IsSameMonth(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year() && t1.Month() == t2.Month()
}

// IsSameYear checks if two times are in the same year,
// ignoring all other components.
func IsSameYear(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year()
}

// Helper function to check if two times have the same date
func isSameDate(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year() && t1.Month() == t2.Month() && t1.Day() == t2.Day()
}
