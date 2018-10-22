/*
 * Copyright 2018 Jonathan Ben-tzur
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package flake

import (
	"encoding/binary"
	"fmt"
)

// Nil is the zero flake id.
const Nil = ID(0)

// Size returns the size (in bytes) of a flake id.
const Size = 8

// These constants define the distribution of bits in a flake id.
const (
	// BucketBits is the number of bits dedicated to the bucket of the id
	BucketBits = 10
	// TimestampBits is the number of bits dedicated to the timestamp (in milliseconds) of the id.
	TimestampBits = 42
	// SequenceBits is the number of bits dedicated to the sequence number of the id.
	SequenceBits = 12
)

// These constants define the upper limits of all the components in a flake id.
const (
	// BucketLimit is the maximum bucket id allowed.
	BucketLimit = 1<<BucketBits - 1
	// TimestampLimit is the maximum timestamp allowed.
	TimestampLimit = 1<<TimestampBits - 1
	// SequenceLimit is the maximum sequence number allowed.
	SequenceLimit = 1<<SequenceBits - 1
)

// ID is a unique 64-bit unsigned integer generated based on time.
type ID uint64

// FromBytes returns the id represented by the 8-byte byte array.
func FromBytes(bytes []byte) (ID, error) {
	if len(bytes) != Size {
		return Nil, fmt.Errorf("unexpected number of bytes for flake id: %d", len(bytes))
	}
	return ID(binary.BigEndian.Uint64(bytes)), nil
}

// Bucket returns the bucket component of the ID.
func (f ID) Bucket() uint64 {
	return f.Uint64() >> (TimestampBits + SequenceBits)
}

// Time returns the time component of the ID.
func (f ID) Time() Time {
	return Time(f.Uint64() << (BucketBits) >> (BucketBits + SequenceBits))
}

// Sequence returns the sequence component of the ID.
func (f ID) Sequence() uint64 {
	return f.Uint64() << (TimestampBits + BucketBits) >> (TimestampBits + BucketBits)
}

// Before returns true if this id comes before the provided id. Ordering is defined as first ordering
// by the timestamp, then by the bucket, and lastly by the sequence. While there is a defined global
// ordering of ids, the primary requirement for ordering is that there is ordering of time and then
// sequence for a bucket. Bucket ordering may not be entirely accurate due to clock skew between
// generators.
func (f ID) Before(o ID) bool {
	fTime, oTime := f.Time(), o.Time()
	if fTime.Equal(oTime) {
		fBucket, oBucket := f.Bucket(), o.Bucket()
		if fBucket == oBucket {
			return f.Sequence() < o.Sequence()
		}
		return fBucket < oBucket
	}
	return fTime.Before(oTime)
}

// After returns true if this id comes after the provided id. See `Before` for a full description of
// id ordering.
func (f ID) After(o ID) bool {
	return !f.Equal(o) && !f.Before(o)
}

// Equal returns true if the two ids are equal and false otherwise.
func (f ID) Equal(o ID) bool {
	return f == o
}

// Uint64 returns the uint64 representation of the ID.
func (f ID) Uint64() uint64 {
	return uint64(f)
}

// Bytes returns a big-endian encoded byte array of the ID.
func (f ID) Bytes() []byte {
	buf := make([]byte, Size)
	f.Put(buf)
	return buf
}

// Put places the id into the specified byte buffer. The buffer must be at least the size of
// a uint64. The buffer is big-endian encoded.
func (f ID) Put(buf []byte) {
	binary.BigEndian.PutUint64(buf, f.Uint64())
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

func (f ID) String() string {
	return fmt.Sprintf("{Bucket: %d, Time: %s, Sequence: %d}", f.Bucket(), f.Time().StandardTime(), f.Sequence())
}
