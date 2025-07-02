# 🗓️ go-dateutils

> A modern, functional date utility library for Go, inspired by the beloved date-fns JavaScript library.

[![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.19-007d9c.svg)](https://golang.org/)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/chmenegatti/go-date-fns)](https://goreportcard.com/report/github.com/chmenegatti/go-date-fns)
[![Test Coverage](https://img.shields.io/badge/coverage-100%25-brightgreen.svg)](https://github.com/chmenegatti/go-date-fns)

**go-dateutils** brings the power and elegance of functional date manipulation to Go. With over **150+ pure, immutable functions**, it provides a comprehensive toolkit for working with dates and times in a safe, predictable, and timezone-aware manner.

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
dateutils.IsEqual(date1, date1)         // true
dateutils.IsBeforeOrEqual(date1, date1) // true

// Date-only comparisons (ignores time)
dateutils.IsBeforeDate(date1, date2)    // true
dateutils.IsAfterDate(date1, date2)     // false

// Same period comparisons
dateutils.IsSameDay(date1, date2)       // false (different days)
dateutils.IsSameWeek(date1, date2)      // true (same week)
dateutils.IsSameMonth(date1, date2)     // true (same month)
dateutils.IsSameYear(date1, date2)      // true (same year)
dateutils.IsSameHour(date1, date2)      // false (different hours)
dateutils.IsSameMinute(date1, date2)    // false (different minutes)

// Sorting and selection
times := []time.Time{date1, date2, time.Now()}
earliest := dateutils.Min(times)        // Get earliest time
latest := dateutils.Max(times)          // Get latest time
closest := dateutils.ClosestTo(date1, times) // Get closest to date1

// Comparison with ordering
result := dateutils.CompareAsc(date1, date2)  // -1 (date1 before date2)
result = dateutils.CompareDesc(date1, date2)  // 1 (desc order)

// Timezone-aware comparisons
est, _ := time.LoadLocation("America/New_York")
dateutils.IsBeforeInTimezone(date1, date2, est)
```

### 🔧 **Manipulation Functions**

Add, subtract, and modify dates immutably:

```go
baseDate := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC) // Monday

// Add/subtract days, weeks, months, years
nextWeek := dateutils.AddDays(baseDate, 7)         // 2024-01-22
lastWeek := dateutils.SubDays(baseDate, 7)         // 2024-01-08
nextMonth := dateutils.AddMonths(baseDate, 1)      // 2024-02-15
lastYear := dateutils.SubYears(baseDate, 1)        // 2023-01-15

// Add/subtract time units
laterToday := dateutils.AddHours(baseDate, 5)      // Add 5 hours
earlierToday := dateutils.SubMinutes(baseDate, 30) // Subtract 30 minutes
futureTime := dateutils.AddSeconds(baseDate, 120)  // Add 2 minutes

// Business days (skips weekends)
nextBusinessDay := dateutils.AddBusinessDays(baseDate, 5)  // Skips weekend, lands on weekday
pastBusinessDay := dateutils.SubBusinessDays(baseDate, 3)  // Goes back 3 business days

// Timezone-aware addition
est, _ := time.LoadLocation("America/New_York")
dateInEST := dateutils.AddDaysWithTimezone(baseDate, 1, est)
```

### 📊 **Get Functions**

Extract specific components from dates:

```go
date := time.Date(2024, 6, 15, 14, 30, 45, 500000000, time.UTC) // Saturday

// Date components
year := dateutils.GetYear(date)           // 2024
month := dateutils.GetMonth(date)         // 6 (June)
day := dateutils.GetDate(date)            // 15
dayOfWeek := dateutils.GetDay(date)       // 6 (Saturday, 0=Sunday)
dayOfYear := dateutils.GetDayOfYear(date) // 167
quarter := dateutils.GetQuarter(date)     // 2 (Q2)

// Time components
hours := dateutils.GetHours(date)         // 14
minutes := dateutils.GetMinutes(date)     // 30
seconds := dateutils.GetSeconds(date)     // 45
milliseconds := dateutils.GetMilliseconds(date) // 500

// Week-related
week := dateutils.GetWeek(date)           // ISO week number
weekOfMonth := dateutils.GetWeekOfMonth(date) // Week of the month
weekYear := dateutils.GetWeekYear(date)   // ISO week year
```

### ⚙️ **Set Functions**

Create new dates by setting specific components:

```go
baseDate := time.Date(2024, 6, 15, 14, 30, 45, 0, time.UTC)

// Set date components
newYear := dateutils.SetYear(baseDate, 2025)      // 2025-06-15 14:30:45
newMonth := dateutils.SetMonth(baseDate, 12)      // 2024-12-15 14:30:45
newDay := dateutils.SetDate(baseDate, 1)          // 2024-06-01 14:30:45
newQuarter := dateutils.SetQuarter(baseDate, 1)   // 2024-01-15 14:30:45

// Set time components
newHours := dateutils.SetHours(baseDate, 9)       // 2024-06-15 09:30:45
newMinutes := dateutils.SetMinutes(baseDate, 0)   // 2024-06-15 14:00:45
newSeconds := dateutils.SetSeconds(baseDate, 0)   // 2024-06-15 14:30:00

// Set by day of week or day of year
monday := dateutils.SetDay(baseDate, 1)           // Set to Monday of the same week
newYearDay := dateutils.SetDayOfYear(baseDate, 100) // Set to 100th day of year
```

### 📏 **Difference Functions**

Calculate time differences with various precision levels:

```go
start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
end := time.Date(2024, 1, 15, 12, 30, 45, 0, time.UTC)

// Time unit differences (integer)
years := dateutils.DifferenceInYears(end, start)                  // 0
months := dateutils.DifferenceInMonths(end, start)                // 0
weeks := dateutils.DifferenceInWeeks(end, start)                  // 2
days := dateutils.DifferenceInDays(end, start)                    // 14
hours := dateutils.DifferenceInHours(end, start)                  // 348
minutes := dateutils.DifferenceInMinutes(end, start)              // 20910
seconds := dateutils.DifferenceInSeconds(end, start)              // 1254645

// Calendar-based differences (ignore time components)
calendarYears := dateutils.DifferenceInCalendarYears(end, start)  // 0
calendarMonths := dateutils.DifferenceInCalendarMonths(end, start)// 0
calendarDays := dateutils.DifferenceInCalendarDays(end, start)    // 14
quarters := dateutils.DifferenceInQuarters(end, start)            // 0
calendarQuarters := dateutils.DifferenceInCalendarQuarters(end, start) // 0

// Business and specialized differences
businessDays := dateutils.DifferenceInBusinessDays(end, start)    // 10 (excluding weekends)

// Floating-point precision
daysPrecise := dateutils.DifferenceInDaysFloat(end, start)        // 14.52...
weeksPrecise := dateutils.DifferenceInWeeksFloat(end, start)      // 2.07...
hoursPrecise := dateutils.DifferenceInHoursFloat(end, start)      // 348.51...

// Absolute differences (always positive)
absDays := dateutils.AbsDifferenceInDays(start, end)              // 14
absHours := dateutils.AbsDifferenceInHours(start, end)            // 348
absMinutes := dateutils.AbsDifferenceInMinutes(start, end)        // 20910
absMonths := dateutils.AbsDifferenceInMonths(start, end)          // 0
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

### 🔄 **Interval Utilities**

**NEW:** Generate arrays of dates within intervals - perfect for reporting, scheduling, and data analysis:

```go
// Define an interval for Q1 2024
q1Interval := dateutils.Interval{
    Start: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
    End:   time.Date(2024, 3, 31, 23, 59, 59, 0, time.UTC),
}

// Get all months in Q1
months := dateutils.EachMonthOfInterval(q1Interval)
// Returns: [2024-01-01, 2024-02-01, 2024-03-01]

// Get all weekends in January 2024
january := dateutils.Interval{
    Start: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
    End:   time.Date(2024, 1, 31, 23, 59, 59, 0, time.UTC),
}
weekends := dateutils.EachWeekendOfInterval(january)
// Returns all Saturday and Sunday dates in January

// Get business days for a work week
workWeek := dateutils.Interval{
    Start: time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC),  // Monday
    End:   time.Date(2024, 1, 12, 23, 59, 59, 0, time.UTC), // Friday
}
businessDays := dateutils.EachBusinessDayOfInterval(workWeek)
// Returns: [Mon, Tue, Wed, Thu, Fri] (5 business days)

// Get all quarters in 2024
year := dateutils.Interval{
    Start: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
    End:   time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC),
}
quarters := dateutils.EachQuarterOfInterval(year)
// Returns: [Q1, Q2, Q3, Q4] start dates

// Get hourly intervals for a workday
workday := dateutils.Interval{
    Start: time.Date(2024, 1, 15, 9, 0, 0, 0, time.UTC),
    End:   time.Date(2024, 1, 15, 17, 0, 0, 0, time.UTC),
}
hours := dateutils.EachHourOfInterval(workday)
// Returns: [09:00, 10:00, 11:00, ..., 17:00]
```

**Available Each Functions:**
- `EachDayOfInterval` - All days in interval
- `EachWeekOfInterval` - Weekly intervals (Monday start)
- `EachWeekOfIntervalSunday` - Weekly intervals (Sunday start)
- `EachMonthOfInterval` - Monthly intervals
- `EachYearOfInterval` - Yearly intervals  
- `EachQuarterOfInterval` - Quarterly intervals
- `EachHourOfInterval` - Hourly intervals
- `EachMinuteOfInterval` - Minute intervals
- `EachWeekendOfInterval` - Weekend days only
- `EachBusinessDayOfInterval` - Business days only

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

// Set specific meeting times using Set functions
meetingBase := time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC)
quarterlyMeeting := dateutils.SetQuarter(meetingBase, 3)      // Q3 meeting
monthlyMeeting := dateutils.SetDate(quarterlyMeeting, 1)      // First of the month
finalTime := dateutils.SetHours(monthlyMeeting, 14)          // 2 PM meeting

// Get components for reporting
quarter := dateutils.GetQuarter(finalTime)                   // Which quarter?
week := dateutils.GetWeek(finalTime)                         // Which week of year?
dayOfYear := dateutils.GetDayOfYear(finalTime)              // Day number in year

// Calculate time spans using new time unit functions
projectStart := dateutils.SubMonths(today, 3)                // 3 months ago
deadline2 := dateutils.AddHours(projectStart, 8*24*30)       // Add work hours
timeLeft := dateutils.SubDays(deadline2, 5)                  // 5 days before deadline
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

### 🔧 Date Component Manipulation

```go
// Extract and manipulate date components
baseDate := time.Date(2024, 8, 15, 16, 45, 30, 0, time.UTC)

// Extract all components
fmt.Printf("Year: %d, Quarter: Q%d\n", dateutils.GetYear(baseDate), dateutils.GetQuarter(baseDate))
fmt.Printf("Month: %d, Day: %d, Day of Year: %d\n", 
    dateutils.GetMonth(baseDate), dateutils.GetDate(baseDate), dateutils.GetDayOfYear(baseDate))
fmt.Printf("Week: %d, Day of Week: %d\n", dateutils.GetWeek(baseDate), dateutils.GetDay(baseDate))
fmt.Printf("Time: %02d:%02d:%02d.%03d\n", 
    dateutils.GetHours(baseDate), dateutils.GetMinutes(baseDate), 
    dateutils.GetSeconds(baseDate), dateutils.GetMilliseconds(baseDate))

// Create variations by setting components
yearEnd := dateutils.SetMonth(dateutils.SetDate(baseDate, 31), 12)  // Dec 31st
startOfQuarter := dateutils.SetQuarter(baseDate, 1)                 // Q1 start
mondayMeeting := dateutils.SetDay(baseDate, 1)                      // Next Monday
nineAM := dateutils.SetHours(dateutils.SetMinutes(baseDate, 0), 9)  // 9:00 AM

// Time comparisons using new functions
date1 := time.Date(2024, 6, 15, 14, 30, 0, 0, time.UTC)
date2 := time.Date(2024, 6, 22, 10, 15, 0, 0, time.UTC)

if dateutils.IsSameWeek(date1, date2) {
    fmt.Println("Same week!")
} else if dateutils.IsSameMonth(date1, date2) {
    fmt.Println("Same month, different week")
}

// Find closest date from a set
candidates := []time.Time{
    dateutils.SubDays(baseDate, 10),
    dateutils.AddDays(baseDate, 5),
    dateutils.AddMonths(baseDate, 1),
}
closest := dateutils.ClosestTo(baseDate, candidates)
fmt.Printf("Closest date: %s\n", dateutils.FormatSafe(closest, dateutils.DateISO, nil))
```

### ⏱️ Comprehensive Time Difference Calculations

```go
// Project timeline analysis
projectStart := time.Date(2024, 1, 1, 9, 0, 0, 0, time.UTC)
projectEnd := time.Date(2024, 6, 15, 17, 30, 45, 0, time.UTC)

// All time unit differences
fmt.Println("=== Project Duration Analysis ===")
fmt.Printf("Duration in years: %d\n", dateutils.DifferenceInYears(projectEnd, projectStart))
fmt.Printf("Duration in months: %d\n", dateutils.DifferenceInMonths(projectEnd, projectStart))
fmt.Printf("Duration in weeks: %d\n", dateutils.DifferenceInWeeks(projectEnd, projectStart))
fmt.Printf("Duration in days: %d\n", dateutils.DifferenceInDays(projectEnd, projectStart))
fmt.Printf("Duration in hours: %d\n", dateutils.DifferenceInHours(projectEnd, projectStart))
fmt.Printf("Duration in minutes: %d\n", dateutils.DifferenceInMinutes(projectEnd, projectStart))
fmt.Printf("Duration in seconds: %d\n", dateutils.DifferenceInSeconds(projectEnd, projectStart))

// High-precision calculations
fmt.Printf("Precise duration: %.2f days\n", dateutils.DifferenceInDaysFloat(projectEnd, projectStart))
fmt.Printf("Working days only: %d\n", dateutils.DifferenceInBusinessDays(projectEnd, projectStart))

// Time remaining calculations
deadline := dateutils.AddMonths(projectStart, 8) // 8-month deadline
now := time.Now()

if dateutils.IsBefore(now, deadline) {
    remaining := dateutils.AbsDifferenceInDays(deadline, now)
    fmt.Printf("Days until deadline: %d\n", remaining)
} else {
    overdue := dateutils.AbsDifferenceInDays(now, deadline)
    fmt.Printf("Days overdue: %d\n", overdue)
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

## Made with 🗓️ and ❤️ for the Go community

[⭐ Star us on GitHub](https://github.com/chmenegatti/go-date-fns) | [🐛 Report Bug](https://github.com/chmenegatti/go-date-fns/issues) | [💡 Request Feature](https://github.com/chmenegatti/go-date-fns/issues)

</div>
