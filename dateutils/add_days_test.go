package dateutils

import (
	"testing"
	"time"
)

func TestAddDays(t *testing.T) {
	utc := time.UTC
	baseTime := time.Date(2023, 12, 25, 15, 30, 45, 123, utc)

	tests := []struct {
		name string
		time time.Time
		days int
		want time.Time
	}{
		{
			name: "add positive days",
			time: baseTime,
			days: 5,
			want: time.Date(2023, 12, 30, 15, 30, 45, 123, utc),
		},
		{
			name: "add negative days (subtract)",
			time: baseTime,
			days: -5,
			want: time.Date(2023, 12, 20, 15, 30, 45, 123, utc),
		},
		{
			name: "add zero days",
			time: baseTime,
			days: 0,
			want: baseTime,
		},
		{
			name: "add days crossing month boundary",
			time: time.Date(2023, 1, 30, 12, 0, 0, 0, utc),
			days: 5,
			want: time.Date(2023, 2, 4, 12, 0, 0, 0, utc),
		},
		{
			name: "add days crossing year boundary",
			time: time.Date(2023, 12, 30, 12, 0, 0, 0, utc),
			days: 5,
			want: time.Date(2024, 1, 4, 12, 0, 0, 0, utc),
		},
		{
			name: "subtract days crossing month boundary",
			time: time.Date(2023, 3, 2, 12, 0, 0, 0, utc),
			days: -5,
			want: time.Date(2023, 2, 25, 12, 0, 0, 0, utc),
		},
		{
			name: "leap year calculation",
			time: time.Date(2024, 2, 28, 12, 0, 0, 0, utc), // 2024 is leap year
			days: 1,
			want: time.Date(2024, 2, 29, 12, 0, 0, 0, utc),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AddDays(tt.time, tt.days)
			if !got.Equal(tt.want) {
				t.Errorf("AddDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddBusinessDays(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name         string
		time         time.Time
		businessDays int
		want         time.Time
	}{
		{
			name:         "add business days from Monday",
			time:         time.Date(2023, 12, 25, 12, 0, 0, 0, utc), // Monday
			businessDays: 5,
			want:         time.Date(2024, 1, 1, 12, 0, 0, 0, utc), // Next Monday
		},
		{
			name:         "add business days from Friday",
			time:         time.Date(2023, 12, 29, 12, 0, 0, 0, utc), // Friday
			businessDays: 3,
			want:         time.Date(2024, 1, 3, 12, 0, 0, 0, utc), // Wednesday (skips weekend)
		},
		{
			name:         "add business days from Saturday (should start from next Monday)",
			time:         time.Date(2023, 12, 30, 12, 0, 0, 0, utc), // Saturday
			businessDays: 1,
			want:         time.Date(2024, 1, 2, 12, 0, 0, 0, utc), // Tuesday (skips Sat, Sun, Mon)
		},
		{
			name:         "subtract business days",
			time:         time.Date(2023, 12, 29, 12, 0, 0, 0, utc), // Friday
			businessDays: -3,
			want:         time.Date(2023, 12, 26, 12, 0, 0, 0, utc), // Tuesday
		},
		{
			name:         "add zero business days",
			time:         time.Date(2023, 12, 25, 12, 0, 0, 0, utc), // Monday
			businessDays: 0,
			want:         time.Date(2023, 12, 25, 12, 0, 0, 0, utc), // Same Monday
		},
		{
			name:         "add business days from Sunday",
			time:         time.Date(2023, 12, 31, 12, 0, 0, 0, utc), // Sunday
			businessDays: 1,
			want:         time.Date(2024, 1, 2, 12, 0, 0, 0, utc), // Tuesday (skips Sun, Mon)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AddBusinessDays(tt.time, tt.businessDays)
			if !got.Equal(tt.want) {
				t.Errorf("AddBusinessDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddWeeks(t *testing.T) {
	utc := time.UTC
	baseTime := time.Date(2023, 12, 25, 15, 30, 45, 123, utc)

	tests := []struct {
		name  string
		time  time.Time
		weeks int
		want  time.Time
	}{
		{
			name:  "add positive weeks",
			time:  baseTime,
			weeks: 2,
			want:  time.Date(2024, 1, 8, 15, 30, 45, 123, utc),
		},
		{
			name:  "add negative weeks (subtract)",
			time:  baseTime,
			weeks: -1,
			want:  time.Date(2023, 12, 18, 15, 30, 45, 123, utc),
		},
		{
			name:  "add zero weeks",
			time:  baseTime,
			weeks: 0,
			want:  baseTime,
		},
		{
			name:  "add weeks crossing year boundary",
			time:  time.Date(2023, 12, 28, 12, 0, 0, 0, utc),
			weeks: 1,
			want:  time.Date(2024, 1, 4, 12, 0, 0, 0, utc),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AddWeeks(tt.time, tt.weeks)
			if !got.Equal(tt.want) {
				t.Errorf("AddWeeks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddDaysWithTimezone(t *testing.T) {
	utc := time.UTC
	est, _ := time.LoadLocation("America/New_York")
	pst, _ := time.LoadLocation("America/Los_Angeles")

	baseTime := time.Date(2023, 12, 25, 15, 30, 45, 0, utc)

	tests := []struct {
		name     string
		time     time.Time
		days     int
		timezone *time.Location
		want     time.Time
	}{
		{
			name:     "add days and convert to EST",
			time:     baseTime,
			days:     5,
			timezone: est,
			want:     time.Date(2023, 12, 30, 15, 30, 45, 0, utc).In(est),
		},
		{
			name:     "add days with nil timezone (preserve original)",
			time:     baseTime,
			days:     3,
			timezone: nil,
			want:     time.Date(2023, 12, 28, 15, 30, 45, 0, utc),
		},
		{
			name:     "add days from EST time and convert to PST",
			time:     time.Date(2023, 12, 25, 10, 30, 45, 0, est),
			days:     7,
			timezone: pst,
			want:     time.Date(2024, 1, 1, 10, 30, 45, 0, est).In(pst),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AddDaysWithTimezone(tt.time, tt.days, tt.timezone)
			if !got.Equal(tt.want) {
				t.Errorf("AddDaysWithTimezone() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test immutability
func TestAddDaysImmutability(t *testing.T) {
	original := time.Date(2023, 12, 25, 15, 30, 45, 123, time.UTC)
	originalCopy := original

	result := AddDays(original, 5)

	// Original should be unchanged
	if !original.Equal(originalCopy) {
		t.Errorf("AddDays() modified original time: got %v, want %v", original, originalCopy)
	}

	// Result should be different
	if result.Equal(original) {
		t.Errorf("AddDays() did not create new time instance")
	}
}

func BenchmarkAddDays(b *testing.B) {
	baseTime := time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC)

	for i := 0; i < b.N; i++ {
		_ = AddDays(baseTime, 5)
	}
}

func BenchmarkAddBusinessDays(b *testing.B) {
	baseTime := time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC) // Monday

	for i := 0; i < b.N; i++ {
		_ = AddBusinessDays(baseTime, 5)
	}
}

func BenchmarkAddWeeks(b *testing.B) {
	baseTime := time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC)

	for i := 0; i < b.N; i++ {
		_ = AddWeeks(baseTime, 2)
	}
}
