package dateutils

import (
	"testing"
	"time"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		time time.Time
		want bool
	}{
		{
			name: "valid time",
			time: time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC),
			want: true,
		},
		{
			name: "zero time",
			time: time.Time{},
			want: false,
		},
		{
			name: "current time",
			time: time.Now(),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValid(tt.time)
			if got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLeapYear(t *testing.T) {
	tests := []struct {
		name string
		year int
		want bool
	}{
		{
			name: "leap year divisible by 4",
			year: 2024,
			want: true,
		},
		{
			name: "not leap year",
			year: 2023,
			want: false,
		},
		{
			name: "century year not divisible by 400",
			year: 1900,
			want: false,
		},
		{
			name: "century year divisible by 400",
			year: 2000,
			want: true,
		},
		{
			name: "leap year 2020",
			year: 2020,
			want: true,
		},
		{
			name: "not leap year 2021",
			year: 2021,
			want: false,
		},
		{
			name: "not leap year 2022",
			year: 2022,
			want: false,
		},
		{
			name: "not leap year 2023",
			year: 2023,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsLeapYear(tt.year)
			if got != tt.want {
				t.Errorf("IsLeapYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsWeekend(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		time time.Time
		want bool
	}{
		{
			name: "Saturday",
			time: time.Date(2023, 12, 30, 15, 30, 0, 0, utc), // Saturday
			want: true,
		},
		{
			name: "Sunday",
			time: time.Date(2023, 12, 31, 15, 30, 0, 0, utc), // Sunday
			want: true,
		},
		{
			name: "Monday",
			time: time.Date(2023, 12, 25, 15, 30, 0, 0, utc), // Monday
			want: false,
		},
		{
			name: "Friday",
			time: time.Date(2023, 12, 29, 15, 30, 0, 0, utc), // Friday
			want: false,
		},
		{
			name: "Wednesday",
			time: time.Date(2023, 12, 27, 15, 30, 0, 0, utc), // Wednesday
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsWeekend(tt.time)
			if got != tt.want {
				t.Errorf("IsWeekend() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsWeekday(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		time time.Time
		want bool
	}{
		{
			name: "Monday",
			time: time.Date(2023, 12, 25, 15, 30, 0, 0, utc), // Monday
			want: true,
		},
		{
			name: "Friday",
			time: time.Date(2023, 12, 29, 15, 30, 0, 0, utc), // Friday
			want: true,
		},
		{
			name: "Saturday",
			time: time.Date(2023, 12, 30, 15, 30, 0, 0, utc), // Saturday
			want: false,
		},
		{
			name: "Sunday",
			time: time.Date(2023, 12, 31, 15, 30, 0, 0, utc), // Sunday
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsWeekday(tt.time)
			if got != tt.want {
				t.Errorf("IsWeekday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsWithinInterval(t *testing.T) {
	utc := time.UTC

	start := time.Date(2023, 12, 25, 10, 0, 0, 0, utc)
	end := time.Date(2023, 12, 25, 18, 0, 0, 0, utc)

	tests := []struct {
		name  string
		time  time.Time
		start time.Time
		end   time.Time
		want  bool
	}{
		{
			name:  "within interval",
			time:  time.Date(2023, 12, 25, 15, 0, 0, 0, utc),
			start: start,
			end:   end,
			want:  true,
		},
		{
			name:  "equal to start",
			time:  start,
			start: start,
			end:   end,
			want:  true,
		},
		{
			name:  "equal to end",
			time:  end,
			start: start,
			end:   end,
			want:  true,
		},
		{
			name:  "before interval",
			time:  time.Date(2023, 12, 25, 9, 0, 0, 0, utc),
			start: start,
			end:   end,
			want:  false,
		},
		{
			name:  "after interval",
			time:  time.Date(2023, 12, 25, 19, 0, 0, 0, utc),
			start: start,
			end:   end,
			want:  false,
		},
		{
			name:  "invalid interval (start after end)",
			time:  time.Date(2023, 12, 25, 15, 0, 0, 0, utc),
			start: end,
			end:   start,
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsWithinInterval(tt.time, tt.start, tt.end)
			if got != tt.want {
				t.Errorf("IsWithinInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFirstDayOfMonth(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		time time.Time
		want bool
	}{
		{
			name: "first day of month",
			time: time.Date(2023, 12, 1, 15, 30, 0, 0, utc),
			want: true,
		},
		{
			name: "not first day of month",
			time: time.Date(2023, 12, 15, 15, 30, 0, 0, utc),
			want: false,
		},
		{
			name: "first day of January",
			time: time.Date(2023, 1, 1, 0, 0, 0, 0, utc),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsFirstDayOfMonth(tt.time)
			if got != tt.want {
				t.Errorf("IsFirstDayOfMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLastDayOfMonth(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		time time.Time
		want bool
	}{
		{
			name: "last day of December",
			time: time.Date(2023, 12, 31, 15, 30, 0, 0, utc),
			want: true,
		},
		{
			name: "last day of February (non-leap year)",
			time: time.Date(2023, 2, 28, 15, 30, 0, 0, utc),
			want: true,
		},
		{
			name: "last day of February (leap year)",
			time: time.Date(2024, 2, 29, 15, 30, 0, 0, utc),
			want: true,
		},
		{
			name: "not last day of month",
			time: time.Date(2023, 12, 15, 15, 30, 0, 0, utc),
			want: false,
		},
		{
			name: "February 28 in leap year (not last day)",
			time: time.Date(2024, 2, 28, 15, 30, 0, 0, utc),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsLastDayOfMonth(tt.time)
			if got != tt.want {
				t.Errorf("IsLastDayOfMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFirstDayOfYear(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		time time.Time
		want bool
	}{
		{
			name: "January 1st",
			time: time.Date(2023, 1, 1, 15, 30, 0, 0, utc),
			want: true,
		},
		{
			name: "not January 1st",
			time: time.Date(2023, 1, 2, 15, 30, 0, 0, utc),
			want: false,
		},
		{
			name: "December 31st",
			time: time.Date(2023, 12, 31, 15, 30, 0, 0, utc),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsFirstDayOfYear(tt.time)
			if got != tt.want {
				t.Errorf("IsFirstDayOfYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLastDayOfYear(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		time time.Time
		want bool
	}{
		{
			name: "December 31st",
			time: time.Date(2023, 12, 31, 15, 30, 0, 0, utc),
			want: true,
		},
		{
			name: "not December 31st",
			time: time.Date(2023, 12, 30, 15, 30, 0, 0, utc),
			want: false,
		},
		{
			name: "January 1st",
			time: time.Date(2023, 1, 1, 15, 30, 0, 0, utc),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsLastDayOfYear(tt.time)
			if got != tt.want {
				t.Errorf("IsLastDayOfYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSameDate(t *testing.T) {
	utc := time.UTC
	est, _ := time.LoadLocation("America/New_York")

	tests := []struct {
		name string
		t1   time.Time
		t2   time.Time
		want bool
	}{
		{
			name: "same date and time",
			t1:   time.Date(2023, 12, 25, 15, 30, 0, 0, utc),
			t2:   time.Date(2023, 12, 25, 15, 30, 0, 0, utc),
			want: true,
		},
		{
			name: "same date, different time",
			t1:   time.Date(2023, 12, 25, 15, 30, 0, 0, utc),
			t2:   time.Date(2023, 12, 25, 8, 45, 0, 0, utc),
			want: true,
		},
		{
			name: "different date",
			t1:   time.Date(2023, 12, 25, 15, 30, 0, 0, utc),
			t2:   time.Date(2023, 12, 26, 15, 30, 0, 0, utc),
			want: false,
		},
		{
			name: "same date, different timezone",
			t1:   time.Date(2023, 12, 25, 15, 30, 0, 0, utc),
			t2:   time.Date(2023, 12, 25, 10, 30, 0, 0, est),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsSameDate(tt.t1, tt.t2)
			if got != tt.want {
				t.Errorf("IsSameDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSameMonth(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		t1   time.Time
		t2   time.Time
		want bool
	}{
		{
			name: "same month and year",
			t1:   time.Date(2023, 12, 25, 15, 30, 0, 0, utc),
			t2:   time.Date(2023, 12, 15, 8, 45, 0, 0, utc),
			want: true,
		},
		{
			name: "different month, same year",
			t1:   time.Date(2023, 12, 25, 15, 30, 0, 0, utc),
			t2:   time.Date(2023, 11, 25, 15, 30, 0, 0, utc),
			want: false,
		},
		{
			name: "same month, different year",
			t1:   time.Date(2023, 12, 25, 15, 30, 0, 0, utc),
			t2:   time.Date(2024, 12, 25, 15, 30, 0, 0, utc),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsSameMonth(tt.t1, tt.t2)
			if got != tt.want {
				t.Errorf("IsSameMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSameYear(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		t1   time.Time
		t2   time.Time
		want bool
	}{
		{
			name: "same year",
			t1:   time.Date(2023, 12, 25, 15, 30, 0, 0, utc),
			t2:   time.Date(2023, 6, 15, 8, 45, 0, 0, utc),
			want: true,
		},
		{
			name: "different year",
			t1:   time.Date(2023, 12, 25, 15, 30, 0, 0, utc),
			t2:   time.Date(2024, 12, 25, 15, 30, 0, 0, utc),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsSameYear(tt.t1, tt.t2)
			if got != tt.want {
				t.Errorf("IsSameYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsLeapYear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsLeapYear(2024)
	}
}

func BenchmarkIsWeekend(b *testing.B) {
	testTime := time.Date(2023, 12, 30, 15, 30, 0, 0, time.UTC) // Saturday

	for i := 0; i < b.N; i++ {
		_ = IsWeekend(testTime)
	}
}
