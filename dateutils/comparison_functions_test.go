package dateutils

import (
	"testing"
	"time"
)

func TestIsEqual(t *testing.T) {
	tests := []struct {
		name     string
		t1       time.Time
		t2       time.Time
		expected bool
	}{
		{
			name:     "same times",
			t1:       time.Date(2024, 1, 15, 12, 30, 45, 0, time.UTC),
			t2:       time.Date(2024, 1, 15, 12, 30, 45, 0, time.UTC),
			expected: true,
		},
		{
			name:     "different times",
			t1:       time.Date(2024, 1, 15, 12, 30, 45, 0, time.UTC),
			t2:       time.Date(2024, 1, 15, 12, 30, 46, 0, time.UTC),
			expected: false,
		},
		{
			name:     "same instant, different timezones",
			t1:       time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC),
			t2:       time.Date(2024, 1, 15, 7, 0, 0, 0, time.FixedZone("EST", -5*3600)),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsEqual(tt.t1, tt.t2)
			if result != tt.expected {
				t.Errorf("IsEqual() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestIsSameDay(t *testing.T) {
	tests := []struct {
		name     string
		t1       time.Time
		t2       time.Time
		expected bool
	}{
		{
			name:     "same day, different times",
			t1:       time.Date(2024, 1, 15, 8, 0, 0, 0, time.UTC),
			t2:       time.Date(2024, 1, 15, 20, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "different days",
			t1:       time.Date(2024, 1, 15, 23, 59, 59, 0, time.UTC),
			t2:       time.Date(2024, 1, 16, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsSameDay(tt.t1, tt.t2)
			if result != tt.expected {
				t.Errorf("IsSameDay() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestIsSameWeek(t *testing.T) {
	tests := []struct {
		name     string
		t1       time.Time
		t2       time.Time
		expected bool
	}{
		{
			name:     "same week (Monday to Sunday)",
			t1:       time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC), // Monday
			t2:       time.Date(2024, 1, 21, 12, 0, 0, 0, time.UTC), // Sunday
			expected: true,
		},
		{
			name:     "different weeks",
			t1:       time.Date(2024, 1, 21, 12, 0, 0, 0, time.UTC), // Sunday
			t2:       time.Date(2024, 1, 22, 12, 0, 0, 0, time.UTC), // Monday (next week)
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsSameWeek(tt.t1, tt.t2)
			if result != tt.expected {
				t.Errorf("IsSameWeek() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestIsSameHour(t *testing.T) {
	tests := []struct {
		name     string
		t1       time.Time
		t2       time.Time
		expected bool
	}{
		{
			name:     "same hour, different minutes",
			t1:       time.Date(2024, 1, 15, 12, 10, 0, 0, time.UTC),
			t2:       time.Date(2024, 1, 15, 12, 50, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "different hours",
			t1:       time.Date(2024, 1, 15, 12, 59, 59, 0, time.UTC),
			t2:       time.Date(2024, 1, 15, 13, 0, 0, 0, time.UTC),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsSameHour(tt.t1, tt.t2)
			if result != tt.expected {
				t.Errorf("IsSameHour() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestIsSameMinute(t *testing.T) {
	tests := []struct {
		name     string
		t1       time.Time
		t2       time.Time
		expected bool
	}{
		{
			name:     "same minute, different seconds",
			t1:       time.Date(2024, 1, 15, 12, 30, 10, 0, time.UTC),
			t2:       time.Date(2024, 1, 15, 12, 30, 50, 0, time.UTC),
			expected: true,
		},
		{
			name:     "different minutes",
			t1:       time.Date(2024, 1, 15, 12, 30, 59, 0, time.UTC),
			t2:       time.Date(2024, 1, 15, 12, 31, 0, 0, time.UTC),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsSameMinute(tt.t1, tt.t2)
			if result != tt.expected {
				t.Errorf("IsSameMinute() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestCompareAsc(t *testing.T) {
	tests := []struct {
		name     string
		t1       time.Time
		t2       time.Time
		expected int
	}{
		{
			name:     "t1 before t2",
			t1:       time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC),
			t2:       time.Date(2024, 1, 16, 12, 0, 0, 0, time.UTC),
			expected: -1,
		},
		{
			name:     "t1 after t2",
			t1:       time.Date(2024, 1, 16, 12, 0, 0, 0, time.UTC),
			t2:       time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "t1 equals t2",
			t1:       time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC),
			t2:       time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC),
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CompareAsc(tt.t1, tt.t2)
			if result != tt.expected {
				t.Errorf("CompareAsc() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestCompareDesc(t *testing.T) {
	tests := []struct {
		name     string
		t1       time.Time
		t2       time.Time
		expected int
	}{
		{
			name:     "t1 before t2 (desc should return 1)",
			t1:       time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC),
			t2:       time.Date(2024, 1, 16, 12, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "t1 after t2 (desc should return -1)",
			t1:       time.Date(2024, 1, 16, 12, 0, 0, 0, time.UTC),
			t2:       time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC),
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CompareDesc(tt.t1, tt.t2)
			if result != tt.expected {
				t.Errorf("CompareDesc() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		name     string
		times    []time.Time
		expected time.Time
	}{
		{
			name: "multiple times",
			times: []time.Time{
				time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC),
				time.Date(2024, 1, 10, 12, 0, 0, 0, time.UTC), // earliest
				time.Date(2024, 1, 20, 12, 0, 0, 0, time.UTC),
			},
			expected: time.Date(2024, 1, 10, 12, 0, 0, 0, time.UTC),
		},
		{
			name:     "empty slice",
			times:    []time.Time{},
			expected: time.Time{},
		},
		{
			name: "single time",
			times: []time.Time{
				time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC),
			},
			expected: time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Min(tt.times)
			if !result.Equal(tt.expected) {
				t.Errorf("Min() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name     string
		times    []time.Time
		expected time.Time
	}{
		{
			name: "multiple times",
			times: []time.Time{
				time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC),
				time.Date(2024, 1, 10, 12, 0, 0, 0, time.UTC),
				time.Date(2024, 1, 20, 12, 0, 0, 0, time.UTC), // latest
			},
			expected: time.Date(2024, 1, 20, 12, 0, 0, 0, time.UTC),
		},
		{
			name:     "empty slice",
			times:    []time.Time{},
			expected: time.Time{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Max(tt.times)
			if !result.Equal(tt.expected) {
				t.Errorf("Max() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestClosestTo(t *testing.T) {
	target := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		target   time.Time
		times    []time.Time
		expected time.Time
	}{
		{
			name:   "closest time",
			target: target,
			times: []time.Time{
				time.Date(2024, 1, 10, 12, 0, 0, 0, time.UTC), // 5 days before
				time.Date(2024, 1, 16, 12, 0, 0, 0, time.UTC), // 1 day after (closest)
				time.Date(2024, 1, 20, 12, 0, 0, 0, time.UTC), // 5 days after
			},
			expected: time.Date(2024, 1, 16, 12, 0, 0, 0, time.UTC),
		},
		{
			name:     "empty slice",
			target:   target,
			times:    []time.Time{},
			expected: time.Time{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ClosestTo(tt.target, tt.times)
			if !result.Equal(tt.expected) {
				t.Errorf("ClosestTo() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

// Benchmarks
func BenchmarkCompareAsc(b *testing.B) {
	t1 := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, 1, 16, 12, 0, 0, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		CompareAsc(t1, t2)
	}
}

func BenchmarkMin(b *testing.B) {
	times := []time.Time{
		time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 10, 12, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 20, 12, 0, 0, 0, time.UTC),
	}
	for i := 0; i < b.N; i++ {
		Min(times)
	}
}
