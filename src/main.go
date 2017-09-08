package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/bwmarrin/discordgo"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error starting GoBot,", err)
		return
	}

	dg.AddHandler(onMessage)
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}

	fmt.Println("Bot is up and running.")
	signalClose := make(chan os.Signal, 1)
	signal.Notify(signalClose, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<- signalClose
	dg.Close()
}

func onMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	if message.Content == "ping" {
		session.ChannelMessageSend(message.ChannelID, "<:twipbbt:351795248961159168>")
	}
}
