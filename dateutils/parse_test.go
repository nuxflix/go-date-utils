package dateutils

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	// Test timezone
	utc := time.UTC
	est, _ := time.LoadLocation("America/New_York")

	tests := []struct {
		name     string
		dateStr  string
		timezone *time.Location
		want     time.Time
		wantErr  bool
	}{
		{
			name:     "RFC3339 format",
			dateStr:  "2023-12-25T10:30:00Z",
			timezone: utc,
			want:     time.Date(2023, 12, 25, 10, 30, 0, 0, utc),
			wantErr:  false,
		},
		{
			name:     "ISO date format",
			dateStr:  "2023-12-25",
			timezone: utc,
			want:     time.Date(2023, 12, 25, 0, 0, 0, 0, utc),
			wantErr:  false,
		},
		{
			name:     "US date format",
			dateStr:  "12/25/2023",
			timezone: utc,
			want:     time.Date(2023, 12, 25, 0, 0, 0, 0, utc),
			wantErr:  false,
		},
		{
			name:     "Date with time",
			dateStr:  "2023-12-25 10:30:00",
			timezone: utc,
			want:     time.Date(2023, 12, 25, 10, 30, 0, 0, utc),
			wantErr:  false,
		},
		{
			name:     "With timezone conversion",
			dateStr:  "2023-12-25T10:30:00Z",
			timezone: est,
			want:     time.Date(2023, 12, 25, 10, 30, 0, 0, time.UTC).In(est),
			wantErr:  false,
		},
		{
			name:     "Empty string",
			dateStr:  "",
			timezone: utc,
			want:     time.Time{},
			wantErr:  true,
		},
		{
			name:     "Invalid format",
			dateStr:  "not-a-date",
			timezone: utc,
			want:     time.Time{},
			wantErr:  true,
		},
		{
			name:     "Nil timezone (should use UTC)",
			dateStr:  "2023-12-25",
			timezone: nil,
			want:     time.Date(2023, 12, 25, 0, 0, 0, 0, utc),
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.dateStr, tt.timezone)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !got.Equal(tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseWithFormat(t *testing.T) {
	utc := time.UTC
	est, _ := time.LoadLocation("America/New_York")

	tests := []struct {
		name     string
		dateStr  string
		format   string
		timezone *time.Location
		want     time.Time
		wantErr  bool
	}{
		{
			name:     "Custom format",
			dateStr:  "25-12-2023",
			format:   "02-01-2006",
			timezone: utc,
			want:     time.Date(2023, 12, 25, 0, 0, 0, 0, utc),
			wantErr:  false,
		},
		{
			name:     "Time format",
			dateStr:  "10:30:45",
			format:   "15:04:05",
			timezone: utc,
			want:     time.Date(0, 1, 1, 10, 30, 45, 0, utc),
			wantErr:  false,
		},
		{
			name:     "With timezone",
			dateStr:  "2023-12-25",
			format:   "2006-01-02",
			timezone: est,
			want:     time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC).In(est),
			wantErr:  false,
		},
		{
			name:     "Empty date string",
			dateStr:  "",
			format:   "2006-01-02",
			timezone: utc,
			want:     time.Time{},
			wantErr:  true,
		},
		{
			name:     "Empty format",
			dateStr:  "2023-12-25",
			format:   "",
			timezone: utc,
			want:     time.Time{},
			wantErr:  true,
		},
		{
			name:     "Invalid format",
			dateStr:  "2023-12-25",
			format:   "invalid",
			timezone: utc,
			want:     time.Time{},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseWithFormat(tt.dateStr, tt.format, tt.timezone)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseWithFormat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !got.Equal(tt.want) {
				t.Errorf("ParseWithFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkParse(b *testing.B) {
	utc := time.UTC
	for i := 0; i < b.N; i++ {
		_, _ = Parse("2023-12-25T10:30:00Z", utc)
	}
}

func BenchmarkParseWithFormat(b *testing.B) {
	utc := time.UTC
	for i := 0; i < b.N; i++ {
		_, _ = ParseWithFormat("2023-12-25", "2006-01-02", utc)
	}
}
