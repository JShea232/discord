package main

// Program Name:	discord.go
// Author Name:		Jordan Edward Shea <jes7923@rit.edu>
// Description:		This is the main program for launching the the Discord bot.

import (
	"os"
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
	"log"
	"fmt"
	"os/signal"
	"syscall"
)

func main() {
	// Bot token should be stored securely as an environment variable
	var fightBotToken = os.Getenv("fightBotToken")
	if fightBotToken == "" {
		log.Println("Error: fightBotToken was not found as an" +
			" environment variable.")
		return
	}
	// Attempts to create fight-bot using provided token
	fightBot, err := discordgo.New("Bot " + fightBotToken)
	if err != nil {
		log.Println("Error creating the Discord session: ",
			errors.WithStack(err))
		return
	}
	// Error will be thrown if an invalid token was provided earlier
	err = fightBot.Open()
	if err != nil {
		log.Println("Error with opening the connection: ",
			errors.WithStack(err))
		return
	}
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	// Info on Go signals can be found at: https://gobyexample.com/signals
	// sc will be notified of 4 types of signals (SIGINT, SIGTERM, Inter, Kill)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	// Cleanly close the Discord session.
	fightBot.Close()

}
