package dateutils

import (
	"math"
	"time"
)

// DifferenceInHours calculates the difference in hours between two times.
// Returns the number of full hour periods between the times.
// The result is positive if t1 is after t2, negative if t1 is before t2.
func DifferenceInHours(t1, t2 time.Time) int {
	duration := t1.Sub(t2)
	hours := duration.Hours()
	return int(math.Trunc(hours))
}

// DifferenceInHoursFloat calculates the difference in hours between two times
// and returns a floating-point result for more precise calculations.
// The result includes fractional hours (e.g., 1.5 hours).
func DifferenceInHoursFloat(t1, t2 time.Time) float64 {
	duration := t1.Sub(t2)
	return duration.Hours()
}

// DifferenceInMinutes calculates the difference in minutes between two times.
// Returns the number of full minute periods between the times.
// The result is positive if t1 is after t2, negative if t1 is before t2.
func DifferenceInMinutes(t1, t2 time.Time) int {
	duration := t1.Sub(t2)
	minutes := duration.Minutes()
	return int(math.Trunc(minutes))
}

// DifferenceInMinutesFloat calculates the difference in minutes between two times
// and returns a floating-point result for more precise calculations.
// The result includes fractional minutes (e.g., 1.5 minutes).
func DifferenceInMinutesFloat(t1, t2 time.Time) float64 {
	duration := t1.Sub(t2)
	return duration.Minutes()
}

// DifferenceInSeconds calculates the difference in seconds between two times.
// Returns the number of full second periods between the times.
// The result is positive if t1 is after t2, negative if t1 is before t2.
func DifferenceInSeconds(t1, t2 time.Time) int {
	duration := t1.Sub(t2)
	seconds := duration.Seconds()
	return int(math.Trunc(seconds))
}

// DifferenceInSecondsFloat calculates the difference in seconds between two times
// and returns a floating-point result for more precise calculations.
// The result includes fractional seconds (e.g., 1.5 seconds).
func DifferenceInSecondsFloat(t1, t2 time.Time) float64 {
	duration := t1.Sub(t2)
	return duration.Seconds()
}

// DifferenceInMonths calculates the difference in full months between two times.
// This function considers the day component and calculates actual full months.
// The result is positive if t1 is after t2, negative if t1 is before t2.
func DifferenceInMonths(t1, t2 time.Time) int {
	if t1.Before(t2) {
		return -DifferenceInMonths(t2, t1)
	}

	years := t1.Year() - t2.Year()
	months := int(t1.Month()) - int(t2.Month())
	totalMonths := years*12 + months

	// Check if we need to subtract a month because t1 day < t2 day
	if t1.Day() < t2.Day() {
		totalMonths--
	}

	return totalMonths
}

// DifferenceInCalendarMonths calculates the difference in calendar months between two times.
// This ignores the day component and only considers year and month.
// The result is positive if t1 is after t2, negative if t1 is before t2.
func DifferenceInCalendarMonths(t1, t2 time.Time) int {
	years := t1.Year() - t2.Year()
	months := int(t1.Month()) - int(t2.Month())
	return years*12 + months
}

// DifferenceInYears calculates the difference in full years between two times.
// This function considers the month and day components and calculates actual full years.
// The result is positive if t1 is after t2, negative if t1 is before t2.
func DifferenceInYears(t1, t2 time.Time) int {
	if t1.Before(t2) {
		return -DifferenceInYears(t2, t1)
	}

	years := t1.Year() - t2.Year()

	// Check if we need to subtract a year because the anniversary hasn't passed yet
	if t1.Month() < t2.Month() || (t1.Month() == t2.Month() && t1.Day() < t2.Day()) {
		years--
	}

	return years
}

// DifferenceInCalendarYears calculates the difference in calendar years between two times.
// This ignores the month and day components and only considers the year.
// The result is positive if t1 is after t2, negative if t1 is before t2.
func DifferenceInCalendarYears(t1, t2 time.Time) int {
	return t1.Year() - t2.Year()
}

// DifferenceInQuarters calculates the difference in quarters between two times.
// Returns the number of full quarters between the times.
// The result is positive if t1 is after t2, negative if t1 is before t2.
func DifferenceInQuarters(t1, t2 time.Time) int {
	months := DifferenceInMonths(t1, t2)
	return months / 3
}

// DifferenceInCalendarQuarters calculates the difference in calendar quarters between two times.
// This ignores the day component and only considers year and quarter.
// The result is positive if t1 is after t2, negative if t1 is before t2.
func DifferenceInCalendarQuarters(t1, t2 time.Time) int {
	years := t1.Year() - t2.Year()
	quarters := GetQuarter(t1) - GetQuarter(t2)
	return years*4 + quarters
}

// AbsDifferenceInHours returns the absolute difference in hours between two times.
// The result is always positive or zero, regardless of which time is earlier.
func AbsDifferenceInHours(t1, t2 time.Time) int {
	diff := DifferenceInHours(t1, t2)
	if diff < 0 {
		return -diff
	}
	return diff
}

// AbsDifferenceInMinutes returns the absolute difference in minutes between two times.
// The result is always positive or zero, regardless of which time is earlier.
func AbsDifferenceInMinutes(t1, t2 time.Time) int {
	diff := DifferenceInMinutes(t1, t2)
	if diff < 0 {
		return -diff
	}
	return diff
}

// AbsDifferenceInSeconds returns the absolute difference in seconds between two times.
// The result is always positive or zero, regardless of which time is earlier.
func AbsDifferenceInSeconds(t1, t2 time.Time) int {
	diff := DifferenceInSeconds(t1, t2)
	if diff < 0 {
		return -diff
	}
	return diff
}

// AbsDifferenceInMonths returns the absolute difference in months between two times.
// The result is always positive or zero, regardless of which time is earlier.
func AbsDifferenceInMonths(t1, t2 time.Time) int {
	diff := DifferenceInMonths(t1, t2)
	if diff < 0 {
		return -diff
	}
	return diff
}

// AbsDifferenceInYears returns the absolute difference in years between two times.
// The result is always positive or zero, regardless of which time is earlier.
func AbsDifferenceInYears(t1, t2 time.Time) int {
	diff := DifferenceInYears(t1, t2)
	if diff < 0 {
		return -diff
	}
	return diff
}
