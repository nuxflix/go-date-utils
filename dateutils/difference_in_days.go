package dateutils

import (
	"math"
	"time"
)

// DifferenceInDays calculates the difference in days between two times.
// Returns the number of full 24-hour periods between the times.
// The result is positive if t1 is after t2, negative if t1 is before t2.
// The calculation is based on the actual duration between the times.
func DifferenceInDays(t1, t2 time.Time) int {
	duration := t1.Sub(t2)
	days := duration.Hours() / 24
	return int(math.Trunc(days))
}

// DifferenceInDaysFloat calculates the difference in days between two times
// and returns a floating-point result for more precise calculations.
// The result includes fractional days (e.g., 1.5 days).
func DifferenceInDaysFloat(t1, t2 time.Time) float64 {
	duration := t1.Sub(t2)
	return duration.Hours() / 24
}

// DifferenceInCalendarDays calculates the difference in calendar days between two times.
// This ignores the time component and only considers the dates.
// For example, 2023-12-25 23:59:59 and 2023-12-26 00:00:01 would have a difference of 1 day.
// The result is positive if t1 is after t2, negative if t1 is before t2.
func DifferenceInCalendarDays(t1, t2 time.Time) int {
	// Normalize both times to the start of their respective days in UTC
	date1 := time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.UTC)
	date2 := time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.UTC)

	duration := date1.Sub(date2)
	days := duration.Hours() / 24
	return int(math.Round(days))
}

// DifferenceInBusinessDays calculates the difference in business days (Monday-Friday)
// between two times. Weekends are excluded from the count.
// The result is positive if t1 is after t2, negative if t1 is before t2.
func DifferenceInBusinessDays(t1, t2 time.Time) int {
	if t1.Equal(t2) {
		return 0
	}

	// Ensure t1 is after t2, and track if we need to negate the result
	start, end := t2, t1
	multiplier := 1

	if t1.Before(t2) {
		start, end = t1, t2
		multiplier = -1
	}

	// Normalize to start of days for consistent calculation
	startDate := time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, time.UTC)
	endDate := time.Date(end.Year(), end.Month(), end.Day(), 0, 0, 0, 0, time.UTC)

	businessDays := 0
	current := startDate

	for current.Before(endDate) {
		weekday := current.Weekday()
		if weekday != time.Saturday && weekday != time.Sunday {
			businessDays++
		}
		current = current.AddDate(0, 0, 1)
	}

	return businessDays * multiplier
}

// DifferenceInWeeks calculates the difference in weeks between two times.
// Returns the number of full 7-day periods between the times.
// The result is positive if t1 is after t2, negative if t1 is before t2.
func DifferenceInWeeks(t1, t2 time.Time) int {
	days := DifferenceInDays(t1, t2)
	return days / 7
}

// DifferenceInWeeksFloat calculates the difference in weeks between two times
// and returns a floating-point result for more precise calculations.
// The result includes fractional weeks (e.g., 1.5 weeks).
func DifferenceInWeeksFloat(t1, t2 time.Time) float64 {
	days := DifferenceInDaysFloat(t1, t2)
	return days / 7
}

// AbsDifferenceInDays returns the absolute difference in days between two times.
// The result is always positive or zero, regardless of which time is earlier.
func AbsDifferenceInDays(t1, t2 time.Time) int {
	diff := DifferenceInDays(t1, t2)
	if diff < 0 {
		return -diff
	}
	return diff
}
