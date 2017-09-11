package main

import (
	"fmt"

	"github.com/Avalander/GoBot/src/modules/uptime"
)

func init() {
	fmt.Println("Hello")
}

func main() {
	sendMessage := func(text string) {
		fmt.Println(text)
	}
	uptime.SendUptime(sendMessage)
}
