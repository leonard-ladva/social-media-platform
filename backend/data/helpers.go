package data

import "time"

func currentTime() int64 {
	return time.Now().UnixMilli()
}
