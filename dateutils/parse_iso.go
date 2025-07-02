package dateutils

import (
	"errors"
	"time"
)

// ParseISO parses an ISO 8601 date string to time.Time.
// Supports various ISO 8601 formats including:
// - 2006-01-02T15:04:05Z
// - 2006-01-02T15:04:05+07:00
// - 2006-01-02T15:04:05.000Z
// - 2006-01-02
// If timezone is provided, the result will be converted to that timezone.
// Returns an error if the string is not a valid ISO 8601 format.
func ParseISO(isoStr string, timezone *time.Location) (time.Time, error) {
	if isoStr == "" {
		return time.Time{}, errors.New("ISO string cannot be empty")
	}

	// If no timezone is provided, use UTC
	if timezone == nil {
		timezone = time.UTC
	}

	// ISO 8601 formats to try
	isoFormats := []string{
		time.RFC3339Nano,
		time.RFC3339,
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05",
		"2006-01-02T15:04:05Z",
		"2006-01-02",
	}

	var lastErr error

	// Try parsing with each ISO format
	for _, format := range isoFormats {
		if parsedTime, err := time.Parse(format, isoStr); err == nil {
			// Convert to the specified timezone
			return parsedTime.In(timezone), nil
		} else {
			lastErr = err
		}
	}

	// If no format worked, return the last error
	return time.Time{}, errors.New("unable to parse ISO string: " + lastErr.Error())
}

// IsValidISO checks if a string is a valid ISO 8601 date format.
// Returns true if the string can be parsed as ISO 8601, false otherwise.
func IsValidISO(isoStr string) bool {
	_, err := ParseISO(isoStr, nil)
	return err == nil
}
