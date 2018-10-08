package flake

import (
	"fmt"
)

const (
	zeroID = FlakeID(0)

	bucketBits    = 10
	timestampBits = 42
	sequenceBits  = 12

	bucketLimit   = 1<<bucketBits - 1
	sequenceLimit = 1<<sequenceBits - 1
)

type FlakeID uint64

func (f FlakeID) HostID() uint64 {
	return f.Uint64() >> (timestampBits + sequenceBits)
}

func (f FlakeID) FlakeTime() FlakeTime {
	return FlakeTime(f.Uint64() << (bucketBits) >> (bucketBits + sequenceBits))
}

func (f FlakeID) Sequence() uint64 {
	return f.Uint64() << (timestampBits + bucketBits) >> (timestampBits + bucketBits)
}

func (f FlakeID) String() string {
	return fmt.Sprintf("{Host: %d, Time: %s, Sequence: %d}", f.HostID(), f.FlakeTime(), f.Sequence())
}

func (f FlakeID) Uint64() uint64 {
	return uint64(f)
}

func (f FlakeID) Binary() string {
	return fmt.Sprintf("%064b", f.Uint64())
}

func (f FlakeID) Octal() string {
	return fmt.Sprintf("%022o", f.Uint64())
}

func (f FlakeID) Hex() string {
	return fmt.Sprintf("%016x", f.Uint64())
}
