# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

## [Unreleased]

---

## [0.1.0] - 2026-02-28

### Added

#### Parsing
- `Parse` — Multi-format date string parsing with automatic format detection
- `ParseISO` — ISO 8601 date string parsing with timezone support
- `ParseWithFormat` — Custom format date string parsing
- `IsValidISO` — Validate ISO 8601 string format

#### Formatting
- `Format` — Standard formatting using Go reference time layout
- `FormatCustom` — User-friendly format strings (`YYYY-MM-DD`, `HH:mm:ss`)
- `FormatSafe` — Error-safe formatting (returns `""` on failure)
- `FormatWithDefault` — Formatting with fallback default value
- `LightFormat` — Lightweight formatting with limited token set
- `FormatDistance` — Human-readable relative time (e.g. "about 2 hours ago")
- `FormatDistanceStrict` — Exact relative time without approximations
- `FormatDistanceToNow` — Convenience wrapper comparing to current time
- `FormatDistanceToNowStrict` — Strict version of `FormatDistanceToNow`

#### Comparison
- `IsBefore`, `IsBeforeOrEqual`, `IsBeforeDate`, `IsBeforeInTimezone`
- `IsAfter`, `IsAfterOrEqual`, `IsAfterDate`, `IsAfterInTimezone`
- `IsEqual` — Exact time equality
- `IsSameDay`, `IsSameWeek`, `IsSameMonth`, `IsSameYear`
- `IsSameHour`, `IsSameMinute`
- `CompareAsc`, `CompareDesc` — Sortable comparison returning -1, 0, 1
- `Min`, `Max`, `ClosestTo` — Selection from slices of dates

#### Manipulation
- `AddDays`, `SubDays`
- `AddWeeks`, `SubWeeks`
- `AddMonths`, `SubMonths`
- `AddYears`, `SubYears`
- `AddHours`, `SubHours`
- `AddMinutes`, `SubMinutes`
- `AddSeconds`, `SubSeconds`
- `AddBusinessDays`, `SubBusinessDays`
- `AddDaysWithTimezone`

#### Differences
- `DifferenceInSeconds`, `DifferenceInSecondsFloat`
- `DifferenceInMinutes`, `DifferenceInMinutesFloat`
- `DifferenceInHours`, `DifferenceInHoursFloat`
- `DifferenceInDays`, `DifferenceInDaysFloat`, `DifferenceInCalendarDays`
- `DifferenceInWeeks`, `DifferenceInWeeksFloat`
- `DifferenceInMonths`, `DifferenceInCalendarMonths`
- `DifferenceInYears`, `DifferenceInCalendarYears`
- `DifferenceInQuarters`, `DifferenceInCalendarQuarters`
- `DifferenceInBusinessDays`
- `AbsDifferenceInSeconds`, `AbsDifferenceInMinutes`, `AbsDifferenceInHours`
- `AbsDifferenceInDays`, `AbsDifferenceInMonths`, `AbsDifferenceInYears`

#### Get Functions
- `GetYear`, `GetMonth`, `GetDate`, `GetDay`, `GetDayOfYear`, `GetQuarter`
- `GetHours`, `GetMinutes`, `GetSeconds`, `GetMilliseconds`
- `GetWeek`, `GetWeekOfMonth`, `GetWeekYear`
- `GetDaysInMonth`, `GetDaysInYear`

#### Set Functions
- `SetYear`, `SetMonth`, `SetDate`, `SetDay`, `SetDayOfYear`, `SetQuarter`
- `SetHours`, `SetMinutes`, `SetSeconds`, `SetMilliseconds`

#### Validation
- `IsValid`, `IsLeapYear`
- `IsWeekend`, `IsWeekday`
- `IsToday`, `IsTomorrow`, `IsYesterday`
- `IsWithinInterval`
- `IsFirstDayOfMonth`, `IsLastDayOfMonth`
- `IsFirstDayOfYear`, `IsLastDayOfYear`
- `IsSameDate`
- `IsMonday`, `IsTuesday`, `IsWednesday`, `IsThursday`, `IsFriday`, `IsSaturday`, `IsSunday`

#### Period Utilities
- `StartOfDay`, `EndOfDay`
- `StartOfWeek`, `StartOfWeekSunday`, `EndOfWeek`, `EndOfWeekSaturday`
- `StartOfMonth`, `EndOfMonth`
- `StartOfYear`, `EndOfYear`
- `StartOfHour`, `EndOfHour`
- `StartOfMinute`, `EndOfMinute`
- `StartOfQuarter`, `EndOfQuarter`
- `StartOfDecade`, `EndOfDecade`, `LastDayOfDecade`
- `StartOfCentury`, `EndOfCentury`, `LastDayOfCentury`

#### Interval Utilities
- `EachDayOfInterval`
- `EachWeekOfInterval`, `EachWeekOfIntervalSunday`
- `EachMonthOfInterval`
- `EachYearOfInterval`
- `EachQuarterOfInterval`
- `EachHourOfInterval`
- `EachMinuteOfInterval`
- `EachWeekendOfInterval`
- `EachBusinessDayOfInterval`

#### Helpers
- `RoundToNearestMinutes`
- `NextDay`, `PreviousDay` (next/previous occurrence of a weekday)
- `IsValidTimestamp` — Check if Unix timestamp is within valid range
- `MaxTime`, `MinTime` — Max/min representable timestamps

#### Constants
- Format constants: `DateISO`, `DateUS`, `DateEU`, `DateTime24`, `DateTime12`, etc.
- Weekday constants: `Sunday` through `Saturday`
- Month constants: `January` through `December`
- Quarter constants: `Q1` through `Q4`
- Time constants: `MillisecondsInSecond`, `DaysInWeek`, etc.

[Unreleased]: https://github.com/chmenegatti/go-date-fns/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/chmenegatti/go-date-fns/releases/tag/v0.1.0
