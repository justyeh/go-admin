package util

import (
	"strconv"
	"time"
)

type TimeStamp time.Time

func (ts TimeStamp) MarshalJSON() ([]byte, error) {
	origin := time.Time(ts)
	return []byte(strconv.FormatInt(origin.UnixNano()/1000000, 10)), nil
}

func (ts *TimeStamp) ToTime() time.Time {
	return time.Time(*ts)
}

func (ts *TimeStamp) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	millis, err := strconv.ParseInt(string(data), 10, 64)

	*ts = TimeStamp(time.Unix(0, millis*int64(time.Millisecond)))
	return err
}

func (ts TimeStamp) ToString() string {
	return ts.ToTime().Format("2006-01-02 15:04:05")
}
