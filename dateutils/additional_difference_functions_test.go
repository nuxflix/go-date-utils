package dateutils

import (
	"math"
	"testing"
	"time"
)

func TestDifferenceInHours(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		t1   time.Time
		t2   time.Time
		want int
	}{
		{
			name: "14 hours difference",
			t1:   time.Date(2024, 1, 1, 20, 0, 0, 0, utc),
			t2:   time.Date(2024, 1, 1, 6, 0, 0, 0, utc),
			want: 14,
		},
		{
			name: "negative difference",
			t1:   time.Date(2024, 1, 1, 6, 0, 0, 0, utc),
			t2:   time.Date(2024, 1, 1, 20, 0, 0, 0, utc),
			want: -14,
		},
		{
			name: "same time",
			t1:   time.Date(2024, 1, 1, 12, 0, 0, 0, utc),
			t2:   time.Date(2024, 1, 1, 12, 0, 0, 0, utc),
			want: 0,
		},
		{
			name: "partial hour (should truncate)",
			t1:   time.Date(2024, 1, 1, 12, 30, 0, 0, utc),
			t2:   time.Date(2024, 1, 1, 11, 0, 0, 0, utc),
			want: 1,
		},
		{
			name: "crossing day boundary",
			t1:   time.Date(2024, 1, 2, 2, 0, 0, 0, utc),
			t2:   time.Date(2024, 1, 1, 22, 0, 0, 0, utc),
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DifferenceInHours(tt.t1, tt.t2)
			if got != tt.want {
				t.Errorf("DifferenceInHours() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifferenceInHoursFloat(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		t1   time.Time
		t2   time.Time
		want float64
	}{
		{
			name: "1.5 hours",
			t1:   time.Date(2024, 1, 1, 13, 30, 0, 0, utc),
			t2:   time.Date(2024, 1, 1, 12, 0, 0, 0, utc),
			want: 1.5,
		},
		{
			name: "0.25 hours (15 minutes)",
			t1:   time.Date(2024, 1, 1, 12, 15, 0, 0, utc),
			t2:   time.Date(2024, 1, 1, 12, 0, 0, 0, utc),
			want: 0.25,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DifferenceInHoursFloat(tt.t1, tt.t2)
			if math.Abs(got-tt.want) > 0.001 {
				t.Errorf("DifferenceInHoursFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifferenceInMinutes(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		t1   time.Time
		t2   time.Time
		want int
	}{
		{
			name: "14 minutes difference",
			t1:   time.Date(2024, 1, 1, 12, 20, 0, 0, utc),
			t2:   time.Date(2024, 1, 1, 12, 6, 0, 0, utc),
			want: 14,
		},
		{
			name: "negative difference",
			t1:   time.Date(2024, 1, 1, 12, 6, 0, 0, utc),
			t2:   time.Date(2024, 1, 1, 12, 20, 0, 0, utc),
			want: -14,
		},
		{
			name: "partial minute (should truncate)",
			t1:   time.Date(2024, 1, 1, 12, 6, 50, 0, utc),
			t2:   time.Date(2024, 1, 1, 12, 6, 10, 0, utc),
			want: 0,
		},
		{
			name: "crossing hour boundary",
			t1:   time.Date(2024, 1, 1, 13, 5, 0, 0, utc),
			t2:   time.Date(2024, 1, 1, 12, 55, 0, 0, utc),
			want: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DifferenceInMinutes(tt.t1, tt.t2)
			if got != tt.want {
				t.Errorf("DifferenceInMinutes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifferenceInSeconds(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		t1   time.Time
		t2   time.Time
		want int
	}{
		{
			name: "12 seconds difference",
			t1:   time.Date(2024, 1, 1, 12, 30, 20, 0, utc),
			t2:   time.Date(2024, 1, 1, 12, 30, 8, 0, utc),
			want: 12,
		},
		{
			name: "negative difference",
			t1:   time.Date(2024, 1, 1, 12, 30, 8, 0, utc),
			t2:   time.Date(2024, 1, 1, 12, 30, 20, 0, utc),
			want: -12,
		},
		{
			name: "crossing minute boundary",
			t1:   time.Date(2024, 1, 1, 12, 31, 5, 0, utc),
			t2:   time.Date(2024, 1, 1, 12, 30, 55, 0, utc),
			want: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DifferenceInSeconds(tt.t1, tt.t2)
			if got != tt.want {
				t.Errorf("DifferenceInSeconds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifferenceInMonths(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		t1   time.Time
		t2   time.Time
		want int
	}{
		{
			name: "12 months (1 year)",
			t1:   time.Date(2024, 7, 2, 18, 0, 0, 0, utc),
			t2:   time.Date(2023, 7, 2, 6, 0, 0, 0, utc),
			want: 12,
		},
		{
			name: "negative difference",
			t1:   time.Date(2023, 7, 2, 6, 0, 0, 0, utc),
			t2:   time.Date(2024, 7, 2, 18, 0, 0, 0, utc),
			want: -12,
		},
		{
			name: "same month and year",
			t1:   time.Date(2024, 7, 15, 12, 0, 0, 0, utc),
			t2:   time.Date(2024, 7, 10, 12, 0, 0, 0, utc),
			want: 0,
		},
		{
			name: "day adjustment needed (Jan 31 to Feb 28)",
			t1:   time.Date(2024, 2, 28, 12, 0, 0, 0, utc),
			t2:   time.Date(2024, 1, 31, 12, 0, 0, 0, utc),
			want: 0, // Not a full month because day 31 doesn't exist in Feb
		},
		{
			name: "full month difference",
			t1:   time.Date(2024, 3, 1, 12, 0, 0, 0, utc),
			t2:   time.Date(2024, 1, 31, 12, 0, 0, 0, utc),
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DifferenceInMonths(tt.t1, tt.t2)
			if got != tt.want {
				t.Errorf("DifferenceInMonths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifferenceInCalendarMonths(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		t1   time.Time
		t2   time.Time
		want int
	}{
		{
			name: "12 calendar months",
			t1:   time.Date(2024, 7, 2, 18, 0, 0, 0, utc),
			t2:   time.Date(2023, 7, 2, 6, 0, 0, 0, utc),
			want: 12,
		},
		{
			name: "negative difference",
			t1:   time.Date(2023, 7, 2, 6, 0, 0, 0, utc),
			t2:   time.Date(2024, 7, 2, 18, 0, 0, 0, utc),
			want: -12,
		},
		{
			name: "calendar months (ignores day)",
			t1:   time.Date(2024, 9, 1, 12, 0, 0, 0, utc),
			t2:   time.Date(2024, 1, 31, 12, 0, 0, 0, utc),
			want: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DifferenceInCalendarMonths(tt.t1, tt.t2)
			if got != tt.want {
				t.Errorf("DifferenceInCalendarMonths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifferenceInYears(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		t1   time.Time
		t2   time.Time
		want int
	}{
		{
			name: "1 full year",
			t1:   time.Date(2024, 2, 11, 12, 0, 0, 0, utc),
			t2:   time.Date(2023, 2, 11, 12, 0, 0, 0, utc),
			want: 1,
		},
		{
			name: "not yet full year (anniversary not passed)",
			t1:   time.Date(2024, 2, 10, 12, 0, 0, 0, utc),
			t2:   time.Date(2023, 2, 11, 12, 0, 0, 0, utc),
			want: 0,
		},
		{
			name: "negative difference",
			t1:   time.Date(2023, 2, 11, 12, 0, 0, 0, utc),
			t2:   time.Date(2024, 2, 11, 12, 0, 0, 0, utc),
			want: -1,
		},
		{
			name: "leap year edge case",
			t1:   time.Date(2025, 2, 28, 12, 0, 0, 0, utc),
			t2:   time.Date(2024, 2, 29, 12, 0, 0, 0, utc),
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DifferenceInYears(tt.t1, tt.t2)
			if got != tt.want {
				t.Errorf("DifferenceInYears() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifferenceInCalendarYears(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		t1   time.Time
		t2   time.Time
		want int
	}{
		{
			name: "1 calendar year",
			t1:   time.Date(2024, 1, 1, 12, 0, 0, 0, utc),
			t2:   time.Date(2023, 12, 31, 12, 0, 0, 0, utc),
			want: 1,
		},
		{
			name: "negative difference",
			t1:   time.Date(2023, 12, 31, 12, 0, 0, 0, utc),
			t2:   time.Date(2024, 1, 1, 12, 0, 0, 0, utc),
			want: -1,
		},
		{
			name: "same year",
			t1:   time.Date(2024, 12, 31, 12, 0, 0, 0, utc),
			t2:   time.Date(2024, 1, 1, 12, 0, 0, 0, utc),
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DifferenceInCalendarYears(tt.t1, tt.t2)
			if got != tt.want {
				t.Errorf("DifferenceInCalendarYears() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifferenceInQuarters(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		t1   time.Time
		t2   time.Time
		want int
	}{
		{
			name: "2 quarters difference",
			t1:   time.Date(2024, 7, 2, 12, 0, 0, 0, utc),
			t2:   time.Date(2024, 1, 2, 12, 0, 0, 0, utc),
			want: 2,
		},
		{
			name: "negative difference",
			t1:   time.Date(2024, 1, 2, 12, 0, 0, 0, utc),
			t2:   time.Date(2024, 7, 2, 12, 0, 0, 0, utc),
			want: -2,
		},
		{
			name: "same quarter",
			t1:   time.Date(2024, 3, 31, 12, 0, 0, 0, utc),
			t2:   time.Date(2024, 1, 1, 12, 0, 0, 0, utc),
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DifferenceInQuarters(tt.t1, tt.t2)
			if got != tt.want {
				t.Errorf("DifferenceInQuarters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifferenceInCalendarQuarters(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		t1   time.Time
		t2   time.Time
		want int
	}{
		{
			name: "4 calendar quarters",
			t1:   time.Date(2024, 7, 2, 18, 0, 0, 0, utc),
			t2:   time.Date(2023, 7, 2, 6, 0, 0, 0, utc),
			want: 4,
		},
		{
			name: "negative difference",
			t1:   time.Date(2023, 7, 2, 6, 0, 0, 0, utc),
			t2:   time.Date(2024, 7, 2, 18, 0, 0, 0, utc),
			want: -4,
		},
		{
			name: "within same year",
			t1:   time.Date(2024, 10, 1, 12, 0, 0, 0, utc), // Q4
			t2:   time.Date(2024, 1, 1, 12, 0, 0, 0, utc),  // Q1
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DifferenceInCalendarQuarters(tt.t1, tt.t2)
			if got != tt.want {
				t.Errorf("DifferenceInCalendarQuarters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbsDifferenceFunctions(t *testing.T) {
	utc := time.UTC
	t1 := time.Date(2024, 1, 1, 12, 0, 0, 0, utc)
	t2 := time.Date(2024, 1, 1, 10, 30, 30, 0, utc)

	// Test absolute functions always return positive values
	if got := AbsDifferenceInHours(t1, t2); got != 1 {
		t.Errorf("AbsDifferenceInHours() = %v, want 1", got)
	}
	if got := AbsDifferenceInHours(t2, t1); got != 1 {
		t.Errorf("AbsDifferenceInHours() = %v, want 1", got)
	}

	if got := AbsDifferenceInMinutes(t1, t2); got != 89 {
		t.Errorf("AbsDifferenceInMinutes() = %v, want 89", got)
	}
	if got := AbsDifferenceInMinutes(t2, t1); got != 89 {
		t.Errorf("AbsDifferenceInMinutes() = %v, want 89", got)
	}

	if got := AbsDifferenceInSeconds(t1, t2); got != 5370 {
		t.Errorf("AbsDifferenceInSeconds() = %v, want 5370", got)
	}
	if got := AbsDifferenceInSeconds(t2, t1); got != 5370 {
		t.Errorf("AbsDifferenceInSeconds() = %v, want 5370", got)
	}
}

// Benchmark tests
func BenchmarkDifferenceInHours(b *testing.B) {
	t1 := time.Date(2024, 1, 1, 20, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, 1, 1, 6, 0, 0, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		DifferenceInHours(t1, t2)
	}
}

func BenchmarkDifferenceInMinutes(b *testing.B) {
	t1 := time.Date(2024, 1, 1, 12, 20, 0, 0, time.UTC)
	t2 := time.Date(2024, 1, 1, 12, 6, 0, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		DifferenceInMinutes(t1, t2)
	}
}

func BenchmarkDifferenceInSeconds(b *testing.B) {
	t1 := time.Date(2024, 1, 1, 12, 30, 20, 0, time.UTC)
	t2 := time.Date(2024, 1, 1, 12, 30, 8, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		DifferenceInSeconds(t1, t2)
	}
}

func BenchmarkDifferenceInMonths(b *testing.B) {
	t1 := time.Date(2024, 7, 2, 18, 0, 0, 0, time.UTC)
	t2 := time.Date(2023, 7, 2, 6, 0, 0, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		DifferenceInMonths(t1, t2)
	}
}

func BenchmarkDifferenceInYears(b *testing.B) {
	t1 := time.Date(2024, 2, 11, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2023, 2, 11, 12, 0, 0, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		DifferenceInYears(t1, t2)
	}
}
