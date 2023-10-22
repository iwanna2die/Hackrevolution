package utils

import (
	"strconv"
	"time"
)

func MakeDuration(tm string) (time.Duration, error) {
	dur, err := strconv.ParseInt(tm, 10, 64)
	if err != nil {
		return 0, err
	}
	return time.Minute * time.Duration(dur), nil
}
