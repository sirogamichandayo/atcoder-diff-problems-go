package util_tool

import "time"

const Iso8601JstLayout = "2006-01-02T15:04:05+09:00"

var jst *time.Location

func ConvertIso8601JstToEpochTime(s string) (int64, error) {
	var err error
	if jst == nil {
		jst, err = time.LoadLocation("Asia/Tokyo")
		if err != nil {
			return 0, err
		}
	}

	t, err := time.ParseInLocation(Iso8601JstLayout, s, jst)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}
