package dateutils

import (
	"fmt"
	"math"
	"time"
)

// FormatDistanceOptions represents options for formatting distance between two times.
type FormatDistanceOptions struct {
	IncludeSeconds bool // Include seconds in output for more precision
	AddSuffix      bool // Add "ago" or "in" suffix to indicate past/future
}

// FormatDistance returns the distance between the given dates in words.
// This function provides human-readable relative time formatting similar to date-fns.
//
// The function supports various time ranges:
// - Less than a minute: "less than a minute", "half a minute"
// - Minutes: "1 minute", "2 minutes", etc.
// - Hours: "about 1 hour", "about 2 hours", etc.
// - Days: "1 day", "2 days", etc.
// - Months: "about 1 month", "2 months", etc.
// - Years: "about 1 year", "over 1 year", etc.
//
// Parameters:
//   - date: The date to compare
//   - baseDate: The date to compare with (reference point)
//   - options: Optional formatting options
//
// Returns a human-readable string describing the distance between the dates.
//
// Example:
//
//	now := time.Now()
//	past := now.Add(-2 * time.Hour)
//	result := FormatDistance(now, past, &FormatDistanceOptions{AddSuffix: true})
//	// Returns: "about 2 hours ago"
func FormatDistance(date, baseDate time.Time, options *FormatDistanceOptions) string {
	if options == nil {
		options = &FormatDistanceOptions{}
	}

	// Determine which date is later
	var laterDate, earlierDate time.Time
	var isInFuture bool

	if date.After(baseDate) {
		laterDate = date
		earlierDate = baseDate
		isInFuture = true
	} else {
		laterDate = baseDate
		earlierDate = date
		isInFuture = false
	}

	// Calculate difference in various units
	seconds := int(math.Abs(laterDate.Sub(earlierDate).Seconds()))
	minutes := int(math.Round(float64(seconds) / 60))

	var result string

	// 0 up to 2 minutes
	if minutes < 2 {
		if options.IncludeSeconds {
			if seconds < 5 {
				result = "less than 5 seconds"
			} else if seconds < 10 {
				result = "less than 10 seconds"
			} else if seconds < 20 {
				result = "less than 20 seconds"
			} else if seconds < 40 {
				result = "half a minute"
			} else if seconds < 60 {
				result = "less than a minute"
			} else {
				result = "1 minute"
			}
		} else {
			if minutes == 0 {
				result = "less than a minute"
			} else {
				result = "1 minute"
			}
		}
	} else if minutes < 45 {
		// 2 minutes up to 45 minutes
		result = formatMinutes(minutes)
	} else if minutes < 90 {
		// 45 minutes up to 1.5 hours
		result = "about 1 hour"
	} else if minutes < 1440 { // 24 hours
		// 1.5 hours up to 24 hours
		hours := int(math.Round(float64(minutes) / 60))
		result = formatHours(hours)
	} else if minutes < 2520 { // 1.75 days
		// 1 day up to 1.75 days
		result = "1 day"
	} else if minutes < 43200 { // 30 days
		// 1.75 days up to 30 days
		days := int(math.Round(float64(minutes) / 1440))
		result = formatDays(days)
	} else if minutes < 86400 { // 60 days
		// 1 month up to 2 months
		result = "about 1 month"
	} else {
		// Calculate months and years for longer periods
		monthsDiff := DifferenceInMonths(laterDate, earlierDate)
		if monthsDiff < 12 {
			result = formatMonths(monthsDiff)
		} else {
			years := monthsDiff / 12
			remainingMonths := monthsDiff % 12

			if remainingMonths < 3 {
				result = formatYears(years, "about")
			} else if remainingMonths < 9 {
				result = formatYears(years, "over")
			} else {
				result = formatYears(years+1, "almost")
			}
		}
	}

	// Add suffix if requested
	if options.AddSuffix {
		if isInFuture {
			result = "in " + result
		} else {
			result = result + " ago"
		}
	}

	return result
}

// FormatDistanceStrict returns the distance between dates in a strict format.
// Unlike FormatDistance, this function always returns exact values without approximations.
//
// Parameters:
//   - date: The date to compare
//   - baseDate: The date to compare with (reference point)
//   - options: Optional formatting options
//
// Returns an exact string describing the distance between the dates.
//
// Example:
//
//	now := time.Now()
//	past := now.Add(-125 * time.Minute)
//	result := FormatDistanceStrict(now, past, &FormatDistanceOptions{AddSuffix: true})
//	// Returns: "2 hours ago" (exact, no "about")
func FormatDistanceStrict(date, baseDate time.Time, options *FormatDistanceOptions) string {
	if options == nil {
		options = &FormatDistanceOptions{}
	}

	// Determine which date is later
	var laterDate, earlierDate time.Time
	var isInFuture bool

	if date.After(baseDate) {
		laterDate = date
		earlierDate = baseDate
		isInFuture = true
	} else {
		laterDate = baseDate
		earlierDate = date
		isInFuture = false
	}

	// Calculate exact differences
	totalSeconds := int(math.Abs(laterDate.Sub(earlierDate).Seconds()))

	var result string
	var value int
	var unit string

	if totalSeconds < 60 {
		value = totalSeconds
		unit = "second"
	} else if totalSeconds < 3600 {
		value = totalSeconds / 60
		unit = "minute"
	} else if totalSeconds < 86400 {
		value = totalSeconds / 3600
		unit = "hour"
	} else if totalSeconds < 604800 {
		value = totalSeconds / 86400
		unit = "day"
	} else if totalSeconds < 2629746 { // ~30.44 days
		value = totalSeconds / 604800
		unit = "week"
	} else if totalSeconds < 31556952 { // ~365.25 days
		value = DifferenceInMonths(laterDate, earlierDate)
		unit = "month"
	} else {
		value = DifferenceInYears(laterDate, earlierDate)
		unit = "year"
	}

	// Format the result
	if value == 1 {
		if unit == "hour" {
			result = "1 hour"
		} else {
			result = fmt.Sprintf("1 %s", unit)
		}
	} else {
		result = fmt.Sprintf("%d %ss", value, unit)
	}

	// Add suffix if requested
	if options.AddSuffix {
		if isInFuture {
			result = "in " + result
		} else {
			result = result + " ago"
		}
	}

	return result
}

// FormatDistanceToNow returns the distance from the given date to now in words.
// This is a convenience function that uses the current time as the base date.
//
// Parameters:
//   - date: The date to compare with now
//   - options: Optional formatting options
//
// Returns a human-readable string describing the distance from now.
//
// Example:
//
//	past := time.Now().Add(-30 * time.Minute)
//	result := FormatDistanceToNow(past, &FormatDistanceOptions{AddSuffix: true})
//	// Returns: "30 minutes ago"
func FormatDistanceToNow(date time.Time, options *FormatDistanceOptions) string {
	return FormatDistance(date, time.Now(), options)
}

// FormatDistanceToNowStrict returns the exact distance from the given date to now.
// This is a convenience function that uses the current time as the base date.
//
// Parameters:
//   - date: The date to compare with now
//   - options: Optional formatting options
//
// Returns an exact string describing the distance from now.
//
// Example:
//
//	past := time.Now().Add(-125 * time.Minute)
//	result := FormatDistanceToNowStrict(past, &FormatDistanceOptions{AddSuffix: true})
//	// Returns: "2 hours ago"
func FormatDistanceToNowStrict(date time.Time, options *FormatDistanceOptions) string {
	return FormatDistanceStrict(date, time.Now(), options)
}

// Helper functions for formatting different time units

func formatMinutes(minutes int) string {
	if minutes == 1 {
		return "1 minute"
	}
	return fmt.Sprintf("%d minutes", minutes)
}

func formatHours(hours int) string {
	if hours == 1 {
		return "about 1 hour"
	}
	return fmt.Sprintf("about %d hours", hours)
}

func formatDays(days int) string {
	if days == 1 {
		return "1 day"
	}
	return fmt.Sprintf("%d days", days)
}

func formatMonths(months int) string {
	if months == 1 {
		return "about 1 month"
	}
	return fmt.Sprintf("%d months", months)
}

func formatYears(years int, prefix string) string {
	if years == 1 {
		return fmt.Sprintf("%s 1 year", prefix)
	}
	return fmt.Sprintf("%s %d years", prefix, years)
}
