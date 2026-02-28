# Contributing to go-date-fns

First off, thank you for considering contributing to **go-date-fns**! 🎉

This document provides guidelines for contributing to keep the project high-quality, consistent, and easy to maintain.

---

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Project Structure](#project-structure)
- [How to Contribute](#how-to-contribute)
- [Adding New Functions](#adding-new-functions)
- [Code Style](#code-style)
- [Testing Guidelines](#testing-guidelines)
- [Commit Messages](#commit-messages)
- [Pull Request Process](#pull-request-process)

---

## Code of Conduct

This project adheres to our [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

---

## Getting Started

1. **Fork** the repository on GitHub.
2. **Clone** your fork locally:
   ```bash
   git clone https://github.com/YOUR_USERNAME/go-date-fns.git
   cd go-date-fns
   ```
3. **Add the upstream remote**:
   ```bash
   git remote add upstream https://github.com/chmenegatti/go-date-fns.git
   ```
4. **Create a branch** for your changes:
   ```bash
   git checkout -b feature/add-my-function
   ```

---

## Development Setup

**Requirements**: Go 1.21+

```bash
# Run all tests
go test ./...

# Run tests with race detector (required before submitting)
go test -race ./...

# Run tests with coverage
go test -cover ./...

# Run benchmarks
go test -bench=. -benchmem ./...

# Lint
go vet ./...
```

---

## Project Structure

```
go-date-fns/
├── dateutils/          # All library functions (one file per feature group)
│   ├── add_days.go              # AddDays, AddWeeks, AddBusinessDays
│   ├── format.go                # Format, FormatCustom, FormatSafe
│   ├── format_distance.go       # FormatDistance, FormatDistanceToNow
│   ├── comparison_functions.go  # IsBefore, IsAfter, IsSameDay...
│   ├── difference_in_days.go    # DifferenceInDays, DifferenceInBusinessDays
│   ├── interval_utilities.go    # EachDayOfInterval, EachWeekOfInterval...
│   ├── doc.go                   # Package-level GoDoc
│   └── *_test.go                # Tests alongside source files
├── examples/           # Runnable example program
├── .github/
│   ├── workflows/       # GitHub Actions CI/CD
│   └── ISSUE_TEMPLATE/  # Issue templates
├── CHANGELOG.md
├── CONTRIBUTING.md
└── README.md
```

---

## How to Contribute

### Reporting Bugs

- Check the [existing issues](https://github.com/chmenegatti/go-date-fns/issues) first.
- Use the **Bug Report** issue template.
- Include the Go version, OS, and a minimal reproducible example.

### Requesting Features

- Check the [existing issues](https://github.com/chmenegatti/go-date-fns/issues) and the [date-fns docs](https://date-fns.org/docs/) for reference.
- Use the **Feature Request** issue template.
- Reference the equivalent `date-fns` function if applicable.

### Fixing Bugs / Implementing Features

1. Comment on the issue to signal you're working on it.
2. Follow the [Adding New Functions](#adding-new-functions) guide.
3. Open a Pull Request using the PR template.

---

## Adding New Functions

Every new function must follow this checklist:

### 1. Choose the Right File

Place your function in the most appropriate existing file, or create a new file if it introduces a new category:

| Category | File |
|---|---|
| Parsing | `parse.go` / `parse_iso.go` |
| Formatting | `format.go` / `format_distance.go` |
| Comparison | `comparison_functions.go` / `is_before.go` / `is_after.go` |
| Manipulation | `add_days.go` / `sub_and_add_time_units.go` |
| Differences | `difference_in_days.go` / `additional_difference_functions.go` |
| Validation | `is_valid.go` / `additional_utility_functions.go` |
| Period Utilities | `start_of_day.go` |
| Interval | `interval_utilities.go` |

### 2. Write the Function

Follow the existing style:

```go
// MyFunction does X with the given date.
// Returns Y. Supports negative values for Z.
//
// Example:
//
//	result := MyFunction(time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC))
//	// Returns: ...
func MyFunction(t time.Time) time.Time {
    // implementation
}
```

Key rules:
- **Pure function**: never modify the input `time.Time`
- **Immutable**: always return a new `time.Time`
- **GoDoc comment**: required for all exported functions
- **Code example** in the GoDoc: required

### 3. Write Tests

Tests go in the corresponding `_test.go` file:

```go
func TestMyFunction(t *testing.T) {
    tests := []struct {
        name     string
        date     time.Time
        expected time.Time
    }{
        {
            name:     "Basic case",
            date:     time.Date(2024, time.January, 15, 0, 0, 0, 0, time.UTC),
            expected: time.Date(2024, time.January, 15, 0, 0, 0, 0, time.UTC),
        },
        // Add edge cases: zero time, leap years, timezone changes, DST, etc.
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := MyFunction(tt.date)
            if !result.Equal(tt.expected) {
                t.Errorf("MyFunction(%v) = %v, expected %v", tt.date, result, tt.expected)
            }
        })
    }
}
```

Required edge cases to consider:
- Zero `time.Time` value
- Dates around DST transitions
- Leap years (especially Feb 29)
- Year boundaries (Dec 31 → Jan 1)

### 4. Write a Benchmark

```go
func BenchmarkMyFunction(b *testing.B) {
    date := time.Date(2024, time.January, 15, 14, 30, 45, 0, time.UTC)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        MyFunction(date)
    }
}
```

### 5. Verify

```bash
go test -race ./...
go vet ./...
```

---

## Code Style

- Follow standard Go conventions (`gofmt`, `go vet`)
- Use clear, descriptive variable names
- Prefer early returns over deep nesting
- Do not use external dependencies — only the Go standard library
- Function names should match the `date-fns` equivalent when possible (e.g., `addDays` → `AddDays`)

---

## Testing Guidelines

- **100% coverage** is the goal for all new functions
- Use **table-driven tests** (see examples above)
- Include **benchmarks** for all new functions
- Tests must pass with `go test -race ./...`
- Test file naming: `myfunc.go` → `myfunc_test.go`

---

## Commit Messages

Follow [Conventional Commits](https://www.conventionalcommits.org/):

```
feat: add RoundToNearestHours function
fix: correct RoundToNearestMinutes threshold for odd intervals
docs: update README with new interval functions
test: add edge cases for DifferenceInBusinessDays
chore: update go.mod to Go 1.21
```

---

## Pull Request Process

1. Ensure all tests pass: `go test -race ./...`
2. Ensure `go vet ./...` reports no issues
3. Fill in the PR template completely
4. Link the PR to the relevant issue (`Fixes #123`)
5. Wait for a review — we aim to respond within 3 business days

**Thank you for contributing!** 🚀
