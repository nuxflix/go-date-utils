package dateutils

// Constants for date/time calculations
// These constants match the date-fns library for compatibility

const (
	// Time constants
	MillisecondsInSecond = 1000
	SecondsInMinute      = 60
	MinutesInHour        = 60
	HoursInDay           = 24

	// Date constants
	DaysInWeek     = 7
	MonthsInYear   = 12
	QuartersInYear = 4

	// Average calculations (for estimation purposes)
	DaysInYear  = 365.2425  // Average considering leap years
	DaysInMonth = 30.436875 // Average days per month (365.2425 / 12)
	WeeksInYear = 52.1775   // Average weeks per year (365.2425 / 7)

	// Duration constants in milliseconds (for compatibility with date-fns)
	MillisecondsInMinute = MillisecondsInSecond * SecondsInMinute // 60,000
	MillisecondsInHour   = MillisecondsInMinute * MinutesInHour   // 3,600,000
	MillisecondsInDay    = MillisecondsInHour * HoursInDay        // 86,400,000
	MillisecondsInWeek   = MillisecondsInDay * DaysInWeek         // 604,800,000
)

// MaxTime represents the maximum time value that can be represented.
// This matches JavaScript's Date maximum value.
var MaxTime = GetMaxTime()

// MinTime represents the minimum time value that can be represented.
// This matches JavaScript's Date minimum value.
var MinTime = GetMinTime()

// GetMaxTime returns the maximum representable time.
// JavaScript Date max: 8,640,000,000,000,000 milliseconds from Unix epoch
func GetMaxTime() int64 {
	return 8640000000000000 // milliseconds
}

// GetMinTime returns the minimum representable time.
// JavaScript Date min: -8,640,000,000,000,000 milliseconds from Unix epoch
func GetMinTime() int64 {
	return -8640000000000000 // milliseconds
}

// IsValidTimestamp checks if a timestamp is within the valid range.
func IsValidTimestamp(timestamp int64) bool {
	return timestamp >= MinTime && timestamp <= MaxTime
}

// Weekday helper constants for better readability
const (
	Sunday    = 0
	Monday    = 1
	Tuesday   = 2
	Wednesday = 3
	Thursday  = 4
	Friday    = 5
	Saturday  = 6
)

// Month helper constants
const (
	January   = 1
	February  = 2
	March     = 3
	April     = 4
	May       = 5
	June      = 6
	July      = 7
	August    = 8
	September = 9
	October   = 10
	November  = 11
	December  = 12
)

// Quarter helper constants
const (
	Q1 = 1
	Q2 = 2
	Q3 = 3
	Q4 = 4
)
