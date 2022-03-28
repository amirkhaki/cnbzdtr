package handler

import (
	dg "github.com/bwmarrin/discordgo"
	"log"
	"fmt"
)

func (h *Handler) MessageCreate(s *dg.Session, m *dg.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	user, err := h.store.GetUserOrCreate(m.Author.ID)
	if err != nil {
		log.Println(err)
		return
	}
	seh := &SEH{s:s, lvls:h.lvls}
	user.OnMostScoreChange("SEH", seh)
	err = user.IncreaseScore(200)
	if err != nil {
		log.Println(err)
		return
	}
	err = h.store.UpdateUser(user)
	if err != nil {
		log.Println(err)
	}
	message := fmt.Sprintf("user %s\nlevel:%s", user.ID, h.lvls.Level(user.Score).Title)
	_, err = s.ChannelMessageSend(m.ChannelID, message)
	if err != nil {
		log.Println(err)
	}
}

