package dateutils

import (
	"time"
)

// IsEqual checks if two times are exactly equal.
// Returns true if both times represent the same instant in time.
// This follows date-fns naming convention.
func IsEqual(t1, t2 time.Time) bool {
	return t1.Equal(t2)
}

// IsSameDay checks if two times are in the same calendar day.
// Time components are ignored - only year, month, and day are compared.
// This is an alias for IsSameDate to match date-fns naming convention.
func IsSameDay(t1, t2 time.Time) bool {
	return IsSameDate(t1, t2)
}

// IsSameWeek checks if two times are in the same calendar week.
// Week starts on Monday (ISO week).
// Returns true if both times fall within the same Monday-to-Sunday period.
func IsSameWeek(t1, t2 time.Time) bool {
	// Convert both times to UTC to ensure consistent comparison
	t1 = t1.UTC()
	t2 = t2.UTC()

	// Get the start of week (Monday) for both dates
	startOfWeek1 := StartOfWeek(t1)
	startOfWeek2 := StartOfWeek(t2)

	return startOfWeek1.Equal(startOfWeek2)
}

// IsSameHour checks if two times are in the same calendar hour.
// Minutes, seconds, and nanoseconds are ignored.
// Returns true if both times fall within the same hour.
func IsSameHour(t1, t2 time.Time) bool {
	// Convert to same timezone for comparison
	t1 = t1.UTC()
	t2 = t2.UTC()

	return t1.Year() == t2.Year() &&
		t1.Month() == t2.Month() &&
		t1.Day() == t2.Day() &&
		t1.Hour() == t2.Hour()
}

// IsSameMinute checks if two times are in the same calendar minute.
// Seconds and nanoseconds are ignored.
// Returns true if both times fall within the same minute.
func IsSameMinute(t1, t2 time.Time) bool {
	// Convert to same timezone for comparison
	t1 = t1.UTC()
	t2 = t2.UTC()

	return t1.Year() == t2.Year() &&
		t1.Month() == t2.Month() &&
		t1.Day() == t2.Day() &&
		t1.Hour() == t2.Hour() &&
		t1.Minute() == t2.Minute()
}

// CompareAsc compares two times and returns:
// -1 if t1 is before t2
//
//	0 if t1 equals t2
//	1 if t1 is after t2
//
// This follows date-fns naming convention (ascending order).
func CompareAsc(t1, t2 time.Time) int {
	if t1.Before(t2) {
		return -1
	}
	if t1.After(t2) {
		return 1
	}
	return 0
}

// CompareDesc compares two times and returns:
//
//	1 if t1 is before t2
//	0 if t1 equals t2
//
// -1 if t1 is after t2
// This follows date-fns naming convention (descending order).
func CompareDesc(t1, t2 time.Time) int {
	return -CompareAsc(t1, t2)
}

// Min returns the earliest time from the provided slice of times.
// Returns zero time if the slice is empty.
// This follows date-fns naming convention.
func Min(times []time.Time) time.Time {
	if len(times) == 0 {
		return time.Time{}
	}

	min := times[0]
	for _, t := range times[1:] {
		if t.Before(min) {
			min = t
		}
	}
	return min
}

// Max returns the latest time from the provided slice of times.
// Returns zero time if the slice is empty.
// This follows date-fns naming convention.
func Max(times []time.Time) time.Time {
	if len(times) == 0 {
		return time.Time{}
	}

	max := times[0]
	for _, t := range times[1:] {
		if t.After(max) {
			max = t
		}
	}
	return max
}

// ClosestTo finds the time in the slice that is closest to the target time.
// Returns zero time if the slice is empty.
// This follows date-fns naming convention.
func ClosestTo(target time.Time, times []time.Time) time.Time {
	if len(times) == 0 {
		return time.Time{}
	}

	closest := times[0]
	minDiff := AbsDuration(target.Sub(closest))

	for _, t := range times[1:] {
		diff := AbsDuration(target.Sub(t))
		if diff < minDiff {
			minDiff = diff
			closest = t
		}
	}
	return closest
}

// AbsDuration returns the absolute value of a duration.
// Helper function for ClosestTo.
func AbsDuration(d time.Duration) time.Duration {
	if d < 0 {
		return -d
	}
	return d
}
