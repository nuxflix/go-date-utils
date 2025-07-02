package dateutils

import (
	"time"
)

// GetISOWeek returns the ISO week number (1-53) for the given date.
// ISO week-numbering year starts on the Monday that is nearest to January 1st.
// The result is based on the ISO 8601 standard.
//
// Example:
//
//	GetISOWeek(time.Date(2005, time.January, 2, 0, 0, 0, 0, time.UTC)) // 53
//	GetISOWeek(time.Date(2005, time.January, 3, 0, 0, 0, 0, time.UTC)) // 1
func GetISOWeek(t time.Time) int {
	_, week := t.ISOWeek()
	return week
}

// GetISOWeekYear returns the ISO week-numbering year for the given date.
// This may differ from the regular year near year boundaries.
//
// Example:
//
//	GetISOWeekYear(time.Date(2005, time.January, 2, 0, 0, 0, 0, time.UTC)) // 2004
//	GetISOWeekYear(time.Date(2005, time.January, 3, 0, 0, 0, 0, time.UTC)) // 2005
func GetISOWeekYear(t time.Time) int {
	year, _ := t.ISOWeek()
	return year
}

// GetISOWeeksInYear returns the number of ISO weeks in the given year.
// Most years have 52 weeks, but some have 53.
//
// Example:
//
//	GetISOWeeksInYear(time.Date(2020, time.June, 15, 0, 0, 0, 0, time.UTC)) // 53
//	GetISOWeeksInYear(time.Date(2021, time.June, 15, 0, 0, 0, 0, time.UTC)) // 52
func GetISOWeeksInYear(t time.Time) int {
	// Get the ISO year
	isoYear := GetISOWeekYear(t)

	// Check the last day of the ISO year to see if week 53 exists
	// Find the last day of the ISO year
	nextYear := time.Date(isoYear+1, time.January, 4, 0, 0, 0, 0, t.Location())
	startOfNextISOYear := StartOfISOWeek(nextYear)
	lastDayOfISOYear := startOfNextISOYear.AddDate(0, 0, -1)

	_, lastWeek := lastDayOfISOYear.ISOWeek()
	return lastWeek
}

// StartOfISOWeek returns the start of the ISO week (Monday) for the given date.
// The time is set to 00:00:00.000.
//
// Example:
//
//	StartOfISOWeek(time.Date(2014, time.September, 2, 11, 55, 0, 0, time.UTC))
//	// Returns: Monday, September 1, 2014 00:00:00.000
func StartOfISOWeek(t time.Time) time.Time {
	// ISO week starts on Monday (weekday 1)
	weekday := int(t.Weekday())
	if weekday == 0 { // Sunday
		weekday = 7
	}

	daysToSubtract := weekday - 1
	start := t.AddDate(0, 0, -daysToSubtract)

	return time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, t.Location())
}

// EndOfISOWeek returns the end of the ISO week (Sunday) for the given date.
// The time is set to 23:59:59.999.
//
// Example:
//
//	EndOfISOWeek(time.Date(2014, time.September, 2, 11, 55, 0, 0, time.UTC))
//	// Returns: Sunday, September 7, 2014 23:59:59.999
func EndOfISOWeek(t time.Time) time.Time {
	start := StartOfISOWeek(t)
	end := start.AddDate(0, 0, 6)

	return time.Date(end.Year(), end.Month(), end.Day(), 23, 59, 59, 999000000, t.Location())
}

// LastDayOfISOWeek returns the last day of the ISO week (Sunday) for the given date.
// The time is set to 00:00:00.000.
//
// Example:
//
//	LastDayOfISOWeek(time.Date(2014, time.September, 2, 11, 55, 0, 0, time.UTC))
//	// Returns: Sunday, September 7, 2014 00:00:00.000
func LastDayOfISOWeek(t time.Time) time.Time {
	start := StartOfISOWeek(t)
	lastDay := start.AddDate(0, 0, 6)

	return time.Date(lastDay.Year(), lastDay.Month(), lastDay.Day(), 0, 0, 0, 0, t.Location())
}

// StartOfISOWeekYear returns the start of the ISO week-numbering year for the given date.
// The ISO week-numbering year always starts on the Monday that contains January 4th.
//
// Example:
//
//	StartOfISOWeekYear(time.Date(2005, time.July, 2, 0, 0, 0, 0, time.UTC))
//	// Returns: Monday, January 3, 2005 00:00:00.000
func StartOfISOWeekYear(t time.Time) time.Time {
	isoYear := GetISOWeekYear(t)

	// January 4th of the ISO year is always in the first ISO week
	jan4th := time.Date(isoYear, time.January, 4, 0, 0, 0, 0, t.Location())
	return StartOfISOWeek(jan4th)
}

// EndOfISOWeekYear returns the end of the ISO week-numbering year for the given date.
// The time is set to 23:59:59.999.
//
// Example:
//
//	EndOfISOWeekYear(time.Date(2005, time.July, 2, 0, 0, 0, 0, time.UTC))
//	// Returns: Sunday, January 2, 2006 23:59:59.999
func EndOfISOWeekYear(t time.Time) time.Time {
	isoYear := GetISOWeekYear(t)

	// The ISO year ends the day before the next ISO year starts
	nextYear := time.Date(isoYear+1, time.January, 4, 0, 0, 0, 0, t.Location())
	startOfNextYear := StartOfISOWeek(nextYear)
	endOfYear := startOfNextYear.AddDate(0, 0, -1)

	return time.Date(endOfYear.Year(), endOfYear.Month(), endOfYear.Day(), 23, 59, 59, 999000000, t.Location())
}

// LastDayOfISOWeekYear returns the last day of the ISO week-numbering year for the given date.
// The time is set to 00:00:00.000.
//
// Example:
//
//	LastDayOfISOWeekYear(time.Date(2005, time.July, 2, 0, 0, 0, 0, time.UTC))
//	// Returns: Sunday, January 2, 2006 00:00:00.000
func LastDayOfISOWeekYear(t time.Time) time.Time {
	isoYear := GetISOWeekYear(t)

	// The ISO year ends the day before the next ISO year starts
	nextYear := time.Date(isoYear+1, time.January, 4, 0, 0, 0, 0, t.Location())
	startOfNextYear := StartOfISOWeek(nextYear)
	lastDay := startOfNextYear.AddDate(0, 0, -1)

	return time.Date(lastDay.Year(), lastDay.Month(), lastDay.Day(), 0, 0, 0, 0, t.Location())
}

// DifferenceInCalendarISOWeeks calculates the number of calendar ISO weeks between two dates.
// This counts the actual ISO week boundaries crossed, not elapsed time.
//
// Example:
//
//	t1 := time.Date(2014, time.September, 1, 0, 0, 0, 0, time.UTC)  // Monday
//	t2 := time.Date(2014, time.September, 8, 0, 0, 0, 0, time.UTC)  // Monday
//	DifferenceInCalendarISOWeeks(t2, t1) // 1
func DifferenceInCalendarISOWeeks(laterDate, earlierDate time.Time) int {
	startOfLaterWeek := StartOfISOWeek(laterDate)
	startOfEarlierWeek := StartOfISOWeek(earlierDate)

	diff := startOfLaterWeek.Sub(startOfEarlierWeek)
	weeks := int(diff.Hours() / (24 * 7))

	return weeks
}

// DifferenceInCalendarISOWeekYears calculates the number of calendar ISO week-numbering years between two dates.
//
// Example:
//
//	t1 := time.Date(2014, time.July, 2, 0, 0, 0, 0, time.UTC)
//	t2 := time.Date(2015, time.July, 2, 0, 0, 0, 0, time.UTC)
//	DifferenceInCalendarISOWeekYears(t2, t1) // 1
func DifferenceInCalendarISOWeekYears(laterDate, earlierDate time.Time) int {
	laterISOYear := GetISOWeekYear(laterDate)
	earlierISOYear := GetISOWeekYear(earlierDate)

	return laterISOYear - earlierISOYear
}

// IsSameISOWeek returns true if the given dates are in the same ISO week.
//
// Example:
//
//	t1 := time.Date(2014, time.August, 31, 0, 0, 0, 0, time.UTC)  // Sunday
//	t2 := time.Date(2014, time.September, 1, 0, 0, 0, 0, time.UTC) // Monday
//	IsSameISOWeek(t1, t2) // false (different ISO weeks)
func IsSameISOWeek(date1, date2 time.Time) bool {
	year1, week1 := date1.ISOWeek()
	year2, week2 := date2.ISOWeek()

	return year1 == year2 && week1 == week2
}

// IsSameISOWeekYear returns true if the given dates are in the same ISO week-numbering year.
//
// Example:
//
//	t1 := time.Date(2005, time.January, 1, 0, 0, 0, 0, time.UTC)  // Saturday
//	t2 := time.Date(2005, time.January, 2, 0, 0, 0, 0, time.UTC)  // Sunday
//	IsSameISOWeekYear(t1, t2) // false (different ISO years: 2004 vs 2005)
func IsSameISOWeekYear(date1, date2 time.Time) bool {
	return GetISOWeekYear(date1) == GetISOWeekYear(date2)
}
