package xtime

import "time"

// Now 获取当前时间YYYY-MM-DD HH:MM:SS
func Now() string {
	t := time.Now()
	formattedTime := t.Format("2006-01-02 15:04:05")
	return formattedTime
}

// Time2Str 转换取当前时间戳
func Time2Str(t time.Time) string {
	formattedTime := t.Format("2006-01-02 15:04:05")
	return formattedTime
}

// Unix2Str 转换时间戳
func Unix2Str(timestamp int64) string {
	t := time.Unix(int64(timestamp), 0)
	formatTime := Time2Str(t)
	return formatTime
}

// Str2Time 转换时间
func Str2Time(t string) (time.Time, error) {
	parse, err := time.Parse("2006-01-02 15:04:05", t)
	if err != nil {
		return time.Now(), err
	}
	return parse, nil
}
