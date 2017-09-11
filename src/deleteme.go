package main

import (
	"fmt"
)

func init() {
	fmt.Println("Hello")
}

func main() {
	sendMessage := func(text string) {
		fmt.Println(text)
	}
	Handle("uptime", sendMessage)
}
