package main

import (
	"GoBot/src/modules/debug"
	"GoBot/src/modules/ping"
	"GoBot/src/modules/uptime"
)

type command struct {
	canHandle func(string) bool
	handle    func(string, func(string))
}

var (
	commands []command
)

func init() {
	commands = append(commands,
		command{debug.CanHandle, debug.DebugMessage},
		command{uptime.CanHandle, uptime.SendUptime},
		command{ping.CanHandle, ping.SendPong})
}

func Handle(message string, SendMessage func(string)) {
	for i := 0; i < len(commands); i++ {
		if commands[i].canHandle(message) {
			commands[i].handle(message, SendMessage)
		}
	}
}
