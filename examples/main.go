// Package main demonstrates practical usage of go-date-fns.
// Run with: go run ./examples/main.go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/chmenegatti/go-date-fns/dateutils"
)

func main() {
	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Println("║       go-date-fns  Examples          ║")
	fmt.Println("╚══════════════════════════════════════╝")
	fmt.Println()

	exampleParsing()
	exampleFormatting()
	exampleManipulation()
	exampleComparisons()
	exampleDifferences()
	exampleValidation()
	examplePeriodUtils()
	exampleTimezones()
	exampleIntervals()
	exampleFormatDistance()
	exampleBusinessScenarios()

	fmt.Println("\n╔══════════════════════════════════════╗")
	fmt.Printf("║  Library: %d+ date utility functions  ║\n", 140)
	fmt.Println("╚══════════════════════════════════════╝")
}

// ---------------------------------------------------------
// 1. Parsing
// ---------------------------------------------------------
func exampleParsing() {
	fmt.Println("── 1. Parsing ────────────────────────────")

	// Multi-format auto-detection
	formats := []string{
		"2024-01-15",
		"01/15/2024",
		"January 15, 2024",
	}
	for _, s := range formats {
		d, err := dateutils.Parse(s, time.UTC)
		if err != nil {
			log.Printf("  parse %q: %v", s, err)
			continue
		}
		fmt.Printf("  Parse(%q) → %s\n", s, dateutils.FormatSafe(d, dateutils.DateISO, nil))
	}

	// ISO 8601 with timezone
	iso, _ := dateutils.ParseISO("2024-07-04T19:30:00-05:00", time.UTC)
	fmt.Printf("  ParseISO(\"2024-07-04T19:30:00-05:00\") → %s (UTC)\n",
		dateutils.FormatSafe(iso, dateutils.DateTimeISO, nil))

	// Custom format
	custom, _ := dateutils.ParseWithFormat("25-12-2024", "02-01-2006", time.UTC)
	fmt.Printf("  ParseWithFormat(\"25-12-2024\", \"02-01-2006\") → %s\n",
		dateutils.FormatSafe(custom, dateutils.DateISO, nil))

	fmt.Println()
}

// ---------------------------------------------------------
// 2. Formatting
// ---------------------------------------------------------
func exampleFormatting() {
	fmt.Println("── 2. Formatting ─────────────────────────")

	d := time.Date(2024, 7, 4, 19, 30, 0, 0, time.UTC)

	examples := []struct{ label, format string }{
		{"DateISO", dateutils.DateISO},
		{"DateUS", dateutils.DateUS},
		{"DateTime24", dateutils.DateTime24},
		{"DateTime12", dateutils.DateTime12},
		{"Readable", dateutils.Readable},
		{"WeekdayLong", dateutils.WeekdayLong},
	}
	for _, e := range examples {
		out, _ := dateutils.Format(d, e.format, nil)
		fmt.Printf("  %-14s → %s\n", e.label, out)
	}

	// User-friendly placeholders
	c1, _ := dateutils.FormatCustom(d, "DD/MM/YYYY HH:mm", nil)
	c2, _ := dateutils.FormatCustom(d, "MMMM DD, YYYY", nil)
	fmt.Printf("  FormatCustom YYYY/MM/DD → %s\n", c1)
	fmt.Printf("  FormatCustom MMMM DD    → %s\n", c2)

	// Safe / with default
	zero := time.Time{}
	fmt.Printf("  FormatSafe(zero)        → %q\n", dateutils.FormatSafe(zero, dateutils.DateISO, nil))
	fmt.Printf("  FormatWithDefault(zero) → %q\n", dateutils.FormatWithDefault(zero, dateutils.DateISO, nil, "N/A"))
	fmt.Println()
}

// ---------------------------------------------------------
// 3. Manipulation
// ---------------------------------------------------------
func exampleManipulation() {
	fmt.Println("── 3. Manipulation ───────────────────────")

	base := time.Date(2024, 1, 15, 9, 0, 0, 0, time.UTC) // Monday
	fmt.Printf("  Base date: %s\n", dateutils.FormatSafe(base, dateutils.WeekdayLong, nil))
	fmt.Printf("  +7 days:   %s\n", dateutils.FormatSafe(dateutils.AddDays(base, 7), dateutils.WeekdayLong, nil))
	fmt.Printf("  +1 month:  %s\n", dateutils.FormatSafe(dateutils.AddMonths(base, 1), dateutils.DateISO, nil))
	fmt.Printf("  +5 biz:    %s\n", dateutils.FormatSafe(dateutils.AddBusinessDays(base, 5), dateutils.WeekdayLong, nil))
	fmt.Printf("  -3 biz:    %s\n", dateutils.FormatSafe(dateutils.SubBusinessDays(base, 3), dateutils.WeekdayLong, nil))
	fmt.Printf("  +8 hours:  %s\n", dateutils.FormatSafe(dateutils.AddHours(base, 8), dateutils.DateTime24, nil))

	// Set functions
	modified := dateutils.SetHours(dateutils.SetDate(base, 1), 14) // First of month, 2 PM
	fmt.Printf("  SetDate(1)+SetHours(14): %s\n", dateutils.FormatSafe(modified, dateutils.DateTime24, nil))
	fmt.Println()
}

// ---------------------------------------------------------
// 4. Comparisons
// ---------------------------------------------------------
func exampleComparisons() {
	fmt.Println("── 4. Comparisons ────────────────────────")

	d1 := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	d2 := time.Date(2024, 1, 22, 15, 0, 0, 0, time.UTC)

	fmt.Printf("  d1=%s  d2=%s\n",
		dateutils.FormatSafe(d1, dateutils.DateISO, nil),
		dateutils.FormatSafe(d2, dateutils.DateISO, nil))
	fmt.Printf("  IsBefore(d1,d2):   %t\n", dateutils.IsBefore(d1, d2))
	fmt.Printf("  IsSameMonth:       %t\n", dateutils.IsSameMonth(d1, d2))
	fmt.Printf("  IsSameWeek:        %t\n", dateutils.IsSameWeek(d1, d2))
	fmt.Printf("  CompareAsc(d1,d2): %d\n", dateutils.CompareAsc(d1, d2))

	// Min / Max / ClosestTo
	times := []time.Time{d2, d1, time.Date(2024, 1, 19, 0, 0, 0, 0, time.UTC)}
	fmt.Printf("  Min(times): %s\n", dateutils.FormatSafe(dateutils.Min(times), dateutils.DateISO, nil))
	fmt.Printf("  Max(times): %s\n", dateutils.FormatSafe(dateutils.Max(times), dateutils.DateISO, nil))
	fmt.Println()
}

// ---------------------------------------------------------
// 5. Differences
// ---------------------------------------------------------
func exampleDifferences() {
	fmt.Println("── 5. Differences ────────────────────────")

	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 6, 15, 12, 30, 0, 0, time.UTC)

	fmt.Printf("  %s → %s\n",
		dateutils.FormatSafe(start, dateutils.DateISO, nil),
		dateutils.FormatSafe(end, dateutils.DateISO, nil))
	fmt.Printf("  DifferenceInMonths:       %d\n", dateutils.DifferenceInMonths(end, start))
	fmt.Printf("  DifferenceInDays:         %d\n", dateutils.DifferenceInDays(end, start))
	fmt.Printf("  DifferenceInDaysFloat:    %.2f\n", dateutils.DifferenceInDaysFloat(end, start))
	fmt.Printf("  DifferenceInBusinessDays: %d\n", dateutils.DifferenceInBusinessDays(end, start))
	fmt.Printf("  DifferenceInHours:        %d\n", dateutils.DifferenceInHours(end, start))
	fmt.Println()
}

// ---------------------------------------------------------
// 6. Validation
// ---------------------------------------------------------
func exampleValidation() {
	fmt.Println("── 6. Validation ─────────────────────────")

	d := time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC) // Leap day
	fmt.Printf("  IsValid(2024-02-29): %t\n", dateutils.IsValid(d))
	fmt.Printf("  IsValid(zero):       %t\n", dateutils.IsValid(time.Time{}))
	fmt.Printf("  IsLeapYear(2024):    %t\n", dateutils.IsLeapYear(2024))
	fmt.Printf("  IsLeapYear(2023):    %t\n", dateutils.IsLeapYear(2023))
	fmt.Printf("  IsWeekend(2024-02-29 Thu): %t\n", dateutils.IsWeekend(d))

	monday := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)
	fmt.Printf("  IsMonday(2024-01-15): %t\n", dateutils.IsMonday(monday))

	jan1 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Printf("  IsFirstDayOfYear(2024-01-01): %t\n", dateutils.IsFirstDayOfYear(jan1))
	fmt.Printf("  IsFirstDayOfMonth(2024-01-01): %t\n", dateutils.IsFirstDayOfMonth(jan1))
	fmt.Println()
}

// ---------------------------------------------------------
// 7. Period Utilities
// ---------------------------------------------------------
func examplePeriodUtils() {
	fmt.Println("── 7. Period Utilities ───────────────────")

	d := time.Date(2024, 6, 15, 14, 30, 45, 0, time.UTC) // Saturday mid-day
	fmt.Printf("  Date: %s\n", dateutils.FormatSafe(d, dateutils.WeekdayLong, nil))
	fmt.Printf("  StartOfDay:    %s\n", dateutils.FormatSafe(dateutils.StartOfDay(d), dateutils.DateTime24, nil))
	fmt.Printf("  EndOfDay:      %s\n", dateutils.FormatSafe(dateutils.EndOfDay(d), dateutils.DateTime24, nil))
	fmt.Printf("  StartOfWeek:   %s\n", dateutils.FormatSafe(dateutils.StartOfWeek(d), dateutils.DateISO, nil))
	fmt.Printf("  EndOfWeek:     %s\n", dateutils.FormatSafe(dateutils.EndOfWeek(d), dateutils.DateISO, nil))
	fmt.Printf("  StartOfMonth:  %s\n", dateutils.FormatSafe(dateutils.StartOfMonth(d), dateutils.DateISO, nil))
	fmt.Printf("  EndOfMonth:    %s\n", dateutils.FormatSafe(dateutils.EndOfMonth(d), dateutils.DateISO, nil))
	fmt.Printf("  StartOfYear:   %s\n", dateutils.FormatSafe(dateutils.StartOfYear(d), dateutils.DateISO, nil))
	fmt.Printf("  StartOfQuarter:%s\n", dateutils.FormatSafe(dateutils.StartOfQuarter(d), dateutils.DateISO, nil))
	fmt.Println()
}

// ---------------------------------------------------------
// 8. Timezones
// ---------------------------------------------------------
func exampleTimezones() {
	fmt.Println("── 8. Timezones ──────────────────────────")

	utcTime := time.Date(2024, 7, 4, 20, 0, 0, 0, time.UTC)
	zones := []string{"America/New_York", "Europe/London", "Asia/Tokyo", "America/Sao_Paulo"}

	for _, zone := range zones {
		loc, err := time.LoadLocation(zone)
		if err != nil {
			continue
		}
		formatted, _ := dateutils.Format(utcTime, dateutils.DateTime24, loc)
		fmt.Printf("  %-25s → %s\n", zone, formatted)
	}
	fmt.Println()
}

// ---------------------------------------------------------
// 9. Interval Utilities
// ---------------------------------------------------------
func exampleIntervals() {
	fmt.Println("── 9. Interval Utilities ─────────────────")

	jan2024 := dateutils.Interval{
		Start: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		End:   time.Date(2024, 1, 31, 23, 59, 59, 0, time.UTC),
	}

	weekends := dateutils.EachWeekendOfInterval(jan2024)
	fmt.Printf("  Weekends in Jan 2024 (%d total):\n", len(weekends))
	for _, d := range weekends {
		fmt.Printf("    %s %s\n", d.Weekday(), dateutils.FormatSafe(d, dateutils.DateISO, nil))
	}

	q1 := dateutils.Interval{
		Start: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		End:   time.Date(2024, 3, 31, 23, 59, 59, 0, time.UTC),
	}
	months := dateutils.EachMonthOfInterval(q1)
	fmt.Printf("  Months in Q1 2024: ")
	for _, m := range months {
		fmt.Printf("%s ", dateutils.FormatSafe(m, "Jan", nil))
	}
	fmt.Println()
	fmt.Println()
}

// ---------------------------------------------------------
// 10. FormatDistance (relative time)
// ---------------------------------------------------------
func exampleFormatDistance() {
	fmt.Println("── 10. FormatDistance (Relative Time) ────")

	now := time.Now()
	cases := []struct {
		label string
		d     time.Time
	}{
		{"30 seconds ago", now.Add(-30 * time.Second)},
		{"5 minutes ago", now.Add(-5 * time.Minute)},
		{"2 hours ago", now.Add(-2 * time.Hour)},
		{"3 days ago", now.Add(-3 * 24 * time.Hour)},
		{"in 2 weeks", now.Add(14 * 24 * time.Hour)},
	}

	opts := &dateutils.FormatDistanceOptions{AddSuffix: true, IncludeSeconds: true}
	for _, c := range cases {
		result := dateutils.FormatDistance(c.d, now, opts)
		fmt.Printf("  %-18s → %q\n", c.label, result)
	}
	fmt.Println()
}

// ---------------------------------------------------------
// 11. Real-World Business Scenarios
// ---------------------------------------------------------
func exampleBusinessScenarios() {
	fmt.Println("── 11. Business Scenarios ────────────────")

	// a) JWT token expiry check
	fmt.Println("  a) JWT-style token expiry:")
	issuedAt := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	tokenTTL := 24 // hours
	expiresAt := dateutils.AddHours(issuedAt, tokenTTL)
	checkAt := time.Date(2024, 1, 15, 11, 30, 0, 0, time.UTC)
	isValid := dateutils.IsBefore(checkAt, expiresAt)
	remaining := dateutils.DifferenceInMinutes(expiresAt, checkAt)
	fmt.Printf("    Issued: %s | Expires: %s\n",
		dateutils.FormatSafe(issuedAt, dateutils.DateTime24, nil),
		dateutils.FormatSafe(expiresAt, dateutils.DateTime24, nil))
	fmt.Printf("    Valid at %s: %t (%d min remaining)\n",
		dateutils.FormatSafe(checkAt, "15:04", nil), isValid, remaining)

	// b) Monthly report stats
	fmt.Println("\n  b) Monthly report (January 2024):")
	jan := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	startM := dateutils.StartOfMonth(jan)
	endM := dateutils.EndOfMonth(jan)
	totalDays := dateutils.DifferenceInCalendarDays(endM, startM) + 1
	bizDays := dateutils.DifferenceInBusinessDays(endM, startM) + 1
	fmt.Printf("    Total days:   %d\n", totalDays)
	fmt.Printf("    Working days: %d\n", bizDays)
	fmt.Printf("    Weekend days: %d\n", totalDays-bizDays)

	// c) Project deadline in business days
	fmt.Println("\n  c) Project deadline (20 business days from today):")
	today := time.Now()
	deadline := dateutils.AddBusinessDays(today, 20)
	fmt.Printf("    Today:    %s\n", dateutils.FormatSafe(today, dateutils.DateISO, nil))
	fmt.Printf("    Deadline: %s (%s)\n",
		dateutils.FormatSafe(deadline, dateutils.DateISO, nil),
		dateutils.FormatSafe(deadline, "Monday", nil))

	// d) Meeting series (weekly for 4 weeks, skip weekends)
	fmt.Println("\n  d) Weekly meeting series (4 weeks, skip weekends):")
	meetingBase := time.Date(2024, 1, 13, 14, 0, 0, 0, time.UTC) // Saturday
	for i := 0; i < 4; i++ {
		m := dateutils.AddWeeks(meetingBase, i)
		if dateutils.IsWeekend(m) {
			m = dateutils.AddDays(m, 2) // skip to Monday
		}
		fmt.Printf("    Meeting %d: %s\n", i+1, dateutils.FormatSafe(m, dateutils.WeekdayLong, nil))
	}
	fmt.Println()
}
