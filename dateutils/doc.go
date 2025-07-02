// Package dateutils provides functional date utility functions for Go, inspired by date-fns.
// It offers pure and immutable functions for date manipulation, parsing, formatting,
// comparison, and validation with timezone support and zero external dependencies.
//
// Key Features:
//   - Pure and immutable functions
//   - Timezone-aware operations
//   - Zero external dependencies
//   - WebAssembly compatible
//   - Comprehensive test coverage
//
// The library is organized into several categories of functions:
//
// Parsing: Parse, ParseISO, ParseWithFormat
// Formatting: Format, FormatCustom, FormatSafe
// Comparison: IsBefore, IsAfter, IsEqual
// Manipulation: AddDays, AddHours, AddMonths, etc.
// Differences: DifferenceInDays, DifferenceInHours, etc.
// Validation: IsValid, IsLeapYear, IsWeekend, etc.
// Period Utils: StartOfDay, EndOfDay, StartOfWeek, etc.
//
// Example usage:
//
//	import "github.com/cesar/go-dateutils/dateutils"
//
//	// Parse and format dates
//	date, err := dateutils.Parse("2023-12-25", time.UTC)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	formatted, _ := dateutils.Format(date, dateutils.DateISO, nil)
//	fmt.Println(formatted) // "2023-12-25"
//
//	// Date manipulation
//	nextWeek := dateutils.AddDays(date, 7)
//	daysDiff := dateutils.DifferenceInDays(nextWeek, date)
//	fmt.Println(daysDiff) // 7
//
//	// Validation and comparison
//	isWeekend := dateutils.IsWeekend(date)
//	isBefore := dateutils.IsBefore(date, nextWeek)
//
//	// Period utilities
//	startOfWeek := dateutils.StartOfWeek(date)
//	endOfMonth := dateutils.EndOfMonth(date)
package dateutils
