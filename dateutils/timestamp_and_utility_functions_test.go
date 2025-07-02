package dateutils

import (
	"testing"
	"time"
)

func TestGetUnixTime(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected int64
	}{
		{
			name:     "February 29, 2012 11:45:05",
			date:     time.Date(2012, time.February, 29, 11, 45, 5, 0, time.UTC),
			expected: 1330515905, // Corrected: 11:45:05 UTC
		},
		{
			name:     "Unix epoch",
			date:     time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected: 0,
		},
		{
			name:     "Year 2000",
			date:     time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected: 946684800,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetUnixTime(tt.date)
			if result != tt.expected {
				t.Errorf("GetUnixTime(%v) = %d, expected %d", tt.date, result, tt.expected)
			}
		})
	}
}

func TestFromUnixTime(t *testing.T) {
	tests := []struct {
		name     string
		unixTime int64
		expected time.Time
	}{
		{
			name:     "February 29, 2012 11:45:05",
			unixTime: 1330515905, // Corrected timestamp
			expected: time.Date(2012, time.February, 29, 11, 45, 5, 0, time.UTC),
		},
		{
			name:     "Unix epoch",
			unixTime: 0,
			expected: time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Year 2000",
			unixTime: 946684800,
			expected: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FromUnixTime(tt.unixTime)
			if !result.Equal(tt.expected) {
				t.Errorf("FromUnixTime(%d) = %v, expected %v", tt.unixTime, result, tt.expected)
			}
		})
	}
}
