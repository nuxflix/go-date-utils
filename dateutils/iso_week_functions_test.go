package dateutils

import (
	"testing"
	"time"
)

func TestGetISOWeek(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected int
	}{
		{
			name:     "January 2, 2005 (week 53 of 2004)",
			date:     time.Date(2005, time.January, 2, 0, 0, 0, 0, time.UTC),
			expected: 53,
		},
		{
			name:     "January 3, 2005 (week 1 of 2005)",
			date:     time.Date(2005, time.January, 3, 0, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "Middle of year",
			date:     time.Date(2024, time.June, 15, 0, 0, 0, 0, time.UTC),
			expected: 24,
		},
		{
			name:     "End of year",
			date:     time.Date(2024, time.December, 30, 0, 0, 0, 0, time.UTC),
			expected: 1, // This is week 1 of 2025
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetISOWeek(tt.date)
			if result != tt.expected {
				t.Errorf("GetISOWeek(%v) = %d, expected %d", tt.date, result, tt.expected)
			}
		})
	}
}

func TestGetISOWeekYear(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected int
	}{
		{
			name:     "January 2, 2005 (ISO year 2004)",
			date:     time.Date(2005, time.January, 2, 0, 0, 0, 0, time.UTC),
			expected: 2004,
		},
		{
			name:     "January 3, 2005 (ISO year 2005)",
			date:     time.Date(2005, time.January, 3, 0, 0, 0, 0, time.UTC),
			expected: 2005,
		},
		{
			name:     "Regular year middle",
			date:     time.Date(2024, time.June, 15, 0, 0, 0, 0, time.UTC),
			expected: 2024,
		},
		{
			name:     "December 30, 2024 (ISO year 2025)",
			date:     time.Date(2024, time.December, 30, 0, 0, 0, 0, time.UTC),
			expected: 2025,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetISOWeekYear(tt.date)
			if result != tt.expected {
				t.Errorf("GetISOWeekYear(%v) = %d, expected %d", tt.date, result, tt.expected)
			}
		})
	}
}

func TestGetISOWeeksInYear(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected int
	}{
		{
			name:     "2020 has 53 weeks",
			date:     time.Date(2020, time.June, 15, 0, 0, 0, 0, time.UTC),
			expected: 53,
		},
		{
			name:     "2021 has 52 weeks",
			date:     time.Date(2021, time.June, 15, 0, 0, 0, 0, time.UTC),
			expected: 52,
		},
		{
			name:     "2015 has 53 weeks",
			date:     time.Date(2015, time.February, 11, 0, 0, 0, 0, time.UTC),
			expected: 53,
		},
		{
			name:     "2019 has 52 weeks",
			date:     time.Date(2019, time.August, 15, 0, 0, 0, 0, time.UTC),
			expected: 52,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetISOWeeksInYear(tt.date)
			if result != tt.expected {
				t.Errorf("GetISOWeeksInYear(%v) = %d, expected %d", tt.date, result, tt.expected)
			}
		})
	}
}

func TestStartOfISOWeek(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected time.Time
	}{
		{
			name:     "Tuesday September 2, 2014",
			date:     time.Date(2014, time.September, 2, 11, 55, 0, 0, time.UTC),
			expected: time.Date(2014, time.September, 1, 0, 0, 0, 0, time.UTC), // Monday
		},
		{
			name:     "Sunday September 7, 2014",
			date:     time.Date(2014, time.September, 7, 15, 30, 0, 0, time.UTC),
			expected: time.Date(2014, time.September, 1, 0, 0, 0, 0, time.UTC), // Monday
		},
		{
			name:     "Monday September 1, 2014",
			date:     time.Date(2014, time.September, 1, 8, 0, 0, 0, time.UTC),
			expected: time.Date(2014, time.September, 1, 0, 0, 0, 0, time.UTC), // Same Monday
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StartOfISOWeek(tt.date)
			if !result.Equal(tt.expected) {
				t.Errorf("StartOfISOWeek(%v) = %v, expected %v", tt.date, result, tt.expected)
			}
		})
	}
}

func TestEndOfISOWeek(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected time.Time
	}{
		{
			name:     "Tuesday September 2, 2014",
			date:     time.Date(2014, time.September, 2, 11, 55, 0, 0, time.UTC),
			expected: time.Date(2014, time.September, 7, 23, 59, 59, 999000000, time.UTC), // Sunday
		},
		{
			name:     "Monday September 1, 2014",
			date:     time.Date(2014, time.September, 1, 8, 0, 0, 0, time.UTC),
			expected: time.Date(2014, time.September, 7, 23, 59, 59, 999000000, time.UTC), // Sunday
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EndOfISOWeek(tt.date)
			if !result.Equal(tt.expected) {
				t.Errorf("EndOfISOWeek(%v) = %v, expected %v", tt.date, result, tt.expected)
			}
		})
	}
}

func TestLastDayOfISOWeek(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected time.Time
	}{
		{
			name:     "Tuesday September 2, 2014",
			date:     time.Date(2014, time.September, 2, 11, 55, 0, 0, time.UTC),
			expected: time.Date(2014, time.September, 7, 0, 0, 0, 0, time.UTC), // Sunday
		},
		{
			name:     "Monday September 1, 2014",
			date:     time.Date(2014, time.September, 1, 8, 0, 0, 0, time.UTC),
			expected: time.Date(2014, time.September, 7, 0, 0, 0, 0, time.UTC), // Sunday
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LastDayOfISOWeek(tt.date)
			if !result.Equal(tt.expected) {
				t.Errorf("LastDayOfISOWeek(%v) = %v, expected %v", tt.date, result, tt.expected)
			}
		})
	}
}

func TestStartOfISOWeekYear(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected time.Time
	}{
		{
			name:     "July 2, 2005",
			date:     time.Date(2005, time.July, 2, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2005, time.January, 3, 0, 0, 0, 0, time.UTC), // Monday
		},
		{
			name:     "January 1, 2009",
			date:     time.Date(2009, time.January, 1, 16, 0, 0, 0, time.UTC),
			expected: time.Date(2008, time.December, 29, 0, 0, 0, 0, time.UTC), // Monday
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StartOfISOWeekYear(tt.date)
			if !result.Equal(tt.expected) {
				t.Errorf("StartOfISOWeekYear(%v) = %v, expected %v", tt.date, result, tt.expected)
			}
		})
	}
}

func TestEndOfISOWeekYear(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected time.Time
	}{
		{
			name:     "July 2, 2005",
			date:     time.Date(2005, time.July, 2, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2006, time.January, 1, 23, 59, 59, 999000000, time.UTC), // Sunday
		},
		{
			name:     "January 1, 2009",
			date:     time.Date(2009, time.January, 1, 16, 0, 0, 0, time.UTC),
			expected: time.Date(2010, time.January, 3, 23, 59, 59, 999000000, time.UTC), // Sunday
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EndOfISOWeekYear(tt.date)
			if !result.Equal(tt.expected) {
				t.Errorf("EndOfISOWeekYear(%v) = %v, expected %v", tt.date, result, tt.expected)
			}
		})
	}
}

func TestLastDayOfISOWeekYear(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected time.Time
	}{
		{
			name:     "July 2, 2005",
			date:     time.Date(2005, time.July, 2, 0, 0, 0, 0, time.UTC),
			expected: time.Date(2006, time.January, 1, 0, 0, 0, 0, time.UTC), // Sunday
		},
		{
			name:     "January 1, 2009",
			date:     time.Date(2009, time.January, 1, 16, 0, 0, 0, time.UTC),
			expected: time.Date(2010, time.January, 3, 0, 0, 0, 0, time.UTC), // Sunday
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LastDayOfISOWeekYear(tt.date)
			if !result.Equal(tt.expected) {
				t.Errorf("LastDayOfISOWeekYear(%v) = %v, expected %v", tt.date, result, tt.expected)
			}
		})
	}
}

func TestDifferenceInCalendarISOWeeks(t *testing.T) {
	tests := []struct {
		name        string
		laterDate   time.Time
		earlierDate time.Time
		expected    int
	}{
		{
			name:        "One ISO week difference",
			laterDate:   time.Date(2014, time.September, 8, 0, 0, 0, 0, time.UTC), // Monday
			earlierDate: time.Date(2014, time.September, 1, 0, 0, 0, 0, time.UTC), // Monday
			expected:    1,
		},
		{
			name:        "Same ISO week",
			laterDate:   time.Date(2014, time.September, 7, 0, 0, 0, 0, time.UTC), // Sunday
			earlierDate: time.Date(2014, time.September, 1, 0, 0, 0, 0, time.UTC), // Monday
			expected:    0,
		},
		{
			name:        "Multiple weeks",
			laterDate:   time.Date(2014, time.September, 22, 0, 0, 0, 0, time.UTC), // Monday
			earlierDate: time.Date(2014, time.September, 1, 0, 0, 0, 0, time.UTC),  // Monday
			expected:    3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DifferenceInCalendarISOWeeks(tt.laterDate, tt.earlierDate)
			if result != tt.expected {
				t.Errorf("DifferenceInCalendarISOWeeks(%v, %v) = %d, expected %d", tt.laterDate, tt.earlierDate, result, tt.expected)
			}
		})
	}
}

func TestDifferenceInCalendarISOWeekYears(t *testing.T) {
	tests := []struct {
		name        string
		laterDate   time.Time
		earlierDate time.Time
		expected    int
	}{
		{
			name:        "One ISO year difference",
			laterDate:   time.Date(2015, time.July, 2, 0, 0, 0, 0, time.UTC),
			earlierDate: time.Date(2014, time.July, 2, 0, 0, 0, 0, time.UTC),
			expected:    1,
		},
		{
			name:        "Same ISO year",
			laterDate:   time.Date(2014, time.December, 31, 0, 0, 0, 0, time.UTC),
			earlierDate: time.Date(2014, time.January, 1, 0, 0, 0, 0, time.UTC),
			expected:    1, // 2014-12-31 is in ISO year 2015, while 2014-01-01 is in ISO year 2014
		},
		{
			name:        "Cross-boundary case",
			laterDate:   time.Date(2005, time.January, 3, 0, 0, 0, 0, time.UTC), // ISO year 2005
			earlierDate: time.Date(2005, time.January, 2, 0, 0, 0, 0, time.UTC), // ISO year 2004
			expected:    1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DifferenceInCalendarISOWeekYears(tt.laterDate, tt.earlierDate)
			if result != tt.expected {
				t.Errorf("DifferenceInCalendarISOWeekYears(%v, %v) = %d, expected %d", tt.laterDate, tt.earlierDate, result, tt.expected)
			}
		})
	}
}

func TestIsSameISOWeek(t *testing.T) {
	tests := []struct {
		name     string
		date1    time.Time
		date2    time.Time
		expected bool
	}{
		{
			name:     "Same ISO week",
			date1:    time.Date(2014, time.September, 1, 0, 0, 0, 0, time.UTC), // Monday
			date2:    time.Date(2014, time.September, 7, 0, 0, 0, 0, time.UTC), // Sunday
			expected: true,
		},
		{
			name:     "Different ISO weeks",
			date1:    time.Date(2014, time.August, 31, 0, 0, 0, 0, time.UTC),   // Sunday
			date2:    time.Date(2014, time.September, 1, 0, 0, 0, 0, time.UTC), // Monday
			expected: false,
		},
		{
			name:     "Same date",
			date1:    time.Date(2014, time.September, 2, 11, 55, 0, 0, time.UTC),
			date2:    time.Date(2014, time.September, 2, 11, 55, 0, 0, time.UTC),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsSameISOWeek(tt.date1, tt.date2)
			if result != tt.expected {
				t.Errorf("IsSameISOWeek(%v, %v) = %t, expected %t", tt.date1, tt.date2, result, tt.expected)
			}
		})
	}
}

func TestIsSameISOWeekYear(t *testing.T) {
	tests := []struct {
		name     string
		date1    time.Time
		date2    time.Time
		expected bool
	}{
		{
			name:     "Same ISO year",
			date1:    time.Date(2014, time.January, 1, 0, 0, 0, 0, time.UTC),
			date2:    time.Date(2014, time.December, 31, 0, 0, 0, 0, time.UTC),
			expected: false, // 2014-01-01 is in ISO year 2014, but 2014-12-31 is in ISO year 2015
		},
		{
			name:     "Different ISO years (boundary case)",
			date1:    time.Date(2005, time.January, 1, 0, 0, 0, 0, time.UTC), // ISO year 2004
			date2:    time.Date(2005, time.January, 2, 0, 0, 0, 0, time.UTC), // ISO year 2004
			expected: true,
		},
		{
			name:     "Different ISO years",
			date1:    time.Date(2005, time.January, 2, 0, 0, 0, 0, time.UTC), // ISO year 2004
			date2:    time.Date(2005, time.January, 3, 0, 0, 0, 0, time.UTC), // ISO year 2005
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsSameISOWeekYear(tt.date1, tt.date2)
			if result != tt.expected {
				t.Errorf("IsSameISOWeekYear(%v, %v) = %t, expected %t", tt.date1, tt.date2, result, tt.expected)
			}
		})
	}
}

// Benchmark tests
func BenchmarkGetISOWeek(b *testing.B) {
	date := time.Date(2024, time.June, 15, 12, 30, 0, 0, time.UTC)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetISOWeek(date)
	}
}

func BenchmarkStartOfISOWeek(b *testing.B) {
	date := time.Date(2024, time.June, 15, 12, 30, 0, 0, time.UTC)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StartOfISOWeek(date)
	}
}

func BenchmarkGetISOWeekYear(b *testing.B) {
	date := time.Date(2024, time.June, 15, 12, 30, 0, 0, time.UTC)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetISOWeekYear(date)
	}
}
