package usecase

import (
	"fmt"
	"time"
)

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	defaultLocation, _ := time.LoadLocation("Asia/Jakarta")
	parsed, err := time.ParseInLocation(time.RFC3339, time.Time(t).String(), defaultLocation)
	if err != nil {
		return []byte(""), err
	}
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", parsed.Format(time.RFC3339))
	return []byte(stamp), nil
}
