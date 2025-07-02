package main

import (
	"fmt"
	"log"
	"time"

	"github.com/chmenegatti/go-date-fns/dateutils"
)

func main() {
	fmt.Println("=== go-dateutils Examples ===")
	fmt.Println()

	// Parsing examples
	fmt.Println("1. Parsing Dates:")
	date, err := dateutils.Parse("2023-12-25", time.UTC)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Parsed date: %v\n", date)

	isoDate, err := dateutils.ParseISO("2023-12-25T15:30:00Z", time.UTC)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Parsed ISO date: %v\n", isoDate)

	// Formatting examples
	fmt.Println("\n2. Formatting Dates:")
	formatted, _ := dateutils.Format(date, dateutils.DateISO, nil)
	fmt.Printf("ISO format: %s\n", formatted)

	customFormatted, _ := dateutils.FormatCustom(isoDate, "DD/MM/YYYY HH:mm", nil)
	fmt.Printf("Custom format: %s\n", customFormatted)

	readable, _ := dateutils.Format(date, dateutils.Readable, nil)
	fmt.Printf("Readable format: %s\n", readable)

	// Date manipulation
	fmt.Println("\n3. Date Manipulation:")
	nextWeek := dateutils.AddDays(date, 7)
	fmt.Printf("Next week: %s\n", dateutils.FormatSafe(nextWeek, dateutils.DateISO, nil))

	nextBusinessDay := dateutils.AddBusinessDays(date, 5)
	fmt.Printf("5 business days later: %s\n", dateutils.FormatSafe(nextBusinessDay, dateutils.DateISO, nil))

	// Date comparison
	fmt.Println("\n4. Date Comparison:")
	today := time.Now()
	isBefore := dateutils.IsBefore(date, today)
	isAfter := dateutils.IsAfter(date, today)
	fmt.Printf("Dec 25, 2023 is before today: %t\n", isBefore)
	fmt.Printf("Dec 25, 2023 is after today: %t\n", isAfter)

	// Date differences
	fmt.Println("\n5. Date Differences:")
	daysDiff := dateutils.DifferenceInDays(today, date)
	fmt.Printf("Days since Christmas 2023: %d\n", daysDiff)

	businessDaysDiff := dateutils.DifferenceInBusinessDays(today, date)
	fmt.Printf("Business days since Christmas 2023: %d\n", businessDaysDiff)

	// Validation
	fmt.Println("\n6. Validation:")
	fmt.Printf("Is 2024 a leap year? %t\n", dateutils.IsLeapYear(2024))
	fmt.Printf("Is Christmas 2023 a weekend? %t\n", dateutils.IsWeekend(date))
	fmt.Printf("Is date valid? %t\n", dateutils.IsValid(date))

	// Period utilities
	fmt.Println("\n7. Period Utilities:")
	startOfWeek := dateutils.StartOfWeek(date)
	endOfWeek := dateutils.EndOfWeek(date)
	fmt.Printf("Start of week: %s\n", dateutils.FormatSafe(startOfWeek, dateutils.DateTime24, nil))
	fmt.Printf("End of week: %s\n", dateutils.FormatSafe(endOfWeek, dateutils.DateTime24, nil))

	startOfMonth := dateutils.StartOfMonth(date)
	endOfMonth := dateutils.EndOfMonth(date)
	fmt.Printf("Start of month: %s\n", dateutils.FormatSafe(startOfMonth, dateutils.DateTime24, nil))
	fmt.Printf("End of month: %s\n", dateutils.FormatSafe(endOfMonth, dateutils.DateTime24, nil))

	// Timezone examples
	fmt.Println("\n8. Timezone Examples:")
	est, _ := time.LoadLocation("America/New_York")
	utcTime := time.Date(2023, 12, 25, 15, 30, 0, 0, time.UTC)

	estFormatted, _ := dateutils.Format(utcTime, dateutils.DateTime24, est)
	fmt.Printf("UTC time in EST: %s\n", estFormatted)

	// Same date comparison across timezones
	sameDate := dateutils.IsSameDate(utcTime, utcTime.In(est))
	fmt.Printf("Same date across timezones: %t\n", sameDate)

	// Time difference calculations
	fmt.Println("\n9. Time Difference Calculations:")
	projectStart := time.Date(2023, 1, 1, 9, 0, 0, 0, time.UTC)
	projectEnd := time.Date(2023, 6, 15, 17, 30, 0, 0, time.UTC)

	fmt.Printf("Project duration in months: %d\n", dateutils.DifferenceInMonths(projectEnd, projectStart))
	fmt.Printf("Project duration in days: %d\n", dateutils.DifferenceInDays(projectEnd, projectStart))
	fmt.Printf("Project duration in hours: %d\n", dateutils.DifferenceInHours(projectEnd, projectStart))
	fmt.Printf("Business days only: %d\n", dateutils.DifferenceInBusinessDays(projectEnd, projectStart))
	fmt.Printf("Precise duration: %.2f days\n", dateutils.DifferenceInDaysFloat(projectEnd, projectStart))

	// Component extraction and manipulation
	fmt.Println("\n10. Get/Set Component Functions:")
	sampleDate := time.Date(2024, 8, 15, 14, 30, 45, 500000000, time.UTC)

	fmt.Printf("Original date: %s\n", dateutils.FormatSafe(sampleDate, dateutils.DateTime24, nil))
	fmt.Printf("Year: %d, Quarter: Q%d, Month: %d\n",
		dateutils.GetYear(sampleDate), dateutils.GetQuarter(sampleDate), dateutils.GetMonth(sampleDate))
	fmt.Printf("Day: %d, Day of Year: %d, Week: %d\n",
		dateutils.GetDate(sampleDate), dateutils.GetDayOfYear(sampleDate), dateutils.GetWeek(sampleDate))

	// Create variations using Set functions
	newYear := dateutils.SetYear(sampleDate, 2025)
	startOfQuarter := dateutils.SetQuarter(sampleDate, 1)
	fmt.Printf("Set to 2025: %s\n", dateutils.FormatSafe(newYear, dateutils.DateISO, nil))
	fmt.Printf("Set to Q1: %s\n", dateutils.FormatSafe(startOfQuarter, dateutils.DateISO, nil))

	// NEW: Interval utilities - Each functions for comprehensive date ranges
	fmt.Println("\n11. Interval Utilities (Each Functions):")

	// Create an interval for Q1 2024
	q1Interval := dateutils.Interval{
		Start: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		End:   time.Date(2024, 3, 31, 23, 59, 59, 0, time.UTC),
	}

	// Get all months in Q1
	months := dateutils.EachMonthOfInterval(q1Interval)
	fmt.Printf("Months in Q1 2024:\n")
	for _, month := range months {
		fmt.Printf("  %s\n", dateutils.FormatSafe(month, "January 2006", nil))
	}

	// Get all weekends in January 2024
	januaryInterval := dateutils.Interval{
		Start: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		End:   time.Date(2024, 1, 31, 23, 59, 59, 0, time.UTC),
	}

	weekends := dateutils.EachWeekendOfInterval(januaryInterval)
	fmt.Printf("\nWeekends in January 2024:\n")
	for _, weekend := range weekends {
		dayName := weekend.Weekday().String()
		fmt.Printf("  %s %s\n", dayName, dateutils.FormatSafe(weekend, dateutils.DateISO, nil))
	}

	// Get business days for a work week
	workWeekInterval := dateutils.Interval{
		Start: time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC),     // Monday
		End:   time.Date(2024, 1, 12, 23, 59, 59, 0, time.UTC), // Friday
	}

	businessDays := dateutils.EachBusinessDayOfInterval(workWeekInterval)
	fmt.Printf("\nBusiness days in work week:\n")
	for _, day := range businessDays {
		dayName := day.Weekday().String()
		fmt.Printf("  %s %s\n", dayName, dateutils.FormatSafe(day, dateutils.DateISO, nil))
	}

	// Get quarters for 2024
	yearInterval := dateutils.Interval{
		Start: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		End:   time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC),
	}

	quarters := dateutils.EachQuarterOfInterval(yearInterval)
	fmt.Printf("\nQuarters in 2024:\n")
	for i, quarter := range quarters {
		fmt.Printf("  Q%d: %s\n", i+1, dateutils.FormatSafe(quarter, "January 2006", nil))
	}

	// Hours in a workday
	workdayInterval := dateutils.Interval{
		Start: time.Date(2024, 1, 15, 9, 0, 0, 0, time.UTC),
		End:   time.Date(2024, 1, 15, 17, 0, 0, 0, time.UTC),
	}

	workHours := dateutils.EachHourOfInterval(workdayInterval)
	fmt.Printf("\nWork hours on Jan 15, 2024:\n")
	for _, hour := range workHours {
		fmt.Printf("  %s\n", dateutils.FormatSafe(hour, "15:04", nil))
	}

	fmt.Println("\n=== Examples Complete ===")
	fmt.Printf("Library now includes %d+ date utility functions!\n", 140)
}
