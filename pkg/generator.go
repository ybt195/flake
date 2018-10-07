package pkg

import (
	"fmt"
	"sync"
)

type Generator struct {
	bucketID         uint64
	currentTimestamp uint64
	currentSequence  uint64
	lock             sync.Mutex
}

func New(bucketID uint64) (*Generator, error) {
	if bucketID > bucketLimit {
		return nil, fmt.Errorf("bucket id must be between 0 and %d: %d provided", bucketID, bucketLimit)
	}
	return &Generator{
		currentTimestamp: now(),
		currentSequence:  uint64(0),
	}, nil
}

func (g *Generator) Next() (FlakeID, error) {
	timestamp := now()

	g.lock.Lock()
	if timestamp < g.currentTimestamp {
		g.lock.Unlock()
		return zeroID, TimeMovedBack{
			BucketID:      g.bucketID,
			LastTimestamp: FlakeTime(g.currentTimestamp),
			Timestamp:     FlakeTime(timestamp),
		}
	} else if timestamp > g.currentTimestamp {
		g.currentTimestamp = timestamp
		g.currentSequence = 0
	}
	sequence := g.currentSequence
	g.currentSequence++
	g.lock.Unlock()

	if sequence > sequenceLimit {
		return zeroID, SequenceUnavailable{BucketID: g.bucketID, Timestamp: FlakeTime(timestamp)}
	}

	return FlakeID((g.bucketID << (timestampBits + sequenceBits)) | (timestamp << sequenceBits) | sequence), nil
}

func (g *Generator) Must() FlakeID {
	for {
		if id, err := g.Next(); err == nil {
			return id
		}
	}
}
