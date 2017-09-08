package main

import (
	"fmt"
	"github.com/Avalander/GoBot/src/modules/uptime"
)

const MINUTE = 60
const HOUR = MINUTE * 60
const DAY = HOUR * 24
const WEEK = DAY * 7

func init() {
	fmt.Println("Hello")
}

func main() {
	var value = uptime.GetUptime()

	fmt.Println(value)

	fmt.Println("Weeks:", value.Weeks)
	fmt.Println("Days:", value.Days)
	fmt.Println("Hours:", value.Hours)
	fmt.Println("Minutes:", value.Minutes)
	fmt.Println("Seconds:", value.Seconds)
}
