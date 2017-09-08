package main

import (
	"fmt"
	"github.com/Avalander/GoBot/src/modules/uptime"
)

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
