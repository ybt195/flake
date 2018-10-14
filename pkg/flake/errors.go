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

import "fmt"

// SequenceUnavailable represents an error when trying to generate an id where the maximum number of
// ids have been generated for that timestamp.
type SequenceUnavailable struct {
	Bucket    uint64
	Timestamp Time
}

func (err SequenceUnavailable) Error() string {
	return fmt.Sprintf("maximum sequences hit for bucket %d: %s", err.Bucket, err.Timestamp)
}

// TimeMovedBack represents an error when trying to generate an id where the system clock has been
// moved back since the last recorded time.
type TimeMovedBack struct {
	Bucket        uint64
	LastTimestamp Time
	Timestamp     Time
}

func (err TimeMovedBack) Error() string {
	return fmt.Sprintf("time moved back for bucket %d: %s to %s", err.Bucket, err.LastTimestamp, err.Timestamp)
}
