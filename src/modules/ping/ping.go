package ping

func CanHandle(message string) bool {
	return message == "ping"
}

func SendPong(message string, SendMessage func(string)) {
	SendMessage("<:twipbbt:351795248961159168>")
}
