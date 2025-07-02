package main

import (
	"fmt"
	"time"
)

func main() {
	// Test dates that are failing
	d1 := time.Date(2014, time.January, 1, 0, 0, 0, 0, time.UTC)
	d2 := time.Date(2014, time.December, 31, 0, 0, 0, 0, time.UTC)

	year1, week1 := d1.ISOWeek()
	year2, week2 := d2.ISOWeek()

	fmt.Printf("2014-01-01: ISO year %d, week %d (weekday: %s)\n", year1, week1, d1.Weekday())
	fmt.Printf("2014-12-31: ISO year %d, week %d (weekday: %s)\n", year2, week2, d2.Weekday())

	// Test unix timestamp issue
	testDate := time.Date(2012, time.February, 29, 11, 45, 5, 0, time.UTC)
	unixTime := testDate.Unix()
	fmt.Printf("2012-02-29 11:45:05 UTC -> Unix timestamp: %d\n", unixTime)

	// Check what the expected value should be
	// Feb 29, 2012 11:45:05 UTC in Unix time
	expectedUnix := int64(1330512305)
	expectedTime := time.Unix(expectedUnix, 0).UTC()
	fmt.Printf("Expected Unix timestamp %d -> %s\n", expectedUnix, expectedTime.Format("2006-01-02 15:04:05 MST"))
	fmt.Printf("Actual Unix timestamp %d -> %s\n", unixTime, time.Unix(unixTime, 0).UTC().Format("2006-01-02 15:04:05 MST"))
}
