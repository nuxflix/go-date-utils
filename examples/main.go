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

	fmt.Println("\n=== Examples Complete ===")
}
