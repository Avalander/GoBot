package debug

import "fmt"

func CanHandle(message string) bool {
	return true
}

func DebugMessage(message string, SendMessage func(string)) {
	fmt.Printf("[Message] %s\n", message)
}
