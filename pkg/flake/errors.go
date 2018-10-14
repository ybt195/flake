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
