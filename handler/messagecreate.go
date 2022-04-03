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
		message := fmt.Sprintf("Level: %s\nScore: %d\nReferral Count: %d", 
				h.lvls.Level(user.Score).Title, user.Score,
				user.ReferralCount)
		_, err = s.ChannelMessageSend(m.ChannelID, message)
		if err != nil {
			log.Println(err)
		}
		return
	}
	seh := &SEH{s: s, lvls: h.lvls, cfg: h.cfg}
	user.OnMostScoreChange("SEH", seh)
	err = user.IncreaseScore(200)
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
