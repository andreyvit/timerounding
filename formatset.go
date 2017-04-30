package timerounding

import (
	"time"
)

// FormatSet defines format strings appropriate for formatting intervals at various time units.
type FormatSet struct {
	Seconds string
	Minutes string
	Hours   string
	Days    string
}

// A set of format strings that are reasonably concise for storage purposes, but are still possible to read at a glance: 20060102-150405, 20060102-1504, 20060102-15, 20060102.
var Concise = &FormatSet{
	Seconds: "20060102-150405",
	Minutes: "20060102-1504",
	Hours:   "20060102-15",
	Days:    "20060102",
}

/*
	Format returns a textual representation of the time value formatted appropriately for the given time unit. E.g. if the unit is Hours, Format uses fs.Hours or, if fs.Hours is empty, tries fs.Minutes, fs.Seconds and fs.Days, in that order.
*/
func (fs *FormatSet) Format(t time.Time, u Unit) string {
	if u >= Days && fs.Days != "" {
		return t.Format(fs.Days)
	}
	if u >= Hours && fs.Hours != "" {
		return t.Format(fs.Hours)
	}
	if u >= Minutes && fs.Minutes != "" {
		return t.Format(fs.Minutes)
	}
	if fs.Seconds != "" {
		return t.Format(fs.Seconds)
	}
	if fs.Minutes != "" {
		return t.Format(fs.Minutes)
	}
	if fs.Hours != "" {
		return t.Format(fs.Hours)
	}
	if fs.Days != "" {
		return t.Format(fs.Days)
	}
	panic("invalid timerounding.Format")
}
