package opendsd

import (
	"strings"
	"time"
)

const (
	APIRoot                      = "http://opendsd.sandiego.gov/api"
	HeaderExtractTimestampFormat = "1/2/2006 3:04:05 PM"
	TimestampFormat              = "2006-01-02T15:04:05"
)

type HeaderExtractTimestamp time.Time

func (aet *HeaderExtractTimestamp) UnmarshalJSON(data []byte) error {
	tstr := strings.Trim(string(data), "\"")
	if tstr == "null" || tstr == "" {
		return nil
	}

	v, err := time.Parse(HeaderExtractTimestampFormat, tstr)
	if err != nil {
		return err
	}
	*aet = HeaderExtractTimestamp(v)

	return nil
}

type Timestamp time.Time

func (at *Timestamp) UnmarshalJSON(data []byte) error {
	tstr := strings.Trim(string(data), "\"")
	if tstr == "null" || tstr == "" {
		return nil
	}

	v, err := time.Parse(TimestampFormat, tstr)
	if err != nil {
		return err
	}
	*at = Timestamp(v)

	return nil
}
