package dateutils

import (
	"testing"
	"time"
)

func TestEachDayOfInterval(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name     string
		interval Interval
		expected []time.Time
	}{
		{
			name: "5 days interval",
			interval: Interval{
				Start: time.Date(2024, 1, 1, 12, 30, 0, 0, utc),
				End:   time.Date(2024, 1, 5, 18, 45, 0, 0, utc),
			},
			expected: []time.Time{
				time.Date(2024, 1, 1, 0, 0, 0, 0, utc),
				time.Date(2024, 1, 2, 0, 0, 0, 0, utc),
				time.Date(2024, 1, 3, 0, 0, 0, 0, utc),
				time.Date(2024, 1, 4, 0, 0, 0, 0, utc),
				time.Date(2024, 1, 5, 0, 0, 0, 0, utc),
			},
		},
		{
			name: "same day interval",
			interval: Interval{
				Start: time.Date(2024, 1, 1, 9, 0, 0, 0, utc),
				End:   time.Date(2024, 1, 1, 17, 0, 0, 0, utc),
			},
			expected: []time.Time{
				time.Date(2024, 1, 1, 0, 0, 0, 0, utc),
			},
		},
		{
			name: "invalid interval (start after end)",
			interval: Interval{
				Start: time.Date(2024, 1, 5, 0, 0, 0, 0, utc),
				End:   time.Date(2024, 1, 1, 0, 0, 0, 0, utc),
			},
			expected: []time.Time{},
		},
		{
			name: "cross month boundary",
			interval: Interval{
				Start: time.Date(2024, 1, 30, 0, 0, 0, 0, utc),
				End:   time.Date(2024, 2, 2, 0, 0, 0, 0, utc),
			},
			expected: []time.Time{
				time.Date(2024, 1, 30, 0, 0, 0, 0, utc),
				time.Date(2024, 1, 31, 0, 0, 0, 0, utc),
				time.Date(2024, 2, 1, 0, 0, 0, 0, utc),
				time.Date(2024, 2, 2, 0, 0, 0, 0, utc),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EachDayOfInterval(tt.interval)
			if len(result) != len(tt.expected) {
				t.Errorf("EachDayOfInterval() length = %v, expected %v", len(result), len(tt.expected))
				return
			}
			for i, day := range result {
				if !day.Equal(tt.expected[i]) {
					t.Errorf("EachDayOfInterval()[%d] = %v, expected %v", i, day, tt.expected[i])
				}
			}
		})
	}
}

func TestEachWeekOfInterval(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name     string
		interval Interval
		expected []time.Time
	}{
		{
			name: "3 weeks interval",
			interval: Interval{
				Start: time.Date(2024, 1, 8, 12, 0, 0, 0, utc),  // Monday
				End:   time.Date(2024, 1, 21, 15, 0, 0, 0, utc), // Sunday
			},
			expected: []time.Time{
				time.Date(2024, 1, 8, 0, 0, 0, 0, utc),  // Monday start of week
				time.Date(2024, 1, 15, 0, 0, 0, 0, utc), // Next Monday
			},
		},
		{
			name: "single week interval",
			interval: Interval{
				Start: time.Date(2024, 1, 10, 12, 0, 0, 0, utc), // Wednesday
				End:   time.Date(2024, 1, 12, 15, 0, 0, 0, utc), // Friday
			},
			expected: []time.Time{
				time.Date(2024, 1, 8, 0, 0, 0, 0, utc), // Monday of that week
			},
		},
		{
			name: "invalid interval",
			interval: Interval{
				Start: time.Date(2024, 1, 21, 0, 0, 0, 0, utc),
				End:   time.Date(2024, 1, 8, 0, 0, 0, 0, utc),
			},
			expected: []time.Time{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EachWeekOfInterval(tt.interval)
			if len(result) != len(tt.expected) {
				t.Errorf("EachWeekOfInterval() length = %v, expected %v", len(result), len(tt.expected))
				return
			}
			for i, week := range result {
				if !week.Equal(tt.expected[i]) {
					t.Errorf("EachWeekOfInterval()[%d] = %v, expected %v", i, week, tt.expected[i])
				}
			}
		})
	}
}

func TestEachWeekOfIntervalSunday(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name     string
		interval Interval
		expected []time.Time
	}{
		{
			name: "2 weeks interval (Sunday start)",
			interval: Interval{
				Start: time.Date(2024, 1, 8, 12, 0, 0, 0, utc),  // Monday
				End:   time.Date(2024, 1, 20, 15, 0, 0, 0, utc), // Saturday
			},
			expected: []time.Time{
				time.Date(2024, 1, 7, 0, 0, 0, 0, utc),  // Sunday start of week
				time.Date(2024, 1, 14, 0, 0, 0, 0, utc), // Next Sunday
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EachWeekOfIntervalSunday(tt.interval)
			if len(result) != len(tt.expected) {
				t.Errorf("EachWeekOfIntervalSunday() length = %v, expected %v", len(result), len(tt.expected))
				return
			}
			for i, week := range result {
				if !week.Equal(tt.expected[i]) {
					t.Errorf("EachWeekOfIntervalSunday()[%d] = %v, expected %v", i, week, tt.expected[i])
				}
			}
		})
	}
}

func TestEachMonthOfInterval(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name     string
		interval Interval
		expected []time.Time
	}{
		{
			name: "4 months interval",
			interval: Interval{
				Start: time.Date(2024, 1, 15, 12, 0, 0, 0, utc),
				End:   time.Date(2024, 4, 20, 15, 0, 0, 0, utc),
			},
			expected: []time.Time{
				time.Date(2024, 1, 1, 0, 0, 0, 0, utc),
				time.Date(2024, 2, 1, 0, 0, 0, 0, utc),
				time.Date(2024, 3, 1, 0, 0, 0, 0, utc),
				time.Date(2024, 4, 1, 0, 0, 0, 0, utc),
			},
		},
		{
			name: "same month interval",
			interval: Interval{
				Start: time.Date(2024, 1, 5, 12, 0, 0, 0, utc),
				End:   time.Date(2024, 1, 25, 15, 0, 0, 0, utc),
			},
			expected: []time.Time{
				time.Date(2024, 1, 1, 0, 0, 0, 0, utc),
			},
		},
		{
			name: "cross year boundary",
			interval: Interval{
				Start: time.Date(2023, 11, 15, 12, 0, 0, 0, utc),
				End:   time.Date(2024, 2, 20, 15, 0, 0, 0, utc),
			},
			expected: []time.Time{
				time.Date(2023, 11, 1, 0, 0, 0, 0, utc),
				time.Date(2023, 12, 1, 0, 0, 0, 0, utc),
				time.Date(2024, 1, 1, 0, 0, 0, 0, utc),
				time.Date(2024, 2, 1, 0, 0, 0, 0, utc),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EachMonthOfInterval(tt.interval)
			if len(result) != len(tt.expected) {
				t.Errorf("EachMonthOfInterval() length = %v, expected %v", len(result), len(tt.expected))
				return
			}
			for i, month := range result {
				if !month.Equal(tt.expected[i]) {
					t.Errorf("EachMonthOfInterval()[%d] = %v, expected %v", i, month, tt.expected[i])
				}
			}
		})
	}
}

func TestEachYearOfInterval(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name     string
		interval Interval
		expected []time.Time
	}{
		{
			name: "3 years interval",
			interval: Interval{
				Start: time.Date(2022, 5, 15, 12, 0, 0, 0, utc),
				End:   time.Date(2024, 8, 20, 15, 0, 0, 0, utc),
			},
			expected: []time.Time{
				time.Date(2022, 1, 1, 0, 0, 0, 0, utc),
				time.Date(2023, 1, 1, 0, 0, 0, 0, utc),
				time.Date(2024, 1, 1, 0, 0, 0, 0, utc),
			},
		},
		{
			name: "same year interval",
			interval: Interval{
				Start: time.Date(2024, 1, 5, 12, 0, 0, 0, utc),
				End:   time.Date(2024, 12, 25, 15, 0, 0, 0, utc),
			},
			expected: []time.Time{
				time.Date(2024, 1, 1, 0, 0, 0, 0, utc),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EachYearOfInterval(tt.interval)
			if len(result) != len(tt.expected) {
				t.Errorf("EachYearOfInterval() length = %v, expected %v", len(result), len(tt.expected))
				return
			}
			for i, year := range result {
				if !year.Equal(tt.expected[i]) {
					t.Errorf("EachYearOfInterval()[%d] = %v, expected %v", i, year, tt.expected[i])
				}
			}
		})
	}
}

func TestEachQuarterOfInterval(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name     string
		interval Interval
		expected []time.Time
	}{
		{
			name: "full year quarters",
			interval: Interval{
				Start: time.Date(2024, 1, 15, 12, 0, 0, 0, utc),
				End:   time.Date(2024, 12, 20, 15, 0, 0, 0, utc),
			},
			expected: []time.Time{
				time.Date(2024, 1, 1, 0, 0, 0, 0, utc),  // Q1
				time.Date(2024, 4, 1, 0, 0, 0, 0, utc),  // Q2
				time.Date(2024, 7, 1, 0, 0, 0, 0, utc),  // Q3
				time.Date(2024, 10, 1, 0, 0, 0, 0, utc), // Q4
			},
		},
		{
			name: "cross year quarters",
			interval: Interval{
				Start: time.Date(2023, 10, 15, 12, 0, 0, 0, utc), // Q4 2023
				End:   time.Date(2024, 7, 20, 15, 0, 0, 0, utc),  // Q3 2024
			},
			expected: []time.Time{
				time.Date(2023, 10, 1, 0, 0, 0, 0, utc), // Q4 2023
				time.Date(2024, 1, 1, 0, 0, 0, 0, utc),  // Q1 2024
				time.Date(2024, 4, 1, 0, 0, 0, 0, utc),  // Q2 2024
				time.Date(2024, 7, 1, 0, 0, 0, 0, utc),  // Q3 2024
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EachQuarterOfInterval(tt.interval)
			if len(result) != len(tt.expected) {
				t.Errorf("EachQuarterOfInterval() length = %v, expected %v", len(result), len(tt.expected))
				return
			}
			for i, quarter := range result {
				if !quarter.Equal(tt.expected[i]) {
					t.Errorf("EachQuarterOfInterval()[%d] = %v, expected %v", i, quarter, tt.expected[i])
				}
			}
		})
	}
}

func TestEachHourOfInterval(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name     string
		interval Interval
		expected []time.Time
	}{
		{
			name: "4 hours interval",
			interval: Interval{
				Start: time.Date(2024, 1, 1, 10, 30, 0, 0, utc),
				End:   time.Date(2024, 1, 1, 13, 45, 0, 0, utc),
			},
			expected: []time.Time{
				time.Date(2024, 1, 1, 10, 0, 0, 0, utc),
				time.Date(2024, 1, 1, 11, 0, 0, 0, utc),
				time.Date(2024, 1, 1, 12, 0, 0, 0, utc),
				time.Date(2024, 1, 1, 13, 0, 0, 0, utc),
			},
		},
		{
			name: "cross day boundary",
			interval: Interval{
				Start: time.Date(2024, 1, 1, 22, 30, 0, 0, utc),
				End:   time.Date(2024, 1, 2, 2, 15, 0, 0, utc),
			},
			expected: []time.Time{
				time.Date(2024, 1, 1, 22, 0, 0, 0, utc),
				time.Date(2024, 1, 1, 23, 0, 0, 0, utc),
				time.Date(2024, 1, 2, 0, 0, 0, 0, utc),
				time.Date(2024, 1, 2, 1, 0, 0, 0, utc),
				time.Date(2024, 1, 2, 2, 0, 0, 0, utc),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EachHourOfInterval(tt.interval)
			if len(result) != len(tt.expected) {
				t.Errorf("EachHourOfInterval() length = %v, expected %v", len(result), len(tt.expected))
				return
			}
			for i, hour := range result {
				if !hour.Equal(tt.expected[i]) {
					t.Errorf("EachHourOfInterval()[%d] = %v, expected %v", i, hour, tt.expected[i])
				}
			}
		})
	}
}

func TestEachMinuteOfInterval(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name     string
		interval Interval
		expected []time.Time
	}{
		{
			name: "5 minutes interval",
			interval: Interval{
				Start: time.Date(2024, 1, 1, 10, 28, 30, 0, utc),
				End:   time.Date(2024, 1, 1, 10, 32, 15, 0, utc),
			},
			expected: []time.Time{
				time.Date(2024, 1, 1, 10, 28, 0, 0, utc),
				time.Date(2024, 1, 1, 10, 29, 0, 0, utc),
				time.Date(2024, 1, 1, 10, 30, 0, 0, utc),
				time.Date(2024, 1, 1, 10, 31, 0, 0, utc),
				time.Date(2024, 1, 1, 10, 32, 0, 0, utc),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EachMinuteOfInterval(tt.interval)
			if len(result) != len(tt.expected) {
				t.Errorf("EachMinuteOfInterval() length = %v, expected %v", len(result), len(tt.expected))
				return
			}
			for i, minute := range result {
				if !minute.Equal(tt.expected[i]) {
					t.Errorf("EachMinuteOfInterval()[%d] = %v, expected %v", i, minute, tt.expected[i])
				}
			}
		})
	}
}

func TestEachWeekendOfInterval(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name     string
		interval Interval
		expected []time.Time
	}{
		{
			name: "week with weekends",
			interval: Interval{
				Start: time.Date(2024, 1, 5, 12, 0, 0, 0, utc), // Friday
				End:   time.Date(2024, 1, 8, 15, 0, 0, 0, utc), // Monday
			},
			expected: []time.Time{
				time.Date(2024, 1, 6, 0, 0, 0, 0, utc), // Saturday
				time.Date(2024, 1, 7, 0, 0, 0, 0, utc), // Sunday
			},
		},
		{
			name: "weekdays only",
			interval: Interval{
				Start: time.Date(2024, 1, 8, 12, 0, 0, 0, utc),  // Monday
				End:   time.Date(2024, 1, 10, 15, 0, 0, 0, utc), // Wednesday
			},
			expected: []time.Time{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EachWeekendOfInterval(tt.interval)
			if len(result) != len(tt.expected) {
				t.Errorf("EachWeekendOfInterval() length = %v, expected %v", len(result), len(tt.expected))
				return
			}
			for i, day := range result {
				if !day.Equal(tt.expected[i]) {
					t.Errorf("EachWeekendOfInterval()[%d] = %v, expected %v", i, day, tt.expected[i])
				}
			}
		})
	}
}

func TestEachBusinessDayOfInterval(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name     string
		interval Interval
		expected []time.Time
	}{
		{
			name: "week with business days",
			interval: Interval{
				Start: time.Date(2024, 1, 5, 12, 0, 0, 0, utc), // Friday
				End:   time.Date(2024, 1, 8, 15, 0, 0, 0, utc), // Monday
			},
			expected: []time.Time{
				time.Date(2024, 1, 5, 0, 0, 0, 0, utc), // Friday
				time.Date(2024, 1, 8, 0, 0, 0, 0, utc), // Monday
			},
		},
		{
			name: "weekend only",
			interval: Interval{
				Start: time.Date(2024, 1, 6, 12, 0, 0, 0, utc), // Saturday
				End:   time.Date(2024, 1, 7, 15, 0, 0, 0, utc), // Sunday
			},
			expected: []time.Time{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EachBusinessDayOfInterval(tt.interval)
			if len(result) != len(tt.expected) {
				t.Errorf("EachBusinessDayOfInterval() length = %v, expected %v", len(result), len(tt.expected))
				return
			}
			for i, day := range result {
				if !day.Equal(tt.expected[i]) {
					t.Errorf("EachBusinessDayOfInterval()[%d] = %v, expected %v", i, day, tt.expected[i])
				}
			}
		})
	}
}

func TestStartOfQuarter(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name     string
		date     time.Time
		expected time.Time
	}{
		{
			name:     "Q1 date (February)",
			date:     time.Date(2024, 2, 15, 14, 30, 45, 0, utc),
			expected: time.Date(2024, 1, 1, 0, 0, 0, 0, utc),
		},
		{
			name:     "Q2 date (May)",
			date:     time.Date(2024, 5, 20, 10, 15, 30, 0, utc),
			expected: time.Date(2024, 4, 1, 0, 0, 0, 0, utc),
		},
		{
			name:     "Q3 date (August)",
			date:     time.Date(2024, 8, 10, 16, 45, 0, 0, utc),
			expected: time.Date(2024, 7, 1, 0, 0, 0, 0, utc),
		},
		{
			name:     "Q4 date (November)",
			date:     time.Date(2024, 11, 25, 22, 30, 15, 0, utc),
			expected: time.Date(2024, 10, 1, 0, 0, 0, 0, utc),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StartOfQuarter(tt.date)
			if !result.Equal(tt.expected) {
				t.Errorf("StartOfQuarter() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestEndOfQuarter(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name     string
		date     time.Time
		expected time.Time
	}{
		{
			name:     "Q1 date (February)",
			date:     time.Date(2024, 2, 15, 14, 30, 45, 0, utc),
			expected: time.Date(2024, 3, 31, 23, 59, 59, 999999999, utc),
		},
		{
			name:     "Q2 date (May)",
			date:     time.Date(2024, 5, 20, 10, 15, 30, 0, utc),
			expected: time.Date(2024, 6, 30, 23, 59, 59, 999999999, utc),
		},
		{
			name:     "Q3 date (August)",
			date:     time.Date(2024, 8, 10, 16, 45, 0, 0, utc),
			expected: time.Date(2024, 9, 30, 23, 59, 59, 999999999, utc),
		},
		{
			name:     "Q4 date (November)",
			date:     time.Date(2024, 11, 25, 22, 30, 15, 0, utc),
			expected: time.Date(2024, 12, 31, 23, 59, 59, 999999999, utc),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EndOfQuarter(tt.date)
			if !result.Equal(tt.expected) {
				t.Errorf("EndOfQuarter() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

// Benchmark tests
func BenchmarkEachDayOfInterval(b *testing.B) {
	interval := Interval{
		Start: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		End:   time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC),
	}
	for i := 0; i < b.N; i++ {
		EachDayOfInterval(interval)
	}
}

func BenchmarkEachWeekOfInterval(b *testing.B) {
	interval := Interval{
		Start: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		End:   time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
	}
	for i := 0; i < b.N; i++ {
		EachWeekOfInterval(interval)
	}
}

func BenchmarkEachMonthOfInterval(b *testing.B) {
	interval := Interval{
		Start: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		End:   time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
	}
	for i := 0; i < b.N; i++ {
		EachMonthOfInterval(interval)
	}
}

func BenchmarkStartOfQuarter(b *testing.B) {
	date := time.Date(2024, 8, 15, 12, 0, 0, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		StartOfQuarter(date)
	}
}
