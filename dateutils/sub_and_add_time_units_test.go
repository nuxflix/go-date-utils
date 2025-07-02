package dateutils

import (
	"testing"
	"time"
)

func TestSubDays(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		days     int
		expected time.Time
	}{
		{
			name:     "subtract positive days",
			date:     time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC),
			days:     7,
			expected: time.Date(2024, 1, 8, 12, 0, 0, 0, time.UTC),
		},
		{
			name:     "subtract zero days",
			date:     time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC),
			days:     0,
			expected: time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC),
		},
		{
			name:     "subtract negative days (effectively add)",
			date:     time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC),
			days:     -7,
			expected: time.Date(2024, 1, 22, 12, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SubDays(tt.date, tt.days)
			if !result.Equal(tt.expected) {
				t.Errorf("SubDays() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestSubWeeks(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		weeks    int
		expected time.Time
	}{
		{
			name:     "subtract one week",
			date:     time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC),
			weeks:    1,
			expected: time.Date(2024, 1, 8, 12, 0, 0, 0, time.UTC),
		},
		{
			name:     "subtract multiple weeks",
			date:     time.Date(2024, 1, 29, 12, 0, 0, 0, time.UTC),
			weeks:    4,
			expected: time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SubWeeks(tt.date, tt.weeks)
			if !result.Equal(tt.expected) {
				t.Errorf("SubWeeks() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestSubBusinessDays(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		days     int
		expected time.Time
	}{
		{
			name:     "subtract business days from Friday",
			date:     time.Date(2024, 1, 12, 12, 0, 0, 0, time.UTC), // Friday
			days:     5,
			expected: time.Date(2024, 1, 5, 12, 0, 0, 0, time.UTC), // Previous Friday
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SubBusinessDays(tt.date, tt.days)
			if !result.Equal(tt.expected) {
				t.Errorf("SubBusinessDays() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestAddHours(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		hours    int
		expected time.Time
	}{
		{
			name:     "add positive hours",
			date:     time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC),
			hours:    5,
			expected: time.Date(2024, 1, 15, 17, 0, 0, 0, time.UTC),
		},
		{
			name:     "add hours crossing day boundary",
			date:     time.Date(2024, 1, 15, 22, 0, 0, 0, time.UTC),
			hours:    4,
			expected: time.Date(2024, 1, 16, 2, 0, 0, 0, time.UTC),
		},
		{
			name:     "subtract hours (negative)",
			date:     time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC),
			hours:    -5,
			expected: time.Date(2024, 1, 15, 7, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AddHours(tt.date, tt.hours)
			if !result.Equal(tt.expected) {
				t.Errorf("AddHours() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestSubHours(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		hours    int
		expected time.Time
	}{
		{
			name:     "subtract hours",
			date:     time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC),
			hours:    5,
			expected: time.Date(2024, 1, 15, 7, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SubHours(tt.date, tt.hours)
			if !result.Equal(tt.expected) {
				t.Errorf("SubHours() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestAddMinutes(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		minutes  int
		expected time.Time
	}{
		{
			name:     "add minutes",
			date:     time.Date(2024, 1, 15, 12, 30, 0, 0, time.UTC),
			minutes:  15,
			expected: time.Date(2024, 1, 15, 12, 45, 0, 0, time.UTC),
		},
		{
			name:     "add minutes crossing hour boundary",
			date:     time.Date(2024, 1, 15, 12, 50, 0, 0, time.UTC),
			minutes:  20,
			expected: time.Date(2024, 1, 15, 13, 10, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AddMinutes(tt.date, tt.minutes)
			if !result.Equal(tt.expected) {
				t.Errorf("AddMinutes() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestAddSeconds(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		seconds  int
		expected time.Time
	}{
		{
			name:     "add seconds",
			date:     time.Date(2024, 1, 15, 12, 30, 30, 0, time.UTC),
			seconds:  15,
			expected: time.Date(2024, 1, 15, 12, 30, 45, 0, time.UTC),
		},
		{
			name:     "add seconds crossing minute boundary",
			date:     time.Date(2024, 1, 15, 12, 30, 50, 0, time.UTC),
			seconds:  20,
			expected: time.Date(2024, 1, 15, 12, 31, 10, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AddSeconds(tt.date, tt.seconds)
			if !result.Equal(tt.expected) {
				t.Errorf("AddSeconds() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestAddMonths(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		months   int
		expected time.Time
	}{
		{
			name:     "add months",
			date:     time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC),
			months:   2,
			expected: time.Date(2024, 3, 15, 12, 0, 0, 0, time.UTC),
		},
		{
			name:     "add months crossing year boundary",
			date:     time.Date(2024, 11, 15, 12, 0, 0, 0, time.UTC),
			months:   3,
			expected: time.Date(2025, 2, 15, 12, 0, 0, 0, time.UTC),
		},
		{
			name:     "add months with day adjustment (Jan 31 + 1 month = Feb 29 in leap year)",
			date:     time.Date(2024, 1, 31, 12, 0, 0, 0, time.UTC),
			months:   1,
			expected: time.Date(2024, 2, 29, 12, 0, 0, 0, time.UTC), // 2024 is leap year
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AddMonths(tt.date, tt.months)
			if !result.Equal(tt.expected) {
				t.Errorf("AddMonths() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestAddYears(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		years    int
		expected time.Time
	}{
		{
			name:     "add years",
			date:     time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC),
			years:    2,
			expected: time.Date(2026, 1, 15, 12, 0, 0, 0, time.UTC),
		},
		{
			name:     "add years with leap year adjustment (Feb 29 -> Feb 28)",
			date:     time.Date(2024, 2, 29, 12, 0, 0, 0, time.UTC), // Leap year
			years:    1,
			expected: time.Date(2025, 2, 28, 12, 0, 0, 0, time.UTC), // Non-leap year
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AddYears(tt.date, tt.years)
			if !result.Equal(tt.expected) {
				t.Errorf("AddYears() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

// Benchmarks
func BenchmarkSubDays(b *testing.B) {
	date := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		SubDays(date, 7)
	}
}

func BenchmarkAddHours(b *testing.B) {
	date := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		AddHours(date, 5)
	}
}

func BenchmarkAddMinutes(b *testing.B) {
	date := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		AddMinutes(date, 30)
	}
}

func BenchmarkAddMonths(b *testing.B) {
	date := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		AddMonths(date, 2)
	}
}
