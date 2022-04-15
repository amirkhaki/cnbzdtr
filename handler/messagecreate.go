package handler

import (
	"context"
	"fmt"
	dg "github.com/bwmarrin/discordgo"
	"log"
)

func (h *Handler) MessageCreate(ctx context.Context, s *dg.Session, m *dg.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}
	user, err := h.store.GetUserOrCreate(ctx, m.Author.ID)
	if err != nil {
		log.Println(err)
		return
	}
	if m.Content == "!stats" {
		message := fmt.Sprintf("Level: %s\nScore: %d\nReferral Count: %d\nMessage Count: %d",
			h.lvls.Level(user.Score).Title, user.Score,
			user.ReferralCount,
			user.MessageCount)
		_, err = s.ChannelMessageSend(m.ChannelID, message)
		if err != nil {
			log.Println(err)
		}
		return
	}
	seh := &SEH{s: s, lvls: h.lvls, cfg: h.cfg}
	user.OnMostScoreChange("SEH", seh)
	user.AddMessage(m.ID)
	err = user.IncreaseScore(h.cfg.MessageScore)
	if err != nil {
		log.Println(err)
		return
	}
	err = h.store.UpdateUser(ctx, user)
	if err != nil {
		log.Println(err)
	}
	message := fmt.Sprintf("user %s\nlevel:%s", user.ID, h.lvls.Level(user.Score).Title)
	_, err = s.ChannelMessageSend(m.ChannelID, message)
	if err != nil {
		log.Println(err)
	}
}
