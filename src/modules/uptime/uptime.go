package uptime

import (
	"time"
	"strconv"
)

var (
	start time.Time
)

type Uptime struct {
	Weeks int
	Days int
	Hours int
	Minutes int
	Seconds int
}

const SECOND = 1
const MINUTE = 60
const HOUR = MINUTE * 60
const DAY = HOUR * 24
const WEEK = DAY * 7

func init() {
	start = time.Now()
}

func GetUptime() Uptime {
	i, _ := strconv.ParseInt("1504010580", 10, 64)
	uptime := int64(time.Now().Sub(time.Unix(i, 0)).Seconds())
	return Uptime{
		convertDurationTo(WEEK, &uptime),
		convertDurationTo(DAY, &uptime),
		convertDurationTo(HOUR, &uptime),
		convertDurationTo(MINUTE, &uptime),
		convertDurationTo(SECOND, &uptime),
	}
}

func convertDurationTo(name int64, value *int64) int {
	newValue := *value / name
	*value = *value % name
	return int(newValue)
}
