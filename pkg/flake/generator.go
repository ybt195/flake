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
	"sync"
)

// Generator generates unique flake ids.
type Generator struct {
	bucketID         uint64
	currentTimestamp uint64
	currentSequence  uint64
	lock             sync.Mutex
}

// New returns a new flake id generator configured with the bucket id.
func New(bucketID uint64) (*Generator, error) {
	if bucketID > bucketLimit {
		return nil, fmt.Errorf("bucket id must be between 0 and %d: %d provided", bucketLimit, bucketID)
	}
	return &Generator{
		bucketID:         bucketID,
		currentTimestamp: now(),
		currentSequence:  uint64(0),
	}, nil
}

// Next returns the next id in the generator. An error is returned if the wall clock moves back in
// time or if too many ids have been generated for the same timestamp.
func (g *Generator) Next() (ID, error) {
	timestamp := now()

	g.lock.Lock()
	if timestamp < g.currentTimestamp {
		g.lock.Unlock()
		return zeroID, TimeMovedBack{
			Bucket:        g.bucketID,
			LastTimestamp: Time(g.currentTimestamp),
			Timestamp:     Time(timestamp),
		}
	} else if timestamp > g.currentTimestamp {
		g.currentTimestamp = timestamp
		g.currentSequence = 0
	}
	sequence := g.currentSequence
	g.currentSequence++
	g.lock.Unlock()

	if sequence > sequenceLimit {
		return zeroID, SequenceUnavailable{Bucket: g.bucketID, Timestamp: Time(timestamp)}
	}

	return ID((g.bucketID << (timestampBits + sequenceBits)) | (timestamp << sequenceBits) | sequence), nil
}

// Must returns the next id in the generator. Must will block until an ID is available.
func (g *Generator) Must() ID {
	for {
		if id, err := g.Next(); err == nil {
			return id
		}
	}
}
