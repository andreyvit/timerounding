package timerounding_test

import (
	"fmt"
	"time"

	"github.com/andreyvit/timerounding"
)

func Example() {
	t := time.Date(2017, 1, 7, 9, 37, 12, 0, time.UTC)
	fmt.Println(timerounding.Round(t, 5*time.Minute))
	fmt.Println(timerounding.Round(t, 24*time.Hour))
	fmt.Println(timerounding.FormatRounded(t, 5*time.Minute, timerounding.Concise))
	fmt.Println(timerounding.FormatRounded(t, 48*time.Hour, timerounding.Concise))
	fmt.Println(timerounding.MakeInterval(20 * time.Minute).Round(t))
	fmt.Println(timerounding.MakeInterval(20*time.Minute).Next(t, 1))
	fmt.Println(timerounding.MakeInterval(20*time.Minute).Prev(t, 1))
	// Output: 2017-01-07 09:35:00 +0000 UTC
	// 2017-01-07 00:00:00 +0000 UTC
	// 20170107-0935
	// 20170106
	// 2017-01-07 09:20:00 +0000 UTC
	// 2017-01-07 09:40:00 +0000 UTC
	// 2017-01-07 09:00:00 +0000 UTC
}

func ExampleFormatRounded() {
	fmt.Println(timerounding.FormatRounded(time.Date(2017, 1, 7, 9, 37, 12, 0, time.UTC), 5*time.Minute, timerounding.Concise))
	fmt.Println(timerounding.FormatRounded(time.Date(2017, 1, 7, 9, 37, 12, 0, time.UTC), 2*time.Hour, timerounding.Concise))
	fmt.Println(timerounding.FormatRounded(time.Date(2017, 1, 7, 9, 37, 12, 0, time.UTC), 24*time.Hour, timerounding.Concise))
	fmt.Println()
	fmt.Println(timerounding.FormatRounded(time.Date(2016, 1, 2, 13, 4, 0, 0, time.UTC), 5*time.Minute, timerounding.Concise))
	fmt.Println(timerounding.FormatRounded(time.Date(2016, 1, 2, 13, 5, 0, 0, time.UTC), 5*time.Minute, timerounding.Concise))
	fmt.Println(timerounding.FormatRounded(time.Date(2016, 1, 2, 13, 6, 0, 0, time.UTC), 5*time.Minute, timerounding.Concise))
	fmt.Println(timerounding.FormatRounded(time.Date(2016, 1, 2, 13, 9, 0, 0, time.UTC), 5*time.Minute, timerounding.Concise))
	fmt.Println(timerounding.FormatRounded(time.Date(2016, 1, 2, 13, 59, 0, 0, time.UTC), 5*time.Minute, timerounding.Concise))
	fmt.Println()
	fmt.Println(timerounding.FormatRounded(time.Date(2016, 1, 2, 13, 9, 0, 0, time.UTC), 15*time.Minute, timerounding.Concise))
	fmt.Println(timerounding.FormatRounded(time.Date(2016, 1, 2, 13, 40, 0, 0, time.UTC), 15*time.Minute, timerounding.Concise))
	// Output: 20170107-0935
	// 20170107-08
	// 20170107
	//
	// 20160102-1300
	// 20160102-1305
	// 20160102-1305
	// 20160102-1305
	// 20160102-1355
	//
	// 20160102-1300
	// 20160102-1330
}

func ExampleMakeInterval() {
	fmt.Println(timerounding.MakeInterval(0))
	fmt.Println(timerounding.MakeInterval(1 * time.Second))
	fmt.Println(timerounding.MakeInterval(5 * time.Second))
	fmt.Println(timerounding.MakeInterval(1 * time.Minute))
	fmt.Println(timerounding.MakeInterval(15 * time.Minute))
	fmt.Println(timerounding.MakeInterval(1 * time.Hour))
	fmt.Println(timerounding.MakeInterval(2 * time.Hour))
	fmt.Println(timerounding.MakeInterval(23 * time.Hour))
	fmt.Println(timerounding.MakeInterval(24 * time.Hour))
	fmt.Println(timerounding.MakeInterval(48 * time.Hour))
	// Output: none
	// 1s
	// 5s
	// 1m
	// 15m
	// 1h
	// 2h
	// 23h
	// 1d
	// 2d
}

func ExampleTryMakeInterval() {
	fmt.Println(timerounding.TryMakeInterval(0))
	fmt.Println(timerounding.TryMakeInterval(15 * time.Minute))
	fmt.Println(timerounding.TryMakeInterval(-15 * time.Minute))
	fmt.Println(timerounding.TryMakeInterval(15*time.Minute + 5*time.Second))
	fmt.Println(timerounding.TryMakeInterval(500 * time.Millisecond))
	// Output: none <nil>
	// 15m <nil>
	// none negative duration
	// 15m duration is a mix of multiple units
	// none duration too small
}

func ExampleInterval_AddTo() {
	t := time.Date(2017, 1, 7, 9, 37, 12, 0, time.UTC)
	fmt.Println(timerounding.Interval5m.AddTo(t, 1))
	fmt.Println(timerounding.Interval5m.AddTo(t, 3))
	// Output: 2017-01-07 09:42:12 +0000 UTC
	// 2017-01-07 09:52:12 +0000 UTC
}

func ExampleInterval_Next() {
	t := time.Date(2017, 1, 7, 9, 37, 12, 0, time.UTC)
	fmt.Println(timerounding.Interval5m.Next(t, 2))
	fmt.Println(timerounding.Interval5m.Next(t, 1))
	fmt.Println(timerounding.Interval5m.Next(t, 0))
	fmt.Println(timerounding.Interval5m.Next(t, -1))
	fmt.Println(timerounding.Interval5m.Next(t, -2))
	// Output: 2017-01-07 09:45:00 +0000 UTC
	// 2017-01-07 09:40:00 +0000 UTC
	// 2017-01-07 09:35:00 +0000 UTC
	// 2017-01-07 09:30:00 +0000 UTC
	// 2017-01-07 09:25:00 +0000 UTC
}

func ExampleInterval_Prev() {
	t := time.Date(2017, 1, 7, 9, 37, 12, 0, time.UTC)
	fmt.Println(timerounding.Interval5m.Prev(t, 1))
	// Output: 2017-01-07 09:30:00 +0000 UTC
}

func ExampleInterval_Round() {
	t := time.Date(2017, 1, 7, 9, 37, 12, 0, time.UTC)
	fmt.Println(timerounding.Interval5m.Round(t))
	// Output: 2017-01-07 09:35:00 +0000 UTC
}

func ExampleRound() {
	t := time.Date(2017, 1, 7, 9, 37, 12, 0, time.UTC)
	fmt.Println(timerounding.Round(t, 5*time.Minute))
	fmt.Println(timerounding.Round(t, 2*24*time.Hour))
	// Output: 2017-01-07 09:35:00 +0000 UTC
	// 2017-01-06 00:00:00 +0000 UTC
}
