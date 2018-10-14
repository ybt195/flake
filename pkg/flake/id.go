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

// FromBytes returns the id represented by the 8-byte byte array.
func FromBytes(bytes []byte) (ID, error) {
	if len(bytes) != 8 {
		return zeroID, fmt.Errorf("unexpected number of bytes for flake id: %d", len(bytes))
	}
	return ID(binary.BigEndian.Uint64(bytes)), nil
}

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

func (f ID) String() string {
	return fmt.Sprintf("{Bucket: %d, Time: %s, Sequence: %d}", f.Bucket(), f.Time(), f.Sequence())
}
