# 🗓️ go-dateutils

> A modern, functional date utility library for Go, inspired by the beloved date-fns JavaScript library.

[![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.19-007d9c.svg)](https://golang.org/)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/chmenegatti/go-date-fns)](https://goreportcard.com/report/github.com/chmenegatti/go-date-fns)
[![Test Coverage](https://img.shields.io/badge/coverage-100%25-brightgreen.svg)](https://github.com/chmenegatti/go-date-fns)

**go-dateutils** brings the power and elegance of functional date manipulation to Go. With over **40+ pure, immutable functions**, it provides a comprehensive toolkit for working with dates and times in a safe, predictable, and timezone-aware manner.

---

## ✨ Why go-dateutils?

In the world of Go development, working with dates can be complex and error-prone. **go-dateutils** solves this by providing:

- 🧮 **Pure Functions**: No side effects, original dates never modified
- 🌍 **Timezone Aware**: Full timezone support across all operations  
- 🚀 **Zero Dependencies**: Only uses Go standard library
- 🎯 **Type Safe**: Leverages Go's strong typing system
- 📦 **WebAssembly Ready**: Compile to WASM without issues
- 🧪 **Battle Tested**: 100% test coverage with comprehensive benchmarks
- 📚 **Well Documented**: Every function has clear examples and documentation

---

## 🚀 Quick Start

### Installation

```bash
go get github.com/chmenegatti/go-date-fns
```

### Basic Usage

```go
package main

import (
    "fmt"
    "time"
    
    "github.com/chmenegatti/go-date-fns/dateutils"
)

func main() {
    // Parse a date
    date, _ := dateutils.Parse("2024-01-15", time.UTC)
    
    // Add 7 days
    nextWeek := dateutils.AddDays(date, 7)
    
    // Format the result
    formatted, _ := dateutils.Format(nextWeek, dateutils.DateISO, nil)
    fmt.Println(formatted) // Output: 2024-01-22
    
    // Check if it's a weekend
    isWeekend := dateutils.IsWeekend(nextWeek)
    fmt.Printf("Is weekend: %t\n", isWeekend)
}
```

---

## 📚 Complete Feature Overview

### 🔍 **Parsing Functions**

Convert strings to time.Time objects with intelligent format detection:

```go
// Smart parsing with common formats
date, err := dateutils.Parse("2024-01-15", time.UTC)
date, err := dateutils.Parse("01/15/2024", time.UTC)
date, err := dateutils.Parse("January 15, 2024", time.UTC)

// ISO 8601 parsing
isoDate, err := dateutils.ParseISO("2024-01-15T14:30:00Z", time.UTC)

// Custom format parsing
customDate, err := dateutils.ParseWithFormat("15-01-2024", "02-01-2006", time.UTC)

// Validation
isValid := dateutils.IsValidISO("2024-01-15T14:30:00Z")
```

### 🎨 **Formatting Functions**

Transform dates into human-readable strings:

```go
date := time.Date(2024, 1, 15, 14, 30, 0, 0, time.UTC)

// Predefined formats
iso, _ := dateutils.Format(date, dateutils.DateISO, nil)           // "2024-01-15"
us, _ := dateutils.Format(date, dateutils.DateUS, nil)             // "01/15/2024"
readable, _ := dateutils.Format(date, dateutils.Readable, nil)     // "January 15, 2024"
dateTime, _ := dateutils.Format(date, dateutils.DateTime24, nil)   // "2024-01-15 14:30:00"

// Custom formatting with user-friendly placeholders
custom, _ := dateutils.FormatCustom(date, "DD/MM/YYYY at HH:mm", nil)  // "15/01/2024 at 14:30"
custom, _ := dateutils.FormatCustom(date, "MMMM DD, YYYY", nil)        // "January 15, 2024"

// Safe formatting (returns empty string on error)
safe := dateutils.FormatSafe(date, dateutils.DateISO, nil)

// Formatting with fallback
withDefault := dateutils.FormatWithDefault(date, dateutils.DateISO, nil, "N/A")
```

### ⚖️ **Comparison Functions**

Compare dates with precision and timezone awareness:

```go
date1 := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
date2 := time.Date(2024, 1, 16, 15, 0, 0, 0, time.UTC)

// Precise comparisons (includes time)
dateutils.IsBefore(date1, date2)        // true
dateutils.IsAfter(date1, date2)         // false
dateutils.IsBeforeOrEqual(date1, date1) // true

// Date-only comparisons (ignores time)
dateutils.IsBeforeDate(date1, date2)    // true
dateutils.IsAfterDate(date1, date2)     // false

// Timezone-aware comparisons
est, _ := time.LoadLocation("America/New_York")
dateutils.IsBeforeInTimezone(date1, date2, est)
```

### 🔧 **Manipulation Functions**

Add, subtract, and modify dates immutably:

```go
baseDate := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC) // Monday

// Add/subtract days
nextWeek := dateutils.AddDays(baseDate, 7)         // 2024-01-22
lastWeek := dateutils.AddDays(baseDate, -7)        // 2024-01-08

// Business days (skips weekends)
nextBusinessDay := dateutils.AddBusinessDays(baseDate, 5)  // Skips weekend, lands on weekday

// Add weeks
nextMonth := dateutils.AddWeeks(baseDate, 4)       // ~4 weeks later

// Timezone-aware addition
est, _ := time.LoadLocation("America/New_York")
dateInEST := dateutils.AddDaysWithTimezone(baseDate, 1, est)
```

### 📏 **Difference Functions**

Calculate time differences with various precision levels:

```go
start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
end := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)

// Integer differences
days := dateutils.DifferenceInDays(end, start)                    // 14
calendarDays := dateutils.DifferenceInCalendarDays(end, start)    // 14
businessDays := dateutils.DifferenceInBusinessDays(end, start)    // 10 (excluding weekends)
weeks := dateutils.DifferenceInWeeks(end, start)                  // 2

// Floating-point precision
daysPrecise := dateutils.DifferenceInDaysFloat(end, start)        // 14.5
weeksPrecise := dateutils.DifferenceInWeeksFloat(end, start)      // 2.07...

// Absolute differences (always positive)
absDays := dateutils.AbsDifferenceInDays(start, end)              // 14
```

### ✅ **Validation Functions**

Comprehensive date validation and checking:

```go
date := time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC) // Leap year date
today := time.Now()

// Basic validation
dateutils.IsValid(date)                    // true
dateutils.IsValid(time.Time{})             // false (zero time)

// Year validation
dateutils.IsLeapYear(2024)                 // true
dateutils.IsLeapYear(2023)                 // false

// Day type validation
dateutils.IsWeekend(date)                  // true/false based on day
dateutils.IsWeekday(date)                  // opposite of IsWeekend

// Time-based validation
dateutils.IsToday(today)                   // true
dateutils.IsTomorrow(dateutils.AddDays(today, 1))   // true
dateutils.IsYesterday(dateutils.AddDays(today, -1)) // true

// Interval validation
start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
end := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)
test := time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC)
dateutils.IsWithinInterval(test, start, end)        // true

// Boundary validation
dateutils.IsFirstDayOfMonth(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC))  // true
dateutils.IsLastDayOfMonth(time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC))  // true
dateutils.IsFirstDayOfYear(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC))   // true
dateutils.IsLastDayOfYear(time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC))  // true

// Comparison validation
date1 := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
date2 := time.Date(2024, 1, 15, 15, 0, 0, 0, time.UTC)
dateutils.IsSameDate(date1, date2)         // true (same date, different time)
dateutils.IsSameMonth(date1, date2)        // true
dateutils.IsSameYear(date1, date2)         // true
```

### 📅 **Period Utility Functions**

Get the start and end of various time periods:

```go
date := time.Date(2024, 1, 15, 14, 30, 45, 0, time.UTC) // Monday, mid-day

// Day boundaries
startOfDay := dateutils.StartOfDay(date)    // 2024-01-15 00:00:00
endOfDay := dateutils.EndOfDay(date)        // 2024-01-15 23:59:59

// Week boundaries (Monday-based)
startOfWeek := dateutils.StartOfWeek(date)  // Previous Monday 00:00:00
endOfWeek := dateutils.EndOfWeek(date)      // Next Sunday 23:59:59

// Week boundaries (Sunday-based)
startOfWeekSunday := dateutils.StartOfWeekSunday(date)  // Previous Sunday 00:00:00
endOfWeekSaturday := dateutils.EndOfWeekSaturday(date)  // Next Saturday 23:59:59

// Month boundaries
startOfMonth := dateutils.StartOfMonth(date) // 2024-01-01 00:00:00
endOfMonth := dateutils.EndOfMonth(date)     // 2024-01-31 23:59:59

// Year boundaries
startOfYear := dateutils.StartOfYear(date)   // 2024-01-01 00:00:00
endOfYear := dateutils.EndOfYear(date)       // 2024-12-31 23:59:59

// Hour boundaries
startOfHour := dateutils.StartOfHour(date)   // 2024-01-15 14:00:00
endOfHour := dateutils.EndOfHour(date)       // 2024-01-15 14:59:59

// Minute boundaries
startOfMinute := dateutils.StartOfMinute(date) // 2024-01-15 14:30:00
endOfMinute := dateutils.EndOfMinute(date)     // 2024-01-15 14:30:59
```

---

## 🌍 Timezone Support

**go-dateutils** provides first-class timezone support across all functions:

```go
// Load timezones
utc := time.UTC
est, _ := time.LoadLocation("America/New_York")
pst, _ := time.LoadLocation("America/Los_Angeles")
tokyo, _ := time.LoadLocation("Asia/Tokyo")

// Parse with timezone
date, _ := dateutils.ParseISO("2024-01-15T10:00:00Z", est)

// Format in different timezones
utcFormat, _ := dateutils.Format(date, dateutils.DateTime24, utc)
estFormat, _ := dateutils.Format(date, dateutils.DateTime24, est)
pstFormat, _ := dateutils.Format(date, dateutils.DateTime24, pst)

// Timezone-aware comparisons
isAfterInEST := dateutils.IsAfterInTimezone(date1, date2, est)

// Add days with timezone conversion
futureInTokyo := dateutils.AddDaysWithTimezone(date, 7, tokyo)
```

---

## 🎯 Real-World Examples

### 📊 Business Date Calculations

```go
// Calculate project deadline (20 business days from today)
today := time.Now()
deadline := dateutils.AddBusinessDays(today, 20)
fmt.Printf("Project deadline: %s\n", 
    dateutils.FormatSafe(deadline, dateutils.Readable, nil))

// Calculate working days between two dates
startDate, _ := dateutils.Parse("2024-01-01", time.UTC)
endDate, _ := dateutils.Parse("2024-01-31", time.UTC)
workingDays := dateutils.DifferenceInBusinessDays(endDate, startDate)
fmt.Printf("Working days in January 2024: %d\n", workingDays)
```

### 📅 Event Planning

```go
// Plan a weekly meeting series
meetingStart, _ := dateutils.ParseISO("2024-01-15T14:00:00Z", time.UTC)
meetings := make([]time.Time, 0, 12)

for i := 0; i < 12; i++ {
    meeting := dateutils.AddWeeks(meetingStart, i)
    
    // Skip if it falls on a weekend
    if dateutils.IsWeekend(meeting) {
        meeting = dateutils.AddDays(meeting, 1) // Move to Monday
    }
    
    meetings = append(meetings, meeting)
}

// Display the meeting schedule
for i, meeting := range meetings {
    formatted, _ := dateutils.FormatCustom(meeting, "MMMM DD, YYYY at HH:mm", nil)
    fmt.Printf("Meeting %d: %s\n", i+1, formatted)
}
```

### 🌐 Multi-Timezone Application

```go
// Handle global user base
users := []struct {
    Name     string
    Timezone string
    LastSeen time.Time
}{
    {"Alice", "America/New_York", time.Now().Add(-2 * time.Hour)},
    {"Bob", "Europe/London", time.Now().Add(-5 * time.Hour)},
    {"Charlie", "Asia/Tokyo", time.Now().Add(-1 * time.Hour)},
}

for _, user := range users {
    userTZ, _ := time.LoadLocation(user.Timezone)
    
    // Format last seen in user's timezone
    lastSeenLocal, _ := dateutils.Format(user.LastSeen, dateutils.DateTime12, userTZ)
    
    // Check if user was active today in their timezone
    wasActiveToday := dateutils.IsSameDate(user.LastSeen.In(userTZ), time.Now().In(userTZ))
    
    fmt.Printf("%s (in %s): Last seen %s, Active today: %t\n",
        user.Name, user.Timezone, lastSeenLocal, wasActiveToday)
}
```

### 📈 Analytics and Reporting

```go
// Generate monthly report data
reportStart := dateutils.StartOfMonth(time.Now())
reportEnd := dateutils.EndOfMonth(time.Now())

// Calculate metrics
totalDays := dateutils.DifferenceInCalendarDays(reportEnd, reportStart) + 1
workingDays := dateutils.DifferenceInBusinessDays(reportEnd, reportStart) + 1
weekends := totalDays - workingDays

fmt.Printf("📊 Monthly Report - %s\n", 
    dateutils.FormatSafe(reportStart, "MMMM YYYY", nil))
fmt.Printf("Total days: %d\n", totalDays)
fmt.Printf("Working days: %d\n", workingDays)
fmt.Printf("Weekend days: %d\n", weekends)

// Find all Mondays in the month
current := reportStart
mondays := []time.Time{}

for !dateutils.IsAfter(current, reportEnd) {
    if current.Weekday() == time.Monday {
        mondays = append(mondays, current)
    }
    current = dateutils.AddDays(current, 1)
}

fmt.Printf("Mondays this month: %d\n", len(mondays))
```

---

## 🏗️ Architecture & Design Principles

### Functional Programming

All functions in **go-dateutils** follow functional programming principles:

- **Pure Functions**: No side effects or global state mutations
- **Immutability**: Original time.Time values are never modified
- **Predictability**: Same inputs always produce same outputs
- **Composability**: Functions can be easily chained and combined

```go
// Example of function composition
date, _ := dateutils.Parse("2024-01-15", time.UTC)
result := dateutils.FormatSafe(
    dateutils.StartOfWeek(
        dateutils.AddBusinessDays(date, 5),
    ),
    dateutils.DateISO,
    nil,
)
```

### Type Safety

Leverages Go's strong type system to prevent common date-related errors:

```go
// Compile-time safety
var days int = 7
var weeks int = 2

nextWeek := dateutils.AddDays(date, days)     // ✅ Correct
nextMonth := dateutils.AddWeeks(date, weeks)  // ✅ Correct
// invalid := dateutils.AddDays(date, "7")    // ❌ Compile error
```

### Performance Optimization

- **Zero Allocations**: Most functions avoid unnecessary memory allocations
- **Efficient Algorithms**: Optimized date calculations
- **Benchmark Tested**: All functions include performance benchmarks

---

## 🧪 Testing & Quality

### Comprehensive Test Coverage

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run benchmarks
go test -bench=. ./...

# Run tests with race detection
go test -race ./...
```

### Example Test Results

```bash
=== RUN   TestAddDays
=== RUN   TestAddBusinessDays
=== RUN   TestFormat
=== RUN   TestParse
...
PASS
coverage: 100.0% of statements

BenchmarkAddDays-12           50000000     23.4 ns/op     0 B/op     0 allocs/op
BenchmarkFormat-12            2000000      654 ns/op     48 B/op     3 allocs/op
BenchmarkParse-12             1000000     1543 ns/op    184 B/op     8 allocs/op
```

---

## 🔧 Advanced Usage

### Custom Format Placeholders

**go-dateutils** supports user-friendly format placeholders:

| Placeholder | Description | Example |
|-------------|-------------|---------|
| `YYYY` | 4-digit year | 2024 |
| `YY` | 2-digit year | 24 |
| `MMMM` | Full month name | January |
| `MMM` | Short month name | Jan |
| `MM` | 2-digit month | 01 |
| `M` | Month without leading zero | 1 |
| `DD` | 2-digit day | 15 |
| `D` | Day without leading zero | 15 |
| `HH` | 24-hour format hour | 14 |
| `H` | 24-hour format hour (no leading zero) | 14 |
| `hh` | 12-hour format hour | 02 |
| `h` | 12-hour format hour (no leading zero) | 2 |
| `mm` | Minutes | 30 |
| `m` | Minutes (no leading zero) | 30 |
| `ss` | Seconds | 45 |
| `s` | Seconds (no leading zero) | 45 |
| `A` | AM/PM | PM |
| `a` | am/pm | pm |

### Error Handling Patterns

```go
// Explicit error handling
date, err := dateutils.Parse("invalid-date", time.UTC)
if err != nil {
    log.Printf("Failed to parse date: %v", err)
    return
}

// Safe functions (no errors returned)
formatted := dateutils.FormatSafe(date, dateutils.DateISO, nil)
if formatted == "" {
    log.Printf("Failed to format date")
}

// Functions with fallbacks
formatted := dateutils.FormatWithDefault(date, dateutils.DateISO, nil, "Unknown Date")
```

### WebAssembly Usage

Compile your Go application with **go-dateutils** to WebAssembly:

```bash
GOOS=js GOARCH=wasm go build -o main.wasm main.go
```

The library works seamlessly in WebAssembly environments with no modifications needed.

---

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

### Development Setup

```bash
# Clone the repository
git clone https://github.com/chmenegatti/go-date-fns.git
cd go-date-fns

# Run tests
go test ./...

# Run benchmarks
go test -bench=. ./...

# Check code coverage
go test -cover ./...
```

### Adding New Functions

1. Create the function in the appropriate file
2. Add comprehensive tests
3. Add benchmarks
4. Update documentation
5. Run all tests and ensure they pass

---

## 📋 API Reference

For complete API documentation, visit [pkg.go.dev](https://pkg.go.dev/github.com/chmenegatti/go-date-fns).

### Format Constants

```go
const (
    DateISO       = "2006-01-02"                 // ISO date format
    DateUS        = "01/02/2006"                 // US date format
    DateTime24    = "2006-01-02 15:04:05"        // 24-hour datetime
    DateTime12    = "2006-01-02 3:04:05 PM"      // 12-hour datetime
    Time24        = "15:04:05"                   // 24-hour time
    Time12        = "3:04:05 PM"                 // 12-hour time
    Readable      = "January 2, 2006"            // Human readable
)
```

---

## 🚀 Performance

**go-dateutils** is designed for high performance:

- **Minimal Allocations**: Most operations perform zero or minimal memory allocations
- **Optimized Algorithms**: Efficient date calculations using Go's time package
- **Benchmarked**: All functions include performance benchmarks

### Benchmark Results

```bash
BenchmarkAddDays-12                50000000    23.4 ns/op      0 B/op    0 allocs/op
BenchmarkAddBusinessDays-12         5000000   245.6 ns/op      0 B/op    0 allocs/op
BenchmarkFormat-12                  2000000   654.3 ns/op     48 B/op    3 allocs/op
BenchmarkParse-12                   1000000  1543.2 ns/op    184 B/op    8 allocs/op
BenchmarkDifferenceInDays-12       20000000    67.8 ns/op      0 B/op    0 allocs/op
BenchmarkIsWeekend-12             100000000    12.1 ns/op      0 B/op    0 allocs/op
```

---

## 🔄 Migration from Other Libraries

### From time Package

```go
// Before (using time package)
t := time.Now()
future := t.AddDate(0, 0, 7)  // Add 7 days
formatted := future.Format("2006-01-02")

// After (using go-dateutils)  
t := time.Now()
future := dateutils.AddDays(t, 7)
formatted, _ := dateutils.Format(future, dateutils.DateISO, nil)
```

### From Other Date Libraries

**go-dateutils** provides a familiar API for developers coming from other ecosystems:

```javascript
// JavaScript (date-fns)
import { addDays, format } from 'date-fns'
const result = format(addDays(new Date(), 7), 'yyyy-MM-dd')
```

```go
// Go (go-dateutils)
result, _ := dateutils.Format(
    dateutils.AddDays(time.Now(), 7), 
    dateutils.DateISO, 
    nil,
)
```

---

## ❓ FAQ

### Q: Why choose go-dateutils over the standard time package?

**A:** While Go's `time` package is excellent, **go-dateutils** provides:

- Higher-level, more intuitive functions
- Consistent error handling patterns  
- Business logic functions (like business days)
- User-friendly formatting with custom placeholders
- Functional programming approach

### Q: Is go-dateutils thread-safe?

**A:** Yes! All functions are pure and immutable, making them inherently thread-safe. You can safely use them in concurrent goroutines without any synchronization.

### Q: Can I use go-dateutils in production?

**A:** Absolutely! The library is production-ready with:

- 100% test coverage
- Comprehensive benchmarks
- Zero external dependencies
- Battle-tested algorithms

### Q: How does performance compare to the standard library?

**A:** **go-dateutils** is built on top of Go's `time` package, so performance is comparable. Some operations may have minimal overhead due to additional safety checks and features.

### Q: Does it support all timezones?

**A:** Yes! **go-dateutils** uses Go's standard timezone database, supporting all IANA timezone identifiers.

---

## 📞 Support

- 📚 **Documentation**: [pkg.go.dev](https://pkg.go.dev/github.com/chmenegatti/go-date-fns)
- 🐛 **Bug Reports**: [GitHub Issues](https://github.com/chmenegatti/go-date-fns/issues)
- 💬 **Discussions**: [GitHub Discussions](https://github.com/chmenegatti/go-date-fns/discussions)
- 📧 **Email**: <support@go-dateutils.com>

---

## 📄 License

MIT License - see the [LICENSE](LICENSE) file for details.

---

## 🙏 Acknowledgments

- Inspired by the excellent [date-fns](https://date-fns.org/) JavaScript library
- Built with ❤️ by the Go community
- Special thanks to all contributors and testers

---

<div align="center">

**Made with 🗓️ and ❤️ for the Go community**

[⭐ Star us on GitHub](https://github.com/chmenegatti/go-date-fns) | [🐛 Report Bug](https://github.com/chmenegatti/go-date-fns/issues) | [💡 Request Feature](https://github.com/chmenegatti/go-date-fns/issues)

</div>
