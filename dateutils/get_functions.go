package dateutils

import (
	"time"
)

// GetDate returns the day of the month for the given time (1-31).
// This function extracts the day component from a date.
// Follows date-fns naming convention.
func GetDate(t time.Time) int {
	return t.Day()
}

// GetDay returns the day of the week for the given time (0-6).
// Sunday = 0, Monday = 1, Tuesday = 2, ..., Saturday = 6.
// This follows JavaScript/date-fns convention where Sunday is 0.
func GetDay(t time.Time) int {
	// Go's Weekday() returns Sunday = 0, Monday = 1, etc., which matches date-fns
	return int(t.Weekday())
}

// GetDayOfYear returns the day of the year for the given time (1-366).
// January 1st is day 1, December 31st is day 365 (366 in leap years).
func GetDayOfYear(t time.Time) int {
	return t.YearDay()
}

// GetHours returns the hour component of the given time (0-23).
// Uses 24-hour format.
func GetHours(t time.Time) int {
	return t.Hour()
}

// GetMinutes returns the minute component of the given time (0-59).
func GetMinutes(t time.Time) int {
	return t.Minute()
}

// GetSeconds returns the second component of the given time (0-59).
func GetSeconds(t time.Time) int {
	return t.Second()
}

// GetMilliseconds returns the millisecond component of the given time (0-999).
// This extracts milliseconds from the nanoseconds component.
func GetMilliseconds(t time.Time) int {
	return t.Nanosecond() / 1000000
}

// GetMonth returns the month component of the given time (1-12).
// January = 1, February = 2, ..., December = 12.
// This follows date-fns convention (1-based months).
func GetMonth(t time.Time) int {
	return int(t.Month())
}

// GetYear returns the year component of the given time.
func GetYear(t time.Time) int {
	return t.Year()
}

// GetQuarter returns the quarter of the year for the given time (1-4).
// Q1 = Jan-Mar (1), Q2 = Apr-Jun (2), Q3 = Jul-Sep (3), Q4 = Oct-Dec (4).
func GetQuarter(t time.Time) int {
	month := t.Month()
	switch {
	case month >= time.January && month <= time.March:
		return 1
	case month >= time.April && month <= time.June:
		return 2
	case month >= time.July && month <= time.September:
		return 3
	default: // October, November, December
		return 4
	}
}

// GetWeek returns the ISO week number for the given time (1-53).
// Week 1 is the first week that contains at least 4 days of the new year.
// This follows the ISO 8601 standard.
func GetWeek(t time.Time) int {
	_, week := t.ISOWeek()
	return week
}

// GetWeekOfMonth returns the week of the month for the given time (1-6).
// The first week is the week containing the first day of the month.
// Weeks start on Monday (ISO convention).
func GetWeekOfMonth(t time.Time) int {
	// Get the first day of the month
	firstDay := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())

	// Calculate days from the first day to current date
	daysDiff := t.Day() - 1

	// Calculate days from Monday of the first week to current date
	// If first day is not Monday, we need to adjust
	firstDayWeekday := int(firstDay.Weekday())
	if firstDayWeekday == 0 { // Sunday
		firstDayWeekday = 7 // Convert Sunday from 0 to 7 for ISO week calculation
	}

	// Days to add to get to the current date from the Monday of the first week
	daysFromMonday := daysDiff + (firstDayWeekday - 1)

	// Week number is calculated by dividing by 7 and adding 1
	return (daysFromMonday / 7) + 1
}

// GetWeekYear returns the ISO week year for the given time.
// The ISO week year can differ from the calendar year for dates near year boundaries.
// For example, January 1, 2021 belongs to ISO week year 2020.
func GetWeekYear(t time.Time) int {
	year, _ := t.ISOWeek()
	return year
}
