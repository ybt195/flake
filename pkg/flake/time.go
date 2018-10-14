package flake

import (
	"time"
)

const (
	epoch = uint64(1262304000000) // Friday, January 1, 2010 12:00:00 AM
)

// Time represents the time component of the flake id.
type Time uint64

// FromStandardTime returns a flake timestamp from standard go time.
func FromStandardTime(standard time.Time) Time {
	return Time(toMillis(standard) - epoch)
}

// StandardTime returns the standard go time from the flake timestamp.
func (f Time) StandardTime() time.Time {
	return time.Unix(0, int64(uint64(f)+epoch)*int64(time.Millisecond)).In(time.UTC)
}

func (f Time) String() string {
	return f.StandardTime().Format("2006-01-02T15:04:05.999Z07:00")
}

func now() uint64 {
	return toMillis(time.Now()) - epoch
}

func toMillis(t time.Time) uint64 {
	return uint64(t.UnixNano()) / uint64(time.Millisecond)
}
