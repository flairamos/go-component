package xtime

import "time"

func Now() string {
	t := time.Now()
	formattedTime := t.Format("2006-01-02 15:04:05")
	return formattedTime
}
