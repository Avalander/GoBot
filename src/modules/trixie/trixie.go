package trixie

import (
	"math/rand"
	"time"
)

var quotes = []string{
	"Hah! You think you're better than the Great and Powerful Trixie? You think you have more magical talent? Well, come on, show Trixie what you've got. Show us all.",
	"Hah! Once again, the Great and Powerful Trixie has proven herself to be the most amazing unicorn in all of Equestria. Was there ever any doubt?",
	"Come on, come all! Come and witness the amazing magic of the Great and Powerful Trixie!",
	"Anything you can do, I can do better!",
	"Ooo, what's the matter? Afraid you'll get a hair out of place in that rat's nest you call a mane?",
	"Pull, you fools! Somepony set off the magic force field, and Trixie intends to punish them!",
	"Tickles? That was supposed to make you writhe in agony!",
	"The Great and Powerful Trixie doesn't trust wheels.",
	"Don't you think the Great and Apologetic Trixie is the most magnificent humble pony you have ever seen?",
	"I am not sad at all! I definitely don't feel as if my heart is breaking into a million pieces!",
	"Cheated? *Moi?*",
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func CanHandle(message string) bool {
	return message == "trixie"
}

func SendQuote(message string, SendMessage func(string)) {
	index := rand.Intn(len(quotes))
	SendMessage(quotes[index])
}
