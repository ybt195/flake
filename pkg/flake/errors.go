package flake

import "fmt"

type SequenceUnavailable struct {
	BucketID  uint64
	Timestamp FlakeTime
}

func (err SequenceUnavailable) Error() string {
	return fmt.Sprintf("maximum sequences hit for bucket %d: %s", err.BucketID, err.Timestamp)
}

type TimeMovedBack struct {
	BucketID      uint64
	LastTimestamp FlakeTime
	Timestamp     FlakeTime
}

func (err TimeMovedBack) Error() string {
	return fmt.Sprintf("time moved back for bucket %d: %s to %s", err.BucketID, err.LastTimestamp, err.Timestamp)
}
