package flake

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var timeFixture = FromStandardTime(time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC))
var formattedTimeFixture = "2015-01-01T00:00:00Z"

func TestSequenceUnavailable_Error(t *testing.T) {
	err := SequenceUnavailable{
		BucketID:  1,
		Timestamp: timeFixture,
	}
	assert.EqualError(t, err, fmt.Sprintf("maximum sequences hit for bucket 1: %s", formattedTimeFixture))
}

func TestTimeMovedBack_Error(t *testing.T) {
	new := FromStandardTime(time.Date(2014, 1, 1, 0, 0, 0, 0, time.UTC))
	err := TimeMovedBack{
		BucketID:      1,
		LastTimestamp: timeFixture,
		Timestamp:     new,
	}
	assert.EqualError(t, err, fmt.Sprintf(
		"time moved back for bucket 1: %s to %s", formattedTimeFixture, new))
}
