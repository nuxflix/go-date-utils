package dateutils

import (
	"time"
)

// NextDay returns the next occurrence of the specified weekday.
// Returns the date of the next occurrence of the given weekday.
// If the current date is already the target weekday, returns the next week's occurrence.
//
// weekday: 0 = Sunday, 1 = Monday, ..., 6 = Saturday
func NextDay(date time.Time, weekday time.Weekday) time.Time {
	currentWeekday := date.Weekday()
	daysUntilTarget := int(weekday - currentWeekday)

	// If the target day is today or in the past this week, move to next week
	if daysUntilTarget <= 0 {
		daysUntilTarget += 7
	}

	return date.AddDate(0, 0, daysUntilTarget)
}

// PreviousDay returns the previous occurrence of the specified weekday.
// Returns the date of the previous occurrence of the given weekday.
// If the current date is already the target weekday, returns the previous week's occurrence.
//
// weekday: 0 = Sunday, 1 = Monday, ..., 6 = Saturday
func PreviousDay(date time.Time, weekday time.Weekday) time.Time {
	currentWeekday := date.Weekday()
	daysSinceTarget := int(currentWeekday - weekday)

	// If the target day is today or in the future this week, move to previous week
	if daysSinceTarget <= 0 {
		daysSinceTarget += 7
	}

	return date.AddDate(0, 0, -daysSinceTarget)
}

// NextSunday returns the next Sunday after the given date.
func NextSunday(date time.Time) time.Time {
	return NextDay(date, time.Sunday)
}

// NextMonday returns the next Monday after the given date.
func NextMonday(date time.Time) time.Time {
	return NextDay(date, time.Monday)
}

// NextTuesday returns the next Tuesday after the given date.
func NextTuesday(date time.Time) time.Time {
	return NextDay(date, time.Tuesday)
}

// NextWednesday returns the next Wednesday after the given date.
func NextWednesday(date time.Time) time.Time {
	return NextDay(date, time.Wednesday)
}

// NextThursday returns the next Thursday after the given date.
func NextThursday(date time.Time) time.Time {
	return NextDay(date, time.Thursday)
}

// NextFriday returns the next Friday after the given date.
func NextFriday(date time.Time) time.Time {
	return NextDay(date, time.Friday)
}

// NextSaturday returns the next Saturday after the given date.
func NextSaturday(date time.Time) time.Time {
	return NextDay(date, time.Saturday)
}

// PreviousSunday returns the previous Sunday before the given date.
func PreviousSunday(date time.Time) time.Time {
	return PreviousDay(date, time.Sunday)
}

// PreviousMonday returns the previous Monday before the given date.
func PreviousMonday(date time.Time) time.Time {
	return PreviousDay(date, time.Monday)
}

// PreviousTuesday returns the previous Tuesday before the given date.
func PreviousTuesday(date time.Time) time.Time {
	return PreviousDay(date, time.Tuesday)
}

// PreviousWednesday returns the previous Wednesday before the given date.
func PreviousWednesday(date time.Time) time.Time {
	return PreviousDay(date, time.Wednesday)
}

// PreviousThursday returns the previous Thursday before the given date.
func PreviousThursday(date time.Time) time.Time {
	return PreviousDay(date, time.Thursday)
}

// PreviousFriday returns the previous Friday before the given date.
func PreviousFriday(date time.Time) time.Time {
	return PreviousDay(date, time.Friday)
}

// PreviousSaturday returns the previous Saturday before the given date.
func PreviousSaturday(date time.Time) time.Time {
	return PreviousDay(date, time.Saturday)
}
