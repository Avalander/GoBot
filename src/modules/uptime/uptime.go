package uptime

import (
	"fmt"
	"strconv"
	"time"
)

var (
	start time.Time
)

const SECOND = 1
const MINUTE = 60
const HOUR = MINUTE * 60
const DAY = HOUR * 24
const WEEK = DAY * 7

func init() {
	start = time.Now()
}

func SendUptime(SendMessage func(string)) {
	i, _ := strconv.ParseInt("1504010580", 10, 64)
	uptime := int64(time.Now().Sub(time.Unix(i, 0)).Seconds())
	text := fmt.Sprintf("I've been up for %d weeks, %d days, %d hours, %d minutes and %d seconds.",
		convertDurationTo(WEEK, &uptime),
		convertDurationTo(DAY, &uptime),
		convertDurationTo(HOUR, &uptime),
		convertDurationTo(MINUTE, &uptime),
		convertDurationTo(SECOND, &uptime))
	SendMessage(text)
}

func convertDurationTo(name int64, value *int64) int {
	newValue := *value / name
	*value = *value % name
	return int(newValue)
}
