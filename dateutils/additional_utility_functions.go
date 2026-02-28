package dateutils

import (
	"time"
)

// LightFormat provides a lightweight date formatting function with basic tokens.
// This is similar to date-fns lightFormat but with a simplified token set.
// Supported tokens:
//
//	YYYY - 4-digit year
//	MM   - 2-digit month (01-12)
//	DD   - 2-digit day (01-31)
//	HH   - 2-digit hour (00-23)
//	mm   - 2-digit minute (00-59)
//	ss   - 2-digit second (00-59)
//	SSS  - 3-digit millisecond (000-999)
//
// Example:
//
//	LightFormat(time.Date(2019, 2, 11, 14, 0, 0, 0, time.UTC), "YYYY-MM-DD HH:mm:ss")
//	// Returns: "2019-02-11 14:00:00"
func LightFormat(t time.Time, format string) string {
	if format == "" {
		return ""
	}

	result := format

	// Replace tokens with actual values
	year := t.Year()
	month := int(t.Month())
	day := t.Day()
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()
	millisecond := t.Nanosecond() / 1000000

	// Replace year
	result = replaceToken(result, "YYYY", padZero(year, 4))

	// Replace month
	result = replaceToken(result, "MM", padZero(month, 2))

	// Replace day
	result = replaceToken(result, "DD", padZero(day, 2))

	// Replace hour
	result = replaceToken(result, "HH", padZero(hour, 2))

	// Replace minute
	result = replaceToken(result, "mm", padZero(minute, 2))

	// Replace second
	result = replaceToken(result, "ss", padZero(second, 2))

	// Replace millisecond
	result = replaceToken(result, "SSS", padZero(millisecond, 3))

	return result
}

// replaceToken replaces all occurrences of token with value in the string
func replaceToken(s, token, value string) string {
	result := ""
	tokenLen := len(token)
	i := 0

	for i < len(s) {
		if i+tokenLen <= len(s) && s[i:i+tokenLen] == token {
			result += value
			i += tokenLen
		} else {
			result += string(s[i])
			i++
		}
	}

	return result
}

// padZero pads a number with leading zeros to the specified width
func padZero(num, width int) string {
	str := ""
	if num == 0 {
		str = "0"
	} else {
		for num > 0 {
			str = string(rune('0'+(num%10))) + str
			num /= 10
		}
	}

	for len(str) < width {
		str = "0" + str
	}

	return str
}

// RoundToNearestMinutes rounds the given date to the nearest specified number of minutes.
// This is equivalent to date-fns roundToNearestMinutes function.
//
// Example:
//
//	RoundToNearestMinutes(time.Date(2014, 6, 10, 12, 7, 30, 0, time.UTC), 5)
//	// Returns: 2014-06-10 12:05:00 (rounded down to nearest 5 minutes)
//	RoundToNearestMinutes(time.Date(2014, 6, 10, 12, 8, 30, 0, time.UTC), 5)
//	// Returns: 2014-06-10 12:10:00 (rounded up to nearest 5 minutes)
func RoundToNearestMinutes(t time.Time, nearestTo int) time.Time {
	if nearestTo <= 0 {
		return t
	}

	minutes := t.Minute()
	remainder := minutes % nearestTo

	var adjustedMinutes int
	// Use float64 comparison to avoid integer division truncation.
	// e.g. for nearestTo=5, threshold should be 2.5, not 2 (int division).
	if float64(remainder) < float64(nearestTo)/2.0 {
		// Round down
		adjustedMinutes = minutes - remainder
	} else {
		// Round up
		adjustedMinutes = minutes + (nearestTo - remainder)
	}

	// Handle overflow
	hourAdjustment := 0
	if adjustedMinutes >= 60 {
		hourAdjustment = adjustedMinutes / 60
		adjustedMinutes = adjustedMinutes % 60
	}

	result := time.Date(t.Year(), t.Month(), t.Day(), t.Hour()+hourAdjustment, adjustedMinutes, 0, 0, t.Location())

	// Handle day overflow
	if result.Day() != t.Day() {
		result = result.AddDate(0, 0, -1)
		result = time.Date(result.Year(), result.Month(), result.Day(), 23, adjustedMinutes, 0, 0, t.Location())
	}

	return result
}

// StartOfDecade returns the start of the decade for the given date.
// The decade starts with years ending in 0 (e.g., 2020, 2030).
//
// Example:
//
//	StartOfDecade(time.Date(1985, 10, 20, 0, 0, 0, 0, time.UTC))
//	// Returns: 1980-01-01 00:00:00
func StartOfDecade(t time.Time) time.Time {
	year := t.Year()
	decadeStart := (year / 10) * 10
	return time.Date(decadeStart, time.January, 1, 0, 0, 0, 0, t.Location())
}

// EndOfDecade returns the end of the decade for the given date.
// The decade ends with years ending in 9 (e.g., 2029, 2039).
//
// Example:
//
//	EndOfDecade(time.Date(1985, 10, 20, 0, 0, 0, 0, time.UTC))
//	// Returns: 1989-12-31 23:59:59.999
func EndOfDecade(t time.Time) time.Time {
	year := t.Year()
	decadeEnd := ((year / 10) * 10) + 9
	return time.Date(decadeEnd, time.December, 31, 23, 59, 59, 999000000, t.Location())
}

// LastDayOfDecade returns the last day of the decade for the given date.
// The time is set to 00:00:00.000.
//
// Example:
//
//	LastDayOfDecade(time.Date(1985, 10, 20, 0, 0, 0, 0, time.UTC))
//	// Returns: 1989-12-31 00:00:00
func LastDayOfDecade(t time.Time) time.Time {
	year := t.Year()
	decadeEnd := ((year / 10) * 10) + 9
	return time.Date(decadeEnd, time.December, 31, 0, 0, 0, 0, t.Location())
}

// StartOfCentury returns the start of the century for the given date.
// The century starts with years ending in 01 (e.g., 1901, 2001).
//
// Example:
//
//	StartOfCentury(time.Date(1985, 10, 20, 0, 0, 0, 0, time.UTC))
//	// Returns: 1901-01-01 00:00:00
func StartOfCentury(t time.Time) time.Time {
	year := t.Year()
	centuryStart := ((year-1)/100)*100 + 1
	return time.Date(centuryStart, time.January, 1, 0, 0, 0, 0, t.Location())
}

// EndOfCentury returns the end of the century for the given date.
// The century ends with years ending in 00 (e.g., 2000, 2100).
//
// Example:
//
//	EndOfCentury(time.Date(1985, 10, 20, 0, 0, 0, 0, time.UTC))
//	// Returns: 2000-12-31 23:59:59.999
func EndOfCentury(t time.Time) time.Time {
	year := t.Year()
	centuryEnd := ((year-1)/100 + 1) * 100
	return time.Date(centuryEnd, time.December, 31, 23, 59, 59, 999000000, t.Location())
}

// LastDayOfCentury returns the last day of the century for the given date.
// The time is set to 00:00:00.000.
//
// Example:
//
//	LastDayOfCentury(time.Date(1985, 10, 20, 0, 0, 0, 0, time.UTC))
//	// Returns: 2000-12-31 00:00:00
func LastDayOfCentury(t time.Time) time.Time {
	year := t.Year()
	centuryEnd := ((year-1)/100 + 1) * 100
	return time.Date(centuryEnd, time.December, 31, 0, 0, 0, 0, t.Location())
}

// GetDaysInYear returns the number of days in the year of the given date.
// This accounts for leap years.
//
// Example:
//
//	GetDaysInYear(time.Date(2020, 6, 15, 0, 0, 0, 0, time.UTC)) // 366 (leap year)
//	GetDaysInYear(time.Date(2021, 6, 15, 0, 0, 0, 0, time.UTC)) // 365 (regular year)
func GetDaysInYear(t time.Time) int {
	if IsLeapYear(t.Year()) {
		return 366
	}
	return 365
}

// GetDaysInMonth returns the number of days in the month of the given date.
// This accounts for leap years when calculating February.
//
// Example:
//
//	GetDaysInMonth(time.Date(2020, 2, 15, 0, 0, 0, 0, time.UTC)) // 29 (February in leap year)
//	GetDaysInMonth(time.Date(2021, 2, 15, 0, 0, 0, 0, time.UTC)) // 28 (February in regular year)
//	GetDaysInMonth(time.Date(2021, 4, 15, 0, 0, 0, 0, time.UTC)) // 30 (April)
func GetDaysInMonth(t time.Time) int {
	// Get the first day of the next month, then subtract a day to get the last day of current month
	year, month, _ := t.Date()
	firstOfNextMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, t.Location())
	lastOfThisMonth := firstOfNextMonth.AddDate(0, 0, -1)
	return lastOfThisMonth.Day()
}

// IsMonday returns true if the given date is a Monday.
//
// Example:
//
//	IsMonday(time.Date(2014, 9, 1, 0, 0, 0, 0, time.UTC)) // true
func IsMonday(t time.Time) bool {
	return t.Weekday() == time.Monday
}

// IsTuesday returns true if the given date is a Tuesday.
//
// Example:
//
//	IsTuesday(time.Date(2014, 9, 2, 0, 0, 0, 0, time.UTC)) // true
func IsTuesday(t time.Time) bool {
	return t.Weekday() == time.Tuesday
}

// IsWednesday returns true if the given date is a Wednesday.
//
// Example:
//
//	IsWednesday(time.Date(2014, 9, 3, 0, 0, 0, 0, time.UTC)) // true
func IsWednesday(t time.Time) bool {
	return t.Weekday() == time.Wednesday
}

// IsThursday returns true if the given date is a Thursday.
//
// Example:
//
//	IsThursday(time.Date(2014, 9, 4, 0, 0, 0, 0, time.UTC)) // true
func IsThursday(t time.Time) bool {
	return t.Weekday() == time.Thursday
}

// IsFriday returns true if the given date is a Friday.
//
// Example:
//
//	IsFriday(time.Date(2014, 9, 5, 0, 0, 0, 0, time.UTC)) // true
func IsFriday(t time.Time) bool {
	return t.Weekday() == time.Friday
}

// IsSaturday returns true if the given date is a Saturday.
//
// Example:
//
//	IsSaturday(time.Date(2014, 9, 6, 0, 0, 0, 0, time.UTC)) // true
func IsSaturday(t time.Time) bool {
	return t.Weekday() == time.Saturday
}

// IsSunday returns true if the given date is a Sunday.
//
// Example:
//
//	IsSunday(time.Date(2014, 9, 7, 0, 0, 0, 0, time.UTC)) // true
func IsSunday(t time.Time) bool {
	return t.Weekday() == time.Sunday
}
