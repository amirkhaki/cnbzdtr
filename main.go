package main

import (
	"github.com/amirkhaki/cnbzdtr/config"
	"github.com/amirkhaki/cnbzdtr/entity"
	"github.com/amirkhaki/cnbzdtr/handler"
	"github.com/amirkhaki/cnbzdtr/store"

	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var h *handler.Handler
var ctx = context.Background()

func init() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	imS, err := store.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	lvls := entity.NewLevels()
	lvls.AddLevel(entity.Level{From: 0, Title: "Zero Level"})
	lvls.AddLevel(entity.Level{From: 300, Title: "First Level"})
	lvls.AddLevel(entity.Level{From: 600, Title: "Second Level"})
	lvls.AddLevel(entity.Level{From: 2600, Title: "Third Level"})
	h = handler.New(lvls, imS)
}
func inviteCreate(s *discordgo.Session, i *discordgo.InviteCreate) {
	h.InviteCreate(ctx, s, i)
}
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	h.MessageCreate(ctx, s, m)
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
