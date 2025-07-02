package dateutils

import (
	"time"
)

// SubDays subtracts the specified number of days from the given time.
// Returns a new time.Time instance, leaving the original unchanged (immutable).
// This is the counterpart to AddDays and follows date-fns naming convention.
func SubDays(t time.Time, days int) time.Time {
	return AddDays(t, -days)
}

// SubWeeks subtracts the specified number of weeks from the given time.
// Returns a new time.Time instance, leaving the original unchanged (immutable).
// This is the counterpart to AddWeeks and follows date-fns naming convention.
func SubWeeks(t time.Time, weeks int) time.Time {
	return AddWeeks(t, -weeks)
}

// SubBusinessDays subtracts the specified number of business days from the given time.
// Weekends (Saturday and Sunday) are skipped.
// Returns a new time.Time instance, leaving the original unchanged (immutable).
// This is the counterpart to AddBusinessDays and follows date-fns naming convention.
func SubBusinessDays(t time.Time, businessDays int) time.Time {
	return AddBusinessDays(t, -businessDays)
}

// AddHours adds the specified number of hours to the given time.
// Returns a new time.Time instance, leaving the original unchanged (immutable).
// Supports negative values to subtract hours.
func AddHours(t time.Time, hours int) time.Time {
	return t.Add(time.Duration(hours) * time.Hour)
}

// SubHours subtracts the specified number of hours from the given time.
// Returns a new time.Time instance, leaving the original unchanged (immutable).
// This is the counterpart to AddHours and follows date-fns naming convention.
func SubHours(t time.Time, hours int) time.Time {
	return AddHours(t, -hours)
}

// AddMinutes adds the specified number of minutes to the given time.
// Returns a new time.Time instance, leaving the original unchanged (immutable).
// Supports negative values to subtract minutes.
func AddMinutes(t time.Time, minutes int) time.Time {
	return t.Add(time.Duration(minutes) * time.Minute)
}

// SubMinutes subtracts the specified number of minutes from the given time.
// Returns a new time.Time instance, leaving the original unchanged (immutable).
// This is the counterpart to AddMinutes and follows date-fns naming convention.
func SubMinutes(t time.Time, minutes int) time.Time {
	return AddMinutes(t, -minutes)
}

// AddSeconds adds the specified number of seconds to the given time.
// Returns a new time.Time instance, leaving the original unchanged (immutable).
// Supports negative values to subtract seconds.
func AddSeconds(t time.Time, seconds int) time.Time {
	return t.Add(time.Duration(seconds) * time.Second)
}

// SubSeconds subtracts the specified number of seconds from the given time.
// Returns a new time.Time instance, leaving the original unchanged (immutable).
// This is the counterpart to AddSeconds and follows date-fns naming convention.
func SubSeconds(t time.Time, seconds int) time.Time {
	return AddSeconds(t, -seconds)
}

// AddMonths adds the specified number of months to the given time.
// Returns a new time.Time instance, leaving the original unchanged (immutable).
// Supports negative values to subtract months.
// If the resulting day doesn't exist in the target month (e.g., Jan 31 + 1 month),
// it will be adjusted to the last day of that month (Feb 28/29).
func AddMonths(t time.Time, months int) time.Time {
	year := t.Year()
	month := t.Month()
	day := t.Day()

	// Calculate target year and month
	targetMonth := int(month) + months
	for targetMonth > 12 {
		year++
		targetMonth -= 12
	}
	for targetMonth < 1 {
		year--
		targetMonth += 12
	}

	// Get the last day of the target month
	lastDayOfMonth := time.Date(year, time.Month(targetMonth+1), 0, 0, 0, 0, 0, time.UTC).Day()

	// Adjust day if it's beyond the last day of the target month
	if day > lastDayOfMonth {
		day = lastDayOfMonth
	}

	// Create the new date with the same time components
	return time.Date(year, time.Month(targetMonth), day,
		t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// SubMonths subtracts the specified number of months from the given time.
// Returns a new time.Time instance, leaving the original unchanged (immutable).
// This is the counterpart to AddMonths and follows date-fns naming convention.
func SubMonths(t time.Time, months int) time.Time {
	return AddMonths(t, -months)
}

// AddYears adds the specified number of years to the given time.
// Returns a new time.Time instance, leaving the original unchanged (immutable).
// Supports negative values to subtract years.
// If the date is Feb 29 and the target year is not a leap year,
// it will be adjusted to Feb 28.
func AddYears(t time.Time, years int) time.Time {
	targetYear := t.Year() + years
	month := t.Month()
	day := t.Day()

	// Handle leap year adjustment for Feb 29
	if month == time.February && day == 29 && !isLeapYear(targetYear) {
		day = 28
	}

	// Create the new date with the same time components
	return time.Date(targetYear, month, day,
		t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// Helper function to check if a year is a leap year
func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// SubYears subtracts the specified number of years from the given time.
// Returns a new time.Time instance, leaving the original unchanged (immutable).
// This is the counterpart to AddYears and follows date-fns naming convention.
func SubYears(t time.Time, years int) time.Time {
	return AddYears(t, -years)
}
