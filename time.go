package chronos

import (
	"bytes"
	"time"
)

// Time contains an embedded time assignment
// The reason for this type is to allow unmarshalling of empty strings into nil time types
type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, []byte(`""`)) != 0 {
		return t.Time.UnmarshalJSON(data)
	}

	return nil
}
