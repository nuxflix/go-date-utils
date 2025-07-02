# go-dateutils - Complete Function Reference

A comprehensive date utility library for Go, inspired by date-fns with functional programming principles.

## Features

✅ **Pure and immutable functions** - Original dates are never modified
✅ **Timezone-aware operations** - Full timezone support
✅ **Zero external dependencies** - Only uses Go standard library
✅ **WebAssembly compatible** - Can be compiled to WASM
✅ **Comprehensive test coverage** - Extensively tested with benchmarks
✅ **140+ functions** - Comprehensive API matching date-fns capabilities

## Installation

```bash
go get github.com/chmenegatti/go-date-fns
```

## Quick Start

```go
import "github.com/chmenegatti/go-date-fns/dateutils"

// Parse and manipulate dates
date, _ := dateutils.Parse("2024-12-25", time.UTC)
nextWeek := dateutils.AddDays(date, 7)
formatted, _ := dateutils.Format(nextWeek, dateutils.DateISO, nil)

// Use Get and Set functions
year := dateutils.GetYear(date)           // 2024
newYear := dateutils.SetYear(date, 2025)  // Change to 2025
```

## Function Categories

## 📅 Parsing Functions

### `Parse(dateStr string, timezone *time.Location) (time.Time, error)`
Parse a date string using common formats automatically.

### `ParseISO(isoStr string, timezone *time.Location) (time.Time, error)`
Parse an ISO 8601 formatted date string.

### `ParseWithFormat(dateStr, format string, timezone *time.Location) (time.Time, error)`
Parse a date string using a specific format.

### `IsValidISO(isoStr string) bool`
Check if a string is valid ISO 8601 format.

---

## 📝 Formatting Functions

### `Format(t time.Time, format string, timezone *time.Location) (string, error)`
Format a time using predefined format constants.

### `FormatCustom(t time.Time, customFormat string, timezone *time.Location) (string, error)`
Format using user-friendly placeholders (YYYY-MM-DD, HH:mm, etc.).

### `FormatSafe(t time.Time, format string, timezone *time.Location) string`
Safe formatting that returns empty string on error.

### `FormatWithDefault(t time.Time, format string, timezone *time.Location, defaultValue string) string`
Formatting with fallback default value.

---

## 📊 Get Functions

### Date Components
- `GetYear(t time.Time) int` - Get year component
- `GetMonth(t time.Time) int` - Get month (1-12)
- `GetDate(t time.Time) int` - Get day of month (1-31)
- `GetDay(t time.Time) int` - Get day of week (0-6, Sunday=0)
- `GetDayOfYear(t time.Time) int` - Get day of year (1-366)
- `GetQuarter(t time.Time) int` - Get quarter (1-4)

### Time Components
- `GetHours(t time.Time) int` - Get hours (0-23)
- `GetMinutes(t time.Time) int` - Get minutes (0-59)
- `GetSeconds(t time.Time) int` - Get seconds (0-59)
- `GetMilliseconds(t time.Time) int` - Get milliseconds (0-999)

### Week Functions
- `GetWeek(t time.Time) int` - Get ISO week number (1-53)
- `GetWeekOfMonth(t time.Time) int` - Get week of month (1-6)
- `GetWeekYear(t time.Time) int` - Get ISO week year

---

## ⚙️ Set Functions

### Date Components
- `SetYear(t time.Time, year int) time.Time` - Set year
- `SetMonth(t time.Time, month int) time.Time` - Set month (1-12)
- `SetDate(t time.Time, day int) time.Time` - Set day of month
- `SetDay(t time.Time, day int) time.Time` - Set day of week
- `SetDayOfYear(t time.Time, dayOfYear int) time.Time` - Set day of year
- `SetQuarter(t time.Time, quarter int) time.Time` - Set quarter

### Time Components
- `SetHours(t time.Time, hours int) time.Time` - Set hours
- `SetMinutes(t time.Time, minutes int) time.Time` - Set minutes
- `SetSeconds(t time.Time, seconds int) time.Time` - Set seconds
- `SetMilliseconds(t time.Time, milliseconds int) time.Time` - Set milliseconds

---

## 📝 Formatting Functions

### `Format(t time.Time, format string, timezone *time.Location) (string, error)`

Format a time using Go's reference time layout.

### `FormatCustom(t time.Time, customFormat string, timezone *time.Location) (string, error)`

Format using user-friendly placeholders (YYYY, MM, DD, etc.).

### `FormatSafe(t time.Time, format string, timezone *time.Location) string`

Format without errors - returns empty string for invalid input.

### `FormatWithDefault(t time.Time, format string, timezone *time.Location, defaultValue string) string`

Format with fallback default value.

**Format Constants Available:**

- `DateISO` - "2006-01-02"
- `DateUS` - "01/02/2006"
- `DateTime24` - "2006-01-02 15:04:05"
- `Readable` - "January 2, 2006"
- And many more...

---

## 🔍 Comparison Functions

### `IsBefore(t1, t2 time.Time) bool`

Check if t1 is before t2.

### `IsBeforeOrEqual(t1, t2 time.Time) bool`

Check if t1 is before or equal to t2.

### `IsBeforeDate(t1, t2 time.Time) bool`

Compare only dates, ignoring time components.

### `IsBeforeInTimezone(t1, t2 time.Time, timezone *time.Location) bool`

Compare times in specific timezone.

### `IsAfter(t1, t2 time.Time) bool`

Check if t1 is after t2.

### `IsAfterOrEqual(t1, t2 time.Time) bool`

Check if t1 is after or equal to t2.

### `IsAfterDate(t1, t2 time.Time) bool`

Compare only dates, ignoring time components.

### `IsAfterInTimezone(t1, t2 time.Time, timezone *time.Location) bool`

Compare times in specific timezone.

---

## ➕ Manipulation Functions

### `AddDays(t time.Time, days int) time.Time`

Add or subtract days (negative values subtract).

### `AddBusinessDays(t time.Time, businessDays int) time.Time`

Add business days, skipping weekends.

### `AddWeeks(t time.Time, weeks int) time.Time`

Add or subtract weeks.

### `AddDaysWithTimezone(t time.Time, days int, timezone *time.Location) time.Time`

Add days and convert to specified timezone.

---

## 📏 Difference Functions

### Time Unit Differences (Integer)

#### `DifferenceInSeconds(t1, t2 time.Time) int`
Calculate full seconds between times.

#### `DifferenceInMinutes(t1, t2 time.Time) int`
Calculate full minutes between times.

#### `DifferenceInHours(t1, t2 time.Time) int`
Calculate full hours between times.

#### `DifferenceInDays(t1, t2 time.Time) int`
Calculate full days between times.

#### `DifferenceInWeeks(t1, t2 time.Time) int`
Calculate full weeks between times.

#### `DifferenceInMonths(t1, t2 time.Time) int`
Calculate full months between times.

#### `DifferenceInYears(t1, t2 time.Time) int`
Calculate full years between times.

#### `DifferenceInQuarters(t1, t2 time.Time) int`
Calculate full quarters between times.

### Calendar-Based Differences

#### `DifferenceInCalendarDays(t1, t2 time.Time) int`
Calculate calendar days (ignores time).

#### `DifferenceInCalendarMonths(t1, t2 time.Time) int`
Calculate calendar months (ignores day).

#### `DifferenceInCalendarYears(t1, t2 time.Time) int`
Calculate calendar years (ignores month/day).

#### `DifferenceInCalendarQuarters(t1, t2 time.Time) int`
Calculate calendar quarters (ignores day).

### Business Days

#### `DifferenceInBusinessDays(t1, t2 time.Time) int`
Calculate business days between times (excludes weekends).

### Floating-Point Precision

#### `DifferenceInSecondsFloat(t1, t2 time.Time) float64`
Calculate seconds with decimal precision.

#### `DifferenceInMinutesFloat(t1, t2 time.Time) float64`
Calculate minutes with decimal precision.

#### `DifferenceInHoursFloat(t1, t2 time.Time) float64`
Calculate hours with decimal precision.

#### `DifferenceInDaysFloat(t1, t2 time.Time) float64`
Calculate days with decimal precision.

#### `DifferenceInWeeksFloat(t1, t2 time.Time) float64`
Calculate weeks with decimal precision.

### Absolute Differences (Always Positive)

#### `AbsDifferenceInSeconds(t1, t2 time.Time) int`
Absolute difference in seconds.

#### `AbsDifferenceInMinutes(t1, t2 time.Time) int`
Absolute difference in minutes.

#### `AbsDifferenceInHours(t1, t2 time.Time) int`
Absolute difference in hours.

#### `AbsDifferenceInDays(t1, t2 time.Time) int`

Absolute difference in days (always positive).

#### `AbsDifferenceInMonths(t1, t2 time.Time) int`

Absolute difference in months (always positive).

#### `AbsDifferenceInYears(t1, t2 time.Time) int`

Absolute difference in years (always positive).

---

## ✅ Validation Functions

### `IsValid(t time.Time) bool`

Check if time is valid (not zero value).

### `IsLeapYear(year int) bool`

Check if year is a leap year.

### `IsWeekend(t time.Time) bool`

Check if time falls on weekend.

### `IsWeekday(t time.Time) bool`

Check if time falls on weekday.

### `IsToday(t time.Time, timezone *time.Location) bool`

Check if time is today in specified timezone.

### `IsTomorrow(t time.Time, timezone *time.Location) bool`

Check if time is tomorrow.

### `IsYesterday(t time.Time, timezone *time.Location) bool`

Check if time is yesterday.

### `IsWithinInterval(t, start, end time.Time) bool`

Check if time falls within interval.

### `IsFirstDayOfMonth(t time.Time) bool`

Check if time is first day of month.

### `IsLastDayOfMonth(t time.Time) bool`

Check if time is last day of month.

### `IsFirstDayOfYear(t time.Time) bool`

Check if time is January 1st.

### `IsLastDayOfYear(t time.Time) bool`

Check if time is December 31st.

### `IsSameDate(t1, t2 time.Time) bool`

Check if times have same date.

### `IsSameMonth(t1, t2 time.Time) bool`

Check if times are in same month.

### `IsSameYear(t1, t2 time.Time) bool`

Check if times are in same year.

---

## 📅 Period Utility Functions

### `StartOfDay(t time.Time) time.Time`

Get start of day (00:00:00).

### `EndOfDay(t time.Time) time.Time`

Get end of day (23:59:59.999999999).

### `StartOfWeek(t time.Time) time.Time`

Get start of week (Monday 00:00:00).

### `StartOfWeekSunday(t time.Time) time.Time`

Get start of week (Sunday 00:00:00).

### `EndOfWeek(t time.Time) time.Time`

Get end of week (Sunday 23:59:59.999999999).

### `EndOfWeekSaturday(t time.Time) time.Time`

Get end of week (Saturday 23:59:59.999999999).

### `StartOfMonth(t time.Time) time.Time`

Get start of month (1st day 00:00:00).

### `EndOfMonth(t time.Time) time.Time`

Get end of month (last day 23:59:59.999999999).

### `StartOfYear(t time.Time) time.Time`

Get start of year (January 1st 00:00:00).

### `EndOfYear(t time.Time) time.Time`

Get end of year (December 31st 23:59:59.999999999).

### `StartOfHour(t time.Time) time.Time`

Get start of hour (XX:00:00).

### `EndOfHour(t time.Time) time.Time`

Get end of hour (XX:59:59.999999999).

### `StartOfMinute(t time.Time) time.Time`

Get start of minute (XX:XX:00).

### `EndOfMinute(t time.Time) time.Time`

Get end of minute (XX:XX:59.999999999).

---

## 🌍 Timezone Support

All functions that accept timezone parameters:

- Use UTC if `nil` is passed
- Convert results to specified timezone when provided
- Preserve original timezone when no timezone specified

## 🚀 Performance

The library is designed for performance with:

- Minimal allocations
- Pure functions (no side effects)
- Efficient algorithms
- Comprehensive benchmarks

## � Interval Utilities

**NEW:** Generate arrays of dates within intervals - perfect for reporting, scheduling, and data analysis.

### `Interval` Type

```go
type Interval struct {
    Start time.Time
    End   time.Time
}
```

### Each Functions

#### `EachDayOfInterval(interval Interval) []time.Time`
Returns all days within the interval (start of each day).

#### `EachWeekOfInterval(interval Interval) []time.Time`
Returns all week starts within the interval (Monday-based).

#### `EachWeekOfIntervalSunday(interval Interval) []time.Time`
Returns all week starts within the interval (Sunday-based).

#### `EachMonthOfInterval(interval Interval) []time.Time`
Returns all month starts within the interval.

#### `EachYearOfInterval(interval Interval) []time.Time`
Returns all year starts within the interval.

#### `EachQuarterOfInterval(interval Interval) []time.Time`
Returns all quarter starts within the interval.

#### `EachHourOfInterval(interval Interval) []time.Time`
Returns all hour starts within the interval.

#### `EachMinuteOfInterval(interval Interval) []time.Time`
Returns all minute starts within the interval.

#### `EachWeekendOfInterval(interval Interval) []time.Time`
Returns all weekend days (Saturday & Sunday) within the interval.

#### `EachBusinessDayOfInterval(interval Interval) []time.Time`
Returns all business days (Monday-Friday) within the interval.

### Quarter Utilities

#### `StartOfQuarter(t time.Time) time.Time`
Get start of quarter (Q1=Jan, Q2=Apr, Q3=Jul, Q4=Oct).

#### `EndOfQuarter(t time.Time) time.Time`
Get end of quarter with precise timing.

**Examples:**

```go
// Get all weekends in January 2024
interval := dateutils.Interval{
    Start: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
    End:   time.Date(2024, 1, 31, 23, 59, 59, 0, time.UTC),
}
weekends := dateutils.EachWeekendOfInterval(interval)

// Get all business days in a work week
workWeek := dateutils.Interval{
    Start: time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC),  // Monday
    End:   time.Date(2024, 1, 12, 23, 59, 59, 0, time.UTC), // Friday
}
businessDays := dateutils.EachBusinessDayOfInterval(workWeek)

// Get all quarters in 2024
year := dateutils.Interval{
    Start: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
    End:   time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC),
}
quarters := dateutils.EachQuarterOfInterval(year)
```

---

## �📋 Testing

Run tests:

```bash
go test ./dateutils
go test -v ./dateutils  # verbose
go test -bench=. ./dateutils  # benchmarks
```

## 📄 License

MIT License - see LICENSE file for details.

---

**Total Functions Implemented: 140+**
**Test Coverage: Comprehensive**
**Dependencies: Zero (only Go standard library)**
**API Compatibility: Close to date-fns with Go idioms**
