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
	"fmt"
	"strings"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

var h *handler.Handler
var ctx = context.Background()
func readLevels(lvls *entity.Levels) {
	i := 1
	for {
		lvl := os.Getenv(fmt.Sprintf("CD_LEVEL_%d", i))
		if lvl == "" {
			break
		}
		lvl_parts := strings.Split(lvl, "|")
		if len(lvl_parts) < 3 {
			break
		}
		from, err := strconv.Atoi(lvl_parts[0])
		if err != nil {
			log.Fatal(err)
		}
		lvls.AddLevel(entity.Level{From:uint64(from), Title:lvl_parts[1], Url: lvl_parts[2]})
		i += 1
	}
}
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
	readLevels(lvls)
	h = handler.New(lvls, imS, cfg)
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
