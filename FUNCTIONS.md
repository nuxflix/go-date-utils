# go-dateutils - Complete Function Reference

A comprehensive date utility library for Go, inspired by date-fns with functional programming principles.

## Features

✅ **Pure and immutable functions** - Original dates are never modified
✅ **Timezone-aware operations** - Full timezone support
✅ **Zero external dependencies** - Only uses Go standard library
✅ **WebAssembly compatible** - Can be compiled to WASM
✅ **Comprehensive test coverage** - Extensively tested with benchmarks

## Installation

```bash
go get github.com/cesar/go-dateutils
```

## Quick Start

```go
import "github.com/cesar/go-dateutils/dateutils"

// Parse and manipulate dates
date, _ := dateutils.Parse("2023-12-25", time.UTC)
nextWeek := dateutils.AddDays(date, 7)
formatted, _ := dateutils.Format(nextWeek, dateutils.DateISO, nil)
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

### `DifferenceInDays(t1, t2 time.Time) int`

Calculate full days between times.

### `DifferenceInDaysFloat(t1, t2 time.Time) float64`

Calculate days with decimal precision.

### `DifferenceInCalendarDays(t1, t2 time.Time) int`

Calculate calendar days (ignores time).

### `DifferenceInBusinessDays(t1, t2 time.Time) int`

Calculate business days between times.

### `DifferenceInWeeks(t1, t2 time.Time) int`

Calculate full weeks between times.

### `DifferenceInWeeksFloat(t1, t2 time.Time) float64`

Calculate weeks with decimal precision.

### `AbsDifferenceInDays(t1, t2 time.Time) int`

Absolute difference in days (always positive).

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

## 📋 Testing

Run tests:

```bash
go test ./dateutils
go test -v ./dateutils  # verbose
go test -bench=. ./dateutils  # benchmarks
```

## 📄 License

MIT License - see LICENSE file for details.

---

**Total Functions Implemented: 40+**
**Test Coverage: Comprehensive**
**Dependencies: Zero (only Go standard library)**
