package tformat

import "time"

func FormatJKTTime(t string) (string, error) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	var tz time.Time
	tz, err := time.ParseInLocation(time.RFC3339, t, loc)

	if err != nil {
		return "", err
	}

	return tz.String(), nil
}
