package flake

import (
	"time"
)

const (
	epoch = uint64(1262304000000)
)

type FlakeTime uint64

func (f FlakeTime) Timestamp() time.Time {
	return time.Unix(0, int64(uint64(f)+epoch)*int64(time.Millisecond)).In(time.UTC)
}

func (f FlakeTime) String() string {
	return f.Timestamp().Format("2006-01-02T15:04:05.999Z07:00")
}

func now() uint64 {
	return toMillis(time.Now()) - epoch
}

func toMillis(t time.Time) uint64 {
	return uint64(t.UnixNano()) / uint64(time.Millisecond)
}
