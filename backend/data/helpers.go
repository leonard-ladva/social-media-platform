package data

import "time"

func CurrentTime() int64 {
	return time.Now().UnixMilli()
}
