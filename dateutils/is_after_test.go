package dateutils

import (
	"testing"
	"time"
)

func TestIsAfter(t *testing.T) {
	utc := time.UTC
	t1 := time.Date(2023, 12, 25, 10, 30, 0, 0, utc)
	t2 := time.Date(2023, 12, 25, 10, 31, 0, 0, utc)
	t3 := time.Date(2023, 12, 24, 10, 30, 0, 0, utc)

	tests := []struct {
		name string
		t1   time.Time
		t2   time.Time
		want bool
	}{
		{
			name: "t2 after t1 (same day, different time)",
			t1:   t2,
			t2:   t1,
			want: true,
		},
		{
			name: "t1 after t3 (different day)",
			t1:   t1,
			t2:   t3,
			want: true,
		},
		{
			name: "t1 before t2",
			t1:   t1,
			t2:   t2,
			want: false,
		},
		{
			name: "same time",
			t1:   t1,
			t2:   t1,
			want: false,
		},
		{
			name: "with nanoseconds difference",
			t1:   time.Date(2023, 12, 25, 10, 30, 0, 124, utc),
			t2:   time.Date(2023, 12, 25, 10, 30, 0, 123, utc),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsAfter(tt.t1, tt.t2)
			if got != tt.want {
				t.Errorf("IsAfter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAfterOrEqual(t *testing.T) {
	utc := time.UTC
	t1 := time.Date(2023, 12, 25, 10, 30, 0, 0, utc)
	t2 := time.Date(2023, 12, 25, 10, 31, 0, 0, utc)

	tests := []struct {
		name string
		t1   time.Time
		t2   time.Time
		want bool
	}{
		{
			name: "t2 after t1",
			t1:   t2,
			t2:   t1,
			want: true,
		},
		{
			name: "t1 equal t2",
			t1:   t1,
			t2:   t1,
			want: true,
		},
		{
			name: "t1 before t2",
			t1:   t1,
			t2:   t2,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsAfterOrEqual(tt.t1, tt.t2)
			if got != tt.want {
				t.Errorf("IsAfterOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAfterDate(t *testing.T) {
	utc := time.UTC
	est, _ := time.LoadLocation("America/New_York")

	tests := []struct {
		name string
		t1   time.Time
		t2   time.Time
		want bool
	}{
		{
			name: "same date, different time - should be false",
			t1:   time.Date(2023, 12, 25, 15, 30, 0, 0, utc),
			t2:   time.Date(2023, 12, 25, 10, 30, 0, 0, utc),
			want: false,
		},
		{
			name: "different date, t1 after t2",
			t1:   time.Date(2023, 12, 25, 0, 0, 1, 0, utc),
			t2:   time.Date(2023, 12, 24, 23, 59, 59, 0, utc),
			want: true,
		},
		{
			name: "different date, t1 before t2",
			t1:   time.Date(2023, 12, 24, 23, 59, 59, 0, utc),
			t2:   time.Date(2023, 12, 25, 0, 0, 0, 0, utc),
			want: false,
		},
		{
			name: "same date in different timezones",
			t1:   time.Date(2023, 12, 25, 15, 0, 0, 0, utc),
			t2:   time.Date(2023, 12, 25, 5, 0, 0, 0, est), // 10:00 UTC (same day)
			want: false,
		},
		{
			name: "cross-timezone date comparison",
			t1:   time.Date(2023, 12, 26, 5, 0, 0, 0, utc),
			t2:   time.Date(2023, 12, 25, 23, 0, 0, 0, est), // Dec 26 04:00 UTC
			want: true,                                      // Dec 26 vs Dec 25
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsAfterDate(tt.t1, tt.t2)
			if got != tt.want {
				t.Errorf("IsAfterDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAfterDateOrEqual(t *testing.T) {
	utc := time.UTC

	tests := []struct {
		name string
		t1   time.Time
		t2   time.Time
		want bool
	}{
		{
			name: "same date, different time - should be true",
			t1:   time.Date(2023, 12, 25, 15, 30, 0, 0, utc),
			t2:   time.Date(2023, 12, 25, 10, 30, 0, 0, utc),
			want: true,
		},
		{
			name: "t1 date after t2 date",
			t1:   time.Date(2023, 12, 25, 0, 0, 1, 0, utc),
			t2:   time.Date(2023, 12, 24, 23, 59, 59, 0, utc),
			want: true,
		},
		{
			name: "t1 date before t2 date",
			t1:   time.Date(2023, 12, 24, 23, 59, 59, 0, utc),
			t2:   time.Date(2023, 12, 25, 0, 0, 0, 0, utc),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsAfterDateOrEqual(tt.t1, tt.t2)
			if got != tt.want {
				t.Errorf("IsAfterDateOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAfterInTimezone(t *testing.T) {
	utc := time.UTC
	est, _ := time.LoadLocation("America/New_York")
	pst, _ := time.LoadLocation("America/Los_Angeles")

	// Times that are the same moment but in different timezones
	utcTime := time.Date(2023, 12, 25, 15, 0, 0, 0, utc)
	estTime := time.Date(2023, 12, 25, 10, 0, 0, 0, est) // Same as 15:00 UTC

	tests := []struct {
		name     string
		t1       time.Time
		t2       time.Time
		timezone *time.Location
		want     bool
	}{
		{
			name:     "same moment in UTC",
			t1:       utcTime,
			t2:       estTime,
			timezone: utc,
			want:     false, // Same moment
		},
		{
			name:     "compare in EST timezone",
			t1:       time.Date(2023, 12, 25, 11, 0, 0, 0, est),
			t2:       utcTime,
			timezone: est,
			want:     true, // 11:00 EST vs 10:00 EST
		},
		{
			name:     "nil timezone should use UTC",
			t1:       time.Date(2023, 12, 25, 15, 0, 0, 0, utc),
			t2:       time.Date(2023, 12, 25, 14, 0, 0, 0, utc),
			timezone: nil,
			want:     true,
		},
		{
			name:     "cross-timezone comparison in PST",
			t1:       time.Date(2023, 12, 25, 10, 0, 0, 0, pst), // 18:00 UTC
			t2:       time.Date(2023, 12, 25, 12, 0, 0, 0, est), // 17:00 UTC
			timezone: pst,
			want:     true, // 10:00 PST vs 09:00 PST
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsAfterInTimezone(tt.t1, tt.t2, tt.timezone)
			if got != tt.want {
				t.Errorf("IsAfterInTimezone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkIsAfter(b *testing.B) {
	t1 := time.Date(2023, 12, 25, 10, 31, 0, 0, time.UTC)
	t2 := time.Date(2023, 12, 25, 10, 30, 0, 0, time.UTC)

	for i := 0; i < b.N; i++ {
		_ = IsAfter(t1, t2)
	}
}

func BenchmarkIsAfterDate(b *testing.B) {
	t1 := time.Date(2023, 12, 25, 0, 0, 1, 0, time.UTC)
	t2 := time.Date(2023, 12, 24, 23, 59, 59, 0, time.UTC)

	for i := 0; i < b.N; i++ {
		_ = IsAfterDate(t1, t2)
	}
}
