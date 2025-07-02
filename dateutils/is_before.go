package dateutils

import (
	"time"
)

// IsBefore checks if the first time is before the second time.
// Returns true if t1 is before t2, false otherwise.
// Both times are compared with their full precision including nanoseconds.
func IsBefore(t1, t2 time.Time) bool {
	return t1.Before(t2)
}

// IsBeforeOrEqual checks if the first time is before or equal to the second time.
// Returns true if t1 is before or equal to t2, false otherwise.
func IsBeforeOrEqual(t1, t2 time.Time) bool {
	return t1.Before(t2) || t1.Equal(t2)
}

// IsBeforeDate checks if the first date is before the second date,
// ignoring the time component. Only the year, month, and day are compared.
// Both times are normalized to the start of their respective days in UTC.
func IsBeforeDate(t1, t2 time.Time) bool {
	// Normalize both times to start of day in UTC for comparison
	date1 := time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.UTC)
	date2 := time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.UTC)

	return date1.Before(date2)
}

// IsBeforeDateOrEqual checks if the first date is before or equal to the second date,
// ignoring the time component. Only the year, month, and day are compared.
func IsBeforeDateOrEqual(t1, t2 time.Time) bool {
	// Normalize both times to start of day in UTC for comparison
	date1 := time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.UTC)
	date2 := time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.UTC)

	return date1.Before(date2) || date1.Equal(date2)
}

// IsBeforeInTimezone checks if the first time is before the second time
// when both are converted to the specified timezone.
// This is useful when comparing times that might be in different timezones.
func IsBeforeInTimezone(t1, t2 time.Time, timezone *time.Location) bool {
	if timezone == nil {
		timezone = time.UTC
	}

	t1InTz := t1.In(timezone)
	t2InTz := t2.In(timezone)

	return t1InTz.Before(t2InTz)
}
