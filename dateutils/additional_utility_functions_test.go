package dateutils

import (
	"testing"
	"time"
)

func TestLightFormat(t *testing.T) {
	date := time.Date(2019, time.February, 11, 14, 0, 5, 123000000, time.UTC)

	tests := []struct {
		name     string
		format   string
		expected string
	}{
		{
			name:     "ISO date format",
			format:   "YYYY-MM-DD",
			expected: "2019-02-11",
		},
		{
			name:     "ISO datetime format",
			format:   "YYYY-MM-DD HH:mm:ss",
			expected: "2019-02-11 14:00:05",
		},
		{
			name:     "With milliseconds",
			format:   "YYYY-MM-DD HH:mm:ss.SSS",
			expected: "2019-02-11 14:00:05.123",
		},
		{
			name:     "US date format",
			format:   "MM/DD/YYYY",
			expected: "02/11/2019",
		},
		{
			name:     "Time only",
			format:   "HH:mm:ss",
			expected: "14:00:05",
		},
		{
			name:     "Empty format",
			format:   "",
			expected: "",
		},
		{
			name:     "No tokens",
			format:   "Today is a good day",
			expected: "Today is a good day",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LightFormat(date, tt.format)
			if result != tt.expected {
				t.Errorf("LightFormat(%v, %q) = %q, expected %q", date, tt.format, result, tt.expected)
			}
		})
	}
}

func TestRoundToNearestMinutes(t *testing.T) {
	tests := []struct {
		name      string
		date      time.Time
		nearestTo int
		expected  time.Time
	}{
		{
			name:      "Round down to nearest 5 minutes",
			date:      time.Date(2014, time.June, 10, 12, 7, 30, 0, time.UTC),
			nearestTo: 5,
			expected:  time.Date(2014, time.June, 10, 12, 5, 0, 0, time.UTC),
		},
		{
			name:      "Round up to nearest 5 minutes",
			date:      time.Date(2014, time.June, 10, 12, 8, 30, 0, time.UTC),
			nearestTo: 5,
			expected:  time.Date(2014, time.June, 10, 12, 10, 0, 0, time.UTC),
		},
		{
			name:      "Round to nearest 15 minutes",
			date:      time.Date(2014, time.June, 10, 12, 22, 30, 0, time.UTC),
			nearestTo: 15,
			expected:  time.Date(2014, time.June, 10, 12, 15, 0, 0, time.UTC),
		},
		{
			name:      "Round to nearest 30 minutes",
			date:      time.Date(2014, time.June, 10, 12, 45, 30, 0, time.UTC),
			nearestTo: 30,
			expected:  time.Date(2014, time.June, 10, 13, 0, 0, 0, time.UTC),
		},
		{
			name:      "Zero nearestTo (no change)",
			date:      time.Date(2014, time.June, 10, 12, 7, 30, 0, time.UTC),
			nearestTo: 0,
			expected:  time.Date(2014, time.June, 10, 12, 7, 30, 0, time.UTC),
		},
		{
			name:      "Negative nearestTo (no change)",
			date:      time.Date(2014, time.June, 10, 12, 7, 30, 0, time.UTC),
			nearestTo: -5,
			expected:  time.Date(2014, time.June, 10, 12, 7, 30, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RoundToNearestMinutes(tt.date, tt.nearestTo)
			if !result.Equal(tt.expected) {
				t.Errorf("RoundToNearestMinutes(%v, %d) = %v, expected %v", tt.date, tt.nearestTo, result, tt.expected)
			}
		})
	}
}

func TestStartOfDecade(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected time.Time
	}{
		{
			name:     "Mid-1980s decade",
			date:     time.Date(1985, time.October, 20, 0, 0, 0, 0, time.UTC),
			expected: time.Date(1980, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Start of decade",
			date:     time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "End of decade",
			date:     time.Date(2029, time.December, 31, 23, 59, 59, 0, time.UTC),
			expected: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StartOfDecade(tt.date)
			if !result.Equal(tt.expected) {
				t.Errorf("StartOfDecade(%v) = %v, expected %v", tt.date, result, tt.expected)
			}
		})
	}
}

func TestEndOfDecade(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected time.Time
	}{
		{
			name:     "Mid-1980s decade",
			date:     time.Date(1985, time.October, 20, 0, 0, 0, 0, time.UTC),
			expected: time.Date(1989, time.December, 31, 23, 59, 59, 999000000, time.UTC),
		},
		{
			name:     "Start of decade",
			date:     time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2029, time.December, 31, 23, 59, 59, 999000000, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EndOfDecade(tt.date)
			if !result.Equal(tt.expected) {
				t.Errorf("EndOfDecade(%v) = %v, expected %v", tt.date, result, tt.expected)
			}
		})
	}
}

func TestLastDayOfDecade(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected time.Time
	}{
		{
			name:     "Mid-1980s decade",
			date:     time.Date(1985, time.October, 20, 0, 0, 0, 0, time.UTC),
			expected: time.Date(1989, time.December, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Start of decade",
			date:     time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2029, time.December, 31, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LastDayOfDecade(tt.date)
			if !result.Equal(tt.expected) {
				t.Errorf("LastDayOfDecade(%v) = %v, expected %v", tt.date, result, tt.expected)
			}
		})
	}
}

func TestStartOfCentury(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected time.Time
	}{
		{
			name:     "Mid-20th century",
			date:     time.Date(1985, time.October, 20, 0, 0, 0, 0, time.UTC),
			expected: time.Date(1901, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Start of 21st century",
			date:     time.Date(2001, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2001, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "End of 20th century",
			date:     time.Date(2000, time.December, 31, 23, 59, 59, 0, time.UTC),
			expected: time.Date(1901, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StartOfCentury(tt.date)
			if !result.Equal(tt.expected) {
				t.Errorf("StartOfCentury(%v) = %v, expected %v", tt.date, result, tt.expected)
			}
		})
	}
}

func TestEndOfCentury(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected time.Time
	}{
		{
			name:     "Mid-20th century",
			date:     time.Date(1985, time.October, 20, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2000, time.December, 31, 23, 59, 59, 999000000, time.UTC),
		},
		{
			name:     "Start of 21st century",
			date:     time.Date(2001, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2100, time.December, 31, 23, 59, 59, 999000000, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EndOfCentury(tt.date)
			if !result.Equal(tt.expected) {
				t.Errorf("EndOfCentury(%v) = %v, expected %v", tt.date, result, tt.expected)
			}
		})
	}
}

func TestLastDayOfCentury(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected time.Time
	}{
		{
			name:     "Mid-20th century",
			date:     time.Date(1985, time.October, 20, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2000, time.December, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Start of 21st century",
			date:     time.Date(2001, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2100, time.December, 31, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LastDayOfCentury(tt.date)
			if !result.Equal(tt.expected) {
				t.Errorf("LastDayOfCentury(%v) = %v, expected %v", tt.date, result, tt.expected)
			}
		})
	}
}

func TestGetDaysInYear(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected int
	}{
		{
			name:     "Leap year 2020",
			date:     time.Date(2020, time.June, 15, 0, 0, 0, 0, time.UTC),
			expected: 366,
		},
		{
			name:     "Regular year 2021",
			date:     time.Date(2021, time.June, 15, 0, 0, 0, 0, time.UTC),
			expected: 365,
		},
		{
			name:     "Leap year 2000",
			date:     time.Date(2000, time.June, 15, 0, 0, 0, 0, time.UTC),
			expected: 366,
		},
		{
			name:     "Non-leap year 1900",
			date:     time.Date(1900, time.June, 15, 0, 0, 0, 0, time.UTC),
			expected: 365,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetDaysInYear(tt.date)
			if result != tt.expected {
				t.Errorf("GetDaysInYear(%v) = %d, expected %d", tt.date, result, tt.expected)
			}
		})
	}
}

func TestGetDaysInMonth(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected int
	}{
		{
			name:     "February in leap year",
			date:     time.Date(2020, time.February, 15, 0, 0, 0, 0, time.UTC),
			expected: 29,
		},
		{
			name:     "February in regular year",
			date:     time.Date(2021, time.February, 15, 0, 0, 0, 0, time.UTC),
			expected: 28,
		},
		{
			name:     "April (30 days)",
			date:     time.Date(2021, time.April, 15, 0, 0, 0, 0, time.UTC),
			expected: 30,
		},
		{
			name:     "January (31 days)",
			date:     time.Date(2021, time.January, 15, 0, 0, 0, 0, time.UTC),
			expected: 31,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetDaysInMonth(tt.date)
			if result != tt.expected {
				t.Errorf("GetDaysInMonth(%v) = %d, expected %d", tt.date, result, tt.expected)
			}
		})
	}
}

func TestWeekdayFunctions(t *testing.T) {
	// Create dates for each day of the week
	monday := time.Date(2014, time.September, 1, 0, 0, 0, 0, time.UTC)    // Monday
	tuesday := time.Date(2014, time.September, 2, 0, 0, 0, 0, time.UTC)   // Tuesday
	wednesday := time.Date(2014, time.September, 3, 0, 0, 0, 0, time.UTC) // Wednesday
	thursday := time.Date(2014, time.September, 4, 0, 0, 0, 0, time.UTC)  // Thursday
	friday := time.Date(2014, time.September, 5, 0, 0, 0, 0, time.UTC)    // Friday
	saturday := time.Date(2014, time.September, 6, 0, 0, 0, 0, time.UTC)  // Saturday
	sunday := time.Date(2014, time.September, 7, 0, 0, 0, 0, time.UTC)    // Sunday

	tests := []struct {
		name     string
		function func(time.Time) bool
		date     time.Time
		expected bool
	}{
		{"IsMonday with Monday", IsMonday, monday, true},
		{"IsMonday with Tuesday", IsMonday, tuesday, false},
		{"IsTuesday with Tuesday", IsTuesday, tuesday, true},
		{"IsTuesday with Monday", IsTuesday, monday, false},
		{"IsWednesday with Wednesday", IsWednesday, wednesday, true},
		{"IsWednesday with Monday", IsWednesday, monday, false},
		{"IsThursday with Thursday", IsThursday, thursday, true},
		{"IsThursday with Monday", IsThursday, monday, false},
		{"IsFriday with Friday", IsFriday, friday, true},
		{"IsFriday with Monday", IsFriday, monday, false},
		{"IsSaturday with Saturday", IsSaturday, saturday, true},
		{"IsSaturday with Monday", IsSaturday, monday, false},
		{"IsSunday with Sunday", IsSunday, sunday, true},
		{"IsSunday with Monday", IsSunday, monday, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.function(tt.date)
			if result != tt.expected {
				t.Errorf("%s = %t, expected %t", tt.name, result, tt.expected)
			}
		})
	}
}

// Benchmark tests
func BenchmarkLightFormat(b *testing.B) {
	date := time.Date(2019, time.February, 11, 14, 0, 5, 123000000, time.UTC)
	format := "YYYY-MM-DD HH:mm:ss.SSS"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LightFormat(date, format)
	}
}

func BenchmarkRoundToNearestMinutes(b *testing.B) {
	date := time.Date(2014, time.June, 10, 12, 7, 30, 0, time.UTC)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RoundToNearestMinutes(date, 5)
	}
}

func BenchmarkStartOfDecade(b *testing.B) {
	date := time.Date(1985, time.October, 20, 0, 0, 0, 0, time.UTC)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StartOfDecade(date)
	}
}

func BenchmarkGetDaysInMonth(b *testing.B) {
	date := time.Date(2020, time.February, 15, 0, 0, 0, 0, time.UTC)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetDaysInMonth(date)
	}
}
