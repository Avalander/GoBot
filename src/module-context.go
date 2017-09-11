package main

import (
	"github.com/Avalander/GoBot/src/modules/uptime"
)

type command struct {
	canHandle func(string) bool
	handle    func(string, func(string))
}

var (
	commands []command
)

func init() {
	commands = append(commands, command{uptime.CanHandle, uptime.SendUptime})
}

func Handle(message string, SendMessage func(string)) {
	for i := 0; i < len(commands); i++ {
		if commands[i].canHandle(message) {
			commands[i].handle(message, SendMessage)
		}
	}
}
