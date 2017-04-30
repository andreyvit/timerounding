package timerounding

import (
	"errors"
	"fmt"
	"time"
)

// An Interval represents a positive multiple of the given time unit (e.g. 5 minutes, or 1 hour, or 2 days). When the unit is set to None (which is the zero value of the Interval), the interval is unspecified and generally means that the caller asks to disable time rounding.
type Interval struct {
	N    int
	Unit Unit
}

var (
	IntervalZero = Interval{0, None}
	Interval1m   = Interval{1, Minutes}
	Interval5m   = Interval{5, Minutes}
	Interval15m  = Interval{15, Minutes}
	Interval1h   = Interval{1, Hours}
	Interval1d   = Interval{1, Days}
)

var (
	ErrDurationNegative   = errors.New("negative duration")
	ErrDurationTooSmall   = errors.New("duration too small")
	ErrDurationMixesUnits = errors.New("duration is a mix of multiple units")
)

var (
	units = []Unit{Days, Hours, Minutes, Seconds}
)

// TryMakeInterval returns an Interval corresponding to the given time.Duration, or an error if the given duration cannot be expressed as an interval. A zero duration results in a zero interval {0, None}.
//
// Examples of invalid durations are -5m (negative intervals are not allowed), 2h5m (intervals cannot mix several units of time) and 350ms (the smallest interval currently supported is one second).
func TryMakeInterval(d time.Duration) (Interval, error) {
	if d == 0 {
		return Interval{0, None}, nil
	}
	if d < 0 {
		return Interval{0, None}, ErrDurationNegative
	}

	for _, u := range units {
		invl := u.Duration()

		if d >= invl {
			v := d / invl

			var err error
			if v*invl != d {
				err = ErrDurationMixesUnits
			}

			return Interval{int(v), u}, err
		}
	}

	return Interval{0, None}, ErrDurationTooSmall
}

// MakeInterval returns an Interval corresponding to the given time.Duration. Panics if the given duration cannot be expressed as an interval. See ParseInterval for examples of invalid durations that would cause a panic.
func MakeInterval(d time.Duration) Interval {
	invl, err := TryMakeInterval(d)
	if err != nil {
		panic(err)
	}
	return invl
}

// IsNone returns whether the given interval has its Unit set to None. None generally means that no rounding should occur.
func (invl Interval) IsNone() bool {
	return invl.Unit == None
}

func (invl Interval) String() string {
	if invl.IsNone() {
		return "none"
	}
	return fmt.Sprintf("%d%s", invl.N, invl.Unit.String())
}

// Duration returns time.Duration corresponding to the given interval, e.g. 5*time.Minute for a {5, Minutes} interval.
func (invl Interval) Duration() time.Duration {
	return invl.Unit.Duration() * time.Duration(invl.N)
}

// AddTo advances the given time value by the given multiple of the specified interval. E.g. adding three 5-minute intervals to 9:37 gives 9:52. Note that, unlike NextInterval and PrevInterval, AddInterval does not round the time to the given interval.
func (invl Interval) AddTo(t time.Time, n int) time.Time {
	switch invl.Unit {
	case None:
		return t
	case Seconds, Minutes, Hours:
		return t.Add(invl.Duration() * time.Duration(n))
	case Days:
		return t.AddDate(0, 0, invl.N*n)
	default:
		panic("Invalid unit")
	}
}

// Next gives the start of the nth interval after the one that the given time value falls into. E.g. the next 5-minute interval after 9:37 is 9:40.
func (invl Interval) Next(t time.Time, n int) time.Time {
	t = invl.Round(t)
	return invl.AddTo(t, n)
}

// Prev gives the start of the nth interval prior to the one that the given time value falls into. E.g. the previous 5-minute interval before 9:37 is 9:30. PrevInterval is equivalent to calling NextInterval with a negative n.
func (invl Interval) Prev(t time.Time, n int) time.Time {
	return invl.Next(t, -n)
}

// Round returns the given time value rounded to this interval. E.g. 9:37 rounded to a 5-minute interval is 9:30; rounded to a 2-hour interval, it's 8:00. If the interval is None, the time value is returned unchanged.
func (invl Interval) Round(t time.Time) time.Time {
	switch invl.Unit {
	case None:
		return t
	case Seconds:
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()/invl.N*invl.N, 0, t.Location())
	case Minutes:
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute()/invl.N*invl.N, 0, 0, t.Location())
	case Hours:
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour()/invl.N*invl.N, 0, 0, 0, t.Location())
	case Days:
		// somewhat questionable
		return time.Date(t.Year(), t.Month(), t.Day()/invl.N*invl.N, 0, 0, 0, 0, t.Location())
	default:
		panic("Invalid unit")
	}
}

// FormatRounded returns the given time value rounded to this interval and formatted appropriately according to the specified set of format strings.
func (invl Interval) FormatRounded(t time.Time, fmt *FormatSet) string {
	return fmt.Format(invl.Round(t), invl.Unit)
}

// Round returns the given time value rounded to the interval specified by time.Duration. E.g. 9:37 rounded to a 5-minute interval is 9:30; rounded to a 2-hour interval, it's 8:00.
func Round(t time.Time, d time.Duration) time.Time {
	return MakeInterval(d).Round(t)
}

// FormatRounded returns the given time value rounded to the interval specified by time.Duration and formatted appropriately according to the specified set of format strings.
func FormatRounded(t time.Time, d time.Duration, fmt *FormatSet) string {
	return MakeInterval(d).FormatRounded(t, fmt)
}
