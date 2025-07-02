package dateutils

import (
	"time"
)

// SetDate sets the day of the month for the given time and returns a new time.Time.
// If the day is invalid for the given month/year, it will be adjusted to the last valid day.
// Returns a new time.Time instance, leaving the original unchanged (immutable).
func SetDate(t time.Time, day int) time.Time {
	year := t.Year()
	month := t.Month()

	// Get the last day of the month to validate the day
	lastDayOfMonth := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()

	// Adjust day if it's invalid
	if day < 1 {
		day = 1
	} else if day > lastDayOfMonth {
		day = lastDayOfMonth
	}

	return time.Date(year, month, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// SetDay sets the day of the week for the given time and returns a new time.Time.
// Sunday = 0, Monday = 1, Tuesday = 2, ..., Saturday = 6.
// This follows JavaScript/date-fns convention.
// Returns a new time.Time instance, leaving the original unchanged (immutable).
func SetDay(t time.Time, day int) time.Time {
	currentDay := int(t.Weekday())
	daysDiff := day - currentDay

	// Handle wrap-around for negative values
	if daysDiff < -3 {
		daysDiff += 7
	} else if daysDiff > 3 {
		daysDiff -= 7
	}

	return t.AddDate(0, 0, daysDiff)
}

// SetDayOfYear sets the day of the year for the given time and returns a new time.Time.
// Day 1 = January 1st, Day 365/366 = December 31st.
// If the day is invalid for the given year, it will be adjusted to the valid range.
// Returns a new time.Time instance, leaving the original unchanged (immutable).
func SetDayOfYear(t time.Time, dayOfYear int) time.Time {
	year := t.Year()

	// Get the total days in the year
	daysInYear := 365
	if isLeapYear(year) {
		daysInYear = 366
	}

	// Adjust dayOfYear if it's out of range
	if dayOfYear < 1 {
		dayOfYear = 1
	} else if dayOfYear > daysInYear {
		dayOfYear = daysInYear
	}

	// Create a date with January 1st and add the required days
	jan1 := time.Date(year, 1, 1, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	return jan1.AddDate(0, 0, dayOfYear-1)
}

// SetHours sets the hour component for the given time and returns a new time.Time.
// Hours should be in 24-hour format (0-23).
// If the hour is invalid, it will be adjusted to the valid range.
// Returns a new time.Time instance, leaving the original unchanged (immutable).
func SetHours(t time.Time, hours int) time.Time {
	// Adjust hours if out of range
	if hours < 0 {
		hours = 0
	} else if hours > 23 {
		hours = 23
	}

	return time.Date(t.Year(), t.Month(), t.Day(), hours, t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// SetMinutes sets the minute component for the given time and returns a new time.Time.
// Minutes should be in range (0-59).
// If the minute is invalid, it will be adjusted to the valid range.
// Returns a new time.Time instance, leaving the original unchanged (immutable).
func SetMinutes(t time.Time, minutes int) time.Time {
	// Adjust minutes if out of range
	if minutes < 0 {
		minutes = 0
	} else if minutes > 59 {
		minutes = 59
	}

	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), minutes, t.Second(), t.Nanosecond(), t.Location())
}

// SetSeconds sets the second component for the given time and returns a new time.Time.
// Seconds should be in range (0-59).
// If the second is invalid, it will be adjusted to the valid range.
// Returns a new time.Time instance, leaving the original unchanged (immutable).
func SetSeconds(t time.Time, seconds int) time.Time {
	// Adjust seconds if out of range
	if seconds < 0 {
		seconds = 0
	} else if seconds > 59 {
		seconds = 59
	}

	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), seconds, t.Nanosecond(), t.Location())
}

// SetMilliseconds sets the millisecond component for the given time and returns a new time.Time.
// Milliseconds should be in range (0-999).
// If the millisecond is invalid, it will be adjusted to the valid range.
// Returns a new time.Time instance, leaving the original unchanged (immutable).
func SetMilliseconds(t time.Time, milliseconds int) time.Time {
	// Adjust milliseconds if out of range
	if milliseconds < 0 {
		milliseconds = 0
	} else if milliseconds > 999 {
		milliseconds = 999
	}

	// Convert milliseconds to nanoseconds, preserving existing nanosecond precision beyond milliseconds
	existingNanos := t.Nanosecond() % 1000000 // Get remaining nanoseconds beyond milliseconds
	newNanos := (milliseconds * 1000000) + existingNanos

	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), newNanos, t.Location())
}

// SetMonth sets the month component for the given time and returns a new time.Time.
// Month should be in range (1-12). January = 1, December = 12.
// If the day is invalid for the new month, it will be adjusted to the last valid day.
// Returns a new time.Time instance, leaving the original unchanged (immutable).
func SetMonth(t time.Time, month int) time.Time {
	// Adjust month if out of range
	if month < 1 {
		month = 1
	} else if month > 12 {
		month = 12
	}

	year := t.Year()
	day := t.Day()

	// Get the last day of the target month to validate the day
	lastDayOfMonth := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Day()

	// Adjust day if it's invalid for the new month
	if day > lastDayOfMonth {
		day = lastDayOfMonth
	}

	return time.Date(year, time.Month(month), day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// SetYear sets the year component for the given time and returns a new time.Time.
// If the date is Feb 29 and the target year is not a leap year, it will be adjusted to Feb 28.
// Returns a new time.Time instance, leaving the original unchanged (immutable).
func SetYear(t time.Time, year int) time.Time {
	month := t.Month()
	day := t.Day()

	// Handle leap year adjustment for Feb 29
	if month == time.February && day == 29 && !isLeapYear(year) {
		day = 28
	}

	return time.Date(year, month, day, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// SetQuarter sets the quarter for the given time and returns a new time.Time.
// Quarter should be in range (1-4). Q1 = Jan-Mar, Q2 = Apr-Jun, Q3 = Jul-Sep, Q4 = Oct-Dec.
// The day will be preserved if valid, otherwise adjusted to the last valid day of the month.
// Returns a new time.Time instance, leaving the original unchanged (immutable).
func SetQuarter(t time.Time, quarter int) time.Time {
	// Adjust quarter if out of range
	if quarter < 1 {
		quarter = 1
	} else if quarter > 4 {
		quarter = 4
	}

	// Calculate target month based on quarter
	targetMonth := ((quarter - 1) * 3) + 1 // Q1=1(Jan), Q2=4(Apr), Q3=7(Jul), Q4=10(Oct)

	return SetMonth(t, targetMonth)
}
