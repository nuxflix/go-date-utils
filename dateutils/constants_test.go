package dateutils

import (
	"testing"
)

func TestConstants(t *testing.T) {
	// Test basic time constants
	if MillisecondsInSecond != 1000 {
		t.Errorf("MillisecondsInSecond = %d, expected 1000", MillisecondsInSecond)
	}

	if SecondsInMinute != 60 {
		t.Errorf("SecondsInMinute = %d, expected 60", SecondsInMinute)
	}

	if MinutesInHour != 60 {
		t.Errorf("MinutesInHour = %d, expected 60", MinutesInHour)
	}

	if HoursInDay != 24 {
		t.Errorf("HoursInDay = %d, expected 24", HoursInDay)
	}

	// Test date constants
	if DaysInWeek != 7 {
		t.Errorf("DaysInWeek = %d, expected 7", DaysInWeek)
	}

	if MonthsInYear != 12 {
		t.Errorf("MonthsInYear = %d, expected 12", MonthsInYear)
	}

	if QuartersInYear != 4 {
		t.Errorf("QuartersInYear = %d, expected 4", QuartersInYear)
	}
}

func TestDurationConstants(t *testing.T) {
	// Test derived constants
	expectedMillisecondsInMinute := MillisecondsInSecond * SecondsInMinute
	if MillisecondsInMinute != expectedMillisecondsInMinute {
		t.Errorf("MillisecondsInMinute = %d, expected %d", MillisecondsInMinute, expectedMillisecondsInMinute)
	}

	expectedMillisecondsInHour := MillisecondsInMinute * MinutesInHour
	if MillisecondsInHour != expectedMillisecondsInHour {
		t.Errorf("MillisecondsInHour = %d, expected %d", MillisecondsInHour, expectedMillisecondsInHour)
	}

	expectedMillisecondsInDay := MillisecondsInHour * HoursInDay
	if MillisecondsInDay != expectedMillisecondsInDay {
		t.Errorf("MillisecondsInDay = %d, expected %d", MillisecondsInDay, expectedMillisecondsInDay)
	}

	expectedMillisecondsInWeek := MillisecondsInDay * DaysInWeek
	if MillisecondsInWeek != expectedMillisecondsInWeek {
		t.Errorf("MillisecondsInWeek = %d, expected %d", MillisecondsInWeek, expectedMillisecondsInWeek)
	}
}

func TestTimeRangeConstants(t *testing.T) {
	// Test max/min time constants
	expectedMaxTime := int64(8640000000000000)
	if MaxTime != expectedMaxTime {
		t.Errorf("MaxTime = %d, expected %d", MaxTime, expectedMaxTime)
	}

	expectedMinTime := int64(-8640000000000000)
	if MinTime != expectedMinTime {
		t.Errorf("MinTime = %d, expected %d", MinTime, expectedMinTime)
	}
}

func TestIsValidTimestamp(t *testing.T) {
	tests := []struct {
		name      string
		timestamp int64
		expected  bool
	}{
		{
			name:      "valid timestamp (zero)",
			timestamp: 0,
			expected:  true,
		},
		{
			name:      "valid timestamp (positive)",
			timestamp: 1640995200000, // 2022-01-01 00:00:00 UTC
			expected:  true,
		},
		{
			name:      "valid timestamp (negative)",
			timestamp: -2208988800000, // 1900-01-01 00:00:00 UTC
			expected:  true,
		},
		{
			name:      "invalid timestamp (too large)",
			timestamp: MaxTime + 1,
			expected:  false,
		},
		{
			name:      "invalid timestamp (too small)",
			timestamp: MinTime - 1,
			expected:  false,
		},
		{
			name:      "max valid timestamp",
			timestamp: MaxTime,
			expected:  true,
		},
		{
			name:      "min valid timestamp",
			timestamp: MinTime,
			expected:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidTimestamp(tt.timestamp)
			if result != tt.expected {
				t.Errorf("IsValidTimestamp(%d) = %v, expected %v", tt.timestamp, result, tt.expected)
			}
		})
	}
}

func TestWeekdayConstants(t *testing.T) {
	if Sunday != 0 {
		t.Errorf("Sunday = %d, expected 0", Sunday)
	}
	if Monday != 1 {
		t.Errorf("Monday = %d, expected 1", Monday)
	}
	if Tuesday != 2 {
		t.Errorf("Tuesday = %d, expected 2", Tuesday)
	}
	if Wednesday != 3 {
		t.Errorf("Wednesday = %d, expected 3", Wednesday)
	}
	if Thursday != 4 {
		t.Errorf("Thursday = %d, expected 4", Thursday)
	}
	if Friday != 5 {
		t.Errorf("Friday = %d, expected 5", Friday)
	}
	if Saturday != 6 {
		t.Errorf("Saturday = %d, expected 6", Saturday)
	}
}

func TestMonthConstants(t *testing.T) {
	months := []struct {
		constant int
		expected int
		name     string
	}{
		{January, 1, "January"},
		{February, 2, "February"},
		{March, 3, "March"},
		{April, 4, "April"},
		{May, 5, "May"},
		{June, 6, "June"},
		{July, 7, "July"},
		{August, 8, "August"},
		{September, 9, "September"},
		{October, 10, "October"},
		{November, 11, "November"},
		{December, 12, "December"},
	}

	for _, month := range months {
		if month.constant != month.expected {
			t.Errorf("%s = %d, expected %d", month.name, month.constant, month.expected)
		}
	}
}

func TestQuarterConstants(t *testing.T) {
	if Q1 != 1 {
		t.Errorf("Q1 = %d, expected 1", Q1)
	}
	if Q2 != 2 {
		t.Errorf("Q2 = %d, expected 2", Q2)
	}
	if Q3 != 3 {
		t.Errorf("Q3 = %d, expected 3", Q3)
	}
	if Q4 != 4 {
		t.Errorf("Q4 = %d, expected 4", Q4)
	}
}

func TestAverageConstants(t *testing.T) {
	// Test that average constants are reasonable
	if DaysInYear < 365 || DaysInYear > 366 {
		t.Errorf("DaysInYear = %f, expected between 365 and 366", DaysInYear)
	}

	if DaysInMonth < 28 || DaysInMonth > 32 {
		t.Errorf("DaysInMonth = %f, expected between 28 and 32", DaysInMonth)
	}

	if WeeksInYear < 52 || WeeksInYear > 53 {
		t.Errorf("WeeksInYear = %f, expected between 52 and 53", WeeksInYear)
	}
}

// Benchmark tests
func BenchmarkIsValidTimestamp(b *testing.B) {
	timestamp := int64(1640995200000)
	for i := 0; i < b.N; i++ {
		IsValidTimestamp(timestamp)
	}
}
