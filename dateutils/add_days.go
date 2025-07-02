package dateutils

import (
	"time"
)

// AddDays adds the specified number of days to the given time.
// Returns a new time.Time instance, leaving the original unchanged (immutable).
// Supports negative values to subtract days.
// The time component (hours, minutes, seconds) remains unchanged.
func AddDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

// AddBusinessDays adds the specified number of business days (Monday-Friday) to the given time.
// Weekends (Saturday and Sunday) are skipped.
// Returns a new time.Time instance, leaving the original unchanged (immutable).
// Supports negative values to subtract business days.
func AddBusinessDays(t time.Time, businessDays int) time.Time {
	if businessDays == 0 {
		return t
	}

	result := t
	remaining := businessDays
	increment := 1

	if businessDays < 0 {
		increment = -1
		remaining = -remaining
	}

	// If starting from a weekend, move to the next/previous business day first
	for IsWeekend(result) {
		result = result.AddDate(0, 0, increment)
	}

	for remaining > 0 {
		result = result.AddDate(0, 0, increment)

		// Skip weekends
		weekday := result.Weekday()
		if weekday != time.Saturday && weekday != time.Sunday {
			remaining--
		}
	}

	return result
}

// AddWeeks adds the specified number of weeks (7 days each) to the given time.
// Returns a new time.Time instance, leaving the original unchanged (immutable).
// Supports negative values to subtract weeks.
func AddWeeks(t time.Time, weeks int) time.Time {
	return AddDays(t, weeks*7)
}

// AddDaysWithTimezone adds days to a time and ensures the result is in the specified timezone.
// This is useful when you want to maintain a specific timezone after the addition.
// If timezone is nil, the original timezone is preserved.
func AddDaysWithTimezone(t time.Time, days int, timezone *time.Location) time.Time {
	result := AddDays(t, days)

	if timezone != nil {
		result = result.In(timezone)
	}

	return result
}
