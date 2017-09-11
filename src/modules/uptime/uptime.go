package uptime

import (
	"fmt"
	"strconv"
	"time"
)

var (
	start time.Time
)

const second = 1
const minute = 60
const hour = minute * 60
const day = hour * 24
const week = day * 7

func init() {
	start = time.Now()
}

func CanHandle(message string) bool {
	return message == "uptime"
}

func SendUptime(message string, SendMessage func(string)) {
	i, _ := strconv.ParseInt("1504010580", 10, 64)
	uptime := int64(time.Now().Sub(time.Unix(i, 0)).Seconds())
	text := fmt.Sprintf("I've been up for %d weeks, %d days, %d hours, %d minutes and %d seconds.",
		convertDurationTo(week, &uptime),
		convertDurationTo(day, &uptime),
		convertDurationTo(hour, &uptime),
		convertDurationTo(minute, &uptime),
		convertDurationTo(second, &uptime))
	SendMessage(text)
}

func convertDurationTo(name int64, value *int64) int {
	newValue := *value / name
	*value = *value % name
	return int(newValue)
}
