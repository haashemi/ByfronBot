package main

import (
	"fmt"
	"log"
	"time"

	"github.com/haashemi/ByfronBot/commands"
	"github.com/haashemi/ByfronBot/pkg/bonbast"
	"github.com/haashemi/ByfronBot/pkg/config"
	"github.com/haashemi/tgo"
	"github.com/haashemi/tgo/filters"
	"github.com/haashemi/tgo/routers/message"
)

func main() {
	config, err := config.ParseConfig()
	if err != nil {
		log.Fatalln("Failed to parse config", err)
		return
	}

	bonbastClient, err := bonbast.NewClient(config.BonbastProxy)
	if err != nil {
		log.Fatalln("Failed to initialize bonbast client", err)
		return
	}

	bot := tgo.NewBot(config.TelegramToken, tgo.Options{
		Host:             config.TelegramHost,
		DefaultParseMode: tgo.ParseModeHTML,
	})

	info, err := bot.GetMe()
	if err != nil {
		log.Fatalln("Failed to fetch the bot info", err)
		return
	}

	cmd := commands.NewCommands(bonbastClient)
	mr := message.NewRouter()
	mr.Handle(filters.Command("me", info.Username), cmd.Me)
	mr.Handle(filters.Command("arz", info.Username), cmd.Arz)
	mr.Handle(filters.Command("stp", info.Username), cmd.STP)
	mr.Handle(filters.Command("ptss", info.Username), cmd.PTSS)
	mr.Handle(filters.Command("time", info.Username), cmd.Time)
	mr.Handle(filters.Command("server", info.Username), cmd.Server)
	bot.AddRouter(mr)

	bot.SetMyCommands(&tgo.SetMyCommands{
		Commands: []*tgo.BotCommand{
			{Command: "me", Description: "Your info"},
			{Command: "arz", Description: "IRT exchange rates"},
			{Command: "stp", Description: "Sticker to picture"},
			{Command: "ptss", Description: "Picture to sticker sized document"},
			{Command: "time", Description: "Current time"},
			{Command: "server", Description: "Server's status"},
		},
	})

	for {
		log.Println("Polling started")
		if err = bot.StartPolling(30); err != nil {
			fmt.Println("POLLING STOPPED", err)
			time.Sleep(5 * time.Second)
		}
	}
}
