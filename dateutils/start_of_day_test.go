package dateutils

import (
	"testing"
	"time"
)

func TestStartOfDay(t *testing.T) {
	utc := time.UTC
	est, _ := time.LoadLocation("America/New_York")

	tests := []struct {
		name string
		time time.Time
		want time.Time
	}{
		{
			name: "middle of day",
			time: time.Date(2023, 12, 25, 15, 30, 45, 123456789, utc),
			want: time.Date(2023, 12, 25, 0, 0, 0, 0, utc),
		},
		{
			name: "already start of day",
			time: time.Date(2023, 12, 25, 0, 0, 0, 0, utc),
			want: time.Date(2023, 12, 25, 0, 0, 0, 0, utc),
		},
		{
			name: "end of day",
			time: time.Date(2023, 12, 25, 23, 59, 59, 999999999, utc),
			want: time.Date(2023, 12, 25, 0, 0, 0, 0, utc),
		},
		{
			name: "with timezone",
			time: time.Date(2023, 12, 25, 15, 30, 45, 123456789, est),
			want: time.Date(2023, 12, 25, 0, 0, 0, 0, est),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StartOfDay(tt.time)
			if !got.Equal(tt.want) {
				t.Errorf("StartOfDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEndOfDay(t *testing.T) {
	utc := time.UTC
	est, _ := time.LoadLocation("America/New_York")

	tests := []struct {
		name string
		time time.Time
		want time.Time
	}{
		{
			name: "middle of day",
			time: time.Date(2023, 12, 25, 15, 30, 45, 123456789, utc),
			want: time.Date(2023, 12, 25, 23, 59, 59, 999999999, utc),
		},
		{
			name: "start of day",
			time: time.Date(2023, 12, 25, 0, 0, 0, 0, utc),
			want: time.Date(2023, 12, 25, 23, 59, 59, 999999999, utc),
		},
		{
			name: "with timezone",
			time: time.Date(2023, 12, 25, 15, 30, 45, 123456789, est),
			want: time.Date(2023, 12, 25, 23, 59, 59, 999999999, est),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EndOfDay(tt.time)
			if !got.Equal(tt.want) {
				t.Errorf("EndOfDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartOfWeek(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		time time.Time
		want time.Time
	}{
		{
			name: "Monday",
			time: time.Date(2023, 12, 25, 15, 30, 45, 0, utc), // Monday
			want: time.Date(2023, 12, 25, 0, 0, 0, 0, utc),    // Same Monday
		},
		{
			name: "Tuesday",
			time: time.Date(2023, 12, 26, 15, 30, 45, 0, utc), // Tuesday
			want: time.Date(2023, 12, 25, 0, 0, 0, 0, utc),    // Previous Monday
		},
		{
			name: "Sunday",
			time: time.Date(2023, 12, 31, 15, 30, 45, 0, utc), // Sunday
			want: time.Date(2023, 12, 25, 0, 0, 0, 0, utc),    // Previous Monday
		},
		{
			name: "Saturday",
			time: time.Date(2023, 12, 30, 15, 30, 45, 0, utc), // Saturday
			want: time.Date(2023, 12, 25, 0, 0, 0, 0, utc),    // Previous Monday
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StartOfWeek(tt.time)
			if !got.Equal(tt.want) {
				t.Errorf("StartOfWeek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartOfWeekSunday(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		time time.Time
		want time.Time
	}{
		{
			name: "Sunday",
			time: time.Date(2023, 12, 31, 15, 30, 45, 0, utc), // Sunday
			want: time.Date(2023, 12, 31, 0, 0, 0, 0, utc),    // Same Sunday
		},
		{
			name: "Monday",
			time: time.Date(2023, 12, 25, 15, 30, 45, 0, utc), // Monday
			want: time.Date(2023, 12, 24, 0, 0, 0, 0, utc),    // Previous Sunday
		},
		{
			name: "Saturday",
			time: time.Date(2023, 12, 30, 15, 30, 45, 0, utc), // Saturday
			want: time.Date(2023, 12, 24, 0, 0, 0, 0, utc),    // Previous Sunday
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StartOfWeekSunday(tt.time)
			if !got.Equal(tt.want) {
				t.Errorf("StartOfWeekSunday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEndOfWeek(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		time time.Time
		want time.Time
	}{
		{
			name: "Monday",
			time: time.Date(2023, 12, 25, 15, 30, 45, 0, utc),         // Monday
			want: time.Date(2023, 12, 31, 23, 59, 59, 999999999, utc), // Following Sunday
		},
		{
			name: "Friday",
			time: time.Date(2023, 12, 29, 15, 30, 45, 0, utc),         // Friday
			want: time.Date(2023, 12, 31, 23, 59, 59, 999999999, utc), // Following Sunday
		},
		{
			name: "Sunday",
			time: time.Date(2023, 12, 31, 15, 30, 45, 0, utc),         // Sunday
			want: time.Date(2023, 12, 31, 23, 59, 59, 999999999, utc), // Same Sunday
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EndOfWeek(tt.time)
			if !got.Equal(tt.want) {
				t.Errorf("EndOfWeek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartOfMonth(t *testing.T) {
	utc := time.UTC
	est, _ := time.LoadLocation("America/New_York")

	tests := []struct {
		name string
		time time.Time
		want time.Time
	}{
		{
			name: "middle of month",
			time: time.Date(2023, 12, 25, 15, 30, 45, 123456789, utc),
			want: time.Date(2023, 12, 1, 0, 0, 0, 0, utc),
		},
		{
			name: "first day of month",
			time: time.Date(2023, 12, 1, 15, 30, 45, 123456789, utc),
			want: time.Date(2023, 12, 1, 0, 0, 0, 0, utc),
		},
		{
			name: "last day of month",
			time: time.Date(2023, 12, 31, 23, 59, 59, 999999999, utc),
			want: time.Date(2023, 12, 1, 0, 0, 0, 0, utc),
		},
		{
			name: "with timezone",
			time: time.Date(2023, 12, 25, 15, 30, 45, 123456789, est),
			want: time.Date(2023, 12, 1, 0, 0, 0, 0, est),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StartOfMonth(tt.time)
			if !got.Equal(tt.want) {
				t.Errorf("StartOfMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEndOfMonth(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		time time.Time
		want time.Time
	}{
		{
			name: "December (31 days)",
			time: time.Date(2023, 12, 25, 15, 30, 45, 123456789, utc),
			want: time.Date(2023, 12, 31, 23, 59, 59, 999999999, utc),
		},
		{
			name: "February non-leap year (28 days)",
			time: time.Date(2023, 2, 15, 15, 30, 45, 123456789, utc),
			want: time.Date(2023, 2, 28, 23, 59, 59, 999999999, utc),
		},
		{
			name: "February leap year (29 days)",
			time: time.Date(2024, 2, 15, 15, 30, 45, 123456789, utc),
			want: time.Date(2024, 2, 29, 23, 59, 59, 999999999, utc),
		},
		{
			name: "April (30 days)",
			time: time.Date(2023, 4, 15, 15, 30, 45, 123456789, utc),
			want: time.Date(2023, 4, 30, 23, 59, 59, 999999999, utc),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EndOfMonth(tt.time)
			if !got.Equal(tt.want) {
				t.Errorf("EndOfMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartOfYear(t *testing.T) {
	utc := time.UTC
	est, _ := time.LoadLocation("America/New_York")

	tests := []struct {
		name string
		time time.Time
		want time.Time
	}{
		{
			name: "middle of year",
			time: time.Date(2023, 6, 15, 15, 30, 45, 123456789, utc),
			want: time.Date(2023, 1, 1, 0, 0, 0, 0, utc),
		},
		{
			name: "end of year",
			time: time.Date(2023, 12, 31, 23, 59, 59, 999999999, utc),
			want: time.Date(2023, 1, 1, 0, 0, 0, 0, utc),
		},
		{
			name: "with timezone",
			time: time.Date(2023, 6, 15, 15, 30, 45, 123456789, est),
			want: time.Date(2023, 1, 1, 0, 0, 0, 0, est),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StartOfYear(tt.time)
			if !got.Equal(tt.want) {
				t.Errorf("StartOfYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEndOfYear(t *testing.T) {
	utc := time.UTC
	est, _ := time.LoadLocation("America/New_York")

	tests := []struct {
		name string
		time time.Time
		want time.Time
	}{
		{
			name: "middle of year",
			time: time.Date(2023, 6, 15, 15, 30, 45, 123456789, utc),
			want: time.Date(2023, 12, 31, 23, 59, 59, 999999999, utc),
		},
		{
			name: "start of year",
			time: time.Date(2023, 1, 1, 0, 0, 0, 0, utc),
			want: time.Date(2023, 12, 31, 23, 59, 59, 999999999, utc),
		},
		{
			name: "with timezone",
			time: time.Date(2023, 6, 15, 15, 30, 45, 123456789, est),
			want: time.Date(2023, 12, 31, 23, 59, 59, 999999999, est),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EndOfYear(tt.time)
			if !got.Equal(tt.want) {
				t.Errorf("EndOfYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartOfHour(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		time time.Time
		want time.Time
	}{
		{
			name: "middle of hour",
			time: time.Date(2023, 12, 25, 15, 30, 45, 123456789, utc),
			want: time.Date(2023, 12, 25, 15, 0, 0, 0, utc),
		},
		{
			name: "start of hour",
			time: time.Date(2023, 12, 25, 15, 0, 0, 0, utc),
			want: time.Date(2023, 12, 25, 15, 0, 0, 0, utc),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StartOfHour(tt.time)
			if !got.Equal(tt.want) {
				t.Errorf("StartOfHour() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEndOfHour(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		time time.Time
		want time.Time
	}{
		{
			name: "middle of hour",
			time: time.Date(2023, 12, 25, 15, 30, 45, 123456789, utc),
			want: time.Date(2023, 12, 25, 15, 59, 59, 999999999, utc),
		},
		{
			name: "start of hour",
			time: time.Date(2023, 12, 25, 15, 0, 0, 0, utc),
			want: time.Date(2023, 12, 25, 15, 59, 59, 999999999, utc),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EndOfHour(tt.time)
			if !got.Equal(tt.want) {
				t.Errorf("EndOfHour() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartOfMinute(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		time time.Time
		want time.Time
	}{
		{
			name: "middle of minute",
			time: time.Date(2023, 12, 25, 15, 30, 45, 123456789, utc),
			want: time.Date(2023, 12, 25, 15, 30, 0, 0, utc),
		},
		{
			name: "start of minute",
			time: time.Date(2023, 12, 25, 15, 30, 0, 0, utc),
			want: time.Date(2023, 12, 25, 15, 30, 0, 0, utc),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StartOfMinute(tt.time)
			if !got.Equal(tt.want) {
				t.Errorf("StartOfMinute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEndOfMinute(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		time time.Time
		want time.Time
	}{
		{
			name: "middle of minute",
			time: time.Date(2023, 12, 25, 15, 30, 45, 123456789, utc),
			want: time.Date(2023, 12, 25, 15, 30, 59, 999999999, utc),
		},
		{
			name: "start of minute",
			time: time.Date(2023, 12, 25, 15, 30, 0, 0, utc),
			want: time.Date(2023, 12, 25, 15, 30, 59, 999999999, utc),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EndOfMinute(tt.time)
			if !got.Equal(tt.want) {
				t.Errorf("EndOfMinute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkStartOfDay(b *testing.B) {
	testTime := time.Date(2023, 12, 25, 15, 30, 45, 123456789, time.UTC)

	for i := 0; i < b.N; i++ {
		_ = StartOfDay(testTime)
	}
}

func BenchmarkStartOfWeek(b *testing.B) {
	testTime := time.Date(2023, 12, 25, 15, 30, 45, 123456789, time.UTC)

	for i := 0; i < b.N; i++ {
		_ = StartOfWeek(testTime)
	}
}

func BenchmarkStartOfMonth(b *testing.B) {
	testTime := time.Date(2023, 12, 25, 15, 30, 45, 123456789, time.UTC)

	for i := 0; i < b.N; i++ {
		_ = StartOfMonth(testTime)
	}
}
