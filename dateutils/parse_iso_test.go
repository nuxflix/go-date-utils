package dateutils

import (
	"testing"
	"time"
)

func TestParseISO(t *testing.T) {
	utc := time.UTC
	est, _ := time.LoadLocation("America/New_York")

	tests := []struct {
		name     string
		isoStr   string
		timezone *time.Location
		want     time.Time
		wantErr  bool
	}{
		{
			name:     "RFC3339 with Z",
			isoStr:   "2023-12-25T10:30:00Z",
			timezone: utc,
			want:     time.Date(2023, 12, 25, 10, 30, 0, 0, utc),
			wantErr:  false,
		},
		{
			name:     "RFC3339 with timezone offset",
			isoStr:   "2023-12-25T10:30:00+05:00",
			timezone: utc,
			want:     time.Date(2023, 12, 25, 5, 30, 0, 0, utc),
			wantErr:  false,
		},
		{
			name:     "RFC3339 with nanoseconds",
			isoStr:   "2023-12-25T10:30:00.123456789Z",
			timezone: utc,
			want:     time.Date(2023, 12, 25, 10, 30, 0, 123456789, utc),
			wantErr:  false,
		},
		{
			name:     "Date only",
			isoStr:   "2023-12-25",
			timezone: utc,
			want:     time.Date(2023, 12, 25, 0, 0, 0, 0, utc),
			wantErr:  false,
		},
		{
			name:     "DateTime without timezone",
			isoStr:   "2023-12-25T10:30:00",
			timezone: utc,
			want:     time.Date(2023, 12, 25, 10, 30, 0, 0, utc),
			wantErr:  false,
		},
		{
			name:     "With timezone conversion",
			isoStr:   "2023-12-25T10:30:00Z",
			timezone: est,
			want:     time.Date(2023, 12, 25, 10, 30, 0, 0, time.UTC).In(est),
			wantErr:  false,
		},
		{
			name:     "Empty string",
			isoStr:   "",
			timezone: utc,
			want:     time.Time{},
			wantErr:  true,
		},
		{
			name:     "Invalid ISO format",
			isoStr:   "25-12-2023",
			timezone: utc,
			want:     time.Time{},
			wantErr:  true,
		},
		{
			name:     "Nil timezone (should use UTC)",
			isoStr:   "2023-12-25T10:30:00Z",
			timezone: nil,
			want:     time.Date(2023, 12, 25, 10, 30, 0, 0, utc),
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseISO(tt.isoStr, tt.timezone)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseISO() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !got.Equal(tt.want) {
				t.Errorf("ParseISO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidISO(t *testing.T) {
	tests := []struct {
		name   string
		isoStr string
		want   bool
	}{
		{
			name:   "Valid RFC3339",
			isoStr: "2023-12-25T10:30:00Z",
			want:   true,
		},
		{
			name:   "Valid date only",
			isoStr: "2023-12-25",
			want:   true,
		},
		{
			name:   "Valid with timezone",
			isoStr: "2023-12-25T10:30:00+05:00",
			want:   true,
		},
		{
			name:   "Valid with nanoseconds",
			isoStr: "2023-12-25T10:30:00.123Z",
			want:   true,
		},
		{
			name:   "Invalid format",
			isoStr: "25-12-2023",
			want:   false,
		},
		{
			name:   "Empty string",
			isoStr: "",
			want:   false,
		},
		{
			name:   "Random string",
			isoStr: "not-a-date",
			want:   false,
		},
		{
			name:   "US format (not ISO)",
			isoStr: "12/25/2023",
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidISO(tt.isoStr)
			if got != tt.want {
				t.Errorf("IsValidISO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkParseISO(b *testing.B) {
	utc := time.UTC
	for i := 0; i < b.N; i++ {
		_, _ = ParseISO("2023-12-25T10:30:00Z", utc)
	}
}

func BenchmarkIsValidISO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsValidISO("2023-12-25T10:30:00Z")
	}
}
