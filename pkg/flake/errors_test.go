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
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var timeFixture = FromStandardTime(time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC))

const formattedTimeFixture = "2015-01-01T00:00:00Z"

func TestSequenceUnavailable_Error(t *testing.T) {
	err := SequenceUnavailable{
		Bucket:    1,
		Timestamp: timeFixture,
	}
	assert.EqualError(t, err, fmt.Sprintf("maximum sequences hit for bucket 1: %s", formattedTimeFixture))
}

func TestTimeMovedBack_Error(t *testing.T) {
	new := FromStandardTime(time.Date(2014, 1, 1, 0, 0, 0, 0, time.UTC))
	err := TimeMovedBack{
		Bucket:        1,
		LastTimestamp: timeFixture,
		Timestamp:     new,
	}
	assert.EqualError(t, err, fmt.Sprintf(
		"time moved back for bucket 1: %s to %s", formattedTimeFixture, new))
}
