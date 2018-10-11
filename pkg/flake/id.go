package flake

import (
	"encoding/binary"
	"fmt"
)

const (
	zeroID = ID(0)

	bucketBits    = 10
	timestampBits = 42
	sequenceBits  = 12

	bucketLimit   = 1<<bucketBits - 1
	sequenceLimit = 1<<sequenceBits - 1
)

type ID uint64

func (f ID) HostID() uint64 {
	return f.Uint64() >> (timestampBits + sequenceBits)
}

func (f ID) Time() Time {
	return Time(f.Uint64() << (bucketBits) >> (bucketBits + sequenceBits))
}

func (f ID) Sequence() uint64 {
	return f.Uint64() << (timestampBits + bucketBits) >> (timestampBits + bucketBits)
}

func (f ID) String() string {
	return fmt.Sprintf("{Host: %d, Time: %s, Sequence: %d}", f.HostID(), f.Time(), f.Sequence())
}

func (f ID) Uint64() uint64 {
	return uint64(f)
}

func (f ID) Bytes() []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, f.Uint64())
	return buf
}

func (f ID) Binary() string {
	return fmt.Sprintf("%064b", f.Uint64())
}

func (f ID) Octal() string {
	return fmt.Sprintf("%022o", f.Uint64())
}

func (f ID) Hex() string {
	return fmt.Sprintf("%016x", f.Uint64())
}
