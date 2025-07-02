package dateutils

import (
	"errors"
	"time"
)

// CommonDateFormats contains commonly used date formats
var CommonDateFormats = []string{
	time.RFC3339,
	time.RFC3339Nano,
	"2006-01-02",
	"2006-01-02 15:04:05",
	"2006-01-02T15:04:05",
	"01/02/2006",
	"01-02-2006",
	"02/01/2006",
	"02-01-2006",
	"2006/01/02",
	"2006-01-02 15:04:05 -0700",
	"January 2, 2006",
	"Jan 2, 2006",
	"2 January 2006",
	"2 Jan 2006",
}

// Parse attempts to parse a date string using common date formats.
// It tries multiple formats and returns the first successful parse.
// If timezone is provided, the result will be converted to that timezone.
// Returns an error if the string cannot be parsed with any known format.
func Parse(dateStr string, timezone *time.Location) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, errors.New("date string cannot be empty")
	}

	// If no timezone is provided, use UTC
	if timezone == nil {
		timezone = time.UTC
	}

	var lastErr error

	// Try parsing with each common format
	for _, format := range CommonDateFormats {
		if parsedTime, err := time.Parse(format, dateStr); err == nil {
			// Convert to the specified timezone
			return parsedTime.In(timezone), nil
		} else {
			lastErr = err
		}
	}

	// If no format worked, return the last error
	return time.Time{}, errors.New("unable to parse date string: " + lastErr.Error())
}

// ParseWithFormat parses a date string using a specific format.
// If timezone is provided, the result will be converted to that timezone.
// Returns an error if the string cannot be parsed with the given format.
func ParseWithFormat(dateStr, format string, timezone *time.Location) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, errors.New("date string cannot be empty")
	}
	if format == "" {
		return time.Time{}, errors.New("format string cannot be empty")
	}

	// If no timezone is provided, use UTC
	if timezone == nil {
		timezone = time.UTC
	}

	parsedTime, err := time.Parse(format, dateStr)
	if err != nil {
		return time.Time{}, err
	}

	// Convert to the specified timezone
	return parsedTime.In(timezone), nil
}
