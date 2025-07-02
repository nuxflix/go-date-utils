package dateutils

import (
	"testing"
	"time"
)

func TestFormatDistance(t *testing.T) {
	baseDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		date     time.Time
		baseDate time.Time
		options  *FormatDistanceOptions
		expected string
	}{
		// Less than a minute tests
		{
			name:     "5 seconds without includeSeconds",
			date:     baseDate.Add(5 * time.Second),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "less than a minute",
		},
		{
			name:     "3 seconds with includeSeconds",
			date:     baseDate.Add(3 * time.Second),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{IncludeSeconds: true},
			expected: "less than 5 seconds",
		},
		{
			name:     "5 seconds with includeSeconds",
			date:     baseDate.Add(5 * time.Second),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{IncludeSeconds: true},
			expected: "less than 10 seconds",
		},
		{
			name:     "15 seconds with includeSeconds",
			date:     baseDate.Add(15 * time.Second),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{IncludeSeconds: true},
			expected: "less than 20 seconds",
		},
		{
			name:     "30 seconds with includeSeconds",
			date:     baseDate.Add(30 * time.Second),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{IncludeSeconds: true},
			expected: "half a minute",
		},
		{
			name:     "50 seconds with includeSeconds",
			date:     baseDate.Add(50 * time.Second),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{IncludeSeconds: true},
			expected: "less than a minute",
		},

		// Minutes tests
		{
			name:     "1 minute",
			date:     baseDate.Add(1 * time.Minute),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "1 minute",
		},
		{
			name:     "2 minutes",
			date:     baseDate.Add(2 * time.Minute),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "2 minutes",
		},
		{
			name:     "15 minutes",
			date:     baseDate.Add(15 * time.Minute),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "15 minutes",
		},
		{
			name:     "30 minutes",
			date:     baseDate.Add(30 * time.Minute),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "30 minutes",
		},

		// Hours tests
		{
			name:     "45 minutes (about 1 hour)",
			date:     baseDate.Add(45 * time.Minute),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "about 1 hour",
		},
		{
			name:     "1 hour 30 minutes",
			date:     baseDate.Add(90 * time.Minute),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "about 2 hours",
		},
		{
			name:     "6 hours",
			date:     baseDate.Add(6 * time.Hour),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "about 6 hours",
		},

		// Days tests
		{
			name:     "1 day",
			date:     baseDate.Add(24 * time.Hour),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "1 day",
		},
		{
			name:     "2 days",
			date:     baseDate.Add(48 * time.Hour),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "2 days",
		},
		{
			name:     "14 days",
			date:     baseDate.Add(14 * 24 * time.Hour),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "14 days",
		},

		// Months tests
		{
			name:     "about 1 month",
			date:     baseDate.AddDate(0, 1, 0),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "about 1 month",
		},
		{
			name:     "2 months",
			date:     baseDate.AddDate(0, 2, 0),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "2 months",
		},
		{
			name:     "6 months",
			date:     baseDate.AddDate(0, 6, 0),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "6 months",
		},

		// Years tests
		{
			name:     "about 1 year",
			date:     baseDate.AddDate(1, 0, 0),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "about 1 year",
		},
		{
			name:     "over 1 year (1 year 6 months)",
			date:     baseDate.AddDate(1, 6, 0),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "over 1 year",
		},
		{
			name:     "almost 2 years (1 year 10 months)",
			date:     baseDate.AddDate(1, 10, 0),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "almost 2 years",
		},
		{
			name:     "about 2 years",
			date:     baseDate.AddDate(2, 0, 0),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "about 2 years",
		},

		// Past dates with suffix
		{
			name:     "1 hour ago",
			date:     baseDate.Add(-1 * time.Hour),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{AddSuffix: true},
			expected: "about 1 hour ago",
		},
		{
			name:     "2 days ago",
			date:     baseDate.Add(-48 * time.Hour),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{AddSuffix: true},
			expected: "2 days ago",
		},

		// Future dates with suffix
		{
			name:     "in 1 hour",
			date:     baseDate.Add(1 * time.Hour),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{AddSuffix: true},
			expected: "in about 1 hour",
		},
		{
			name:     "in 2 days",
			date:     baseDate.Add(48 * time.Hour),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{AddSuffix: true},
			expected: "in 2 days",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatDistance(tt.date, tt.baseDate, tt.options)
			if result != tt.expected {
				t.Errorf("FormatDistance() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFormatDistanceStrict(t *testing.T) {
	baseDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		date     time.Time
		baseDate time.Time
		options  *FormatDistanceOptions
		expected string
	}{
		{
			name:     "5 seconds strict",
			date:     baseDate.Add(5 * time.Second),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "5 seconds",
		},
		{
			name:     "1 second strict",
			date:     baseDate.Add(1 * time.Second),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "1 second",
		},
		{
			name:     "1 minute strict",
			date:     baseDate.Add(1 * time.Minute),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "1 minute",
		},
		{
			name:     "2 minutes strict",
			date:     baseDate.Add(2 * time.Minute),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "2 minutes",
		},
		{
			name:     "1 hour strict",
			date:     baseDate.Add(1 * time.Hour),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "1 hour",
		},
		{
			name:     "2 hours strict",
			date:     baseDate.Add(2 * time.Hour),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "2 hours",
		},
		{
			name:     "1 day strict",
			date:     baseDate.Add(24 * time.Hour),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "1 day",
		},
		{
			name:     "2 days strict",
			date:     baseDate.Add(48 * time.Hour),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "2 days",
		},
		{
			name:     "1 week strict",
			date:     baseDate.Add(7 * 24 * time.Hour),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "1 week",
		},
		{
			name:     "2 weeks strict",
			date:     baseDate.Add(14 * 24 * time.Hour),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "2 weeks",
		},
		{
			name:     "1 month strict",
			date:     baseDate.AddDate(0, 1, 0),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "1 month",
		},
		{
			name:     "2 months strict",
			date:     baseDate.AddDate(0, 2, 0),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "2 months",
		},
		{
			name:     "1 year strict",
			date:     baseDate.AddDate(1, 0, 0),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "1 year",
		},
		{
			name:     "2 years strict",
			date:     baseDate.AddDate(2, 0, 0),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "2 years",
		},

		// With suffix tests
		{
			name:     "1 hour ago strict",
			date:     baseDate.Add(-1 * time.Hour),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{AddSuffix: true},
			expected: "1 hour ago",
		},
		{
			name:     "in 2 days strict",
			date:     baseDate.Add(48 * time.Hour),
			baseDate: baseDate,
			options:  &FormatDistanceOptions{AddSuffix: true},
			expected: "in 2 days",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatDistanceStrict(tt.date, tt.baseDate, tt.options)
			if result != tt.expected {
				t.Errorf("FormatDistanceStrict() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFormatDistanceToNow(t *testing.T) {
	// Test that FormatDistanceToNow works correctly
	past := time.Now().Add(-1 * time.Hour)
	result := FormatDistanceToNow(past, &FormatDistanceOptions{AddSuffix: true})

	// The result should contain "hour ago"
	if result != "about 1 hour ago" {
		t.Errorf("FormatDistanceToNow() = %v, want about 1 hour ago", result)
	}
}

func TestFormatDistanceToNowStrict(t *testing.T) {
	// Test that FormatDistanceToNowStrict works correctly
	past := time.Now().Add(-1 * time.Hour)
	result := FormatDistanceToNowStrict(past, &FormatDistanceOptions{AddSuffix: true})

	// The result should contain "hour ago"
	if result != "1 hour ago" {
		t.Errorf("FormatDistanceToNowStrict() = %v, want 1 hour ago", result)
	}
}

func TestFormatDistanceNilOptions(t *testing.T) {
	baseDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	date := baseDate.Add(1 * time.Hour)

	// Test with nil options
	result := FormatDistance(date, baseDate, nil)
	expected := "about 1 hour"

	if result != expected {
		t.Errorf("FormatDistance() with nil options = %v, want %v", result, expected)
	}
}

func TestFormatDistanceStrictNilOptions(t *testing.T) {
	baseDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	date := baseDate.Add(1 * time.Hour)

	// Test with nil options
	result := FormatDistanceStrict(date, baseDate, nil)
	expected := "1 hour"

	if result != expected {
		t.Errorf("FormatDistanceStrict() with nil options = %v, want %v", result, expected)
	}
}

func TestFormatDistanceEdgeCases(t *testing.T) {
	baseDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		date     time.Time
		baseDate time.Time
		options  *FormatDistanceOptions
		expected string
	}{
		{
			name:     "same time",
			date:     baseDate,
			baseDate: baseDate,
			options:  &FormatDistanceOptions{},
			expected: "less than a minute",
		},
		{
			name:     "same time with includeSeconds",
			date:     baseDate,
			baseDate: baseDate,
			options:  &FormatDistanceOptions{IncludeSeconds: true},
			expected: "less than 5 seconds",
		},
		{
			name:     "same time with suffix",
			date:     baseDate,
			baseDate: baseDate,
			options:  &FormatDistanceOptions{AddSuffix: true},
			expected: "less than a minute ago",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatDistance(tt.date, tt.baseDate, tt.options)
			if result != tt.expected {
				t.Errorf("FormatDistance() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// Benchmark tests
func BenchmarkFormatDistance(b *testing.B) {
	baseDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	date := baseDate.Add(2 * time.Hour)
	options := &FormatDistanceOptions{AddSuffix: true}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FormatDistance(date, baseDate, options)
	}
}

func BenchmarkFormatDistanceStrict(b *testing.B) {
	baseDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	date := baseDate.Add(2 * time.Hour)
	options := &FormatDistanceOptions{AddSuffix: true}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FormatDistanceStrict(date, baseDate, options)
	}
}

func BenchmarkFormatDistanceToNow(b *testing.B) {
	date := time.Now().Add(-2 * time.Hour)
	options := &FormatDistanceOptions{AddSuffix: true}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FormatDistanceToNow(date, options)
	}
}
