package dateutils

import (
	"errors"
	"strings"
	"time"
)

// Common format constants following Go's reference time
const (
	// Date formats
	DateISO   = "2006-01-02"
	DateUS    = "01/02/2006"
	DateEU    = "02/01/2006"
	DateSlash = "2006/01/02"
	DateDash  = "2006-01-02"
	DateDot   = "02.01.2006"

	// Time formats
	Time24      = "15:04:05"
	Time12      = "3:04:05 PM"
	TimeShort   = "15:04"
	Time12Short = "3:04 PM"

	// DateTime formats
	DateTime24  = "2006-01-02 15:04:05"
	DateTime12  = "2006-01-02 3:04:05 PM"
	DateTimeISO = "2006-01-02T15:04:05Z07:00"
	DateTimeRFC = time.RFC3339

	// Named formats
	Readable      = "January 2, 2006"
	ReadableShort = "Jan 2, 2006"
	WeekdayLong   = "Monday, January 2, 2006"
	WeekdayShort  = "Mon, Jan 2, 2006"
)

// Format formats a time.Time to string using the specified format.
// The format parameter uses Go's reference time layout: Mon Jan 2 15:04:05 MST 2006.
// If timezone is provided, the time will be converted to that timezone before formatting.
// Returns an error if the time is zero value.
func Format(t time.Time, format string, timezone *time.Location) (string, error) {
	if t.IsZero() {
		return "", errors.New("cannot format zero time")
	}
	if format == "" {
		return "", errors.New("format string cannot be empty")
	}

	// Convert to specified timezone if provided
	if timezone != nil {
		t = t.In(timezone)
	}

	return t.Format(format), nil
}

// FormatSafe is like Format but returns an empty string instead of an error for zero time.
// This is useful when you want to handle zero times gracefully.
func FormatSafe(t time.Time, format string, timezone *time.Location) string {
	result, err := Format(t, format, timezone)
	if err != nil {
		return ""
	}
	return result
}

// FormatWithDefault formats a time.Time to string, returning a default value if formatting fails.
// This is useful for displaying dates in templates or UI where you want fallback text.
func FormatWithDefault(t time.Time, format string, timezone *time.Location, defaultValue string) string {
	result, err := Format(t, format, timezone)
	if err != nil {
		return defaultValue
	}
	return result
}

// FormatCustom provides a more user-friendly formatting function with common placeholders.
// It replaces common date/time placeholders with Go's reference time format:
// - YYYY -> 2006 (4-digit year)
// - YY -> 06 (2-digit year)
// - MM -> 01 (2-digit month)
// - DD -> 02 (2-digit day)
// - HH -> 15 (24-hour format)
// - hh -> 03 (12-hour format)
// - mm -> 04 (minutes)
// - ss -> 05 (seconds)
// - AM/PM -> PM (meridiem)
func FormatCustom(t time.Time, customFormat string, timezone *time.Location) (string, error) {
	if t.IsZero() {
		return "", errors.New("cannot format zero time")
	}
	if customFormat == "" {
		return "", errors.New("format string cannot be empty")
	}

	// Convert to specified timezone if provided
	if timezone != nil {
		t = t.In(timezone)
	}

	// Replace custom placeholders with Go format
	// Order matters - replace longer patterns first to avoid conflicts
	goFormat := customFormat

	// Replace in specific order to avoid conflicts (longer patterns first)
	goFormat = strings.ReplaceAll(goFormat, "YYYY", "2006")
	goFormat = strings.ReplaceAll(goFormat, "AM/PM", "PM")
	goFormat = strings.ReplaceAll(goFormat, "YY", "06")
	goFormat = strings.ReplaceAll(goFormat, "MM", "01")
	goFormat = strings.ReplaceAll(goFormat, "DD", "02")
	goFormat = strings.ReplaceAll(goFormat, "HH", "15")
	goFormat = strings.ReplaceAll(goFormat, "hh", "03")
	goFormat = strings.ReplaceAll(goFormat, "mm", "04")
	goFormat = strings.ReplaceAll(goFormat, "ss", "05")

	return t.Format(goFormat), nil
}
