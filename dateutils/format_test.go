package dateutils

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	utc := time.UTC
	est, _ := time.LoadLocation("America/New_York")
	testTime := time.Date(2023, 12, 25, 15, 30, 45, 0, utc)
	zeroTime := time.Time{}

	tests := []struct {
		name     string
		time     time.Time
		format   string
		timezone *time.Location
		want     string
		wantErr  bool
	}{
		{
			name:     "ISO date format",
			time:     testTime,
			format:   DateISO,
			timezone: nil,
			want:     "2023-12-25",
			wantErr:  false,
		},
		{
			name:     "US date format",
			time:     testTime,
			format:   DateUS,
			timezone: nil,
			want:     "12/25/2023",
			wantErr:  false,
		},
		{
			name:     "24-hour time format",
			time:     testTime,
			format:   Time24,
			timezone: nil,
			want:     "15:30:45",
			wantErr:  false,
		},
		{
			name:     "12-hour time format",
			time:     testTime,
			format:   Time12,
			timezone: nil,
			want:     "3:30:45 PM",
			wantErr:  false,
		},
		{
			name:     "DateTime format",
			time:     testTime,
			format:   DateTime24,
			timezone: nil,
			want:     "2023-12-25 15:30:45",
			wantErr:  false,
		},
		{
			name:     "Readable format",
			time:     testTime,
			format:   Readable,
			timezone: nil,
			want:     "December 25, 2023",
			wantErr:  false,
		},
		{
			name:     "With timezone conversion",
			time:     testTime,
			format:   DateTime24,
			timezone: est,
			want:     "2023-12-25 10:30:45", // EST is UTC-5
			wantErr:  false,
		},
		{
			name:     "Zero time",
			time:     zeroTime,
			format:   DateISO,
			timezone: nil,
			want:     "",
			wantErr:  true,
		},
		{
			name:     "Empty format",
			time:     testTime,
			format:   "",
			timezone: nil,
			want:     "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Format(tt.time, tt.format, tt.timezone)
			if (err != nil) != tt.wantErr {
				t.Errorf("Format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatSafe(t *testing.T) {
	testTime := time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC)
	zeroTime := time.Time{}

	tests := []struct {
		name     string
		time     time.Time
		format   string
		timezone *time.Location
		want     string
	}{
		{
			name:     "Valid time",
			time:     testTime,
			format:   DateISO,
			timezone: nil,
			want:     "2023-12-25",
		},
		{
			name:     "Zero time should return empty string",
			time:     zeroTime,
			format:   DateISO,
			timezone: nil,
			want:     "",
		},
		{
			name:     "Empty format should return empty string",
			time:     testTime,
			format:   "",
			timezone: nil,
			want:     "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatSafe(tt.time, tt.format, tt.timezone)
			if got != tt.want {
				t.Errorf("FormatSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatWithDefault(t *testing.T) {
	testTime := time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC)
	zeroTime := time.Time{}
	defaultValue := "N/A"

	tests := []struct {
		name         string
		time         time.Time
		format       string
		timezone     *time.Location
		defaultValue string
		want         string
	}{
		{
			name:         "Valid time",
			time:         testTime,
			format:       DateISO,
			timezone:     nil,
			defaultValue: defaultValue,
			want:         "2023-12-25",
		},
		{
			name:         "Zero time should return default",
			time:         zeroTime,
			format:       DateISO,
			timezone:     nil,
			defaultValue: defaultValue,
			want:         defaultValue,
		},
		{
			name:         "Empty format should return default",
			time:         testTime,
			format:       "",
			timezone:     nil,
			defaultValue: defaultValue,
			want:         defaultValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatWithDefault(tt.time, tt.format, tt.timezone, tt.defaultValue)
			if got != tt.want {
				t.Errorf("FormatWithDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatCustom(t *testing.T) {
	testTime := time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC)
	zeroTime := time.Time{}

	tests := []struct {
		name         string
		time         time.Time
		customFormat string
		timezone     *time.Location
		want         string
		wantErr      bool
	}{
		{
			name:         "YYYY-MM-DD format",
			time:         testTime,
			customFormat: "YYYY-MM-DD",
			timezone:     nil,
			want:         "2023-12-25",
			wantErr:      false,
		},
		{
			name:         "DD/MM/YYYY format",
			time:         testTime,
			customFormat: "DD/MM/YYYY",
			timezone:     nil,
			want:         "25/12/2023",
			wantErr:      false,
		},
		{
			name:         "HH:mm:ss format",
			time:         testTime,
			customFormat: "HH:mm:ss",
			timezone:     nil,
			want:         "15:30:45",
			wantErr:      false,
		},
		{
			name:         "12-hour format with AM/PM",
			time:         testTime,
			customFormat: "hh:mm:ss AM/PM",
			timezone:     nil,
			want:         "03:30:45 PM",
			wantErr:      false,
		},
		{
			name:         "Mixed format",
			time:         testTime,
			customFormat: "DD-MM-YYYY HH:mm",
			timezone:     nil,
			want:         "25-12-2023 15:30",
			wantErr:      false,
		},
		{
			name:         "Zero time",
			time:         zeroTime,
			customFormat: "YYYY-MM-DD",
			timezone:     nil,
			want:         "",
			wantErr:      true,
		},
		{
			name:         "Empty format",
			time:         testTime,
			customFormat: "",
			timezone:     nil,
			want:         "",
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FormatCustom(tt.time, tt.customFormat, tt.timezone)
			if (err != nil) != tt.wantErr {
				t.Errorf("FormatCustom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FormatCustom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFormat(b *testing.B) {
	testTime := time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		_, _ = Format(testTime, DateISO, nil)
	}
}

func BenchmarkFormatCustom(b *testing.B) {
	testTime := time.Date(2023, 12, 25, 15, 30, 45, 0, time.UTC)
	for i := 0; i < b.N; i++ {
		_, _ = FormatCustom(testTime, "YYYY-MM-DD HH:mm:ss", nil)
	}
}
