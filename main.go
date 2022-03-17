package main

import (
	"github.com/amirkhaki/cnbzdtr/config"
	"github.com/amirkhaki/cnbzdtr/store"
	"github.com/amirkhaki/cnbzdtr/entity"

	"log"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)


var imS = store.New()
func inviteCreate(s *discordgo.Session, i *discordgo.InviteCreate) {
	user, err := imS.GetUserByID(i.Inviter.ID)
	if err != nil {
		user = &entity.User{ID:i.Inviter.ID}
		err = imS.AddUser(user)
		if err != nil {
			log.Println(err)
			return
		}
	}
	err = user.IncreaseScore(2000)
	if err != nil {
		log.Println(err)
		return
	}
	err = imS.UpdateUser(user)
	if err != nil {
		log.Println(err)
	}
}
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	user, err := imS.GetUserByID(m.Author.ID)
	if err != nil {
		user = &entity.User{ID:m.Author.ID}
		err = imS.AddUser(user)
		if err != nil {
			log.Println(err)
			return
		}
	}
	err = user.IncreaseScore(200)
	if err != nil {
		log.Println(err)
		return
	}
	err = imS.UpdateUser(user)
	if err != nil {
		log.Println(err)
	}
	message := fmt.Sprintf("user %s\n score: %d", user.ID, user.Score)
	_, err = s.ChannelMessageSend(m.ChannelID, message)
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
	if err != nil {
		log.Fatal(err)
	}
	dg.AddHandler(messageCreate)
	dg.AddHandler(inviteCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentGuildInvites
	err = dg.Open()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
