package dateutils

import (
	"math"
	"testing"
	"time"
)

func TestDifferenceInDays(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		t1   time.Time
		t2   time.Time
		want int
	}{
		{
			name: "same time",
			t1:   time.Date(2023, 12, 25, 12, 0, 0, 0, utc),
			t2:   time.Date(2023, 12, 25, 12, 0, 0, 0, utc),
			want: 0,
		},
		{
			name: "one day difference (positive)",
			t1:   time.Date(2023, 12, 26, 12, 0, 0, 0, utc),
			t2:   time.Date(2023, 12, 25, 12, 0, 0, 0, utc),
			want: 1,
		},
		{
			name: "one day difference (negative)",
			t1:   time.Date(2023, 12, 25, 12, 0, 0, 0, utc),
			t2:   time.Date(2023, 12, 26, 12, 0, 0, 0, utc),
			want: -1,
		},
		{
			name: "partial day (less than 24 hours)",
			t1:   time.Date(2023, 12, 25, 23, 59, 59, 0, utc),
			t2:   time.Date(2023, 12, 25, 0, 0, 0, 0, utc),
			want: 0, // Less than 24 hours
		},
		{
			name: "exactly 24 hours",
			t1:   time.Date(2023, 12, 26, 0, 0, 0, 0, utc),
			t2:   time.Date(2023, 12, 25, 0, 0, 0, 0, utc),
			want: 1,
		},
		{
			name: "multiple days",
			t1:   time.Date(2023, 12, 30, 12, 0, 0, 0, utc),
			t2:   time.Date(2023, 12, 25, 12, 0, 0, 0, utc),
			want: 5,
		},
		{
			name: "cross month boundary",
			t1:   time.Date(2024, 1, 2, 12, 0, 0, 0, utc),
			t2:   time.Date(2023, 12, 30, 12, 0, 0, 0, utc),
			want: 3,
		},
		{
			name: "cross year boundary",
			t1:   time.Date(2024, 1, 5, 12, 0, 0, 0, utc),
			t2:   time.Date(2023, 12, 25, 12, 0, 0, 0, utc),
			want: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DifferenceInDays(tt.t1, tt.t2)
			if got != tt.want {
				t.Errorf("DifferenceInDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifferenceInDaysFloat(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		t1   time.Time
		t2   time.Time
		want float64
	}{
		{
			name: "half day difference",
			t1:   time.Date(2023, 12, 25, 12, 0, 0, 0, utc),
			t2:   time.Date(2023, 12, 25, 0, 0, 0, 0, utc),
			want: 0.5,
		},
		{
			name: "quarter day difference",
			t1:   time.Date(2023, 12, 25, 6, 0, 0, 0, utc),
			t2:   time.Date(2023, 12, 25, 0, 0, 0, 0, utc),
			want: 0.25,
		},
		{
			name: "one and half days",
			t1:   time.Date(2023, 12, 26, 12, 0, 0, 0, utc),
			t2:   time.Date(2023, 12, 25, 0, 0, 0, 0, utc),
			want: 1.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DifferenceInDaysFloat(tt.t1, tt.t2)
			if math.Abs(got-tt.want) > 0.001 { // Allow small floating point differences
				t.Errorf("DifferenceInDaysFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifferenceInCalendarDays(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		t1   time.Time
		t2   time.Time
		want int
	}{
		{
			name: "same calendar day",
			t1:   time.Date(2023, 12, 25, 23, 59, 59, 0, utc),
			t2:   time.Date(2023, 12, 25, 0, 0, 0, 0, utc),
			want: 0,
		},
		{
			name: "different calendar days (just over midnight)",
			t1:   time.Date(2023, 12, 26, 0, 0, 1, 0, utc),
			t2:   time.Date(2023, 12, 25, 23, 59, 59, 0, utc),
			want: 1,
		},
		{
			name: "multiple calendar days",
			t1:   time.Date(2023, 12, 28, 5, 30, 0, 0, utc),
			t2:   time.Date(2023, 12, 25, 18, 45, 0, 0, utc),
			want: 3,
		},
		{
			name: "negative difference",
			t1:   time.Date(2023, 12, 25, 18, 45, 0, 0, utc),
			t2:   time.Date(2023, 12, 28, 5, 30, 0, 0, utc),
			want: -3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DifferenceInCalendarDays(tt.t1, tt.t2)
			if got != tt.want {
				t.Errorf("DifferenceInCalendarDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifferenceInBusinessDays(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		t1   time.Time
		t2   time.Time
		want int
	}{
		{
			name: "same day",
			t1:   time.Date(2023, 12, 25, 12, 0, 0, 0, utc), // Monday
			t2:   time.Date(2023, 12, 25, 8, 0, 0, 0, utc),  // Same Monday
			want: 0,
		},
		{
			name: "Monday to Friday",
			t1:   time.Date(2023, 12, 29, 12, 0, 0, 0, utc), // Friday
			t2:   time.Date(2023, 12, 25, 12, 0, 0, 0, utc), // Monday
			want: 4,                                         // Mon, Tue, Wed, Thu (Fri not counted as it's the same as t1)
		},
		{
			name: "Friday to Monday (skip weekend)",
			t1:   time.Date(2024, 1, 1, 12, 0, 0, 0, utc),   // Monday
			t2:   time.Date(2023, 12, 29, 12, 0, 0, 0, utc), // Friday
			want: 1,                                         // Only Monday is counted (weekend skipped)
		},
		{
			name: "Including weekend",
			t1:   time.Date(2024, 1, 2, 12, 0, 0, 0, utc),   // Tuesday
			t2:   time.Date(2023, 12, 29, 12, 0, 0, 0, utc), // Friday
			want: 2,                                         // Monday and Tuesday
		},
		{
			name: "negative difference",
			t1:   time.Date(2023, 12, 25, 12, 0, 0, 0, utc), // Monday
			t2:   time.Date(2023, 12, 29, 12, 0, 0, 0, utc), // Friday
			want: -4,
		},
		{
			name: "Saturday to Sunday (no business days)",
			t1:   time.Date(2023, 12, 31, 12, 0, 0, 0, utc), // Sunday
			t2:   time.Date(2023, 12, 30, 12, 0, 0, 0, utc), // Saturday
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DifferenceInBusinessDays(tt.t1, tt.t2)
			if got != tt.want {
				t.Errorf("DifferenceInBusinessDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifferenceInWeeks(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		t1   time.Time
		t2   time.Time
		want int
	}{
		{
			name: "exactly one week",
			t1:   time.Date(2024, 1, 1, 12, 0, 0, 0, utc),
			t2:   time.Date(2023, 12, 25, 12, 0, 0, 0, utc),
			want: 1,
		},
		{
			name: "two weeks",
			t1:   time.Date(2024, 1, 8, 12, 0, 0, 0, utc),
			t2:   time.Date(2023, 12, 25, 12, 0, 0, 0, utc),
			want: 2,
		},
		{
			name: "less than a week",
			t1:   time.Date(2023, 12, 30, 12, 0, 0, 0, utc),
			t2:   time.Date(2023, 12, 25, 12, 0, 0, 0, utc),
			want: 0,
		},
		{
			name: "negative weeks",
			t1:   time.Date(2023, 12, 25, 12, 0, 0, 0, utc),
			t2:   time.Date(2024, 1, 8, 12, 0, 0, 0, utc),
			want: -2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DifferenceInWeeks(tt.t1, tt.t2)
			if got != tt.want {
				t.Errorf("DifferenceInWeeks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifferenceInWeeksFloat(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		t1   time.Time
		t2   time.Time
		want float64
	}{
		{
			name: "one and a half weeks",
			t1:   time.Date(2024, 1, 4, 12, 0, 0, 0, utc), // 10.5 days from Dec 25
			t2:   time.Date(2023, 12, 25, 0, 0, 0, 0, utc),
			want: 1.5,
		},
		{
			name: "half a week",
			t1:   time.Date(2023, 12, 28, 12, 0, 0, 0, utc), // 3.5 days from Dec 25
			t2:   time.Date(2023, 12, 25, 0, 0, 0, 0, utc),
			want: 0.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DifferenceInWeeksFloat(tt.t1, tt.t2)
			if math.Abs(got-tt.want) > 0.001 {
				t.Errorf("DifferenceInWeeksFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbsDifferenceInDays(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		t1   time.Time
		t2   time.Time
		want int
	}{
		{
			name: "positive difference",
			t1:   time.Date(2023, 12, 30, 12, 0, 0, 0, utc),
			t2:   time.Date(2023, 12, 25, 12, 0, 0, 0, utc),
			want: 5,
		},
		{
			name: "negative difference (should become positive)",
			t1:   time.Date(2023, 12, 25, 12, 0, 0, 0, utc),
			t2:   time.Date(2023, 12, 30, 12, 0, 0, 0, utc),
			want: 5,
		},
		{
			name: "same time",
			t1:   time.Date(2023, 12, 25, 12, 0, 0, 0, utc),
			t2:   time.Date(2023, 12, 25, 12, 0, 0, 0, utc),
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AbsDifferenceInDays(tt.t1, tt.t2)
			if got != tt.want {
				t.Errorf("AbsDifferenceInDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkDifferenceInDays(b *testing.B) {
	t1 := time.Date(2023, 12, 30, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2023, 12, 25, 12, 0, 0, 0, time.UTC)

	for i := 0; i < b.N; i++ {
		_ = DifferenceInDays(t1, t2)
	}
}

func BenchmarkDifferenceInBusinessDays(b *testing.B) {
	t1 := time.Date(2023, 12, 29, 12, 0, 0, 0, time.UTC) // Friday
	t2 := time.Date(2023, 12, 25, 12, 0, 0, 0, time.UTC) // Monday

	for i := 0; i < b.N; i++ {
		_ = DifferenceInBusinessDays(t1, t2)
	}
}
