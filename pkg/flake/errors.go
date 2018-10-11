package flake

import "fmt"

type SequenceUnavailable struct {
	Bucket    uint64
	Timestamp Time
}

func (err SequenceUnavailable) Error() string {
	return fmt.Sprintf("maximum sequences hit for bucket %d: %s", err.Bucket, err.Timestamp)
}

type TimeMovedBack struct {
	Bucket        uint64
	LastTimestamp Time
	Timestamp     Time
}

func (err TimeMovedBack) Error() string {
	return fmt.Sprintf("time moved back for bucket %d: %s to %s", err.Bucket, err.LastTimestamp, err.Timestamp)
}
