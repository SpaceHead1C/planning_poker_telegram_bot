package main

import (
	"fmt"
	"sync"
	"telegram-bot-planning-poker/cmd/bot"
)

func main() {
	conf := bot.NewConfig()
	if err := bot.ParseConfig(conf); err != nil {
		panic(err.Error())
	}

	var wg sync.WaitGroup
	wg.Add(1)
	
	b := bot.NewBot(conf)
	fmt.Printf("Bot %s OK", b)

	wg.Wait()
}
