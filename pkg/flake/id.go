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

// ID is a unique 64-bit unsigned integer generated based on time.
type ID uint64

// Bucket returns the bucket component of the ID.
func (f ID) Bucket() uint64 {
	return f.Uint64() >> (timestampBits + sequenceBits)
}

// Time returns the time component of the ID.
func (f ID) Time() Time {
	return Time(f.Uint64() << (bucketBits) >> (bucketBits + sequenceBits))
}

// Sequence returns the sequence component of the ID.
func (f ID) Sequence() uint64 {
	return f.Uint64() << (timestampBits + bucketBits) >> (timestampBits + bucketBits)
}

func (f ID) String() string {
	return fmt.Sprintf("{Bucket: %d, Time: %s, Sequence: %d}", f.Bucket(), f.Time(), f.Sequence())
}

// Uint64 returns the uint64 representation of the ID.
func (f ID) Uint64() uint64 {
	return uint64(f)
}

// Bytes returns a big-endian encoded byte array of the ID.
func (f ID) Bytes() []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, f.Uint64())
	return buf
}

// Binary returns the base-2 representation of the ID.
func (f ID) Binary() string {
	return fmt.Sprintf("%064b", f.Uint64())
}

// Octal returns the base-8 representation of the ID.
func (f ID) Octal() string {
	return fmt.Sprintf("%022o", f.Uint64())
}

// Hex returns the base-16 representation of the ID.
func (f ID) Hex() string {
	return fmt.Sprintf("%016x", f.Uint64())
}
