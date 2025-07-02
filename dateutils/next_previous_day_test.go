package dateutils

import (
	"testing"
	"time"
)

func TestNextDay(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		weekday  time.Weekday
		expected time.Time
	}{
		{
			name:     "Monday to Friday",
			date:     time.Date(2024, 1, 8, 12, 0, 0, 0, time.UTC), // Monday
			weekday:  time.Friday,
			expected: time.Date(2024, 1, 12, 12, 0, 0, 0, time.UTC), // Friday same week
		},
		{
			name:     "Friday to Monday",
			date:     time.Date(2024, 1, 12, 12, 0, 0, 0, time.UTC), // Friday
			weekday:  time.Monday,
			expected: time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC), // Monday next week
		},
		{
			name:     "Sunday to Sunday (next week)",
			date:     time.Date(2024, 1, 7, 12, 0, 0, 0, time.UTC), // Sunday
			weekday:  time.Sunday,
			expected: time.Date(2024, 1, 14, 12, 0, 0, 0, time.UTC), // Sunday next week
		},
		{
			name:     "Wednesday to Monday (next week)",
			date:     time.Date(2024, 1, 10, 12, 0, 0, 0, time.UTC), // Wednesday
			weekday:  time.Monday,
			expected: time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC), // Monday next week
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NextDay(tt.date, tt.weekday)
			if !result.Equal(tt.expected) {
				t.Errorf("NextDay() = %v, expected %v", result, tt.expected)
			}
			// Check that the weekday is correct
			if result.Weekday() != tt.weekday {
				t.Errorf("NextDay() weekday = %v, expected %v", result.Weekday(), tt.weekday)
			}
		})
	}
}

func TestPreviousDay(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		weekday  time.Weekday
		expected time.Time
	}{
		{
			name:     "Friday to Monday",
			date:     time.Date(2024, 1, 12, 12, 0, 0, 0, time.UTC), // Friday
			weekday:  time.Monday,
			expected: time.Date(2024, 1, 8, 12, 0, 0, 0, time.UTC), // Monday same week
		},
		{
			name:     "Monday to Friday",
			date:     time.Date(2024, 1, 8, 12, 0, 0, 0, time.UTC), // Monday
			weekday:  time.Friday,
			expected: time.Date(2024, 1, 5, 12, 0, 0, 0, time.UTC), // Friday previous week
		},
		{
			name:     "Sunday to Sunday (previous week)",
			date:     time.Date(2024, 1, 14, 12, 0, 0, 0, time.UTC), // Sunday
			weekday:  time.Sunday,
			expected: time.Date(2024, 1, 7, 12, 0, 0, 0, time.UTC), // Sunday previous week
		},
		{
			name:     "Wednesday to Friday (previous week)",
			date:     time.Date(2024, 1, 10, 12, 0, 0, 0, time.UTC), // Wednesday
			weekday:  time.Friday,
			expected: time.Date(2024, 1, 5, 12, 0, 0, 0, time.UTC), // Friday previous week
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PreviousDay(tt.date, tt.weekday)
			if !result.Equal(tt.expected) {
				t.Errorf("PreviousDay() = %v, expected %v", result, tt.expected)
			}
			// Check that the weekday is correct
			if result.Weekday() != tt.weekday {
				t.Errorf("PreviousDay() weekday = %v, expected %v", result.Weekday(), tt.weekday)
			}
		})
	}
}

func TestSpecificWeekdayFunctions(t *testing.T) {
	baseDate := time.Date(2024, 1, 10, 12, 0, 0, 0, time.UTC) // Wednesday

	tests := []struct {
		name     string
		function func(time.Time) time.Time
		expected time.Weekday
	}{
		{"NextSunday", NextSunday, time.Sunday},
		{"NextMonday", NextMonday, time.Monday},
		{"NextTuesday", NextTuesday, time.Tuesday},
		{"NextWednesday", NextWednesday, time.Wednesday},
		{"NextThursday", NextThursday, time.Thursday},
		{"NextFriday", NextFriday, time.Friday},
		{"NextSaturday", NextSaturday, time.Saturday},
		{"PreviousSunday", PreviousSunday, time.Sunday},
		{"PreviousMonday", PreviousMonday, time.Monday},
		{"PreviousTuesday", PreviousTuesday, time.Tuesday},
		{"PreviousWednesday", PreviousWednesday, time.Wednesday},
		{"PreviousThursday", PreviousThursday, time.Thursday},
		{"PreviousFriday", PreviousFriday, time.Friday},
		{"PreviousSaturday", PreviousSaturday, time.Saturday},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.function(baseDate)
			if result.Weekday() != tt.expected {
				t.Errorf("%s() weekday = %v, expected %v", tt.name, result.Weekday(), tt.expected)
			}
			// Ensure time components are preserved
			if result.Hour() != baseDate.Hour() || result.Minute() != baseDate.Minute() {
				t.Errorf("%s() time not preserved", tt.name)
			}
		})
	}
}

func TestNextDayImmutability(t *testing.T) {
	originalDate := time.Date(2024, 1, 10, 12, 30, 45, 0, time.UTC)
	NextDay(originalDate, time.Friday)

	// Original date should be unchanged
	expected := time.Date(2024, 1, 10, 12, 30, 45, 0, time.UTC)
	if !originalDate.Equal(expected) {
		t.Errorf("NextDay() modified original date")
	}
}

func TestPreviousDayImmutability(t *testing.T) {
	originalDate := time.Date(2024, 1, 10, 12, 30, 45, 0, time.UTC)
	PreviousDay(originalDate, time.Monday)

	// Original date should be unchanged
	expected := time.Date(2024, 1, 10, 12, 30, 45, 0, time.UTC)
	if !originalDate.Equal(expected) {
		t.Errorf("PreviousDay() modified original date")
	}
}

// Benchmark tests
func BenchmarkNextDay(b *testing.B) {
	date := time.Date(2024, 1, 10, 12, 0, 0, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		NextDay(date, time.Friday)
	}
}

func BenchmarkPreviousDay(b *testing.B) {
	date := time.Date(2024, 1, 10, 12, 0, 0, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		PreviousDay(date, time.Monday)
	}
}
