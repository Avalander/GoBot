package uptime

import (
	"fmt"
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
	//i, _ := strconv.ParseInt("1504010580", 10, 64)
	uptime := int64(time.Now().Sub(start).Seconds())
	text := fmt.Sprintf("I've been up for %s%s%s%s%d seconds.",
		durationToString(convertDurationTo(week, &uptime), "%d weeks, "),
		durationToString(convertDurationTo(day, &uptime), "%d days, "),
		durationToString(convertDurationTo(hour, &uptime), "%d hours, "),
		durationToString(convertDurationTo(minute, &uptime), "%d minutes and "),
		convertDurationTo(second, &uptime))
	SendMessage(text)
}

func convertDurationTo(name int64, value *int64) int {
	newValue := *value / name
	*value = *value % name
	return int(newValue)
}

func durationToString(value int, template string) string {
	if value == 0 {
		return ""
	}
	return fmt.Sprintf(template, value)
}
