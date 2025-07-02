package dateutils

import (
	"testing"
	"time"
)

func TestSetDate(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		day      int
		expected time.Time
	}{
		{
			name:     "set valid day",
			date:     time.Date(2024, 1, 15, 12, 30, 45, 0, time.UTC),
			day:      20,
			expected: time.Date(2024, 1, 20, 12, 30, 45, 0, time.UTC),
		},
		{
			name:     "set day beyond month limit (adjust to last day)",
			date:     time.Date(2024, 2, 15, 12, 30, 45, 0, time.UTC), // February
			day:      31,
			expected: time.Date(2024, 2, 29, 12, 30, 45, 0, time.UTC), // 2024 is leap year
		},
		{
			name:     "set day to 1",
			date:     time.Date(2024, 6, 15, 12, 30, 45, 0, time.UTC),
			day:      1,
			expected: time.Date(2024, 6, 1, 12, 30, 45, 0, time.UTC),
		},
		{
			name:     "set invalid day (below 1)",
			date:     time.Date(2024, 6, 15, 12, 30, 45, 0, time.UTC),
			day:      0,
			expected: time.Date(2024, 6, 1, 12, 30, 45, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SetDate(tt.date, tt.day)
			if !result.Equal(tt.expected) {
				t.Errorf("SetDate() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestSetDay(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		day      int
		expected int // Expected day of week (0=Sunday)
	}{
		{
			name:     "Monday to Friday",
			date:     time.Date(2024, 1, 8, 12, 30, 45, 0, time.UTC), // Monday
			day:      5,                                              // Friday
			expected: 5,
		},
		{
			name:     "Friday to Sunday",
			date:     time.Date(2024, 1, 12, 12, 30, 45, 0, time.UTC), // Friday
			day:      0,                                               // Sunday
			expected: 0,
		},
		{
			name:     "Sunday to Wednesday",
			date:     time.Date(2024, 1, 7, 12, 30, 45, 0, time.UTC), // Sunday
			day:      3,                                              // Wednesday
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SetDay(tt.date, tt.day)
			if int(result.Weekday()) != tt.expected {
				t.Errorf("SetDay() weekday = %v, expected %v", int(result.Weekday()), tt.expected)
			}
			// Check that time components are preserved
			if result.Hour() != tt.date.Hour() || result.Minute() != tt.date.Minute() {
				t.Errorf("SetDay() time not preserved: got %v:%v, expected %v:%v",
					result.Hour(), result.Minute(), tt.date.Hour(), tt.date.Minute())
			}
		})
	}
}

func TestSetDayOfYear(t *testing.T) {
	tests := []struct {
		name          string
		date          time.Time
		dayOfYear     int
		expectedDay   int
		expectedMonth time.Month
	}{
		{
			name:          "set to January 1st",
			date:          time.Date(2024, 6, 15, 12, 30, 45, 0, time.UTC),
			dayOfYear:     1,
			expectedDay:   1,
			expectedMonth: time.January,
		},
		{
			name:          "set to February 29th (leap year)",
			date:          time.Date(2024, 6, 15, 12, 30, 45, 0, time.UTC),
			dayOfYear:     60, // Feb 29 in leap year
			expectedDay:   29,
			expectedMonth: time.February,
		},
		{
			name:          "set beyond year limit (adjust to last day)",
			date:          time.Date(2024, 6, 15, 12, 30, 45, 0, time.UTC),
			dayOfYear:     400,
			expectedDay:   31,
			expectedMonth: time.December,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SetDayOfYear(tt.date, tt.dayOfYear)
			if result.Day() != tt.expectedDay || result.Month() != tt.expectedMonth {
				t.Errorf("SetDayOfYear() = %v/%v, expected %v/%v",
					result.Month(), result.Day(), tt.expectedMonth, tt.expectedDay)
			}
			// Check that time components are preserved
			if result.Hour() != tt.date.Hour() || result.Minute() != tt.date.Minute() {
				t.Errorf("SetDayOfYear() time not preserved")
			}
		})
	}
}

func TestSetHours(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		hours    int
		expected int
	}{
		{
			name:     "set to noon",
			date:     time.Date(2024, 1, 15, 8, 30, 45, 0, time.UTC),
			hours:    12,
			expected: 12,
		},
		{
			name:     "set to midnight",
			date:     time.Date(2024, 1, 15, 14, 30, 45, 0, time.UTC),
			hours:    0,
			expected: 0,
		},
		{
			name:     "set beyond limit (adjust to 23)",
			date:     time.Date(2024, 1, 15, 14, 30, 45, 0, time.UTC),
			hours:    25,
			expected: 23,
		},
		{
			name:     "set below limit (adjust to 0)",
			date:     time.Date(2024, 1, 15, 14, 30, 45, 0, time.UTC),
			hours:    -1,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SetHours(tt.date, tt.hours)
			if result.Hour() != tt.expected {
				t.Errorf("SetHours() = %v, expected %v", result.Hour(), tt.expected)
			}
			// Check that other components are preserved
			if result.Minute() != tt.date.Minute() || result.Second() != tt.date.Second() {
				t.Errorf("SetHours() other time components not preserved")
			}
		})
	}
}

func TestSetMinutes(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		minutes  int
		expected int
	}{
		{
			name:     "set to 30 minutes",
			date:     time.Date(2024, 1, 15, 12, 15, 45, 0, time.UTC),
			minutes:  30,
			expected: 30,
		},
		{
			name:     "set beyond limit (adjust to 59)",
			date:     time.Date(2024, 1, 15, 12, 15, 45, 0, time.UTC),
			minutes:  65,
			expected: 59,
		},
		{
			name:     "set below limit (adjust to 0)",
			date:     time.Date(2024, 1, 15, 12, 15, 45, 0, time.UTC),
			minutes:  -5,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SetMinutes(tt.date, tt.minutes)
			if result.Minute() != tt.expected {
				t.Errorf("SetMinutes() = %v, expected %v", result.Minute(), tt.expected)
			}
		})
	}
}

func TestSetSeconds(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		seconds  int
		expected int
	}{
		{
			name:     "set to 30 seconds",
			date:     time.Date(2024, 1, 15, 12, 30, 15, 0, time.UTC),
			seconds:  30,
			expected: 30,
		},
		{
			name:     "set beyond limit (adjust to 59)",
			date:     time.Date(2024, 1, 15, 12, 30, 15, 0, time.UTC),
			seconds:  65,
			expected: 59,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SetSeconds(tt.date, tt.seconds)
			if result.Second() != tt.expected {
				t.Errorf("SetSeconds() = %v, expected %v", result.Second(), tt.expected)
			}
		})
	}
}

func TestSetMilliseconds(t *testing.T) {
	tests := []struct {
		name         string
		date         time.Time
		milliseconds int
		expected     int
	}{
		{
			name:         "set to 500 milliseconds",
			date:         time.Date(2024, 1, 15, 12, 30, 45, 100*1000000, time.UTC),
			milliseconds: 500,
			expected:     500,
		},
		{
			name:         "set beyond limit (adjust to 999)",
			date:         time.Date(2024, 1, 15, 12, 30, 45, 100*1000000, time.UTC),
			milliseconds: 1500,
			expected:     999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SetMilliseconds(tt.date, tt.milliseconds)
			actualMs := result.Nanosecond() / 1000000
			if actualMs != tt.expected {
				t.Errorf("SetMilliseconds() = %v, expected %v", actualMs, tt.expected)
			}
		})
	}
}

func TestSetMonth(t *testing.T) {
	tests := []struct {
		name          string
		date          time.Time
		month         int
		expectedMonth time.Month
		expectedDay   int
	}{
		{
			name:          "set to June",
			date:          time.Date(2024, 1, 15, 12, 30, 45, 0, time.UTC),
			month:         6,
			expectedMonth: time.June,
			expectedDay:   15,
		},
		{
			name:          "set month with day adjustment (Jan 31 -> Feb 29)",
			date:          time.Date(2024, 1, 31, 12, 30, 45, 0, time.UTC),
			month:         2,
			expectedMonth: time.February,
			expectedDay:   29, // 2024 is leap year
		},
		{
			name:          "set beyond limit (adjust to December)",
			date:          time.Date(2024, 6, 15, 12, 30, 45, 0, time.UTC),
			month:         15,
			expectedMonth: time.December,
			expectedDay:   15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SetMonth(tt.date, tt.month)
			if result.Month() != tt.expectedMonth || result.Day() != tt.expectedDay {
				t.Errorf("SetMonth() = %v/%v, expected %v/%v",
					result.Month(), result.Day(), tt.expectedMonth, tt.expectedDay)
			}
		})
	}
}

func TestSetYear(t *testing.T) {
	tests := []struct {
		name         string
		date         time.Time
		year         int
		expectedYear int
		expectedDay  int
	}{
		{
			name:         "set to 2025",
			date:         time.Date(2024, 6, 15, 12, 30, 45, 0, time.UTC),
			year:         2025,
			expectedYear: 2025,
			expectedDay:  15,
		},
		{
			name:         "leap year adjustment (Feb 29 -> Feb 28)",
			date:         time.Date(2024, 2, 29, 12, 30, 45, 0, time.UTC), // Leap year
			year:         2023,                                            // Non-leap year
			expectedYear: 2023,
			expectedDay:  28,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SetYear(tt.date, tt.year)
			if result.Year() != tt.expectedYear || result.Day() != tt.expectedDay {
				t.Errorf("SetYear() = %v/%v, expected %v/%v",
					result.Year(), result.Day(), tt.expectedYear, tt.expectedDay)
			}
		})
	}
}

func TestSetQuarter(t *testing.T) {
	tests := []struct {
		name          string
		date          time.Time
		quarter       int
		expectedMonth time.Month
	}{
		{
			name:          "set to Q1",
			date:          time.Date(2024, 6, 15, 12, 30, 45, 0, time.UTC),
			quarter:       1,
			expectedMonth: time.January,
		},
		{
			name:          "set to Q2",
			date:          time.Date(2024, 1, 15, 12, 30, 45, 0, time.UTC),
			quarter:       2,
			expectedMonth: time.April,
		},
		{
			name:          "set to Q3",
			date:          time.Date(2024, 1, 15, 12, 30, 45, 0, time.UTC),
			quarter:       3,
			expectedMonth: time.July,
		},
		{
			name:          "set to Q4",
			date:          time.Date(2024, 1, 15, 12, 30, 45, 0, time.UTC),
			quarter:       4,
			expectedMonth: time.October,
		},
		{
			name:          "set beyond limit (adjust to Q4)",
			date:          time.Date(2024, 1, 15, 12, 30, 45, 0, time.UTC),
			quarter:       5,
			expectedMonth: time.October,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SetQuarter(tt.date, tt.quarter)
			if result.Month() != tt.expectedMonth {
				t.Errorf("SetQuarter() month = %v, expected %v", result.Month(), tt.expectedMonth)
			}
		})
	}
}

// Test immutability
func TestSetFunctionsImmutability(t *testing.T) {
	original := time.Date(2024, 6, 15, 12, 30, 45, 123456789, time.UTC)

	// Test each set function doesn't modify the original
	SetDate(original, 20)
	SetDay(original, 0)
	SetDayOfYear(original, 100)
	SetHours(original, 15)
	SetMinutes(original, 45)
	SetSeconds(original, 30)
	SetMilliseconds(original, 500)
	SetMonth(original, 12)
	SetYear(original, 2025)
	SetQuarter(original, 4)

	// Original should be unchanged
	expected := time.Date(2024, 6, 15, 12, 30, 45, 123456789, time.UTC)
	if !original.Equal(expected) {
		t.Errorf("Original time was modified: got %v, expected %v", original, expected)
	}
}

// Benchmark tests
func BenchmarkSetDate(b *testing.B) {
	date := time.Date(2024, 6, 15, 12, 0, 0, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		SetDate(date, 20)
	}
}

func BenchmarkSetMonth(b *testing.B) {
	date := time.Date(2024, 6, 15, 12, 0, 0, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		SetMonth(date, 12)
	}
}
