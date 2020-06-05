package internal

import (
	"time"
)

func JstTime() time.Time {
	now := time.Now().UTC()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	return now.In(jst)
}
