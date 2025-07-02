package dateutils

import (
	"time"
)

// GetUnixTime returns the Unix timestamp (seconds since January 1, 1970 UTC) for the given date.
// This is equivalent to date-fns getUnixTime function.
//
// Example:
//
//	GetUnixTime(time.Date(2012, time.February, 29, 11, 45, 5, 0, time.UTC)) // 1330512305
func GetUnixTime(t time.Time) int64 {
	return t.Unix()
}

// FromUnixTime creates a time.Time from a Unix timestamp (seconds since January 1, 1970 UTC).
// This is equivalent to date-fns fromUnixTime function.
// Decimal values will be discarded.
//
// Example:
//
//	FromUnixTime(1330515905) // Wed Feb 29 2012 11:45:05
func FromUnixTime(unixTime int64) time.Time {
	return time.Unix(unixTime, 0).UTC()
}

// FromUnixTimeInLocation creates a time.Time from a Unix timestamp in the specified location.
// If location is nil, UTC is used.
//
// Example:
//
//	loc, _ := time.LoadLocation("America/New_York")
//	FromUnixTimeInLocation(1330515905, loc)
func FromUnixTimeInLocation(unixTime int64, location *time.Location) time.Time {
	if location == nil {
		location = time.UTC
	}
	return time.Unix(unixTime, 0).In(location)
}

// GetTime returns the timestamp in milliseconds since January 1, 1970 UTC.
// This is equivalent to JavaScript's Date.getTime() or Date.now().
//
// Example:
//
//	GetTime(time.Date(2012, time.February, 29, 11, 45, 5, 123000000, time.UTC)) // 1330512305123
func GetTime(t time.Time) int64 {
	return t.UnixMilli()
}

// FromTime creates a time.Time from a timestamp in milliseconds since January 1, 1970 UTC.
// This is the inverse of GetTime.
//
// Example:
//
//	FromTime(1330512305123) // Wed Feb 29 2012 11:45:05.123
func FromTime(milliseconds int64) time.Time {
	return time.UnixMilli(milliseconds).UTC()
}

// FromTimeInLocation creates a time.Time from a timestamp in milliseconds in the specified location.
// If location is nil, UTC is used.
//
// Example:
//
//	loc, _ := time.LoadLocation("America/New_York")
//	FromTimeInLocation(1330512305123, loc)
func FromTimeInLocation(milliseconds int64, location *time.Location) time.Time {
	if location == nil {
		location = time.UTC
	}
	return time.UnixMilli(milliseconds).In(location)
}

// IsThisSecond returns true if the given date is in the current second.
//
// Example:
//
//	IsThisSecond(time.Now()) // true
func IsThisSecond(t time.Time) bool {
	now := time.Now()
	return now.Year() == t.Year() &&
		now.YearDay() == t.YearDay() &&
		now.Hour() == t.Hour() &&
		now.Minute() == t.Minute() &&
		now.Second() == t.Second()
}

// IsThisMinute returns true if the given date is in the current minute.
//
// Example:
//
//	IsThisMinute(time.Now()) // true
func IsThisMinute(t time.Time) bool {
	now := time.Now()
	return now.Year() == t.Year() &&
		now.YearDay() == t.YearDay() &&
		now.Hour() == t.Hour() &&
		now.Minute() == t.Minute()
}

// IsThisHour returns true if the given date is in the current hour.
//
// Example:
//
//	IsThisHour(time.Now()) // true
func IsThisHour(t time.Time) bool {
	now := time.Now()
	return now.Year() == t.Year() &&
		now.YearDay() == t.YearDay() &&
		now.Hour() == t.Hour()
}

// IsThisWeek returns true if the given date is in the current week (Monday-Sunday).
//
// Example:
//
//	IsThisWeek(time.Now()) // true
func IsThisWeek(t time.Time) bool {
	now := time.Now()
	return IsSameWeek(t, now)
}

// IsThisMonth returns true if the given date is in the current month.
//
// Example:
//
//	IsThisMonth(time.Now()) // true
func IsThisMonth(t time.Time) bool {
	now := time.Now()
	return IsSameMonth(t, now)
}

// IsThisQuarter returns true if the given date is in the current quarter.
//
// Example:
//
//	IsThisQuarter(time.Now()) // true
func IsThisQuarter(t time.Time) bool {
	now := time.Now()
	return GetQuarter(t) == GetQuarter(now) && t.Year() == now.Year()
}

// IsThisYear returns true if the given date is in the current year.
//
// Example:
//
//	IsThisYear(time.Now()) // true
func IsThisYear(t time.Time) bool {
	now := time.Now()
	return IsSameYear(t, now)
}

// IsThisISOWeek returns true if the given date is in the current ISO week.
//
// Example:
//
//	IsThisISOWeek(time.Now()) // true
func IsThisISOWeek(t time.Time) bool {
	now := time.Now()
	return IsSameISOWeek(t, now)
}

// IsThisISOWeekYear returns true if the given date is in the current ISO week-numbering year.
//
// Example:
//
//	IsThisISOWeekYear(time.Now()) // true
func IsThisISOWeekYear(t time.Time) bool {
	now := time.Now()
	return IsSameISOWeekYear(t, now)
}
