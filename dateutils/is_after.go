package dateutils

import (
	"time"
)

// IsAfter checks if the first time is after the second time.
// Returns true if t1 is after t2, false otherwise.
// Both times are compared with their full precision including nanoseconds.
func IsAfter(t1, t2 time.Time) bool {
	return t1.After(t2)
}

// IsAfterOrEqual checks if the first time is after or equal to the second time.
// Returns true if t1 is after or equal to t2, false otherwise.
func IsAfterOrEqual(t1, t2 time.Time) bool {
	return t1.After(t2) || t1.Equal(t2)
}

// IsAfterDate checks if the first date is after the second date,
// ignoring the time component. Only the year, month, and day are compared.
// Both times are normalized to the start of their respective days in UTC.
func IsAfterDate(t1, t2 time.Time) bool {
	// Normalize both times to start of day in UTC for comparison
	date1 := time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.UTC)
	date2 := time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.UTC)

	return date1.After(date2)
}

// IsAfterDateOrEqual checks if the first date is after or equal to the second date,
// ignoring the time component. Only the year, month, and day are compared.
func IsAfterDateOrEqual(t1, t2 time.Time) bool {
	// Normalize both times to start of day in UTC for comparison
	date1 := time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.UTC)
	date2 := time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.UTC)

	return date1.After(date2) || date1.Equal(date2)
}

// IsAfterInTimezone checks if the first time is after the second time
// when both are converted to the specified timezone.
// This is useful when comparing times that might be in different timezones.
func IsAfterInTimezone(t1, t2 time.Time, timezone *time.Location) bool {
	if timezone == nil {
		timezone = time.UTC
	}

	t1InTz := t1.In(timezone)
	t2InTz := t2.In(timezone)

	return t1InTz.After(t2InTz)
}
