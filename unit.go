package timerounding

import (
	"time"
)

// Unit represents a unit of time: Days, Hours, Minutes or Seconds. (Months, years, milliseconds would be possible as well, but not currently implemented.)
type Unit int

const (
	None Unit = iota
	Seconds
	Minutes
	Hours
	Days
)

func (u Unit) String() string {
	switch u {
	case None:
		return "none"
	case Seconds:
		return "s"
	case Minutes:
		return "m"
	case Hours:
		return "h"
	case Days:
		return "d"
	default:
		panic("Invalid unit")
	}
}

// Duration returns time.Duration corresponding to the given time unit, e.g. time.Minute for timerounding.Minutes.
func (u Unit) Duration() time.Duration {
	switch u {
	case None:
		return 0
	case Seconds:
		return time.Second
	case Minutes:
		return time.Minute
	case Hours:
		return time.Hour
	case Days:
		return 24 * time.Hour
	default:
		panic("Invalid unit")
	}
}
