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

// Before returns true if this time is before the provided time and false otherwise.
func (f Time) Before(o Time) bool {
	return f < o
}

// After returns true if this time is after the provided time and false otherwise.
func (f Time) After(o Time) bool {
	return f > o
}

// Equal returns true if the two times are equal to one another and false otherwise.
func (f Time) Equal(o Time) bool {
	return f == o
}

func now() uint64 {
	return toMillis(time.Now()) - epoch
}

func toMillis(t time.Time) uint64 {
	return uint64(t.UnixNano()) / uint64(time.Millisecond)
}
