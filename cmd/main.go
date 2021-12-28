package main

import (
	"fmt"
	"telegram-bot-planning-poker/cmd/bot"
)

func main() {
	conf := bot.NewConfig()
	if err := bot.ParseConfig(conf); err != nil {
		panic(err.Error())
	}

	c := make(chan string)
	b := bot.NewBot(conf)
	go b.Listen(c)

	for s := range c {
		fmt.Println(s)
	}
}
