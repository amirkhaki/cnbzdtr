package handler

import (
	dg "github.com/bwmarrin/discordgo"
	"log"
	"context"
)

func (h *Handler) InviteCreate(ctx context.Context, s *dg.Session, i *dg.InviteCreate) {
	user, err := h.store.GetUserOrCreate(ctx, i.Inviter.ID)
	if err != nil {
		log.Println(err)
		return
	}
	seh := &SEH{s:s, lvls:h.lvls}
	user.OnMostScoreChange("SEH", seh)
	err = user.IncreaseScore(2000)
	if err != nil {
		log.Println(err)
		return
	}
	err = h.store.UpdateUser(ctx, user)
	if err != nil {
		log.Println(err)
	}
}
