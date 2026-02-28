// Package dateutils provides functional date utility functions for Go, inspired by date-fns.
// It offers pure and immutable functions for date manipulation, parsing, formatting,
// comparison, and validation with timezone support and zero external dependencies.
//
// # Key Features
//
//   - Pure and immutable functions (original dates are never modified)
//   - Timezone-aware operations across all functions
//   - Zero external dependencies (only Go standard library)
//   - WebAssembly compatible
//   - 100% test coverage with benchmarks
//
// # Function Categories
//
// Parsing: [Parse], [ParseISO], [ParseWithFormat]
// Formatting: [Format], [FormatCustom], [FormatSafe], [FormatDistance]
// Comparison: [IsBefore], [IsAfter], [IsEqual], [IsSameDay], [IsSameWeek]
// Manipulation: [AddDays], [AddHours], [AddMonths], [SubDays]
// Differences: [DifferenceInDays], [DifferenceInHours], [DifferenceInBusinessDays]
// Validation: [IsValid], [IsLeapYear], [IsWeekend], [IsWithinInterval]
// Period Utils: [StartOfDay], [EndOfDay], [StartOfWeek], [StartOfMonth]
// Interval: [EachDayOfInterval], [EachWeekOfInterval], [EachBusinessDayOfInterval]
//
// # Example
//
//	import "github.com/chmenegatti/go-date-fns/dateutils"
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
//	fmt.Println(isWeekend, isBefore)
//
//	// Period utilities
//	startOfWeek := dateutils.StartOfWeek(date)
//	endOfMonth := dateutils.EndOfMonth(date)
//	fmt.Println(startOfWeek, endOfMonth)
package dateutils
