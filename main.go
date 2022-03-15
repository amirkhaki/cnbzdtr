package main 

import (
	"github.com/amirkhaki/cnbzdtr/config"

	"log"
	"os"
	"os/signal"
	"syscall"
	"strconv"

	"github.com/bwmarrin/discordgo"
)
var counter int
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	counter += 1
	_, err := s.ChannelMessageSend(m.ChannelID, "message"+strconv.Itoa(counter))
	if err != nil {
		log.Println(err)
	}
}


func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	dg, err := discordgo.New("Bot " + cfg.Token)
	if err != nil { log.Fatal(err)  } 
	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages
	err = dg.Open()
	if err != nil { log.Fatal(err)  } 
	log.Println("Bot is now running. Press CTRL-C to exit.")
    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <-sc

    // Cleanly close down the Discord session.
    dg.Close()
}

