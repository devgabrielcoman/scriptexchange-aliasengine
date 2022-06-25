package utils

import "time"

func DateFormat(d int64) string {
	t := time.Unix(d, 0)
	layout := "2006-01-02 15:04:05"
	return t.Format(layout)
}

func CurrentTime() int64 {
	return time.Now().Unix()
}
